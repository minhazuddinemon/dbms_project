---
name: Academic Path Design System
colors:
  surface: '#f7f9fb'
  surface-dim: '#d8dadc'
  surface-bright: '#f7f9fb'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f2f4f6'
  surface-container: '#eceef0'
  surface-container-high: '#e6e8ea'
  surface-container-highest: '#e0e3e5'
  on-surface: '#191c1e'
  on-surface-variant: '#464554'
  inverse-surface: '#2d3133'
  inverse-on-surface: '#eff1f3'
  outline: '#767586'
  outline-variant: '#c7c4d7'
  surface-tint: '#494bd6'
  primary: '#4648d4'
  on-primary: '#ffffff'
  primary-container: '#6063ee'
  on-primary-container: '#fffbff'
  inverse-primary: '#c0c1ff'
  secondary: '#4e45d5'
  on-secondary: '#ffffff'
  secondary-container: '#6860ef'
  on-secondary-container: '#fffbff'
  tertiary: '#006c49'
  on-tertiary: '#ffffff'
  tertiary-container: '#00885d'
  on-tertiary-container: '#000703'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#e1e0ff'
  primary-fixed-dim: '#c0c1ff'
  on-primary-fixed: '#07006c'
  on-primary-fixed-variant: '#2f2ebe'
  secondary-fixed: '#e3dfff'
  secondary-fixed-dim: '#c3c0ff'
  on-secondary-fixed: '#100069'
  on-secondary-fixed-variant: '#372abf'
  tertiary-fixed: '#6ffbbe'
  tertiary-fixed-dim: '#4edea3'
  on-tertiary-fixed: '#002113'
  on-tertiary-fixed-variant: '#005236'
  background: '#f7f9fb'
  on-background: '#191c1e'
  surface-variant: '#e0e3e5'
typography:
  display-lg:
    fontFamily: Inter
    fontSize: 48px
    fontWeight: '700'
    lineHeight: 60px
    letterSpacing: -0.02em
  headline-lg:
    fontFamily: Inter
    fontSize: 32px
    fontWeight: '700'
    lineHeight: 40px
    letterSpacing: -0.01em
  headline-md:
    fontFamily: Inter
    fontSize: 24px
    fontWeight: '600'
    lineHeight: 32px
  body-lg:
    fontFamily: Inter
    fontSize: 18px
    fontWeight: '400'
    lineHeight: 28px
  body-md:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '400'
    lineHeight: 24px
  label-lg:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '600'
    lineHeight: 20px
    letterSpacing: 0.02em
  label-sm:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '500'
    lineHeight: 16px
    letterSpacing: 0.04em
  headline-lg-mobile:
    fontFamily: Inter
    fontSize: 28px
    fontWeight: '700'
    lineHeight: 36px
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  margin-page: 2rem
  gutter-grid: 1.5rem
  padding-card: 1.5rem
  stack-sm: 0.5rem
  stack-md: 1rem
  stack-lg: 2rem
---

## Brand & Style

This design system is engineered for a university admission and route matching platform, focusing on high-level professionalism mixed with an approachable, student-centric energy. The aesthetic is **Corporate Modern with Glassmorphic accents**, characterized by deep indigos, expansive white space, and refined translucency. 

The goal is to evoke a sense of **clarity, momentum, and trust**. By utilizing a sophisticated color palette and soft-edged geometry, the UI reduces the inherent stress of university applications, transforming a complex bureaucratic process into a streamlined, high-tech journey. The visual language balances institutional authority with a forward-thinking, digital-first experience.

## Colors

The palette is anchored by a **Deep Indigo** primary, representing the transition from student life to professional higher education. 

- **Primary & Secondary:** Used for high-emphasis actions, navigation states, and brand-building headers.
- **Surface & Background:** A tiered system of cool grays (Slate) and pure whites to maintain high legibility and a "breathable" interface.
- **Semantic Colors:** Success (Emerald) and Warning (Amber) are softened to avoid visual fatigue, used primarily for application statuses and deadline alerts.
- **Gradients:** Subtle linear gradients (Primary to Secondary) are reserved for hero sections and high-value conversion points to guide the eye toward "Apply" actions.

## Typography

This design system utilizes **Inter** exclusively to ensure a clean, geometric, and highly readable interface across all data densities.

- **Scale:** A tight 1.25x typographic scale ensures contrast between informational body text and functional headings.
- **Hierarchy:** Generous tracking (letter spacing) is applied to small labels to improve scanability in dense dashboard views.
- **Emphasis:** Bold weights are used sparingly for user-specific data (e.g., University Names, Application Counts) to create a clear visual hierarchy.
- **Mobile:** Headlines are reduced by 15-20% on smaller viewports to ensure high-impact text does not break across multiple lines or crowd the interface.

## Layout & Spacing

The layout utilizes a **12-column fluid grid** optimized for a 1920x1080 desktop experience. 

- **Grid Model:** The sidebar remains fixed at 280px, while the main content area expands. Cards typically span 3, 4, or 6 columns depending on content priority.
- **Rhythm:** An 8px base grid drives all spacing decisions.
- **Breakpoints:**
  - **Desktop (1440px+):** Full 12-column visibility with 24px gutters.
  - **Tablet (768px - 1439px):** Content reflows to a 2-column card stack; sidebar collapses into a hamburger menu.
  - **Mobile (<767px):** Single column layout with 16px horizontal margins.
- **Safe Areas:** Maintain a minimum 32px padding around the viewport edges on large displays to prevent visual clutter.

## Elevation & Depth

Visual depth is achieved through **Glassmorphism and Ambient Shadows**, creating a layered "HUD" (Heads-Up Display) effect.

- **Surface Layers:** The background uses a flat slate gray. Cards and containers sit above this using a semi-transparent white (80-90% opacity) with a `20px` backdrop blur.
- **Shadows:** Use extra-diffused shadows with a primary-tinted color (`rgba(99, 102, 241, 0.08)`) to make elements feel light and integrated rather than heavy or "floating."
- **Borders:** Subtle, 1px solid borders in a very light neutral-200 are used to define edges where glass opacity is high, ensuring components don't bleed into each other.

## Shapes

The design system uses a **Rounded** shape language to reinforce the student-centric, friendly brand identity.

- **Core Components:** Standard cards, input fields, and main containers utilize a `1.25rem` (20px) radius as seen in the reference dashboard.
- **Small Elements:** Buttons and tags use a `0.5rem` radius for a more precise, functional appearance.
- **Interactive States:** On hover, card elevation increases slightly, but border-radius remains constant to maintain grid integrity.

## Components

- **Buttons:** Primary buttons use a solid indigo fill with white text. Secondary buttons use a light-indigo ghost style with a 1px border. Hover states should include a soft glow shadow (`primary_color` at 20% opacity).
- **Cards:** Glassmorphic by default. Cards should include a 24px padding and utilize a clear vertical stack: Icon/Category (Top), Primary Data (Middle), and "View All" or "Action" (Bottom).
- **Input Fields:** Large, clean fields with `16px` internal padding and a soft `1px` border that glows primary indigo upon focus.
- **Chips/Badges:** Small, high-contrast badges for "Match %" or "Deadline" status, using a light tint of the semantic color as a background and a darker tint for text.
- **Navigation:** A clean left-hand sidebar with stroke-based icons. The active state is indicated by a background highlight and a primary-colored vertical indicator line.
- **Progress Indicators:** Horizontal timelines for application stages, using rounded nodes and thin connector lines to show flow.