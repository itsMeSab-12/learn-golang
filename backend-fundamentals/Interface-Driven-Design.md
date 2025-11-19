# **1️⃣ What is Interface-Driven Design in Go?**

**Interface-Driven Design (IDD)** means:

> You define **what your code should do** (the behavior) first using interfaces, **not how it will do it**.

Then you write implementations that satisfy those interfaces.

In Go, this is **idiomatic** because:

* Interfaces are **implicit** — a type satisfies an interface just by implementing its methods.
* You can swap implementations without changing the code that depends on the interface.
* It allows **testing with mocks** easily.

---

## **2️⃣ Why use it for repositories?**

A repository is **an abstraction over your storage layer**.

### Without interfaces

```go
type UserRepo struct {
    db *sql.DB
}

func (r *UserRepo) CreateUser(u User) { ... }
```

* Every service directly depends on **one concrete implementation**
* Hard to swap storage layer
* Hard to test without the actual DB

### With interfaces

```go
type UserRepository interface {
    Create(ctx context.Context, u User) (User, error)
    GetByID(ctx context.Context, id int) (User, error)
}
```

* Services depend on **UserRepository**, not the concrete type
* You can implement:

  * `InMemoryUserRepo`
  * `FileUserRepo`
  * `SQLiteUserRepo`
  * `DynamoUserRepo`
* Swapping storage is as simple as changing the injected implementation

---

## **3️⃣ How to design your repository interfaces**

1. **Think in terms of **behavior**, not storage**

   * Don’t expose “SQL queries” in the interface
   * Ask: *What does the service layer need to do?*
     Example:

   ```go
   type SnippetRepository interface {
       Create(ctx context.Context, s Snippet) (Snippet, error)
       GetByID(ctx context.Context, id int) (Snippet, error)
       ListByUser(ctx context.Context, userID int) ([]Snippet, error)
       Delete(ctx context.Context, id int) error
   }
   ```

2. **Use `context.Context` everywhere**

   * Supports cancellation, deadlines, and tracing

3. **Return typed errors**

   ```go
   var ErrNotFound = errors.New("not found")
   ```

4. **Keep interfaces small and focused**

   * One interface per entity (e.g., UserRepository, SnippetRepository)
   * Don’t mix unrelated methods

---

## **4️⃣ Implementations (concrete types)**

## In-memory repo (fast iteration)

```go
type InMemorySnippetRepo struct {
    data   map[int]Snippet
    nextID int
}

func NewInMemorySnippetRepo() *InMemorySnippetRepo {
    return &InMemorySnippetRepo{data: make(map[int]Snippet), nextID: 1}
}

func (r *InMemorySnippetRepo) Create(ctx context.Context, s Snippet) (Snippet, error) {
    s.ID = r.nextID
    r.nextID++
    r.data[s.ID] = s
    return s, nil
}

func (r *InMemorySnippetRepo) GetByID(ctx context.Context, id int) (Snippet, error) {
    s, ok := r.data[id]
    if !ok {
        return Snippet{}, ErrNotFound
    }
    return s, nil
}
```

### File-based repo

```go

type FileSnippetRepo struct {
    path string
    data map[int]Snippet
    nextID int
}

// Methods implement same interface
```

### SQLite repo

```go

type SQLiteSnippetRepo struct {
    db *sql.DB
}

// Methods implement same interface using SQL
```

### DynamoDB repo

```go

type DynamoSnippetRepo struct {
    client *dynamodb.Client
}

// Methods implement same interface using AWS SDK
```

**The service layer doesn’t care which repo is used.**

---

## **5️⃣ How to wire it in Go (Dependency Injection)**

```go
var snippetRepo SnippetRepository

// Dev mode: in-memory
snippetRepo = NewInMemorySnippetRepo()

// Dev mode: file
snippetRepo = NewFileSnippetRepo("data.json")

// Dev mode: SQLite
db, _ := db.NewSQLite("dev.db")
snippetRepo = NewSQLiteSnippetRepo(db)

service := NewSnippetService(snippetRepo)
handler := NewSnippetHandler(service)

```

✅ **No code changes in handlers/services** — only swap the repository implementation.

---

## **6️⃣ Extra Go Best Practices for Interfaces**

1. **Define interfaces at the consumer**

   * In Go, the interface should live where it is **used**, not where it is implemented.
   * Example: `SnippetService` depends on `SnippetRepository` — the interface should live in `services` package, not `repositories`.

2. **Keep interfaces small**

   * Single responsibility → easier to mock in tests

3. **Use `ctx context.Context`**

   * Standard practice in Go for all DB/IO/network operations

4. **Error types**

   * Define sentinel errors (`ErrNotFound`) in interface package or domain models

5. **Testing benefit**

   * In-memory repo or mock repo → full test coverage **without DB**

---

## **7️⃣ Summary / Mental Model**

* **Interface = contract / behavior**
* **Implementation = concrete storage**
* **Service + Handler** = depend only on interface
* Swap implementation anytime

```txt
Service -> SnippetRepository (interface) -> InMemory / File / SQLite / DynamoDB
```

* Start simple (InMemory) → replace storage with minimal friction
* Testable, maintainable, idiomatic Go

---

## Use case of IDD

Starting with **in-memory storage** allows you to focus on building **handlers, services, and repositories** without worrying about DB setup, migrations, or I/O errors. Then you can progressively replace the storage layer with **file-based**, **SQLite**, and eventually **cloud databases**.

Here’s how to do it **idiomatically in Go**, following your bottom-up approach:

---

## **1️⃣ Temporary In-Memory Storage**

### ✅ Goals

* Implement repository interfaces
* Use Go maps/slices to store data temporarily
* Enable full service/handler development without a DB

### Example

```go
// internal/models/user.go
package models

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
}

// internal/repositories/user_repo.go
package repositories

import (
    "context"
    "errors"
    "gosnippet/internal/models"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
    Create(ctx context.Context, u models.User) (models.User, error)
    GetByID(ctx context.Context, id int) (models.User, error)
}

// in-memory implementation
type InMemoryUserRepo struct {
    data   map[int]models.User
    nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{
        data:   make(map[int]models.User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepo) Create(ctx context.Context, u models.User) (models.User, error) {
    u.ID = r.nextID
    r.nextID++
    r.data[u.ID] = u
    return u, nil
}

func (r *InMemoryUserRepo) GetByID(ctx context.Context, id int) (models.User, error) {
    user, ok := r.data[id]
    if !ok {
        return models.User{}, ErrUserNotFound
    }
    return user, nil
}
```

### ✅ Advantages

* **Zero external dependencies**
* Fast feedback
* Perfect for **unit testing** services/handlers

---

## **2️⃣ File-Based Storage**

### ✅ Goal

* Replace the in-memory map with **persistent JSON/CSV/YAML files**
* Keep repository interface unchanged → services/handlers don’t need to know

### Example2

```go
type FileUserRepo struct {
    path string
    data map[int]models.User
    nextID int
}

func NewFileUserRepo(path string) *FileUserRepo {
    return &FileUserRepo{
        path: path,
        data: make(map[int]models.User),
        nextID: 1,
    }
}

// On Create / Update: serialize `data` map to JSON file
// On GetByID: read map from file (or keep in-memory cache + flush)
```

### ✅ Advantages

* You get **persistence across restarts**
* Still simple to implement for dev
* Can gradually add **locking / concurrency handling**

---

## **3️⃣ SQLite Database (Dev)**

### ✅ Goal2

* Introduce **SQL queries and DB connections**
* Implement the same repository interface → services remain unchanged

### Steps

1. Add SQLite driver:

```bash
go get modernc.org/sqlite
```

2. Create DB connection:

```go
package db

import (
    "database/sql"
    _ "modernc.org/sqlite"
)

func NewSQLite(path string) (*sql.DB, error) {
    return sql.Open("sqlite", path)
}
```

3. Implement repository with SQL:

```go
type SQLiteUserRepo struct {
    db *sql.DB
}

func NewSQLiteUserRepo(db *sql.DB) *SQLiteUserRepo {
    return &SQLiteUserRepo{db: db}
}

func (r *SQLiteUserRepo) Create(ctx context.Context, u models.User) (models.User, error) {
    res, err := r.db.ExecContext(ctx, "INSERT INTO users (name,email,password) VALUES (?,?,?)", u.Name, u.Email, u.Password)
    if err != nil { return models.User{}, err }
    id, _ := res.LastInsertId()
    u.ID = int(id)
    return u, nil
}
```

4. Services/handlers **don’t change** — only the repository implementation changes.

---

## **4️⃣ Cloud Databases (SQL / DynamoDB)**

### ✅ Goal3

* Replace SQLite repo with cloud-specific repositories:

  * SQL: MySQL/Postgres → similar queries
  * DynamoDB: implement repository interface using AWS SDK

### ✅ Advantage

* No code changes in **handlers or services**
* Only dependency injected repository changes

---

## **5️⃣ Idiomatic Go Practices for this Approach**

1. **Repository Interface First:**
   Define your repository interface **once**, then implement multiple backends.

2. **Dependency Injection:**

   * Inject `repo` into service, `service` into handler
   * Swap storage layers easily

```go
var repo UserRepository

// Dev mode: In-memory
repo = repositories.NewInMemoryUserRepo()

// Dev+Persistence: File
repo = repositories.NewFileUserRepo("data.json")

// Dev+SQL: SQLite
db, _ := db.NewSQLite("dev.db")
repo = repositories.NewSQLiteUserRepo(db)
```

3. **Unit-Testability:**

   * Services and handlers can be tested **without DB**
   * Just use in-memory repo

4. **Progressive Enhancement:**

   * Start small (fast iteration) → replace storage layer progressively
   * Keep interface consistent → no refactor required in higher layers

---

💡 **Opinionated Pattern Summary**

* **Interface-driven design** for repositories
* **Start in-memory → file → local DB → cloud DB**
* **Service & handler code remains unchanged**
* **Inject storage layer via DI**
* **Context everywhere**: all methods take `ctx context.Context`
* **Typed errors**: return `ErrNotFound`, etc.

---
