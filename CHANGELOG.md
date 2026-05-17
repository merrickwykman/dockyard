# Changelog

<!--
Update after meaningful completed tasks — not every minor commit.
Record what changed, why it changed, and any decisions made.
Prevents relitigating decisions in future sessions.
-->

## [Unreleased]
<!--
Staging area. Move to a dated entry on each push to production.
-->

## Format

<!--
[DATE] — [feature or change]
- What changed
- Why it changed
- Decisions made and reasoning
- Anything tried that didn't work

Example:
[2026-05-07] — habit completion logic
- Created entries table in Supabase
- Derived streak count on read rather than storing it — storing 
  creates sync issues if entries are deleted
- Tried optimistic UI update — reverted due to flicker on 
  slow connections
-->
