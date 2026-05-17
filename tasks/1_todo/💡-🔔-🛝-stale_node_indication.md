---
id: task-20260517-health
title: Distinguish stale nodes from offline nodes
status: 1_todo
type: feature
priority: normal
effort: walk
owner: ""
created: 2026-05-17
creator: "copilot"
---

## Problem

Offline nodes and recently stale nodes can look too similar. Operators should quickly see whether a node just went quiet or has been offline long enough to treat as gone.

## Done when

- Fresh, stale, and offline states have distinct rendering
- The staleness threshold is controlled by the origin node (each node controls the 'time' it'd like to be known for)
- Terminal and web views use the same state calculation
- Existing offline purge/version behavior still works
