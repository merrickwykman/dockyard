# Design Direction

<!--
Fill this in during the planning session alongside spec.md.
Reference it explicitly in the design pass prompt:
"Read AGENTS.md and docs/design.md. Apply the design direction across all pages."
-->

## Aesthetic
<!--
Describe the feeling in plain English.
Example: "Dark, atmospheric, focused — like Obsidian or Linear. 
Not clinical, not corporate, not generic SaaS."
-->

## Principles
<!--
3-5 rules Claude must follow throughout.
Example:
- Dark backgrounds, no harsh whites
- Typography does the heavy lifting over colour
- Information dense but not cluttered
- Subtle over flashy — no gradients, minimal animation
- Consistent spacing — generous padding, clear hierarchy
-->

## Reference points
<!--
Apps, sites, or tools whose aesthetic you want to borrow from.
Example: "Linear for layout density, Obsidian for dark atmosphere, 
Raycast for clean typography."
These give Claude a shared visual vocabulary to work from.
-->

## Colour palette
<!--
Define before building — Claude applies these globally.
Example:
- Background: #0a0a0a
- Surface: #141414
- Border: #2a2a2a
- Text primary: #e8e8e8
- Text secondary: #666666
- Accent: #f59e0b (amber)
- Success: #22c55e
- Destructive: #ef4444
-->

## Typography
<!--
Font choices and scale.
Example:
- Font: Geist (Next.js default) or Inter
- Headings: semibold, tight tracking
- Body: regular, comfortable line height
- Labels: small, muted, uppercase tracking for categories
-->

## Component style
<!--
How UI elements should feel.
Example:
- Buttons: solid, no rounded pills — use rounded-md not rounded-full
- Cards: subtle border, no drop shadows
- Inputs: dark background, border focus ring in accent colour
- No gradients
- No stock imagery
-->

## Layout rules
<!--
Define page structure and spacing.
Example:
- Max content width: 1200px
- Main pages use a centered container with px-6 py-8
- Dashboards use a sidebar + content layout
- Mobile-first, but desktop should feel intentional
- Avoid full-width content unless the feature requires it
-->

## shadcn/ui theme overrides
<!--
Specific CSS variable overrides for the global theme.
Claude applies these to globals.css at the start of the design pass.
Example:
--background: 0 0% 4%;
--foreground: 0 0% 91%;
--card: 0 0% 8%;
--border: 0 0% 16%;
--primary: 38 92% 50%;
-->

## What to avoid
<!--
Explicit exclusions prevent Claude defaulting to generic choices.
Example:
- No white or light grey backgrounds
- No purple — the generic SaaS colour
- No pill-shaped buttons
- No card drop shadows
- No hero gradients
- No default shadcn light theme aesthetics
-->

## Design pass prompt
<!--
Copy this into Claude Code when ready for the visual pass.
Paste after all features are complete and tested.
-->

\`\`\`
Read AGENTS.md and docs/design.md.
Apply the design direction consistently across all pages.
Use the colour palette and CSS variable overrides defined in design.md.
Replace all default shadcn styling with the custom theme.
Ensure typography scale and spacing are consistent across all components.
Do not change any functionality — visual changes only.
After completing, run npx tsc --noEmit and manually verify 
every page before committing.
After completing, list every file changed and what was updated.
\`\`\`