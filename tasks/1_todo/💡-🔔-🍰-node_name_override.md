---
id: task-20260517-node-name-override
title: Add optional node name override
status: 1_todo
type: feature
priority: normal
effort: cake
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

psstd uses the OS hostname as the node identity. That is good for zero-config installs, but cloned hosts, containers, or multiple test instances can collide and overwrite each other in the ledger.

## Done when

- Optional `PSSTD_NODE_NAME` overrides `os.Hostname()`
- Empty override preserves current behavior
- Node name is validated enough to avoid empty names and confusing whitespace
- README documents when to use the override
- Tests cover default hostname and override behavior where practical
