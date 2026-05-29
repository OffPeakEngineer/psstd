# Terminology

This file keeps the concept vocabulary grounded enough to implement without
turning it into a named product taxonomy too early.

## Node

A node is an addressable observer. Today that is usually a running `pulsed`
process on a host. In the broader model, a node can also represent a pod,
container, service, device, bus, orchestrator, availability zone, region, or
other scoped body.

Nodes may have child nodes, but the relationship is hierarchical. A node path is
a rooted parent/child address, not an arbitrary graph.

## Body Node

A body node is an external-facing boundary for a scope. Examples include a pod,
host, orchestrator, zone, or region. Body nodes decide which internal details
are exposed, summarized, hidden, or forwarded.

## Internal Node

An internal node is a child inside a local scope. Internal nodes may share more
detail through a local loopback mesh because they are inside the same boundary.

## Observation

An observation is something a node knows, senses, sees, hears, or chooses to
share. Observations should carry source identity, time, category, path, and
freshness metadata.

## Enumeration

Enumeration answers "what exists?" before a caller asks for a payload. The core
index dimensions are:

- when: timeslice or freshness window
- who: source node identity
- where: node path or scoped body
- what: structure, parameter, metric, tag, or resource kind
- how: current, historical, stale, direct, heard, re-shared, signed, unsigned

## Queryable Resource

A queryable resource is a tagged thing a node can advertise and another node can
request. Themes are the smallest useful example: a node can enumerate available
themes and share one when policy allows it.

## Trust

Trust is local receiver interpretation, not global truth. Signature status,
source distance, scope, freshness, and operator policy should remain separate in
data even when the UI summarizes them.
