package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/robert-nix/ansihtml"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/cockroachdb/pebble/v2"
)

//go:embed templates/dashboard.html
var templateFS embed.FS

var pageTmpl = template.Must(
	template.ParseFS(templateFS, "templates/dashboard.html"),
)

// ── Stats collection ──────────────────────────────────────────────────────────

func collectStats(hostname string) (NodeStats, error) {
	cpuPcts, err := cpu.Percent(200*time.Millisecond, true)
	if err != nil {
		return NodeStats{}, err
	}
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return NodeStats{}, err
	}
	loadStat, err := load.Avg()
	if err != nil {
		return NodeStats{}, err
	}
	return NodeStats{
		Name:      hostname,
		CPU:       cpuPcts,
		MemUsed:   vmStat.Used,
		MemTotal:  vmStat.Total,
		Load:      [3]float64{loadStat.Load1, loadStat.Load5, loadStat.Load15},
		UpdatedAt: time.Now().UnixNano(),
	}, nil
}

// ── ANSI rendering ────────────────────────────────────────────────────────────

const barWidth = 20

var (
	styleGreen  = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	styleBlue   = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	styleYellow = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	styleRed    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	styleDim    = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)

func pctBar(pct float64, width int, style lipgloss.Style) string {
	filled := int(math.Round(pct / 100.0 * float64(width)))
	if filled > width {
		filled = width
	}
	return style.Render(strings.Repeat("█", filled)) +
		styleDim.Render(strings.Repeat("░", width-filled))
}

func renderANSI(s NodeStats) string {
	var sb strings.Builder
	age := time.Since(time.Unix(0, s.UpdatedAt))
	offline := s.UpdatedAt == 0 || age > 15*time.Second

	status := styleGreen.Render("●")
	if offline {
		status = styleRed.Render("●")
	}
	sb.WriteString(fmt.Sprintf("%s %s\n", status, s.Name))
	if offline {
		sb.WriteString(styleDim.Render("  offline\n"))
		return sb.String()
	}
	sb.WriteString(fmt.Sprintf("  updated %.0fs ago\n", age.Seconds()))
	sb.WriteString(styleDim.Render(strings.Repeat("─", barWidth+14) + "\n"))

	for i, pct := range s.CPU {
		bar := pctBar(pct, barWidth, styleGreen)
		sb.WriteString(fmt.Sprintf("CPU%-2d [%s] %5.1f%%\n", i, bar, pct))
	}

	memPct := 0.0
	if s.MemTotal > 0 {
		memPct = float64(s.MemUsed) / float64(s.MemTotal) * 100
	}
	sb.WriteString(fmt.Sprintf("Mem   [%s] %5.1f%%\n", pctBar(memPct, barWidth, styleBlue), memPct))
	sb.WriteString(fmt.Sprintf("      %s / %s\n", fmtBytes(s.MemUsed), fmtBytes(s.MemTotal)))

	loadStyle := styleGreen
	if s.Load[0] > 2.0 {
		loadStyle = styleYellow
	}
	if s.Load[0] > 4.0 {
		loadStyle = styleRed
	}
	sb.WriteString(fmt.Sprintf("Load  %s  %.2f  %.2f  %.2f\n",
		loadStyle.Render("▶"), s.Load[0], s.Load[1], s.Load[2]))

	return sb.String()
}

func fmtBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// ── Layout ────────────────────────────────────────────────────────────────────

type layoutParams struct {
	CellWidth float64
	FontSize  float64
}

func computeLayout(nodeCount, winW, winH int) layoutParams {
	if nodeCount == 0 {
		nodeCount = 1
	}
	aspect := 16.0 / 9.0
	if winW > 0 && winH > 0 {
		aspect = float64(winW) / float64(winH)
	}
	cols := int(math.Round(math.Sqrt(float64(nodeCount) * aspect)))
	if cols < 1 {
		cols = 1
	}
	if cols > nodeCount {
		cols = nodeCount
	}
	cw := 100.0 / float64(cols)
	return layoutParams{CellWidth: cw, FontSize: cw * 0.016}
}
func avgCPU(s NodeStats) float64 {
	if len(s.CPU) == 0 {
		return 0
	}
	var sum float64
	for _, v := range s.CPU {
		sum += v
	}
	return sum / float64(len(s.CPU))
}

func nodeOnline(s NodeStats) bool {
	return s.UpdatedAt != 0 && time.Since(time.Unix(0, s.UpdatedAt)) <= 15*time.Second
}

func computeRefreshIntervalMs(nodes []NodeStats) int {
	if len(nodes) == 0 {
		return 3000
	}
	maxCPU := 0.0
	maxLoad := 0.0
	for _, s := range nodes {
		if !nodeOnline(s) {
			continue
		}
		if cpu := avgCPU(s); cpu > maxCPU {
			maxCPU = cpu
		}
		if s.Load[0] > maxLoad {
			maxLoad = s.Load[0]
		}
	}

	switch {
	case maxCPU < 30 && maxLoad < 1.0:
		return 2000
	case maxCPU < 55 && maxLoad < 2.0:
		return 3500
	case maxCPU < 80 && maxLoad < 4.0:
		return 6500
	default:
		return 11000
	}
}

func findBestNodeHint(nodes []NodeStats) string {
	best := ""
	bestScore := math.MaxFloat64
	for _, s := range nodes {
		if !nodeOnline(s) {
			continue
		}
		score := avgCPU(s) + s.Load[0]*10
		if score < bestScore {
			bestScore = score
			best = fmt.Sprintf("%s (%.0f%% cpu, %.2f load)", s.Name, avgCPU(s), s.Load[0])
		}
	}
	if best == "" {
		return "no responsive peers yet"
	}
	return "lowest-load node: " + best
}

// ── HTTP handler ──────────────────────────────────────────────────────────────

type cellData struct {
	Name    string
	URL     string
	HTML    template.HTML
	Offline bool
	Focused bool
}

type pageData struct {
	Layout        layoutParams
	Nodes         []cellData
	RefreshMs     int
	RefreshLabel  string
	BestHint      string
	Focus         string
	ClearFocusURL string
}

func makeHandler(db *pebble.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		winW, winH := 0, 0
		fmt.Sscanf(r.URL.Query().Get("w"), "%d", &winW)
		fmt.Sscanf(r.URL.Query().Get("h"), "%d", &winH)

		nodes, err := dbScanAll(db)
		if err != nil {
			http.Error(w, "db error", 500)
			return
		}

		focus := r.URL.Query().Get("focus")
		layout := computeLayout(len(nodes), winW, winH)
		cells := make([]cellData, 0, len(nodes))
		for _, s := range nodes {
			htmlBytes := ansihtml.ConvertToHTML([]byte(renderANSI(s)))
			offline := s.UpdatedAt == 0 || time.Since(time.Unix(0, s.UpdatedAt)) > 15*time.Second

			query := r.URL.Query()
			copyQuery := make(map[string][]string, len(query))
			for k, v := range query {
				copyQuery[k] = append([]string(nil), v...)
			}
			queryValues := url.Values(copyQuery)
			queryValues.Set("focus", s.Name)
			queryValues.Set("w", fmt.Sprintf("%d", winW))
			queryValues.Set("h", fmt.Sprintf("%d", winH))
			urlStr := "?" + queryValues.Encode()

			cells = append(cells, cellData{
				Name:    s.Name,
				URL:     urlStr,
				HTML:    template.HTML(htmlBytes),
				Offline: offline,
				Focused: focus != "" && focus == s.Name,
			})
		}

		filtered := cells
		if focus != "" {
			filtered = nil
			for _, cell := range cells {
				if cell.Focused {
					filtered = append(filtered, cell)
				}
			}
			if len(filtered) == 0 {
				filtered = cells
			}
		}

		refreshMs := computeRefreshIntervalMs(nodes)
		bestHint := findBestNodeHint(nodes)
		clearQuery := r.URL.Query()
		copyClearQuery := make(map[string][]string, len(clearQuery))
		for k, v := range clearQuery {
			copyClearQuery[k] = append([]string(nil), v...)
		}
		clearValues := url.Values(copyClearQuery)
		clearValues.Del("focus")
		clearFocusURL := "/"
		if enc := clearQuery.Encode(); enc != "" {
			clearFocusURL = "?" + enc
		}

		var buf bytes.Buffer
		if err := pageTmpl.Execute(&buf, pageData{
			Layout:        layout,
			Nodes:         filtered,
			RefreshMs:     refreshMs,
			RefreshLabel:  fmt.Sprintf("%.1fs", float64(refreshMs)/1000),
			BestHint:      bestHint,
			Focus:         focus,
			ClearFocusURL: clearFocusURL,
		}); err != nil {
			http.Error(w, "template error", 500)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(buf.Bytes())
	}
}
