# Plan

<!--
This is your build bible.
One task at a time. Each task is a work order for Claude.
Backlog lives at the bottom — promote to Active when ready.

Reference tasks explicitly in Claude Code:
"Read AGENTS.md and docs/plan.md Task 1. Implement Task 1 only."
-->

## Active task

<!--
Only ever one active task.
Move to Completed when done, promote next from Backlog.
-->

<!--
Choose one format below. Delete the other before filling in.
Simple: for single-file changes, UI tweaks, obvious scope.
Full: for anything touching multiple systems, auth, data model, AI features, or payments.

### Task [N]: [Name] (simple)
**What:** One sentence — what exists after this that didn't before
**Files:** List files to touch
**Done when:** One or two acceptance checks
**Do not:** One constraint if needed
**Depends on:** Task number and name, or leave blank

---

### Task [N]: [Name]

#### Context
Why does this task exist?
What user problem does it solve?
What came before it that this depends on?

#### Objective
One sentence. What should exist after this task that doesn't exist now?

#### Files to create or edit
Be specific. Claude will only touch these files.
- src/app/[route]/page.tsx — create
- src/components/[Component]/index.tsx — create
- src/lib/supabase.ts — edit

#### Requirements
Specific, testable statements.
- User can do X
- Clicking Y does Z
- Data is saved to [table]

#### Do not do
Explicit constraints. Prevents Claude gold-plating.
- Do not add animations yet
- Do not implement [future feature]
- Do not modify [unrelated file]

#### Future considerations
Optional. Flag things deliberately deferred.
- Auth will be added in Task [N]
- Database persistence planned for Task [N]
- Mobile layout addressed in polish phase

#### Depends on (if applicable)
List any tasks that must be complete before this one starts.
Leave blank if none.

#### Acceptance checks
How do you verify this is done correctly?
- [ ] Check 1
- [ ] Check 2
- [ ] Check 3
-->

---

## Backlog

<!--
Future tasks in rough priority order.
Promote to Active when current task is complete.

### Task [N]: [Name]
One line description.
-->

---

## Completed

<!--
Move finished tasks here.
Gives a record of progress and context for future tasks.

### Task [N]: [Name] ✓
Brief note on what was built and any decisions made.
-->
