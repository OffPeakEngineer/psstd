---
id: task-20260517-cluster-summary-band
title: Add a compact cluster summary band
status: 1_todo
type: feature
priority: normal
effort: walk
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

The dashboard is strongest as an at-a-glance cluster htop. A small summary can answer the first operator questions before scanning individual nodes: how many nodes are online, which node is hottest, and whether anything is stale or offline.

## Done when

- Web toolbar shows online, stale, and offline counts
- Summary identifies the hottest online node by CPU or load
- Terminal mirror includes the same counts in its header
- The summary uses existing node snapshots only and adds no API surface
