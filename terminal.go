package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/cockroachdb/pebble/v2"
)

const (
	terminalCellWidth = 42
	terminalCellGap   = 2
)

var terminalCellStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("8")).
	Padding(0, 1).
	Width(terminalCellWidth)

func terminalRenderLoop(db *pebble.DB, listMode bool) {
	if terminalInteractive(os.Stdin, os.Stdout) {
		err := terminalInteractiveRenderLoop(db, listMode, os.Stdin)
		if err == nil {
			return
		}
		log.Printf("terminal interactive unavailable: %v", err)
	}
	terminalPlainRenderLoop(db, listMode)
}

func terminalPlainRenderLoop(db *pebble.DB, listMode bool) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		renderTerminalSnapshot(db, listMode, false)
		<-ticker.C
	}
}

func terminalInteractive(stdin, stdout *os.File) bool {
	return term.IsTerminal(stdin.Fd()) && term.IsTerminal(stdout.Fd())
}

func terminalInteractiveRenderLoop(db *pebble.DB, listMode bool, stdin *os.File) error {
	oldState, err := term.MakeRaw(stdin.Fd())
	if err != nil {
		return err
	}
	defer term.Restore(stdin.Fd(), oldState)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signals)

	keys := make(chan byte, 4)
	go readTerminalKeys(stdin, keys)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	renderTerminalSnapshot(db, listMode, true)
	for {
		select {
		case <-signals:
			fmt.Print("\r\n")
			return nil
		case key := <-keys:
			switch terminalKeyAction(key) {
			case terminalActionQuit:
				fmt.Print("\r\n")
				return nil
			case terminalActionRefresh:
				renderTerminalSnapshot(db, listMode, true)
			}
		case <-ticker.C:
			renderTerminalSnapshot(db, listMode, true)
		}
	}
}

func readTerminalKeys(stdin *os.File, keys chan<- byte) {
	var buf [1]byte
	for {
		n, err := stdin.Read(buf[:])
		if err != nil {
			return
		}
		if n == 1 {
			keys <- buf[0]
		}
	}
}

type terminalAction int

const (
	terminalActionNone terminalAction = iota
	terminalActionQuit
	terminalActionRefresh
)

func terminalKeyAction(key byte) terminalAction {
	switch key {
	case 'q', 'Q', 3:
		return terminalActionQuit
	case 'r', 'R':
		return terminalActionRefresh
	default:
		return terminalActionNone
	}
}

func renderTerminalSnapshot(db *pebble.DB, listMode bool, interactive bool) {
	nodes, err := dbScanAll(db)
	if err != nil {
		log.Printf("terminal render: %v", err)
		return
	}

	fmt.Print("\033[H\033[2J")
	if len(nodes) == 0 {
		fmt.Println(terminalHeader(0, time.Now(), interactive))
		fmt.Println(styleDim.Render("waiting for cluster state..."))
		return
	}

	fmt.Println(terminalHeader(len(nodes), time.Now(), interactive))
	fmt.Println(summarizeCluster(nodes).TerminalHeader())
	fmt.Println()
	if listMode {
		fmt.Print(renderTerminalNodes(nodes))
		return
	}
	fmt.Print(renderTerminalGrid(nodes, terminalWidth()))
}

func terminalHeader(nodeCount int, now time.Time, interactive bool) string {
	controls := "refreshes every 2s"
	if interactive {
		controls = "q quit - r refresh - Ctrl-C quit"
	}
	return fmt.Sprintf("pulsed terminal mirror (viewer-only) - %d node(s) - %s - %s",
		nodeCount, now.Format(time.RFC3339), controls)
}

func renderTerminalNodes(nodes []NodeStats) string {
	sortNodesByName(nodes)

	var sb strings.Builder
	for i, node := range nodes {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(renderANSI(node))
	}
	return sb.String()
}

func renderTerminalGrid(nodes []NodeStats, width int) string {
	sortNodesByName(nodes)
	if len(nodes) == 0 {
		return ""
	}

	cols := terminalColumns(width)
	rows := make([]string, 0, (len(nodes)+cols-1)/cols)
	for start := 0; start < len(nodes); start += cols {
		end := start + cols
		if end > len(nodes) {
			end = len(nodes)
		}
		cells := make([]string, 0, end-start)
		for _, node := range nodes[start:end] {
			cells = append(cells, renderTerminalCell(node))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cells...))
	}
	return strings.Join(rows, "\n") + "\n"
}

func renderTerminalCell(node NodeStats) string {
	return terminalCellStyle.Render(strings.TrimRight(renderANSI(node), "\n"))
}

func terminalColumns(width int) int {
	cellAndGap := terminalCellWidth + terminalCellGap
	if width < cellAndGap {
		return 1
	}
	cols := (width + terminalCellGap) / cellAndGap
	if cols < 1 {
		return 1
	}
	return cols
}

func terminalWidth() int {
	if value := os.Getenv("COLUMNS"); value != "" {
		if width, err := strconv.Atoi(value); err == nil && width > 0 {
			return width
		}
	}
	return 100
}

func sortNodesByName(nodes []NodeStats) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})
}
