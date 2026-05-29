---
id: task-20260517-node-commands
title: Do not add remote command execution
type: security
priority: normal
effort: cake
owner: ""
created: 2026-05-17
creator: "copilot"
---

## Problem

Remote command execution would turn pulsed from a read-only dashboard into a privileged control surface. That is a large security and product shift, and it does not fit the current lightweight model.

## Done when

- README states that pulsed does not execute remote commands
- Any future command/control idea is explicitly tracked as a separate security design, not a casual dashboard feature
- No command execution endpoint, button, or gossip message is added as part of routine dashboard work
