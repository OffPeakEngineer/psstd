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

## Guidance

- Prefer a small terminal UI library rather than custom raw terminal handling.
- `github.com/gdamore/tcell` is a good candidate: it supports cross-platform input, screen drawing, clean shutdown, and works well for minimally interactive tools.
- Keep the experience simple: use `tcell` only for a compact header, key events, and graceful teardown.
- Fall back to plain periodic output when the environment is non-interactive or terminal features are unavailable.

## Done when

- Terminal mirror shows a compact header explaining viewer-only mode
- Ctrl-C exits cleanly without noisy shutdown output
- If interactive key handling is added, it stays minimal, such as `q` to quit and `r` to refresh
- Non-interactive terminals still get readable periodic output
- Interactive mode uses a terminal library like `tcell` to avoid brittle raw terminal handling
