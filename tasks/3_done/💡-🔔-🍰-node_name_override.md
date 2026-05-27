---
id: task-20260517-node-name-override
title: Add optional node name override
status: 3_done
type: feature
priority: normal
effort: cake
creator: codex
owner: ""
created: 2026-05-17
---

## Problem

pulsed uses the OS hostname as the node identity. That is good for zero-config installs, but cloned hosts, containers, or multiple test instances can collide and overwrite each other in the ledger.

## Done when

- Optional `PULSED_NODE_NAME` overrides `os.Hostname()`
- Empty override preserves current behavior
- Node name is validated enough to avoid empty names and confusing whitespace
- README documents when to use the override
- Tests cover default hostname and override behavior where practical

## Result

- Implemented `PULSED_NODE_NAME` for memberlist identity, mDNS registration, stats heartbeat, and terminal mirror fallback naming
- Added validation and README coverage
- Added focused unit tests for default, override, and whitespace rejection
