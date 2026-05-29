---
id: task-20260517-discovery-diagnostics
title: Add lightweight discovery diagnostics
type: feature
priority: normal
effort: night
creator: codex
owner: ""
created: 2026-05-17
---

## Problem 1

When a node does not appear, the likely causes are mDNS, seeds, bind address, firewall, or version mismatch. pulsed should help operators reason about that without becoming a network troubleshooting suite.

## Done when

- Join failures include the peer address and enough context to act
- Startup logs distinguish explicit seeds from mDNS-discovered peers

## Problem 2

Furthermore, to facilitate visualization and depth filtering, any node should
also tell what nodes it can directly see or is actively talking to. This is the
topology side of the "what I see" observation category, not just a dashboard
decoration.

## Done when

- Dashboard or terminal summary can indicate when the node is running solo
- Direct neighbor visibility can feed observation/share-policy work without
  requiring nodes to trust or re-share every second-hand report

## Related

- Feeds `observation_and_share_policy.md`
