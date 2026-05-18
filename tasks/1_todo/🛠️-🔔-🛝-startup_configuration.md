---
id: task-20260517-cli
title: Print a clear startup configuration
status: 1_todo
type: maintenance
priority: normal
effort: walk
creator: codex
owner: ""
created: 2026-05-17
---

## Problem 1

psstd is easiest to like when it feels obvious what it is doing: which database it owns, which HTTP URL it advertises, which gossip address it listens on, and whether it joined peers. Today that information exists in logs, but it is not shaped as a concise operator-facing summary.

## Done when

- Startup logs show DB path, HTTP listen address, advertised URL, gossip listen address, web enabled state, and version
- Seed and mDNS discovery results are summarized without noisy repetition

## Problem 2
Just a general cleanup of cli flags and provide some of the basics like help text.

## Done when
- Terminal mirror mode should be default mode (-t)
- If a -v flag is provided, then it prints the log instead
- If a -tv flag is provided, then we print both the terminal rendering and a log stream (similar to how screen does window splits)
- A -r flag does "read only" and '-t' is still default, and -v and -tv still do the same, just as a read-only 'observer' of the cluster
- Existing log lines remain useful for troubleshooting
