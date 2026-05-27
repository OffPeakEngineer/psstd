---
id: task-20260517-recent-cluster-events
title: Show recent cluster events locally
status: 0_planning
type: feature
priority: low
effort: night
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

Operators sometimes need lightweight context: a node joined, left, went stale, or recovered. Logs contain this, but the dashboard could show a very short local event trail without turning pulsed into an audit log.

## Done when

- Each process keeps a bounded in-memory ring of recent local cluster events
- Web and terminal views can show the last few events compactly
- Events are not persisted and are not gossiped as a new data product
- The feature can be disabled or omitted from cramped layouts
