---
id: task-20260517-terminal-mirror-controls
title: Polish terminal mirror controls
type: feature
priority: normal
effort: walk
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

The terminal mirror is useful because it gives a second local invocation an immediate read-only view. It currently redraws forever and relies on normal process interruption, which is serviceable but rough for an interactive terminal.

## Done when

- Terminal mirror shows a compact header explaining viewer-only mode
- Ctrl-C exits cleanly without noisy shutdown output
- If interactive key handling is added, it stays minimal, such as `q` to quit and `r` to refresh
- Non-interactive terminals still get readable periodic output
