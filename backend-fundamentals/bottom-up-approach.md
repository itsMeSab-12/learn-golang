
# Bottom Up Approach

In **Go backend development**, especially when building something **idiomatic and testable**, the recommended approach is usually **bottom-up**, i.e., start from the **storage/data layer → services → handlers → router → server**.

Here’s why and how:

---

## 1. Why Bottom-Up (Storage → Router)

| Layer                            | Reason to start here                                                                                                                            |
| -------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| **Storage (DB / files / cache)** | You define schemas, migrations, and connection logic first. Everything else depends on data.                                                    |
| **Repository / DAO**             | Repositories are just thin wrappers around storage. Once implemented, you can **unit test them** independently.                                 |
| **Service (Business Logic)**     | Services depend on repositories. By building them next, you can inject repositories and **write pure logic tests** without worrying about HTTP. |
| **Handler / Controller**         | Once services are ready, handlers become simple adapters from HTTP → service → JSON.                                                            |
| **Router / Middleware / Server** | Router is just “plumbing” that connects HTTP paths to handlers. No logic goes here, so it’s easier to plug in once everything else works.       |

✅ **Benefits:**

* Each layer can be **unit tested in isolation**.
* Avoids writing HTTP handlers that fail because DB logic isn’t ready.
* You naturally apply **dependency injection**: services depend on repositories, handlers depend on services, router depends on handlers.

---

## 2. Recommended Implementation Order

1. **Storage Layer**

   * Setup SQLite connection
   * Define schema & migrations
   * Write a simple DB ping test

2. **Repository Layer**

   * Define interfaces (`UserRepository`, `SnippetRepository`)
   * Implement CRUD operations with DB
   * Write repository unit tests (mock DB optional)

3. **Service Layer**

   * Implement business rules (`CreateUser`, `GetSnippet`, `DeleteSnippet`)
   * Inject repository interfaces
   * Write service unit tests (mock repositories)

4. **Handler / Controller Layer**

   * Implement HTTP handlers for each route
   * Inject service layer
   * Map JSON input → service input → JSON output
   * Write handler tests using `httptest`

5. **Middleware Layer**

   * Logging, panic recovery, auth, rate-limiting
   * Wrap handlers with middleware

6. **Router / Server Layer**

   * Map routes to handlers
   * Compose middleware
   * Start HTTP server

---

## 3. Extra Tips for Go

* **Dependency injection** flows naturally bottom-up:

```go
db := db.NewSQLite("app.db")
repo := repositories.NewUserRepo(db)
svc := services.NewUserService(repo)
handler := handlers.NewUserHandler(svc)
```

* **Top-down (Router → Storage)** is tempting because you see HTTP quickly, but it’s **hard to test**, often leads to “stubbed DB logic,” and can produce brittle code.

* **Bottom-up + Test-Driven**: write repository/service first, then handlers, then router. You’ll have **solid, testable code** and minimal runtime surprises.

---

💡 **Opinionated takeaway:**

> Always **start from storage → repository → service → handler → router** in Go. Build and test incrementally.

---
