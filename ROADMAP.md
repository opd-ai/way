# ROADMAP: Way

## Project Overview

Way is a 100% procedurally generated battle-kart racing game built in Go 1.24+ with Ebiten v2, forming the W-Series successor to V-Series projects (`opd-ai/venture`, `opd-ai/vania`, `opd-ai/violence`, `opd-ai/velocity`). Inspired by Diddy Kong Racing, Mario Kart, Rocket League, and SuperTuxKart, Way exceeds V-Series depth with procedural 3D-ish tracks featuring elevation, loops, shortcuts, and weather; emergent vehicle physics with drifting and boost; competitive multiplayer with item targeting and rubber-banding AI; and dynamic procedural audio including engine sounds, Doppler effects, and crowd noise. All content — tracks, karts, items, audio, visuals, and narrative — is generated at runtime from a deterministic seed with zero embedded assets, fully themed across five genres: **fantasy, sci-fi, horror, cyberpunk, post-apocalyptic**.

## Core Architecture

**ECS Design:** Entities are integer IDs. Components (`TransformComponent`, `VehiclePhysicsComponent`, `ItemHolderComponent`, `RacePositionComponent`) are pure structs stored in typed slices. Systems (`PhysicsSystem`, `ItemSystem`, `RaceSystem`, `RenderSystem`, `NetworkSystem`) iterate components and contain all logic — no methods on component types.

**Key Systems:**

| System | Responsibility |
|--------|---------------|
| `PhysicsSystem` | Vehicle movement, drifting, boost, collision with track geometry |
| `RaceSystem` | Lap counting, position ranking, rubber-banding AI, finish detection |
| `ItemSystem` | Item box pickup, targeting logic, hit validation, cooldowns |
| `TrackSystem` | Waypoint graph traversal, shortcut activation, hazard triggers |
| `RenderSystem` | Runtime sprite/tile/particle/post-processing generation pipeline |
| `AudioSystem` | Procedural engine sounds, Doppler, adaptive music, SFX |
| `NetworkSystem` | Authoritative server, client prediction, reconciliation, interpolation |
| `PCGSystem` | Drives all generators via `Generator.Generate(seed, params)` interface |

**V-Series Reuse:** `pkg/procgen` PCG interface and `GenerationParams` (with `GenreID`), `pkg/audio` oscillators/envelopes/SFX genre modifications, `pkg/rendering` runtime sprite pipeline, `pkg/network` TCP authoritative model with delta compression and lag compensation, `pkg/engine` ECS base and world tick loop.

## Implementation Phases

### Phase 1 — Foundation (Weeks 1–4)
**Focus:** Project scaffold, ECS core, deterministic PCG seed system, single-binary build.

1. **Project scaffold** — `cmd/client/`, `cmd/server/`, `pkg/engine/`, `pkg/procgen/`, `pkg/rendering/`, `pkg/audio/`, `pkg/network/` mirroring V-Series layout. Single `go.mod`. *AC: `go build ./...` succeeds with zero assets on disk.*
2. **ECS world** — `World` type with entity creation/deletion, typed component stores, system registration, and fixed-timestep tick loop at 60 Hz. *AC: 10 000 entities tick in < 2 ms on reference hardware; `go test ./pkg/engine/...` passes.*
3. **PCG seed pipeline** — `Generator` interface (`Generate(seed int64, params GenerationParams) (interface{}, error)` + `Validate()`), `GenreID` constants for all five genres, deterministic PRNG wrapper. *AC: Same seed + genre produces byte-identical output across 100 consecutive calls.*
4. **Single-binary build** — All generated content embedded via `go:generate` hooks or computed at init; no external file reads at runtime. *AC: Running `way` from an otherwise empty working directory under a syscall tracer (`strace`/`dtruss`/platform equivalent) shows zero asset file `open`/`openat` calls; optional OS-specific dependency checks (`ldd way`, `otool -L way`, `dumpbin /DEPENDENTS way.exe`) show only code libraries, and the binary runs from any working directory.*

### Phase 2 — Track & Vehicle PCG (Weeks 5–9)
**Focus:** Procedural track generation, vehicle physics, kart PCG.

1. **Track generator** — Spline-based circuit with control-point jitter, elevation changes (±30 m), loop/ramp segments, shortcut branches, and hazard zones. Output: waypoint graph + mesh descriptor. *AC: All five genres produce visually distinct tracks; no self-intersecting geometry at any seed.*
2. **Vehicle physics** — Arcade kart model: acceleration, braking, lateral grip, drift (oversteer on button hold), boost bar (fills on drift). Runs server-side authoritatively. *AC: Drift boost activates in < 3 s of sustained drift; `go test ./pkg/engine/physics/...` covers edge cases.*
3. **Kart PCG** — Procedural kart body shapes, colour palettes, and names keyed to genre. Stats (speed, handling, weight) vary per seed. *AC: Five distinct kart silhouettes per genre; stat ranges documented and enforced.*
4. **Runtime renderer** — Pixel-art sprite generation for track tiles, kart sprites, particles (dust, boost flame, item explosion). No PNG/BMP files. *AC: `grep -R "go:embed" .` returns zero matches in `pkg/rendering/` (any asset embedding is confined to explicitly documented packages such as `pkg/font/...`).*

### Phase 3 — Items, Hazards & Audio (Weeks 10–14)
**Focus:** Battle items, track hazards, procedural audio.

1. **Item generator** — 8 item archetypes (projectile, mine, shield, speed, EMP, swap, pull, area-bomb) with genre-reskinned names and visuals. Item boxes placed on track by PCG. *AC: Each genre maps all 8 archetypes to distinct flavour; `Validate()` rejects invalid param combinations.*
2. **Hazard generator** — Per-genre dynamic hazards placed in hazard zones: weather overlays (snow/rain/fog), moving obstacles, timed barriers. *AC: Hazard density scales with `Difficulty`; hazards never block the only valid path.*
3. **Procedural audio** — Engine pitch tracks RPM (200–8000 Hz range), Doppler shift on passing karts (±15%), crowd noise scales with race lead changes, adaptive music layers motif intensity on position rank. Genre modifiers applied per V-Series patterns. *AC: No `.wav`/`.ogg` files; `AudioSystem` benchmark ≤ 1 ms per frame.*
4. **SFX generator** — Collision, item use, boost, drift squeal, finish fanfare — all synthesised. *AC: Each SFX is reproducible from its seed; genre timbre differences audible in blind A/B.*

### Phase 4 — Multiplayer & Netcode (Weeks 15–20)
**Focus:** Authoritative server, client prediction, high-latency support.

1. **Authoritative race server** — TCP server ticks race state at 20 Hz; clients send inputs, receive delta-compressed snapshots. *AC: 8-player race desync rate < 1 in 10 000 ticks at 0 ms simulated latency.*
2. **Client-side prediction** — Clients simulate physics locally; server reconciles on snapshot receipt; mispredicted inputs replayed. *AC: No visible snap/stutter at ≤ 100 ms RTT in integration test.*
3. **High-latency tolerance** — Entity interpolation buffer (3–5 snapshots), lag compensation for item-hit validation (100 ms history), adaptive input queue length from 200 ms to 5000 ms. *AC: Game is playable (inputs processed, race finishes correctly) at simulated 2000 ms RTT.*
4. **Item hit validation** — Server validates hits using lag-compensated positions; clients show speculative effects immediately; server confirms/rolls back. *AC: False-positive hit rate < 0.1% at 500 ms RTT in simulation.*
5. **Race state sync** — Lap counts, item states, race timer, and finish order are server-authoritative. *AC: Finish order matches server log in 100% of simulated races.*

### Phase 5 — AI, Ranked Play & Polish (Weeks 21–26)
**Focus:** AI drivers, ranked matchmaking, genre completeness, release readiness.

1. **AI driver system** — Waypoint-following with rubber-banding (AI speed scales to maintain ±2 positions of last human), item usage heuristics keyed to difficulty, genre-flavoured names. *AC: AI completes any generated track; rubber-band engagement measurable via telemetry.*
2. **Ranked matchmaking** — ELO-based rating, race result submission to server, season leaderboard. Federation stub for cross-server events. *AC: Rating converges correctly on 1000-race simulation; API spec documented.*
3. **Genre completeness audit** — Each genre reviewed for track aesthetics, kart naming, item flavour, hazard types, audio timbre, visual palette, and post-processing effect. All five must score ≥ 4/5 on internal differentiation rubric. *AC: Checklist signed off by two reviewers; zero shared asset paths across genres.*
4. **Single-binary release build** — `make release` produces a stripped, statically linked binary for Linux/Windows/macOS. *AC: Binary size ≤ 25 MB; cold-start to main-menu ≤ 3 s on reference hardware.*

## PCG Systems Inventory

| Generator | Algorithm / Approach | Genre Impact |
|-----------|---------------------|--------------|
| Track layout | Spline + control-point jitter, elevation noise (Perlin), shortcut branch probability | Fantasy: rolling hills, stone bridges; Sci-fi: neon tunnels, zero-G loops; Horror: fog swamps, narrow passes; Cyberpunk: neon city canyons, rooftop shortcuts; Post-apoc: rubble fields, collapsed overpasses |
| Track surface | Tile palette noise-mapped to segment type | Fantasy: cobblestone/grass; Sci-fi: metal grid/plasma; Horror: mud/bone; Cyberpunk: asphalt/hologram; Post-apoc: cracked concrete/sand |
| Kart body | Parametric polygon silhouette, colour palette, wheel shape | Fantasy: wood-cart; Sci-fi: pod-racer; Horror: hearse; Cyberpunk: chrome cruiser; Post-apoc: junk buggy |
| Item archetypes | Weighted table by genre; projectile physics parameterised | Fantasy: magic bolt/potion; Sci-fi: plasma shot/EMP; Horror: curse/ghost; Cyberpunk: virus/hack; Post-apoc: molotov/chain |
| Hazard placement | Density map on waypoint graph; hazard type pool by genre | Fantasy: enchanted vines; Sci-fi: gravity wells; Horror: creeping fog/jumpscare; Cyberpunk: police drones; Post-apoc: sandstorms/debris |
| Engine audio | Oscillator bank (sawtooth + pulse) modulated by RPM seed | Fantasy: +5% pitch, lute timbre; Sci-fi: +30% pitch, sine purity; Horror: −30% pitch + vibrato; Cyberpunk: +40% pitch + hard clip; Post-apoc: −10% pitch + noise layer |
| Race music | Motif generator: base BPM, chord progression, layered by rank | Fantasy: orchestral; Sci-fi: synth-wave; Horror: dissonant strings; Cyberpunk: drum-machine EDM; Post-apoc: industrial percussion |
| Visual palette | HSL rotation + saturation scale per genre seed | Fantasy: warm earth tones; Sci-fi: cool blues/whites; Horror: desaturated greens; Cyberpunk: neon cyan/magenta; Post-apoc: sepia/rust |
| Post-processing | Shader parameters: bloom, chromatic aberration, grain | Fantasy: soft glow; Sci-fi: scanlines; Horror: vignette + grain; Cyberpunk: heavy bloom + aberration; Post-apoc: dust overlay |
| Kart/driver names | Markov chain trained on genre vocabulary seed | Fantasy: elvish/knightly; Sci-fi: alphanumeric callsigns; Horror: gothic surnames; Cyberpunk: handle@tag; Post-apoc: scavenger epithets |
| Weather overlay | Particle system parameterised by precipitation type | Fantasy: cherry blossom; Sci-fi: ion storm; Horror: blood rain; Cyberpunk: acid rain; Post-apoc: ash fall |

## Multiplayer Design

**Architecture:** TCP authoritative server at 20 Hz tick; clients at 60 Hz local simulation.

**Position Sync:** Server sends delta-compressed `RaceSnapshot` (positions, velocities, item states, lap data) every 50 ms. Clients interpolate between last two confirmed snapshots for remote karts; local kart uses client-side prediction replayed on reconcile.

**Item Hit Validation:** Server maintains 100 ms position history per entity. On `ItemFireEvent`, server rewinds to `(server_tick − client_latency_ticks)`, checks AABB overlap, confirms or rejects. Client shows speculative hit effect; corrects on server rejection.

**Race State:** Lap counts and finish order are server-authoritative. `RaceStateMessage` broadcasts on each lap crossing and on race end. Clients display optimistic position until next broadcast.

**Rubber-banding:** Server computes position spread each tick. If last-place human is > 2 laps behind first, `RubberBandFactor` applied to AI kart speeds (max ×1.4); never applied to human karts.

**Anti-cheat:** Server validates all physics inputs against max-acceleration envelope; input sequences deviating > 3σ from physics model are discarded. Item fire rate capped server-side. Speed exploit detection via position-delta sanity check each tick.

**Lag Compensation Parameters:**

| Parameter | Default | Max (Tor) |
|-----------|---------|-----------|
| Input queue depth | 4 frames | 100 frames |
| Interpolation buffer | 3 snapshots | 5 snapshots |
| Hit validation history | 100 ms | 250 ms |
| Snapshot send interval | 50 ms | 50 ms |

## Genre Differentiation Matrix

| Aspect | Fantasy | Sci-fi | Horror | Cyberpunk | Post-Apocalyptic |
|--------|---------|--------|--------|-----------|-----------------|
| Track theme | Enchanted forest, castle ramparts | Orbital station, plasma tunnels | Haunted swamp, catacombs | Neon megacity rooftops | Wasteland ruins, collapsed highway |
| Hazards | Vines, magic storms, moving statues | Gravity wells, ion surges, laser gates | Creeping fog, ghost karts, jumpscare pillars | Police drone swarms, holographic walls | Sandstorm walls, falling debris, mine fields |
| Kart style | Wood-cart with arcane runes | Sleek pod with thruster glow | Hearse with bone trim | Chrome low-rider with neon strips | Scrap buggy with jury-rigged armour |
| Items / weapons | Magic bolt, speed potion, dragon shield, vine trap | Plasma shot, EMP burst, phase shield, gravity mine | Curse bolt, ghost possession, fog shroud, corpse mine | Virus hack (inverts controls), stun drone, firewall shield, spike strip | Molotov cocktail, chain whip, scrap shield, caltrops |
| Engine sound | Warm oscillator, slight flutter (+5% pitch) | Clean sine, high pitch (+30%) | Sub-bass growl (−30% pitch + vibrato) | Distorted, clipped (+40% pitch + hard clip) | Rough noise, chugging (−10% pitch + noise) |
| Music style | Orchestral motifs, 120 BPM | Synth-wave arpeggios, 140 BPM | Dissonant strings, 90 BPM | EDM drum machine, 160 BPM | Industrial percussion, 110 BPM |
| Visual palette | Warm earth: amber, forest green, gold | Cool: electric blue, white, cyan | Desaturated: grey-green, fog white | Neon: magenta, cyan, black | Sepia: rust, sand, ash grey |
| Post-processing | Soft bloom, no aberration | Scanlines, mild bloom | Heavy vignette, film grain | Heavy bloom + chromatic aberration | Dust particle overlay, reduced saturation |
| Genre-specific mechanic | Spell-charge boost (fill by hitting item boxes in sequence) | Boost from drafting at high speed (slipstream) | Fear meter — high fear = reduced grip | Hack combo — chain 3 items for overclock boost | Scavenge — break obstacles for scrap items |

## Success Criteria

| Indicator | Target | Measurement |
|-----------|--------|-------------|
| No embedded assets (fonts excepted) | 0 `//go:embed` directives outside `pkg/font/` | `grep -R "go:embed" pkg/ cmd/` returns zero matches outside `pkg/font/` |
| Deterministic PCG | Identical output for same seed + genre | 100-run idempotency test in `go test` |
| 8-player multiplayer | Race completes with 8 clients at 0 ms simulated latency | Integration test in `cmd/server/` |
| High-latency playable | Race finishes correctly at 2000 ms simulated RTT | Netcode integration test with `tc netem` |
| Five genres distinct | ≥ 4/5 differentiation score per genre | Internal rubric checklist |
| Frame budget | Client renders at 60 fps with 8 karts | `go test -bench=BenchmarkRenderFrame` ≤ 16 ms |
| Audio budget | `AudioSystem.Tick()` ≤ 1 ms per frame | `go test -bench=BenchmarkAudioTick` |
| Binary size | ≤ 25 MB stripped static binary | `wc -c way` after `make release` |
| Cold-start time | ≤ 3 s to main menu | Manual timing on reference hardware |
| Test coverage | ≥ 40 % per package | `go test -cover ./pkg/...` |

## Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| Procedural track self-intersection at extreme seeds | Medium | High | Add AABB sweep validator in generator; reject and re-seed up to 10 times |
| Netcode complexity causes desync at high latency | High | High | Port V-Series reconciliation loop first; add determinism tests before Phase 4 |
| Five-genre audio differentiation inaudible | Medium | Medium | Bind genre mod parameters to named constants; A/B test in CI audio benchmark |
| Binary size exceeds target from PCG tables | Low | Low | Profile with `go tool pprof`; compress lookup tables with zstd at init |
| Physics model divergence between client and server | High | High | Run identical physics tick on both; seed all randomness from server; fuzz test |
| Rubber-banding feels unfair to skilled players | Medium | Medium | Make `RubberBandFactor` a server-side config flag; expose in lobby settings |
| ELO ranking manipulation via disconnect | Medium | Medium | Only commit rating on clean race finish; detect rage-quit pattern |

## Plan Log

- 2026-02-28 ROADMAP.md created, all eight required sections complete.
