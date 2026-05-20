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

Base path: `/api/v1` · Swagger UI: `http://localhost:8080/swagger/index.html`

| Method | Path                                       | Description                        |
|--------|--------------------------------------------|------------------------------------|
| GET    | `/students/{studentId}`                    | Student profile                    |
| GET    | `/students/{studentId}/academic`           | Academic record (GPA, credits, …)  |
| GET    | `/students/{studentId}/exams`              | Passed exams                       |
| GET    | `/students/{studentId}/exams/upcoming`     | Upcoming exams                     |
| GET    | `/students/{studentId}/study-plan`         | Full study plan grouped by year    |
| GET    | `/students/{studentId}/documents`          | Uploaded documents                 |
| GET    | `/students/{studentId}/calendar`           | Calendar events (exams, deadlines) |
| GET    | `/swagger/*`                               | Swagger UI                         |

All responses are `application/json`. Errors follow `{"error": "<message>"}`.

> **Example:** `GET /api/v1/students/S1234567`

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
