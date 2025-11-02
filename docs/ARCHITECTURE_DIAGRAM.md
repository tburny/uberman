# Uberman Testing Architecture

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                        Uberman Testing System                       │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                    ┌───────────────┴───────────────┐
                    │                               │
            ┌───────▼────────┐            ┌────────▼────────┐
            │  Native Tests  │            │  Docker Tests   │
            │   (Default)    │            │   (Optional)    │
            └───────┬────────┘            └────────┬────────┘
                    │                               │
        ┌───────────┼──────────┐        ┌──────────┼──────────────┐
        │           │          │        │          │              │
    ┌───▼───┐  ┌───▼────┐ ┌──▼────┐ ┌─▼─────┐ ┌──▼───────┐ ┌────▼─────┐
    │ Unit  │  │Property│ │Integr.│ │Isolate│ │Container │ │  tmpfs   │
    │ Tests │  │ Tests  │ │ Tests │ │  FS   │ │ Cleanup  │ │ (RAM-FS) │
    │ 45ms  │  │ 185ms  │ │ 4.8s  │ │ +2-3s │ │  Auto    │ │  1GB     │
    └───────┘  └────────┘ └───────┘ └───────┘ └──────────┘ └──────────┘
```

## Test Execution Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                     Developer Workflow                          │
└─────────────────────────────────────────────────────────────────┘

┌──────────────────────┐
│  Write/Modify Tests  │
└──────────┬───────────┘
           │
           ▼
   ┌───────────────┐        Fast?         ┌─────────────────┐
   │  Run Locally  ├────────Yes───────────►  go test        │
   └───────┬───────┘                       │  (native)      │
           │                                └────────┬────────┘
           │                                         │
           │                                         ▼
           │                                  ┌──────────────┐
           │                                  │  Tests Pass? │
           │                                  └──────┬───────┘
           │                                         │
           │                                    Yes  │  No
           │                                         │
           │◄────────────────────────────────────────┘
           │                                   (iterate)
           │
           ▼
      Isolated?
           │
           No───────────────────────────────────────┐
           │                                        │
           Yes                                      │
           │                                        │
           ▼                                        ▼
   ┌──────────────────┐                    ┌──────────────┐
   │  make test-docker│                    │ git commit   │
   │   -properties    │                    │              │
   └────────┬─────────┘                    └──────────────┘
            │
            ▼
     ┌─────────────┐
     │ Tests Pass? │
     └──────┬──────┘
            │
        Yes │  No
            │
            ├────────►(Fix issues)
            │
            ▼
    ┌───────────────┐
    │  git commit   │
    └───────────────┘
```

## Docker Test Architecture

```
┌────────────────────────────────────────────────────────────────────┐
│                  Docker Compose Test Services                      │
└────────────────────────────────────────────────────────────────────┘
        │
        ├─────► test-all          (All tests)
        ├─────► test-properties   (Property-based tests)  ⭐ Recommended
        ├─────► test-integration  (Integration tests + testcontainers)
        ├─────► test-coverage     (Coverage report generation)
        ├─────► test-appdir       (Specific: appdir package)
        └─────► test-config       (Specific: config package)

                        │
                        ▼

┌────────────────────────────────────────────────────────────────────┐
│                     Container Environment                          │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│  ┌──────────────────────────────────────────────────────────┐    │
│  │  Volumes:                                                 │    │
│  │    .:/app:ro           (Source code, read-only)          │    │
│  │    /tmp                (Anonymous volume)                │    │
│  │                                                           │    │
│  │  tmpfs:                                                   │    │
│  │    /workspace:1GB      (In-memory filesystem)            │    │
│  │                                                           │    │
│  │  Image: golang:1.24-alpine  (~350MB)                     │    │
│  └──────────────────────────────────────────────────────────┘    │
│                                                                    │
│  ┌──────────────────────────────────────────────────────────┐    │
│  │  Test Execution:                                          │    │
│  │    $ cd /app                                              │    │
│  │    $ go test ./...                                        │    │
│  │                                                           │    │
│  │  Filesystem Isolation:                                    │    │
│  │    ✓ All writes to tmpfs (RAM)                           │    │
│  │    ✓ Automatic cleanup on exit                           │    │
│  │    ✓ No host pollution                                    │    │
│  └──────────────────────────────────────────────────────────┘    │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘
```

## Test Categories & Strategy

```
┌─────────────────────────────────────────────────────────────────────┐
│                        Test Pyramid                                 │
└─────────────────────────────────────────────────────────────────────┘

                         ╱╲
                        ╱  ╲
                       ╱ E2E╲                (Future)
                      ╱──────╲
                     ╱        ╲
                    ╱Integration╲           ~5s (Docker)
                   ╱──────────────╲         MySQL, Services
                  ╱                ╲
                 ╱   Property-Based╲       ~185ms native
                ╱────────────────────╲     ~3.2s Docker ⭐
               ╱                      ╲    500+ directories
              ╱                        ╲
             ╱        Unit Tests        ╲  ~45ms (Fast!)
            ╱──────────────────────────────╲ Most tests here
           ╱                                ╲
          ╱──────────────────────────────────╲

Legend:
━━━ Recommended for Docker isolation
─── Can run natively (default)
```

## Performance Comparison

```
┌──────────────────────────────────────────────────────────────┐
│              Native vs Docker Performance                    │
└──────────────────────────────────────────────────────────────┘

Test Type          │ Native    Docker    Overhead    Use Case
───────────────────┼────────────────────────────────────────────
Unit tests         │  45ms    │ 2.1s  │  +2.0s   │ Development
Property tests     │ 185ms    │ 3.2s  │  +3.0s   │ CI/CD ⭐
Integration        │ 4.8s     │ 5.1s  │  +0.3s   │ Always Docker
───────────────────┴──────────────────────────────────────────

Overhead Breakdown:
  Container startup:    1.5s  ████████████████
  Volume mounting:      0.3s  ███
  tmpfs initialization: 0.2s  ██
                             ──────────────────
  Total overhead:       ~2-3s
```

## CI/CD Pipeline Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                   GitHub Actions Pipeline                       │
└─────────────────────────────────────────────────────────────────┘

┌────────────┐
│ git push   │
└──────┬─────┘
       │
       ▼
┌──────────────────────────────────────────┐
│          Parallel Jobs                    │
├───────────────────┬───────────────────────┤
│                   │                       │
▼                   ▼                       ▼
┌─────────────┐  ┌──────────────┐  ┌──────────────┐
│  unit-tests │  │property-tests│  │integration   │
│   (Native)  │  │  (Docker)    │  │  (Docker)    │
│             │  │              │  │              │
│  ~2 min     │  │  ~5 min      │  │  ~8 min      │
└──────┬──────┘  └──────┬───────┘  └──────┬───────┘
       │                │                  │
       └────────────────┴──────────────────┘
                        │
                        ▼
                ┌───────────────┐
                │   coverage    │
                │   (Docker)    │
                │               │
                │   ~3 min      │
                └───────┬───────┘
                        │
                        ▼
                ┌───────────────┐
                │Upload Codecov │
                └───────────────┘

Total Time: ~10 min (parallelized)
```

## Filesystem Isolation Strategy

```
┌─────────────────────────────────────────────────────────────┐
│           Property Test: Creating 500+ Directories          │
└─────────────────────────────────────────────────────────────┘

Native Execution:                Docker Execution:
┌────────────────┐              ┌──────────────────────┐
│  Host Machine  │              │  Container (Alpine)  │
│                │              │                      │
│  /tmp/         │              │  /workspace/  (tmpfs)│
│   uberman-*    │              │   [RAM-based FS]     │
│   ├─ test-1/   │              │   ├─ app-1/          │
│   ├─ test-2/   │              │   ├─ app-2/          │
│   ├─ ...       │              │   ├─ ...             │
│   └─ test-500/ │              │   └─ app-500/        │
│                │              │                      │
│  Cleanup:      │              │  Cleanup:            │
│  t.Cleanup()   │              │  Container exit      │
│  os.RemoveAll()│              │  (automatic)         │
│                │              │  tmpfs destroyed     │
└────────────────┘              └──────────────────────┘

Risk: Manual cleanup       Benefit: Automatic cleanup
```

## Container Lifecycle

```
┌─────────────────────────────────────────────────────────────────┐
│                 Docker Test Container Lifecycle                 │
└─────────────────────────────────────────────────────────────────┘

  make test-docker-properties
           │
           ▼
  ┌────────────────┐
  │ Build/Pull     │  (First time: ~2min, cached: ~5s)
  │ Image          │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Create         │  (~0.5s)
  │ Container      │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Mount Volumes  │  (~0.3s)
  │ - Source (ro)  │
  │ - tmpfs (1GB)  │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Start          │  (~0.2s)
  │ Container      │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Run Tests      │  (~185ms for property tests)
  │ go test ...    │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Collect        │  (~0.1s)
  │ Results        │
  └───────┬────────┘
          │
          ▼
  ┌────────────────┐
  │ Cleanup        │  (Automatic with --rm flag)
  │ - Stop         │
  │ - Remove       │
  │ - Destroy tmpfs│
  └────────────────┘

Total Time: ~3.2s
```

## Testcontainers Integration

```
┌─────────────────────────────────────────────────────────────────┐
│       Harmonious Coexistence: Docker Compose + Testcontainers  │
└─────────────────────────────────────────────────────────────────┘

                    Uberman Tests
                         │
        ┌────────────────┼────────────────┐
        │                │                │
        ▼                ▼                ▼
┌───────────────┐ ┌──────────────┐ ┌─────────────────┐
│ Filesystem    │ │  Property    │ │  Integration    │
│ Tests         │ │  Tests       │ │  Tests          │
│               │ │              │ │                 │
│ Native/Docker │ │ Docker Compose│ │ Testcontainers │
│ (Your choice) │ │ (Optional)   │ │ (Required)      │
└───────────────┘ └──────────────┘ └────────┬────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │  MySQL Container│
                                    │  (testcontainers│
                                    │   manages it)   │
                                    └─────────────────┘

Separation of Concerns:
- Docker Compose: Test execution environment
- Testcontainers: Service dependencies
```

## Memory Layout

```
┌─────────────────────────────────────────────────────────────────┐
│              Container Memory Allocation                        │
└─────────────────────────────────────────────────────────────────┘

Host Machine (16GB RAM)
├─ System:          4GB
├─ Applications:    8GB
└─ Docker:          4GB
    ├─ Alpine Container:
    │   ├─ Base:         100MB
    │   ├─ Go Runtime:   200MB
    │   ├─ Test Binary:  50MB
    │   ├─ tmpfs:        1GB   ⭐ Test filesystem
    │   └─ Overhead:     150MB
    │   ─────────────────────
    │   Total:          1.5GB per container
    │
    └─ MySQL Container (if running):
        └─ Database:     512MB

Recommendation: 4GB RAM minimum for comfortable testing
```

## Directory Structure in Container

```
┌─────────────────────────────────────────────────────────────────┐
│              Container Filesystem Layout                        │
└─────────────────────────────────────────────────────────────────┘

/                              (Alpine root)
├─ app/                        (Source code, read-only mount)
│  ├─ cmd/
│  ├─ internal/
│  │  ├─ appdir/
│  │  ├─ config/
│  │  └─ testutil/
│  ├─ go.mod
│  ├─ go.sum
│  └─ Makefile
│
├─ workspace/                  (tmpfs, 1GB, RAM-based)
│  └─ [test creates dirs here]
│
├─ tmp/                        (Anonymous volume)
│  └─ [go test temp files]
│
└─ usr/local/go/               (Go installation)
   └─ bin/
      └─ go

Key Points:
- /app: Read-only (protects source)
- /workspace: tmpfs (fast, isolated)
- /tmp: Anonymous volume (auto-cleanup)
```

## Summary: When to Use What

```
┌─────────────────────────────────────────────────────────────────┐
│                   Testing Strategy Decision Tree                │
└─────────────────────────────────────────────────────────────────┘

                 Need to run tests?
                        │
                        ▼
           ┌────────────┴───────────┐
           │                        │
     TDD Development?          CI/CD Pipeline?
           │                        │
           Yes                      Yes
           │                        │
           ▼                        ▼
    ┌─────────────┐          ┌──────────────┐
    │   Native    │          │    Docker    │
    │ go test ... │          │ make test-   │
    │             │          │   docker-*   │
    │ Fast: 45ms  │          │              │
    │             │          │ Isolated: 3s │
    └─────────────┘          └──────────────┘
           │                        │
           │                        │
      Need isolation?               │
           │                        │
           Yes                      │
           │                        │
           ▼                        │
    ┌─────────────┐                │
    │   Docker    │                │
    │ make test-  │                │
    │   docker-   │                │
    │ properties  │                │
    │             │◄───────────────┘
    │ Verified: 3s│
    └─────────────┘
           │
           ▼
      git commit
```

## Legend

```
Symbol Key:
  ⭐ Recommended approach
  ✓  Implemented feature
  ▼  Process flow
  │  Connection
  ├─ Branch
  └─ End of branch
  ╱╲ Pyramid shape
```

---

**Note**: All diagrams created with ASCII art for maximum compatibility and clarity.
