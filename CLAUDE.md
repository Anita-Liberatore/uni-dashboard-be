# CLAUDE.md — uni-dashboard-be

Guida per chiunque (o qualsiasi AI) lavori su questo progetto.
Leggi tutto prima di toccare il codice.

---

## Cos'è questo progetto

Backend REST API per [uni-dashboard-fe](../uni-dashboard-fe), una dashboard universitaria
che espone dati accademici di uno studente: profilo, esami, piano di studi, documenti, calendario.

Stack: **Go 1.22** · **net/http** (nessun framework) · **Swagger/OpenAPI** via swaggo

---

## Architettura: Hexagonal (Ports & Adapters)

```
cmd/server/
  main.go               ← wiring: crea adapter, inietta nel service, avvia HTTP

internal/
  domain/               ← NUCLEO. Zero dipendenze esterne.
    student.go          entity Student, value object AcademicRecord
    exam.go             Exam, UpcomingExam
    studyplan.go        YearPlan, Course
    document.go         Document
    calendar.go         CalendarEvent, EventType
    ports.go            interfacce dei repository (outbound ports)

  application/
    service.go          Orchestrazione use case. Dipende SOLO da domain ports.

  adapters/
    http/               ← Driving adapter (lato sinistro)
      handler.go        HTTP → use case → JSON + annotazioni Swagger
      router.go         routing ServeMux + CORS middleware
    memory/             ← Driven adapter (lato destro, sostituibile)
      repository.go     Implementazione in-memory di tutti i port

docs/                   ← GENERATO da swag. Non modificare a mano.
  docs.go
  swagger.json
  swagger.yaml
```

### Regola d'oro delle dipendenze

```
domain        ← non dipende da nulla
application   ← dipende solo da domain (tramite interfacce)
adapters      ← dipende da domain + application
cmd/server    ← dipende da tutto (è il punto di composizione)
```

**Non rompere mai questa gerarchia.**
Se un adapter importa un altro adapter → problema architetturale.
Se application importa net/http → problema architetturale.

---

## API Endpoints

Base path: `/api/v1`

| Method | Path                                        | Handler              |
|--------|---------------------------------------------|----------------------|
| GET    | `/students/{studentId}`                     | GetStudent           |
| GET    | `/students/{studentId}/academic`            | GetAcademic          |
| GET    | `/students/{studentId}/exams`               | GetExamsPassed       |
| GET    | `/students/{studentId}/exams/upcoming`      | GetExamsUpcoming     |
| GET    | `/students/{studentId}/study-plan`          | GetStudyPlan         |
| GET    | `/students/{studentId}/documents`           | GetDocuments         |
| GET    | `/students/{studentId}/calendar`            | GetCalendar          |
| GET    | `/swagger/*`                                | Swagger UI           |

---

## Come aggiungere un nuovo endpoint

1. **Domain** — se serve una nuova entità, aggiungila in `internal/domain/`.
   Se serve un nuovo metodo repository, aggiungilo in `ports.go`.

2. **Application** — aggiungi il metodo in `application/service.go`.
   Il metodo deve accettare `ctx context.Context` e `studentID string` come primi parametri.

3. **Memory adapter** — implementa il nuovo metodo in `adapters/memory/repository.go`.

4. **HTTP adapter** — aggiungi il handler in `handler.go` con annotazioni Swagger complete,
   poi registra la route in `router.go`.

5. **Rigenera i docs Swagger**:
   ```bash
   swag init -g ./cmd/server/main.go --parseDependency --parseInternal -o docs
   ```

6. Aggiungi il relativo `useCase:` commit (vedi sezione Commit).

---

## Sviluppo locale

```bash
# Avvia il server
go run ./cmd/server
# → http://localhost:8080
# → http://localhost:8080/swagger/index.html  (Swagger UI)

# Rigenera docs Swagger dopo aver modificato le annotazioni
swag init -g ./cmd/server/main.go --parseDependency --parseInternal -o docs

# Build
go build ./...

# Test
go test ./...
```

---

## Sostituire l'in-memory adapter con un database reale

1. Crea `internal/adapters/postgres/repository.go`
2. Implementa le stesse interfacce di `internal/domain/ports.go`
3. In `cmd/server/main.go`, sostituisci le istanze `memory.New*()` con `postgres.New*()`

Zero modifiche a `domain/` e `application/`. Questo è il vantaggio dell'architettura esagonale.

---

## Convenzioni di commit

Ogni commit deve avere un prefisso che indica la natura della modifica.

| Prefisso    | Quando usarlo                                                    |
|-------------|------------------------------------------------------------------|
| `useCase:`  | Nuovo endpoint, nuova logica di business, nuovo use case         |
| `fix:`      | Bug fix                                                          |
| `refactor:` | Ristrutturazione del codice senza cambi di comportamento         |
| `docs:`     | Solo documentazione (README, CLAUDE.md, commenti)                |
| `chore:`    | Dipendenze, tooling, CI, configurazione                          |
| `test:`     | Aggiunta o modifica di test                                      |

### Formato del messaggio

```
<prefisso>: <descrizione breve in inglese, imperativo, max 72 char>

<corpo opzionale: cosa hai fatto e perché, non come>
```

### Esempi

```
useCase: add student profile retrieval endpoint
useCase: add exam history filtered by academic year
useCase: add document upload via multipart/form-data
fix: correct CORS headers for production origin
refactor: extract studentID validation into middleware
docs: add CLAUDE.md with project guidelines
chore: upgrade swaggo/swag to v1.16
test: add unit tests for StudentService use cases
```

### Regole

- **Un commit = una cosa sola.** Non mescolare fix e feature.
- Il corpo del commit spiega il *perché*, non il *cosa* (il diff mostra già il cosa).
- Dopo ogni modifica alle annotazioni Swagger → rigenera i docs → includi `docs/` nel commit.
- Non committare `go.sum` senza aver eseguito `go mod tidy` prima.

---

## Linee guida del codice

- **Nessun framework HTTP.** Si usa `net/http` standard (Go 1.22+) con method-qualified patterns.
- **Errori come valori.** Niente panic fuori da `main`. Gestisci sempre l'errore o propagalo esplicitamente.
- **Nessuna logica nei handler.** Il handler estrae i parametri HTTP, chiama il service, serializza la risposta. Fine.
- **Nessun ORM.** Query SQL a mano quando si aggiunge il DB adapter.
- **Commenti in inglese.** Testo visibile all'utente (es. messaggi di errore JSON) in italiano.
- **`context.Context` sempre come primo parametro** nei metodi di domain, application e adapter.
