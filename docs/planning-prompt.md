# Planning Session Prompt

Paste everything below this line into a new Claude general chat 
session at the start of every new project.

---

I'm starting a new project using my vibe boilerplate. 
Please run a structured planning session with me.

## Boilerplate structure

The repo contains these files that need filling in:

- docs/spec.md — product specification
- docs/architecture.md — stack, data model, technical decisions
- docs/plan.md — task-by-task build plan with work orders
- docs/features/[name].md — one file per feature (created during planning)

AGENTS.md, README.md, CHANGELOG.md and .env.example 
are already complete and do not need filling in.

## My default stack
The default for web apps is:
Next.js (App Router), TypeScript, Tailwind CSS, shadcn/ui, 
Supabase (PostgreSQL + Auth), Zod, React Hook Form, Vercel.

Only suggest deviating from this stack if there is a specific 
reason based on what I'm building. 
If you're building something other than a web app — 
a game, mobile app, CLI tool, API service — tell me and 
I'll recommend the appropriate stack instead.

## How to run this session

Ask me questions one section at a time — do not ask everything 
at once. Work through each document in this order, feel free to suggest answers:

1. docs/spec.md
   - What is the product?
   - Who specifically is it for?
   - What problem does it solve and what do they do today instead?
   - What are the v1 core features (maximum 5)?
   - What does success look like?
   - What is explicitly out of scope for v1?
   - Is there a monetisation model?

2. docs/architecture.md
   - Does this need auth? If so what type?
   - Define every data entity and its fields
   - Any reason to deviate from the default stack?
   - What is explicitly excluded from this version architecturally?

3. docs/plan.md
   - Break v1 into specific build tasks
   - Each task should be completable in one Claude Code session
   - Format each as a work order with: Context, Objective, 
     Files to create or edit, Requirements, Do not do, 
     Future considerations, Acceptance checks

4. docs/features/[name].md
   - One file per core feature
   - Ask me to describe each feature in plain English
   - Format using the feature doc template

5. docs/design.md
**Aesthetic and feeling**
- Describe the emotional tone in 2-3 words (e.g. "calm, focused, premium" 
  or "urgent, dense, analytical")
- What should it feel like to use this app? 
  (e.g. "like a quiet notebook" vs "like a Bloomberg terminal")
- Is this a tool you use quickly and close, or something you linger in?

**Visual references**
- Name 2-3 apps, sites or products whose visual style you want to borrow from
- For each one, say specifically what you're borrowing 
  (e.g. "Linear — the density and muted colours, not the purple")
- Name 1-2 things you've seen that you absolutely don't want 
  (e.g. "the gradient hero of every AI startup landing page")

**Colour direction**
- Light, dark, or system-adaptive?
- One accent colour — what feeling should it convey? 
  (e.g. "amber — warm, not clinical" or "teal — focused, trustworthy")
- Should the UI feel warm or cool overall?

**Typography and density**
- Dense (lots of information visible at once) or spacious (breathing room)?
- Should text do the heavy lifting or should it be more visual/icon-driven?
- Any font preferences or is default fine?

**Component feel**
- Buttons: subtle or prominent? Solid or outlined?
- Cards: bordered, shadowed, or flat?
- Forms and inputs: minimal or clearly defined?
- Any specific UI patterns you love or hate?

**What to explicitly avoid**
- List 3-5 things Claude must never do visually
  (e.g. "no rounded pill buttons, no hero gradients, 
  no white backgrounds, no purple, no stock imagery")

## Output

At the end of the session, output the complete filled-in contents 
of every file, ready to copy directly into the repo.

Files to output:
- docs/spec.md
- docs/architecture.md
- docs/plan.md
- docs/features/[name].md (one per core feature)
- docs/design.md

Label each file clearly:

FILE: docs/spec.md
[contents]

And so on for every file.

Do not output anything that isn't a completed file — 
no summaries, no explanations after the files.
