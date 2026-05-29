---
id: task-20260517-version-mismatch-explanation
title: Explain version mismatch behavior
type: maintenance
priority: normal
effort: walk
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

The code already treats stale records from different versions carefully, but the operator-facing behavior is not obvious. During upgrades, users should understand why old offline nodes disappear or why mismatched live nodes may remain visible until they age out.

## Done when

- README describes version compatibility behavior during rolling upgrades
- Startup logs include the local app version
- Dashboard or terminal display makes mismatched versions visible without alarm
- Existing purge behavior has focused tests if coverage is missing
