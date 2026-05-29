# Tasks

This folder is a Patchboard task board. Tasks are Markdown files, and the
folder containing a task is its workflow state. Keep this board lightweight:
the Patchboard tooling owns strict validation, metadata repair, and deductions
from git history.

## States
- -1_anti-feature
- 0_planning
- 1_todo
- 2_doing
- 3_done

Move a task file between folders to change its state. Git history is the audit
trail.

`-1_anti-feature` is for explicit non-goals: ideas that may sound useful but
would make pulsed heavier, riskier, or less focused. Keep these as reference
points so future planning can explain why the project is not taking that path.

## Task Shape
Use plain, tab-completion-friendly filenames:

~~~text
title_in_snake_case.md
~~~

Store sorting and planning metadata in frontmatter:
- `type`: bug, feature, security, maintenance
- `priority`: fire, urgent, normal, idle, low
- `effort`: huge, trip, night, cake, walk

### File Body
~~~markdown
---
id: task-YYYYMMDD-short-name
title: Short, concrete task title
type: feature
priority: normal
effort: walk
creator: your-name
owner: your-name
created: YYYY-MM-DD
---

## Problem

What needs to change, and why?

## Done when

- The expected behavior is implemented
- Relevant tests or checks pass
~~~

The folder is authoritative for status. Move a file between state folders
instead of editing status metadata.

## Code Annotations

Link code comments back to tasks with square brackets:

~~~text
TODO[task-YYYYMMDD-short-name]: describe the follow-up
FIXME[task-YYYYMMDD-short-name]: describe the known problem
~~~

Unlinked annotations such as `TODO:`, `XXX:`, and `WARN:` are useful inventory,
but they do not fail lint until they reference a task ID.
