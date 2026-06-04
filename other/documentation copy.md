# documentation.ubuntu.com Search — Project Plan

## Project Overview

### Vision

Replace the built-in ReadTheDocs/Sphinx search on documentation.ubuntu.com with a modern, OpenSearch-based semantic search system. The current search suffers from poor relevance and has no capability for suggestions or AI-assisted exploration. The new system will serve both AI agents and human users, with domain-aware scoping and high-quality results.

### Priorities

| Priority | Focus Area |
|----------|-----------|
| P1 | AI search support |
| P2 | Search quality improvements |
| P3 | Suggestions |
| P4 | Other niceties |

### Timeline

Delivery target: **26.10 cycle** (October 2026)

### Team

| Person | Role | Responsibilities |
|--------|------|-----------------|
| Michael Park | Project Lead | Implementation, architecture, integration |
| Nicholas Bello | Web Developer | Search UI prototype, suggestions UI |
| Mehdi Bendriss | OpenSearch Engineer | OpenSearch deployment, indexing, performance, COS, production ops |
| Vladimir Izmalkov | Ulwazi Theme Maintainer | Theme integration, review |

### Cross-cutting Acceptance Criterion

> **All search queries must return results within 2 seconds. If the 2-second threshold is exceeded, the system must automatically fall back to ReadTheDocs search.**

This applies to all search entrypoints (AI API, human UI, suggestions).

---

## Architecture Summary

- **Data source**: Documentation builds from readthedocs.com (.md files per page)
- **Ingestion**: Scheduled/triggered pipeline packaged as a Canonical charm, hosted with IS
- **Processing**: Semantic chunking with embeddings, metadata extraction, domain-tier tagging
- **Storage**: OpenSearch cluster (deployed from scratch) with full-text + vector search
- **Entrypoints**: (1) AI Agent API returning structured .md content, (2) Human search UI in the Ulwazi Sphinx theme
- **Domain model**: Each product defines tier-1 (included by default) and tier-2 (available on-demand) relationships
- **Observability**: Canonical Observability Stack (COS) tracking search performance
- **Fallback**: ReadTheDocs search if latency exceeds 2 seconds

---

## Phasing

| Phase | Stories | Focus | Dependencies |
|-------|---------|-------|--------------|
| Phase 1 — Foundation | 1, 2, 3 | Domain model, ingestion pipeline, OpenSearch setup | None — run in parallel |
| Phase 2 — Entrypoints & Validation | 4, 5, 6, 10 | AI API, human UI, pilot benchmarking, COS | Depends on Phase 1 |
| Phase 3 — Polish & Production | 7, 8, 9 | Quality tuning, production deployment, suggestions | Depends on Phase 2 |

---

## Root Epic

**Replace documentation.ubuntu.com search with OpenSearch-based semantic search**

---

## Stories

### Story 1: Define Documentation Domain Structure & Key Topics

**Owner**: Michael Park  
**Priority**: P1  
**Phase**: 1

Each product defines tier-1 and tier-2 relationships that control which related documentation is included in search results by default (tier-1) or available on-demand (tier-2).

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 1.1 | Design tier-1/tier-2 relationship schema | Schema documented; supports per-product definition of tier-1 (included by default) and tier-2 (on-demand) related products |
| 1.2 | List top 10 search terms and current performance |
| 1.3 | List top 5 topics that should be findable and are not covered by search terms |

---

### Story 2: Documentation Ingestion Pipeline

**Owner**: Nicholas Bello (Michael Park consulting)
**Priority**: P1  
**Phase**: 1

Build a pipeline to regularly fetch documentation builds from readthedocs.com and feed them into the processing layer.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 2.1 | Design ingestion architecture | Architecture doc covering trigger mechanism, frequency, error handling |
| 2.2 | Implement ReadTheDocs build fetcher | Can fetch latest build artifacts (.md files) for a configured product; handles failures gracefully |
| 2.3 | Implement scheduled/triggered ingestion | Runs on configurable schedule; can be triggered manually; all ingestion runs are logged |
| 2.4 | Package as Canonical charm | Charm builds, deploys, and runs ingestion on a test environment; passes charm review guidelines |

---

### Story 3: Semantic Processing & OpenSearch Indexing

**Owner**: Nicholas Bello (Mehdi Bendriss consulting)  
**Priority**: P1  
**Phase**: 1

Deploy OpenSearch, define how documents are semantically processed, and build the indexing pipeline.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 3.1 | Deploy OpenSearch instance | OpenSearch cluster running; accessible from ingestion pipeline; basic monitoring in place |
| 3.2 | Define semantic processing approach | Decision doc covering embedding model, chunking strategy, metadata extraction; benchmarked on sample docs |
| 3.3 | Design OpenSearch index schema | Index mapping defined; supports full-text + vector search; includes domain/tier metadata fields |
| 3.4 | Implement document processor | Takes doc files (html+md), produces chunks with embeddings and metadata; unit tested |
| 3.5 | Implement indexing pipeline | Processed documents indexed into OpenSearch; supports incremental updates and full re-index |

---

### Story 4: AI Agent Search Entrypoint

**Owner**: Michael Park  
**Priority**: P1  
**Phase**: 2

Provide a search API for AI agents that returns easily parseable, structured data. Agents can use the search to explore related information from any documentation file.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 4.1 | Define API contract | OpenAPI spec published; returns structured, parseable responses; supports domain-scoped queries (tier-1 default, tier-2 on-demand) |
| 4.2 | Implement search API | Endpoint operational; returns relevant .md content for queries; respects domain scoping; responds within 2 seconds |
| 4.3 | Integration testing with sample agent | Demonstrated with at least one AI agent consuming search results successfully |

---

### Story 5: Human Search UI

**Owner**: Nicholas Bello (prototype) → Michael Park + Vladimir (Ulwazi integration)  
**Priority**: P1  
**Phase**: 2

Build a search UI for human users that replaces the current Sphinx search and is integrated into the Ulwazi documentation theme.

| # | Sub-task | Owner | Acceptance Criteria |
|---|----------|-------|-------------------|
| 5.1 | Build search UI prototype | Nicholas Bello | Working prototype demonstrating search results display, basic filtering; stakeholder-approved design |
| 5.2 | Implement ReadTheDocs fallback in UI | Nicholas Bello | If search exceeds 2 seconds, UI gracefully falls back to ReadTheDocs search; user is not blocked |
| 5.3 | Integrate into Ulwazi Sphinx theme | Michael Park + Vladimir | Search UI embedded in documentation.ubuntu.com theme; works across all docs pages; replaces old search |
| 5.4 | Cross-browser/responsive testing | Nicholas Bello | Works on Chrome, Firefox, Safari; responsive on mobile |

---

### Story 7: Pilot & Benchmarking

**Owner**: Michael Park + stakeholder  
**Priority**: P2  
**Phase**: 2

Select pilot products, run a manual benchmarking exercise, and use results to drive quality improvements.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 6.1 | Run manual benchmarking exercise | All test queries executed; results scored against rubric; comparison with current Sphinx search documented |
| 6.2 | Produce benchmarking report | Actionable improvement items identified |

---

### Story 8: Deployment & Operations

**Owner**: Mehdi Bendriss  
**Priority**: P1  
**Phase**: 3

Deploy the full stack to production.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 8.1 | Define deployment architecture | Architecture doc approved; covers networking, security, scaling |
| 8.2 | Deploy OpenSearch to production | Production cluster running; secured; backed up |
| 8.3 | Deploy ingestion charm to production | Charm running in production; ingesting docs on schedule |

---

### Story 9: AI Suggestions

**Owner**: Nicholas Bello  
**Priority**: P3  
**Phase**: 3

Infer suggestions from collected search data and doc-set analysis; integrate into the search UI.

| # | Sub-task | Owner | Acceptance Criteria |
|---|----------|-------|-------------------|
| 9.1 | Design suggestion inference | Nicholas Bello | Architecture defined for deriving suggestions from search query data and doc-set analysis; covers data sources, algorithms, update frequency |
| 9.2 | Implement search-data-based suggestions | Michael Park | Suggestions generated from aggregated search patterns; refreshed regularly; contextually relevant to user queries |
| 9.3 | Implement doc-set analysis for suggestions | Michael Park | Whole doc-set analyzed to surface related content/topics; produces suggestion candidates independent of query history |
| 9.4 | Integrate suggestions into search UI | Nicholas Bello | Suggestions displayed in human search UI; non-blocking; stakeholder-approved UX |
| 9.5 | OpenSearch performance optimization for suggestions | Mehdi Bendriss | Suggestion queries return within 2-second latency budget; OpenSearch tuned (caching, query structure, index optimization) |

---

### Story 10: COS (Canonical Observability Stack) Setup

**Owner**: Mehdi Bendriss  
**Priority**: P2  
**Phase**: 2  
**Effort**: Low (deployment only)

Deploy COS to track search performance from OpenSearch.

| # | Sub-task | Acceptance Criteria |
|---|----------|-------------------|
| 10.1 | Deploy COS and integrate with OpenSearch | COS deployed; OpenSearch metrics (query latency, throughput, error rate) flowing into observability dashboards |
| 10.2 | Configure search performance tracking | Dashboard showing p50/p95/p99 search latency; alerting triggered if p95 exceeds 2-second threshold |

---

## Dependencies

```
Story 1 ─┐
Story 2 ─┼─► Story 4 (AI API)
Story 3 ─┘   Story 5 (Human UI)
             Story 6 (Pilot)    ──► Story 7 (Quality)
             Story 10 (COS)         Story 8 (Production deploy)
                                    Story 9 (Suggestions)
```

- Stories 4, 5, 6, 10 require Phase 1 completion (Stories 1, 2, 3)
- Story 7 requires Story 6 (benchmarking results)
- Stories 8, 9 require Phase 2 completion
- Story 9 depends on Story 5 (search UI exists to integrate into)

---

## Open Questions

- [ ] Which 1–2 pilot products to use for benchmarking? (Story 6.1)
- [ ] Which embedding model to use for semantic processing? (Story 3.2)
- [ ] Exact chunking strategy for documentation pages? (Story 3.2)
- [ ] ReadTheDocs API mechanism for fetching builds — webhook vs. polling? (Story 2.1)

---

## Revision History

| Date | Author | Change |
|------|--------|--------|
| 2025-05-14 | — | Initial project plan |
