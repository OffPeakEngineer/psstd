package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/pebble/v2"
)

func terminalRenderLoop(db *pebble.DB) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		renderTerminalSnapshot(db)
		<-ticker.C
	}
}

func renderTerminalSnapshot(db *pebble.DB) {
	nodes, err := dbScanAll(db)
	if err != nil {
		log.Printf("terminal render: %v", err)
		return
	}

	fmt.Print("\033[H\033[2J")
	if len(nodes) == 0 {
		fmt.Println("psstd terminal mirror")
		fmt.Println(styleDim.Render("waiting for cluster state..."))
		return
	}

	fmt.Printf("psstd terminal mirror - %d node(s) - %s\n", len(nodes), time.Now().Format(time.RFC3339))
	fmt.Println(summarizeCluster(nodes).TerminalHeader())
	fmt.Println()
	fmt.Print(renderTerminalNodes(nodes))
}

func renderTerminalNodes(nodes []NodeStats) string {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})

	var sb strings.Builder
	for i, node := range nodes {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(renderANSI(node))
	}
	return sb.String()
}
