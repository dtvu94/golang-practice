# Eight Weeks to Go

> Twenty hours a week, six sessions, three projects. This is not a plan for learning to
> program — it is a plan for learning one small language, its concurrency model, and the
> idioms that separate someone who *writes Go* from someone writing C# in Go's syntax.

**160 hours · 48 sessions · 3 projects · targets Go 1.27**

| | |
|---|---|
| **Budget** | 160 hours over 8 weeks |
| **Weekdays** | Mon, Tue, Thu, Fri — 2 h each |
| **Weekends** | Sat, Sun — 6 h each |
| **Starting point** | 8 years of senior engineering in TypeScript/Node, Python, C#, C++ |
| **Goal** | Hireable as a senior Go engineer on infrastructure-shaped teams |

---

## Contents

- [What 160 hours actually buys you](#what-160-hours-actually-buys-you)
- [How to use the two block sizes](#how-to-use-the-two-block-sizes)
- [Repository layout](#repository-layout)
- [The project ladder](#the-project-ladder)
- [Week 1: Language core](#week-1-language-core)
- [Week 2: Idiom](#week-2-idiom)
- [Week 3: Concurrency](#week-3-concurrency)
- [Week 4: Craft](#week-4-craft)
- [Week 5: Services](#week-5-services)
- [Week 6: Capstone, build](#week-6-capstone-build)
- [Week 7: Capstone, harden](#week-7-capstone-harden)
- [Week 8: Proof](#week-8-proof)
- [What each of your languages will get wrong](#what-each-of-your-languages-will-get-wrong)
- [Reading list](#reading-list)
- [Signals that it worked](#signals-that-it-worked)
- [One correction to the premise](#one-correction-to-the-premise)

---

## What 160 hours actually buys you

Go's specification is short enough to read in an afternoon — that is the entire design goal.
With C++ and C# behind you, the syntax costs a week, not a month. The real distance sits in
three places: **errors as values** instead of exceptions, the **goroutine/channel/context
model**, and a body of **idiom** that a Go reviewer spots violations of within thirty seconds.

At 20 hours a week this is a genuinely serious budget — roughly four times what a casual
learner puts in, and enough to arrive at week eight with three working projects, measured
performance numbers, and a merged upstream contribution.

What it does *not* buy is five years of production scar tissue. That is fine, because it is not
what the role screens for. Eight years across four companies and six-plus products already
bought the expensive things: system decomposition, knowing which tests matter, operating what
you shipped, reading code you did not write. Go is a notation for those skills, and this plan
exists to prove you have the notation.

**One honest warning.** 20 hours on top of full-time work is sustainable for eight weeks and
not much longer. Treat it as a sprint with a defined end date. When a week goes badly — and one
will — cut weekday depth first and protect the weekend project blocks. The projects are what
you will be interviewed about; the weekday sessions are the scaffolding.

---

## How to use the two block sizes

**Weekday · 2 hours.** One concept, one small program, finished. Never start something you
cannot close out in the session — an unfinished weekday exercise carries no memory into the
next day. Read for at most 30 minutes, then type code for 90.

**Weekend · 6 hours.** Project work. Spend the first 20 minutes re-reading your own code from
the last session before writing anything. Break at the 3-hour mark. End every block with a
commit that builds and a one-line note about where you stopped.

**Everything is public.** Commit on every session, CI green from week one. Eight weeks of
dense, visible commit history is itself a hiring signal — and one nobody can fake
retroactively.

**Never read passively.** No Go article gets read without the code being typed into a file and
run. Reading Go and writing Go build different instincts, and interviews only test the second.
When you use a standard library function, open its source — it is the canonical style guide.

---

## Repository layout

Every folder is created up front with a `.gitkeep`, so week one starts with somewhere to put
things rather than with a decision about where things go.

```
golang-practice/
├── README.md              # this plan, and your progress tracker
├── tour/                  # week 1 · every Tour snippet, typed by hand
├── mechanics/             # week 1 · proof-tests: slices, defer, zero values, nil traps
├── treesum/               # weeks 1-4 · filesystem snapshot & diff tool
├── exercism/              # week 2 · solutions plus notes on community idiom
├── concurrency/           # week 3 · rate limiter, singleflight, parallel map, worker pool
├── linkr/                 # week 5 · production JSON API
├── kvd/                   # weeks 6-7 · replicated key-value store
└── notes/                 # weeks 2 & 8 · design docs, codebase walkthrough, interview answers
```

Weeks 1-5 belong here, where a dense and visible eight-week commit history is itself the point.

**Consider splitting `kvd` into its own repository** when you start week six. It is the
interview centrepiece, and a capstone buried inside a repo called `golang-practice` reads as an
exercise rather than as a system. The folder is here so the decision is not blocking — delete it
if you split it out.

---

## The project ladder

Three projects, each deliberately harder than the last. The first is carried across four weeks
and grows a new capability with every skill you acquire, so you can watch your own Go improving
inside a single codebase.

### `treesum` — filesystem snapshot & diff tool · weeks 1–4

Walks a directory tree, hashes every file, writes a JSON manifest, and diffs two manifests.
Starts sequential and naive in week one; becomes idiomatic in week two, concurrent in week
three, and tested, profiled and released in week four. **Zero external dependencies** —
everything comes from the standard library.

### `linkr` — a production JSON API · week 5

A link-shortener and snippet service with Postgres, migrations, token auth, per-client rate
limiting, structured logging, Prometheus metrics, graceful shutdown and integration tests
against a real database. The reference shape of a Go web service.

### `kvd` — a replicated key-value store · weeks 6–7

The capstone. In-memory store behind a write-ahead log with crash recovery, gRPC and HTTP
interfaces, single-leader asynchronous replication to N followers, TTL expiry, chaos tests,
load numbers, and a three-node cluster you can actually run. Deliberately the shape of problem
Go infrastructure teams hire for.

---

## Week 1: Language core

**20 hours.** You are not learning to program this week, you are scanning for surprises. Move
through the Tour quickly and spend the saved time on the handful of mechanics that your
existing languages will actively mislead you about.

- [ ] **Mon · 2 h — Toolchain and Tour, part one.** Install Go 1.27, `gopls`, `dlv`,
  `golangci-lint`. Configure your editor for format-on-save. `go mod init`. Work the Tour's
  Basics and Flow-control sections — typing every snippet into a local file, never the
  playground, because you also need the edit-compile-run loop in muscle memory.
- [ ] **Tue · 2 h — Tour, part two: methods, interfaces, generics.** Then close the browser and
  write ten tiny programs from memory: a struct with both a value-receiver and a
  pointer-receiver method that demonstrably differ; a variadic function; a closure-based
  counter; a map with struct values you cannot assign to; a type switch. Recall without
  reference is the point.
- [ ] **Thu · 2 h — Slices, properly.** The highest-return two hours of the week. Read the Go
  blog's *Arrays, slices and the mechanics of append*, then write a test file that **proves**,
  with assertions: append aliasing into a shared backing array; the exact capacity growth
  pattern; `copy` semantics; three-index slicing `s[a:b:c]` and why it exists; and the
  sub-slice memory-retention leak where a 3-byte slice pins a 10 MB array.
- [ ] **Fri · 2 h — Zero values, defer, panic.** Experiments, each in its own test: reading a
  nil map versus writing to one; appending to a nil slice; a zero `sync.Mutex` being
  immediately usable; `defer` argument evaluation happening at defer time, not call time; LIFO
  ordering; `defer` inside a loop accumulating until function exit; a named return value
  mutated by a deferred closure; `recover` failing to catch a panic from a different goroutine.
- [ ] **Sat · 6 h — `treesum` v0.1, walk and hash.** Build the `snapshot` subcommand: traverse
  a directory with `filepath.WalkDir`, SHA-256 every file, and emit a JSON manifest of path,
  size, mode, modtime and hash. Sequential and unclever. Subcommands via `flag.NewFlagSet`, no
  CLI library. Handle permission errors without aborting the whole walk. Standard library only.
- [ ] **Sun · 6 h — `treesum` v0.2, diff.** Load two manifests and classify every path as
  added, removed, modified, or renamed (same hash, different path). Human-readable output plus
  a `--json` flag. Add exclude patterns, symlink handling, and a stable sort so output is
  deterministic. Write your first table-driven tests against `testing/fstest.MapFS` so no test
  touches the real disk — this constraint will quietly force better design.

> **Ship:** `treesum` runs end to end on a real directory tree, diffs two snapshots correctly,
> and has tests that run in under a second without disk I/O.

**Read:** A Tour of Go · Effective Go · Go Code Review Comments · Rob Pike, *Go Proverbs* (20 min talk)

---

## Week 2: Idiom

**20 hours.** This is the week that decides whether a Go reviewer reads your code as Go. Almost
nothing here is enforced by the compiler — it is all convention, which is precisely why
transplanted engineers get caught by it.

- [ ] **Mon · 2 h — Errors are values.** `errors.New`, `fmt.Errorf` with `%w`, `errors.Is`,
  `errors.As`, `errors.Join`. Build a three-level call chain that wraps at each level, then
  assert on the root cause from the top. Write your own error type with an `Unwrap` method and
  a field the caller can actually use. Decide, in writing, when you would use a sentinel versus
  a typed error.
- [ ] **Tue · 2 h — Interfaces and the nil trap.** Implicit satisfaction; defining interfaces at
  the consumer, not shipping them with the implementation; accept interfaces, return structs;
  one or two methods wide. Write a *failing* test that reproduces a nil `*MyError` stored in an
  `error` interface comparing non-nil, then explain it in a comment in your own words. Study
  `io.Reader`, `io.Writer` and how `io` composes decorators.
- [ ] **Thu · 2 h — Package design.** Read the layout of two standard library packages and one
  popular repository. Learn `internal/`, doc comment conventions, `go doc`, naming that reads at
  the call site (`chi.NewRouter`, not `chi.NewChiRouter`), and why `utils`/`helpers`/`common`
  packages are a smell. Restructure `treesum` into `cmd/treesum` plus `internal/manifest`,
  `internal/walk`, `internal/diff`.
- [ ] **Fri · 2 h — Generics, and when not to.** Type parameters, constraints, the `~`
  approximation, `comparable`, and where inference gives up. Go 1.27's new generic methods. Read
  the `slices` and `maps` standard library packages end to end — they are the best available
  example of tasteful generic Go. Write one generic helper you would genuinely keep, then
  honestly assess whether a plain loop would have been clearer.
- [ ] **Sat · 6 h — `treesum` v0.3, the idiomatic refactor.** Rewrite against everything above:
  consumer-defined `Hasher` and filesystem interfaces; hashing that takes an `io.Reader` so
  tests never touch disk; a coherent error set with sentinels and wrapping; configuration
  through a struct or functional options rather than a widening parameter list. Every exported
  symbol documented. `golangci-lint run` clean with no suppressions.
- [ ] **Sun · 6 h — Write it down, then drill.** First 2 hours: write `DESIGN.md` for `treesum`
  explaining every interface boundary and why it sits where it does — articulating design
  decisions in prose is directly rehearsing the interview. Remaining 4 hours: ten Exercism Go
  exercises on strings, slices, maps and interfaces, and after each one read three community
  solutions. That comparison is the single fastest way to absorb idiom.

> **Ship:** `treesum` restructured, documented, lint-clean, with a design document. Tag `v0.3.0`.

**Read:** Go blog *Working with Errors in Go 1.13* · *100 Go Mistakes* ch. 1–7 · Learn Go with Tests

---

## Week 3: Concurrency

**20 hours.** Every Go interview goes here. Coming from `async/await` and `Task`, the shift is
that goroutines are cheap blocking threads of control, not futures — you do not compose them
with combinators, you coordinate them with channels and `context`.

- [ ] **Mon · 2 h — Goroutines and channels.** Unbuffered handoff versus buffered capacity;
  directional types `chan<-` and `<-chan`; closing as a broadcast; `range` over a channel; a nil
  channel blocking forever inside `select`, and why that is useful. Build: ping-pong between two
  goroutines, a generator function returning a receive-only channel, and a done-channel shutdown.
- [ ] **Tue · 2 h — `select` and `context`.** Cancellation, deadlines, `WithCancelCause` and
  `context.Cause`. The rules: first parameter, named `ctx`, never stored in a struct, never nil.
  Take one blocking function you have already written and make it respect cancellation at every
  single blocking point — then write a test that cancels mid-flight and asserts it returned
  promptly.
- [ ] **Thu · 2 h — `sync` and `atomic`.** `Mutex`, `RWMutex`, `WaitGroup`, `Once`, `sync.Map`
  and its two narrow use cases, `atomic.Int64`, and `errgroup` with `SetLimit`. Build a
  concurrency-safe cache behind a mutex, then deliberately break it and watch `go test -race`
  catch you. Write down the rule you will use for choosing a mutex over a channel.
- [ ] **Fri · 2 h — Patterns and leaks.** Implement pipeline, fan-out/fan-in, worker pool and a
  semaphore-bounded limiter. Then deliberately write four goroutine leaks — an abandoned
  receiver, a blocked send with no reader, `time.After` inside a hot loop, a missing context
  check — and catch each one using Go 1.27's `goroutineleak` pprof profile. Being able to
  demonstrate this on demand is a strong senior signal.
- [ ] **Sat · 6 h — `treesum` v0.4, make it concurrent.** Parallel hashing through a bounded
  worker pool sized from `GOMAXPROCS`, with results merged back in deterministic order. Context
  cancellation on Ctrl-C that drains cleanly. First error cancels the remaining work via
  `errgroup`. Must pass `go test -race`. Measure wall-clock against a large tree before and
  after, and commit the numbers — this is your first quantified claim.
- [ ] **Sun · 6 h — The interview exercise set.** Four self-contained implementations, each with
  tests, each roughly 90 minutes:
  1. a token-bucket rate limiter written from scratch, then compared against `x/time/rate`;
  2. a `singleflight`-style duplicate-request suppressor;
  3. a bounded parallel `Map` helper that respects context;
  4. a worker pool that drains in-flight work on shutdown instead of dropping it.

  These four come up constantly in Go interviews.

> **Ship:** concurrent `treesum` with a measured speedup, plus four standalone concurrency
> primitives you wrote yourself and can rebuild from memory.

**Read:** *Concurrency in Go* (Cox-Buday) · Go blog *Pipelines and cancellation*, *Context*

---

## Week 4: Craft

**20 hours.** Go's tooling is a real differentiator and it is unusually easy to demonstrate. A
candidate who can read a pprof profile out loud reads as senior immediately, because most
candidates cannot.

- [ ] **Mon · 2 h — Test style.** Table-driven tests with `t.Run` subtests, `t.Helper`,
  `t.Cleanup`, `t.Parallel` and its shared-state hazards, golden files, and `go-cmp` with custom
  comparers. Understand why the community leans on the standard library rather than an assertion
  framework, so you can answer the question when it is asked.
- [ ] **Tue · 2 h — Deterministic concurrency tests.** `testing/synctest`, stable since Go 1.25,
  gives you a fake clock and a bubble that knows when all goroutines are blocked. Rewrite two of
  week three's timing-dependent tests to be instant and flake-free — no `time.Sleep` anywhere.
  Also cover `httptest`, `iotest` and `fstest`.
- [ ] **Thu · 2 h — Fuzzing.** `go test -fuzz` against the manifest parser and the diff function,
  with a seed corpus. Let it run while you read, then fix everything it finds and commit the
  crashers as regression tests. Any code that parses untrusted bytes should be fuzzed, and saying
  so in an interview lands well.
- [ ] **Fri · 2 h — Benchmarks.** `testing.B`, `b.ResetTimer`, `b.ReportAllocs`, `-benchmem`, and
  a package-level sink variable to defeat compiler elimination. Use `benchstat` to compare runs
  with statistical honesty rather than eyeballing one number. Benchmark hashing throughput and
  manifest encode/decode.
- [ ] **Sat · 6 h — Profiling day.** Take CPU, heap, allocation, mutex and block profiles of
  `treesum` against a large tree. Read the flame graphs. Run `go tool trace` and watch the
  scheduler's actual behaviour. Inspect escape analysis with `go build -gcflags='-m'`. Find and
  fix at least two real problems — an avoidable allocation in the hot path, an oversized buffer,
  or lock contention — and record before-and-after `benchstat` output in the README.
  Evidence-backed optimisation is the whole point.
- [ ] **Sun · 6 h — Toolchain and release day.** `go vet`, staticcheck and a tuned
  `golangci-lint` config. Modules in depth: `go mod tidy/why/graph`, minimal version selection,
  `replace`, `go work`, vendoring. Build tags, `//go:embed`, `-ldflags` version stamping, and
  cross-compilation to linux/arm64 and darwin/arm64. Then a GitHub Actions pipeline: matrix
  build, vet, lint, race tests, coverage, and a GoReleaser binary release. Tag `treesum v1.0.0`.

> **Ship:** `treesum v1.0.0` — released binaries for three platforms, green CI, fuzz corpus, and
> documented performance work with real numbers. Project one is done.

**Read:** Go blog *Profiling Go Programs*, *Testing concurrent code with testing/synctest*, *Fuzzing*

---

## Week 5: Services

**20 hours.** Go's culture leans hard on the standard library. Reaching for a framework where
`net/http` would do is read as a tell that someone has not internalised the language, so build
this one the way Go teams actually build it.

- [ ] **Mon · 2 h — `net/http` from the ground up.** The `http.Server` struct and every timeout
  field that matters in production; the method-and-wildcard `ServeMux` patterns added in 1.22;
  `Handler` versus `HandlerFunc`; middleware as `func(http.Handler) http.Handler`;
  request-scoped context values behind unexported key types. Build a middleware chain: request
  ID, structured logging, panic recovery, timeout.
- [ ] **Tue · 2 h — JSON, done properly.** Struct tags, `omitempty` versus `omitzero` (new in
  1.24), custom `MarshalJSON`, decoding with `DisallowUnknownFields`, streaming via
  `json.Decoder`, and a validation layer that returns field-level 4xx errors instead of a bare
  400. Then read up on Go 1.27's rewritten `encoding/json/v2` — what changed, why, and what it
  means for your code.
- [ ] **Thu · 2 h — Postgres.** `database/sql` versus pgx native; pool sizing and lifetime
  settings; `QueryContext` everywhere; `sql.Null*` and scanning; transactions with correct
  rollback via `defer`; and `sqlc` to generate typed queries from SQL. Set up migrations with
  `golang-migrate` or `goose` and make them run on startup behind a flag.
- [ ] **Fri · 2 h — Observability and lifecycle.** `log/slog` with a JSON handler and
  request-scoped attributes; the Prometheus client with RED metrics (rate, errors, duration) as
  a histogram; OpenTelemetry tracing basics; configuration from environment and flags with no
  framework; and `signal.NotifyContext` plus `srv.Shutdown(ctx)` for a shutdown that actually
  drains.
- [ ] **Sat · 6 h — `linkr` v0.1, the API.** `POST /v1/links` creating a short code,
  `GET /{code}` redirecting, plus list with pagination, update, and soft delete. Expiry times.
  API-token authentication middleware with hashed tokens in Postgres. Per-token rate limiting.
  Consistent JSON error envelopes with field-level validation detail. Every request logged with
  a correlation ID that flows through to the database layer.
- [ ] **Sun · 6 h — `linkr` v0.2, make it operable.** Integration tests with testcontainers-go
  against a real Postgres; handler tests with `httptest`; `/healthz`, `/readyz` and `/metrics`;
  graceful shutdown verified under load; an OpenAPI spec; a multi-stage Dockerfile onto
  distroless with the final image size stated in the README; and a `docker compose up` that
  brings the whole thing to life in one command.

> **Ship:** `linkr v1.0.0` — a service someone else can clone, run with one command, and hit
> with curl inside five minutes. Project two is done.

**Read:** *Let's Go* and *Let's Go Further* (Alex Edwards) — work through them alongside this week

---

## Week 6: Capstone, build

**20 hours.** A replicated key-value store: chosen because it forces storage, protocol design,
concurrency and failure handling into one codebase, and because it is exactly the class of
problem Go infrastructure teams interview for.

> **Alternative.** If you are aiming at product-backend roles rather than infrastructure, swap
> in a **distributed job queue** with at-least-once delivery, visibility timeouts and a dead
> letter queue. Same skills, different vocabulary.

- [ ] **Mon · 2 h — Design first, in writing.** Before any code: the API surface, the on-disk
  format, the replication model, the failure modes you will handle, and — most importantly — the
  consistency guarantees you will **not** provide. Scope it explicitly: hand-rolled single-leader
  replication with a configured leader is the target; Raft leader election is a stretch goal, not
  the plan. Deciding what to leave out and saying so is senior behaviour.
- [ ] **Tue · 2 h — Storage engine.** An in-memory `map[string][]byte` behind an `RWMutex`,
  fronted by a write-ahead log of length-prefixed records with CRC32 checksums. Startup replays
  the WAL to rebuild state. Write the crash-recovery test *first*: kill the process mid-write,
  restart, assert the store is consistent and the torn tail record was discarded.
- [ ] **Thu · 2 h — Compaction and expiry.** Snapshot the live state and truncate the WAL so it
  does not grow without bound. A background sweeper for TTL expiry. Make the `fsync` policy
  configurable — per-write, batched, or never — then measure all three, because the
  durability-versus-throughput tradeoff is one of the best things you can have numbers for.
- [ ] **Fri · 2 h — Wire protocol.** A protobuf service definition with Get, Set, Delete and a
  server-streaming Watch, served over gRPC, plus a thin HTTP/JSON gateway for curl-ability.
  Interceptors for logging, metrics and panic recovery. Get comfortable with `protoc`, generated
  code checked into the repo, and streaming semantics under cancellation.
- [ ] **Sat · 6 h — Replication.** One leader, N asynchronous followers. The leader appends to
  its WAL, then fans out entries; followers acknowledge; the leader tracks a per-follower offset.
  A follower that reconnects catches up by streaming from its last known offset. Reads are
  servable from a follower (possibly stale) or forced to the leader (fresh) via a request flag —
  and the README must say plainly which one gives you what.
- [ ] **Sun · 6 h — Cluster runtime.** Membership from configuration, health checking between
  nodes, follower reconnection with exponential backoff and jitter, and read-your-writes via an
  offset token the client carries. Per-node Prometheus metrics for replication lag, WAL size and
  apply rate. Finish with a `docker compose` three-node cluster that starts, replicates and
  survives you killing a container.

> **Ship:** a three-node cluster running locally, replicating writes, surviving a follower
> restart, and exposing metrics. Rough edges are fine — week seven is for hardening.

---

## Week 7: Capstone, harden

**20 hours.** The entire difference between a bootcamp portfolio project and a senior one lives
in this week. Nobody is impressed that it works. They are looking for evidence that you know how
it fails.

- [ ] **Mon · 2 h — Deterministic protocol tests.** Use `testing/synctest` to test replication
  timing, follower catch-up and heartbeat behaviour with a fake clock — no sleeps, no flakes,
  milliseconds to run. A distributed system with a fast deterministic test suite is a genuinely
  uncommon thing to show up with.
- [ ] **Tue · 2 h — Chaos tests.** Kill a follower mid-write and assert it catches up. Partition
  the leader and assert the documented behaviour actually happens. Restart a node from a cold
  WAL. Corrupt the WAL's tail bytes and assert clean recovery rather than a panic. Each of these
  is a test you can point at and say "here is the failure I designed for".
- [ ] **Thu · 2 h — Fuzz the codecs.** The WAL record decoder and the protocol framing both parse
  bytes you did not write. Fuzz both, fix everything found, and keep the crashers as a regression
  corpus. Add a maximum-record-size guard so a corrupt length prefix cannot make you allocate
  four gigabytes.
- [ ] **Fri · 2 h — Load test.** Write a Go load generator — a concurrency ladder driving
  increasing client counts. Record writes per second, p50/p95/p99 latency, memory at steady
  state, replication lag under load, and the delta between fsync-per-write and batched fsync. Put
  the table in the README. Measured numbers separate you from every candidate who says "it's
  fast".
- [ ] **Sat · 6 h — Optimise with evidence only.** Profile first, then act: buffer reuse or
  `sync.Pool` to cut allocations in the hot path; a mutex profile pointing at contention,
  answered with a sharded map; batched WAL writes to amortise syscalls. Every single change
  justified by `benchstat` before and after. Finish with a clean `goroutineleak` profile and
  green `-race` in CI. Never optimise anything you have not measured.
- [ ] **Sun · 6 h — The README, and Kubernetes.** Architecture diagram; the guarantees you
  provide and explicitly those you do not; the measured numbers; the tradeoffs you chose and why;
  what you would change at a hundred times the load; what you knowingly left out and what it
  would cost to add. Then a StatefulSet with liveness and readiness probes, resource limits and a
  PVC. Tag `v1.0.0`.

> **Ship:** that README is the highest-leverage artefact of all 160 hours. Most interviewers will
> read it closely and skim the code — write it as though a staff engineer is assessing your
> judgement, because that is exactly what is happening.

---

## Week 8: Proof

**20 hours.** Two goals: convert reading fluency into a credible external claim, and make sure
the market can actually find you. The open-source work starts on Monday because review turnaround
is measured in days, not hours.

- [ ] **Mon · 2 h — Read production Go, part one.** Pick one codebase — Prometheus, NATS, Caddy,
  or etcd — and trace a single request or message end to end through it, taking notes as you go.
  Being able to say "here is how Prometheus handles a scrape" is worth more in an interview than
  another project of your own.
- [ ] **Tue · 2 h — Read production Go, part two.** Same repository, now studying its concurrency
  structure and its testing strategy specifically. How do they avoid leaks? What do they mock and
  what do they run for real? Write a thousand-word walkthrough and publish it — a written artefact
  of reading fluency is unusual and easy to link.
- [ ] **Thu · 2 h — Claim an issue.** Find a good-first-issue in an active Go project, ideally the
  one you just read. A documentation correction, a missing test case, or a small bug all count
  equally; the merge is the signal, not the diff size. Read the contributing guide properly and
  open the PR *today* so the review clock starts.
- [ ] **Fri · 2 h — Land it.** Respond to review comments, rebase, get it merged. Being reviewed
  by a Go maintainer is the highest-quality feedback available to you and it costs nothing. If the
  review is still open, use the time on a second, smaller PR elsewhere.
- [ ] **Sat · 6 h — Interview drilling.** Three hours writing out answers and saying them aloud:
  slice aliasing after append; nil interface versus nil pointer; buffered versus unbuffered
  semantics; how `select` chooses among ready cases; what context cancellation actually does; when
  a mutex beats a channel; the GMP scheduler and what happens on a blocking syscall; escape
  analysis; GC tuning with `GOGC` and `GOMEMLIMIT`; `sync.Pool` semantics. Then three timed
  live-coding reps at 30 minutes each, no documentation open: a worker pool, a rate limiter, an
  LRU cache.
- [ ] **Sun · 6 h — Positioning, and apply.** Rewrite your CV and LinkedIn around distributed
  backend and infrastructure engineering, with Go named explicitly and all three projects linked.
  Polish every README until a stranger can run the project in under five minutes. Write your "why
  Go, why now" narrative connecting eight years, four companies and six-plus products to platform
  work — you will tell this story in every first-round call. Then send the first ten applications.
  Do not wait to feel ready; you will not, and interviewing is the fastest feedback loop remaining.

> **Ship:** one merged upstream PR, one published codebase walkthrough, three polished projects,
> a rewritten CV, and ten applications sent.

---

## What each of your languages will get wrong

Fluency in five languages is mostly an advantage, but each one has installed a reflex Go will
punish. Keep this beside you for the first month.

| Coming from | The reflex to unlearn |
|---|---|
| **TypeScript / Node** | There is no async colouring — everything is blocking code plus goroutines, so stop looking for `await`. Do not build a `Promise.all` abstraction; that is `errgroup`. No `undefined`, no union types, and no structural typing at runtime — interface satisfaction is a compile-time check. And you now have real shared-memory data races, which the single-threaded event loop used to prevent for you. |
| **Python** | No exceptions: every error is a returned value, and ignoring one is a review comment. No duck typing at runtime. Write the explicit loop rather than hunting for a comprehension. No GIL means genuine parallelism, which means genuine races. And the package system has nothing resembling relative imports or `__init__` magic — the import path is the directory path. |
| **C#** | No classes and no inheritance — embedding promotes fields and methods but there is no virtual dispatch and no `base`. No DI container: you wire dependencies by hand in `main()`, and Go teams consider that a feature rather than a gap. `any` is not `object` — it requires a type assertion. No LINQ, no attributes, no `try/catch/finally`; `defer` covers only the last of those. |
| **C++** | No RAII: `defer` runs at *function* exit, not scope exit, so deferring inside a loop accumulates until the function returns. Garbage collection means you stop managing lifetimes and start caring about allocation rate and escape analysis. Generics are deliberately far weaker than templates — no specialisation, no metaprogramming. No operator overloading, no move semantics, no headers, and a compiler that refuses to build on an unused variable. |
| **Everyone** | Slices share backing arrays, so `append` sometimes mutates the caller's data and sometimes does not, depending entirely on capacity. **This is the bug you will write.** It is why week one spends a whole session on it, and why week four revisits it. |

---

## Reading list

Six entries, ordered by when they earn their keep. Resist adding more — collecting resources is
the classic way to feel productive while learning nothing.

| When | What | Why |
|---|---|---|
| Week 1 | **The Tour, Effective Go, Go Code Review Comments** | Free, official and short. The latter two are effectively the criteria your interview code will be judged against, so read them properly rather than skimming. |
| Weeks 1–2 | **Learn Go with Tests** | Free online. Teaches the language and Go's testing culture at the same time, which is exactly the sequencing you want. |
| Weeks 2–5 | **100 Go Mistakes and How to Avoid Them** — Teiva Harsanyi | The single highest-value book for this situation. Written for competent engineers new to Go, and a catalogue of precisely the mistakes an experienced polyglot makes. Read a few entries in every weekday session rather than in one pass. |
| Week 3 | **Concurrency in Go** — Katherine Cox-Buday | Slightly dated on tooling, still the clearest explanation of *why* the model is shaped the way it is. Skim the pattern chapters, read the rationale chapters closely. |
| Week 5 | **Let's Go** & **Let's Go Further** — Alex Edwards | The most directly applicable guides to building a real Go web service with the standard library. Working through them *is* most of week five. |
| Throughout | **The Go Blog** | Especially the posts on errors, context, pipelines, slice internals, profiling, fuzzing and `testing/synctest`. Primary sources, written by the people who made the decisions. |

---

## Signals that it worked

Not a feeling of readiness — that arrives late and unreliably. These are observable, and most
should be tickable by week eight.

- [ ] You considered a channel, then used a mutex instead, and can explain in one sentence why that was the better call.
- [ ] You have read a panic stack trace from inside a goroutine and found the cause without reaching for a debugger.
- [ ] You can explain what happens when `append` exceeds capacity, and you have a committed test that proves the aliasing behaviour.
- [ ] You can name three optimisations you made because a profile told you to, each with before-and-after numbers.
- [ ] You wrote an interface, then deleted it, because the concrete type was the right answer.
- [ ] Someone else's Go — a standard library file, a Prometheus package — reads at normal speed rather than being decoded line by line.
- [ ] You have an opinion about when generics are worth it, and that opinion is "rarely".
- [ ] Your test suite for a distributed system runs in under two seconds and has never flaked.
- [ ] A stranger can clone any of the three projects and have it running within five minutes.

---

## One correction to the premise

> **Worth knowing before committing 160 hours.**
>
> Go does not have more open roles than JavaScript. In raw volume it is not close — JS and
> TypeScript postings outnumber Go postings by a wide margin in every major market. If total job
> count is the goal, this plan is the wrong move.
>
> What is true, and is the better reason to do this: the **ratio** of openings to qualified
> candidates is far healthier in Go, the roles skew senior and infrastructural rather than junior
> and product-shaped, compensation bands sit higher, and the ecosystem churns dramatically less —
> the Go you learn this year is still the Go you use in five years, which is not a claim the
> JavaScript ecosystem can make. For an engineer eight years in who wants to move toward platform
> and backend infrastructure rather than compete in the largest and most crowded pool, that is a
> good trade.
>
> The practical consequence lands in week eight: target **infrastructure, observability,
> developer tooling, fintech and cloud-native teams** specifically. Those teams will trade Go
> tenure for systems experience, and these three projects are pitched directly at them. Applying
> to generic product-backend roles as "the Go candidate with two months of Go" is the version of
> this that does not work.

---

<sub>Targets Go 1.27 (August 2026) — generic methods, `encoding/json/v2`, and the `goroutineleak`
pprof profile. `testing/synctest` stable since Go 1.25. 48 sessions · 160 hours · 3 projects.</sub>
