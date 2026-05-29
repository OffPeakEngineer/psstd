---
id: task-20260517-duplicate-instance-tests
title: Test duplicate-instance terminal fallback
type: maintenance
priority: normal
effort: night
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

The duplicate-instance path is important because it protects the Pebble store and avoids publishing duplicate local nodes. It was added pragmatically, but it deserves tests around lock detection, address-in-use detection, and mirror seed selection.

## Done when

- Lock-detection helper tests cover the expected Pebble lock error shapes
- Address-in-use helper tests cover Unix and Windows-style bind errors
- Mirror seed construction includes configured seeds plus local gossip candidates
- Tests avoid opening real long-running memberlist processes
