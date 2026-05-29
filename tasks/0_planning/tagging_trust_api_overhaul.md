---
id: task-20260517-labels
title: Overhaul node tagging, identity, and trust metadata
type: feature
priority: normal
effort: epic
owner: ""
created: 2026-05-17
revised: 2026-05-29
creator: "copilot"
---

## Problem

Large heterogeneous clusters are harder to reason about when every host is just a
display name plus current metrics. pulsed needs an API-level tagging and trust
model that can describe stable node identity, operator labels, signed claims,
and rings of trust without turning the dashboard into a command/control plane.

This is deeper than `PULSED_ROLE`. Tags become the shared vocabulary for
filtering, history capture, share-policy decisions, and UI focus. Identity and
authenticity need to be explicit enough that pulsed can distinguish a stable
node, a renamed node, an unsigned observation, a direct observation, and data
heard through the mesh.

## Architecture concept

- **Stable identity**: a UUID-backed node ID is distinct from the human-facing
  `server_name`
- **Reserved tags**: well-known tags cover role, environment, trust ring,
  source, observation category, and operator-defined groupings
- **Observation categories**: payloads can distinguish what a node knows, feels,
  sees, hears, and cares to share
- **Signed observations**: HMAC support lets nodes sign gossip/history payloads
  where an operator has configured a shared secret
- **Rings of trust**: trust is modeled as local receiver interpretation of
  source, distance, confidence, and signature status rather than a universal
  truth
- **Care/share policy**: nodes can describe what they are willing to send,
  forward, summarize, suppress, or rebuff
- **API shape**: node and history APIs expose tags, identity, trust, and
  signature status consistently so downstream UI and history logic do not infer
  them from display strings

## Done when

- Optional stable node identity is available as `PULSED_NODE_ID`/UUID while
  `server_name` remains the human-facing display name
- Optional `PULSED_ROLE` maps into a reserved `role` tag rather than a one-off
  display field
- Node tags travel with gossip broadcasts as part of the node state/API payload
- Reserved tag names and user-defined tag names have a documented collision
  strategy
- HMAC signing and verification are designed for gossip/history payloads without
  requiring every deployment to configure shared secrets
- Rings of trust are represented in API data, including unsigned or untrusted
  observations
- Observation category and source path are represented without requiring the UI
  to infer whether data was known, sensed, seen, heard, or re-shared
- Share-policy and rebuff/backoff fields are designed so peers can avoid
  repeatedly sending information a neighbor does not care to receive
- Dashboard and terminal views can render compact role/trust/tag hints without
  owning trust decisions
- Empty tags and unsigned payloads preserve current lightweight behavior

## Related

- Feeds capture, recall, and re-sharing decisions in
  `history_and_navigable_dashboard.md`
- Defines the metadata used by `observation_and_share_policy.md`
