---
id: task-20260518-cli-mode-flags
title: Design explicit CLI mode flags
type: maintenance
priority: normal
effort: night
creator: codex
owner: ""
created: 2026-05-18
---

## Problem

The requested `-t`, `-v`, `-tv`, and `-r` modes change pulsed startup semantics:
today the primary process owns the local database, participates in gossip, serves
HTTP by default, and only falls back to terminal mirror mode when another local
instance already owns the store or gossip port.

Making terminal mirror mode the default needs a clearer mode model so service
deployments, web serving, read-only cluster observation, and log streaming remain
predictable.

## Done when

- CLI modes are explicitly named and documented
- Default behavior for interactive shells and services is decided
- `-t`, `-v`, `-tv`, and `-r` have non-overlapping behavior
- Read-only observer mode cannot accidentally publish a duplicate node
- Tests cover flag parsing and mode selection
