# Tasks

This folder is a Patchboard task board. Tasks are Markdown files, and the
folder containing a task is its workflow state.

## States

- `backlog/`
- `ready/`
- `doing/`
- `blocked/`
- `done/`
- `archived/`

Move a task file between folders to change its state. Git history is the audit
trail.

## Task Shape

~~~markdown
---
id: task-YYYYMMDD-short-name
title: Short, concrete task title
status: backlog
priority: medium
owner: your-name
created: YYYY-MM-DD
---

## Problem

What needs to change, and why?

## Done when

- The expected behavior is implemented
- Relevant tests or checks pass
~~~

The folder is authoritative for status. If frontmatter includes `status`,
it should match the parent folder.

## Code Annotations

Link code comments back to tasks with square brackets:

~~~text
TODO[task-YYYYMMDD-short-name]: describe the follow-up
FIXME[task-YYYYMMDD-short-name]: describe the known problem
~~~

Unlinked annotations such as `TODO:`, `XXX:`, and `WARN:` are useful inventory,
but they do not fail lint until they reference a task ID.
