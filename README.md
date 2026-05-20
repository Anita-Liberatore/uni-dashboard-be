# uni-dashboard-be

REST API backend for [uni-dashboard-fe](../uni-dashboard-fe) — a university student
dashboard that exposes profile, academic records, exams, study plan, documents and
calendar events for a single authenticated student.

---

## Architecture

The project follows **Hexagonal Architecture** (Ports & Adapters) with lightweight
**DDD tactical patterns**:

```
cmd/server/           ← entry point: wires all layers together
internal/
  domain/             ← core: entities, value objects, port interfaces
  application/        ← use cases: orchestrates domain ports
  adapters/
    http/             ← driving adapter: HTTP handlers + router
    memory/           ← driven adapter: in-memory repository (swap with DB)
```

| Layer       | Depends on          | Never depends on             |
|-------------|---------------------|------------------------------|
| domain      | nothing             | application, adapters        |
| application | domain (ports only) | adapters, net/http           |
| adapters    | domain + application| each other                   |

Replacing the in-memory store with PostgreSQL means adding a new driven adapter
and changing one line in `main.go`. Nothing else changes.

---

## API Endpoints

Base path: `/api/v1`

| Method | Path                    | Description                        |
|--------|-------------------------|------------------------------------|
| GET    | `/api/v1/me`            | Student profile                    |
| GET    | `/api/v1/me/academic`   | Academic record (GPA, credits, …)  |
| GET    | `/api/v1/me/exams`      | Passed exams                       |
| GET    | `/api/v1/me/exams/upcoming` | Upcoming exams                 |
| GET    | `/api/v1/me/study-plan` | Full study plan grouped by year    |
| GET    | `/api/v1/me/documents`  | Uploaded documents                 |
| GET    | `/api/v1/me/calendar`   | Calendar events (exams, deadlines) |

All responses are `application/json`. Errors follow `{"error": "<message>"}`.

---

## Running locally

```bash
go run ./cmd/server
# Server starts on :8080
```

The frontend dev server proxies `/api` to `:8080` — no CORS issues in development.

---

## Connecting the frontend

In `uni-dashboard-fe`, open `src/app/services/student.service.ts` and replace each
`of(MOCK_*)` call with the corresponding `http.get<T>()`:

```typescript
// before
getProfile(): Observable<Student> {
  return of(MOCK_STUDENT);
}

// after
getProfile(): Observable<Student> {
  return this.http.get<Student>(`${this.baseUrl}/me`);
}
```

Then uncomment the `HttpClient` injection at the top of the service.

---

## Roadmap

- [ ] PostgreSQL driven adapter  
- [ ] JWT authentication middleware  
- [ ] `POST /api/v1/me/documents` — document upload  
- [ ] `GET /api/v1/me/exams/{id}` — single exam detail
