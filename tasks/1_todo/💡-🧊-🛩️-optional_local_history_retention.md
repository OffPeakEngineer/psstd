---
id: task-20260517-persist-metrics
title: Consider optional local history retention
status: 0_planning
type: feature
priority: low
effort: trip
owner: ""
created: 2026-05-17
creator: "copilot"
---

## Problem

The current model keeps recent gossip state only. Operators sometimes want to know whether a node was hot earlier, but persistent history can grow into a metrics product if the scope is not constrained.

## Done when

- A command line flag --csv="out.csv" file is created, with a header, and simply streams the data it recieves/generates into it.