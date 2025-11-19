# Idiomatic, production-ready Go patterns

Go has its own philosophy: **simplicity, composition, explicitness, small packages, and interface-driven design**. Let’s go layer by layer and highlight **opinionated Go practices**.

---

## 🟦 1. Router Layer

**Idiomatic Go Tips:**

* Use **`http.ServeMux`** for small apps, or **third-party routers** like [Chi](https://github.com/go-chi/chi) for medium/large apps.
* Keep routes **declarative and simple**:

  ```go
  r := chi.NewRouter()
  r.Get("/users/{id}", handlers.GetUser)
  r.Post("/users", handlers.CreateUser)
  ```

* Use **sub-routers** to group related routes:

  ```go
  r.Route("/users", func(r chi.Router) {
      r.Get("/", handlers.ListUsers)
      r.Post("/", handlers.CreateUser)
  })
  ```

* Avoid “fat” routers that include logic — routing is only about **path → handler mapping**.

---

## 🟩 2. Middleware Layer

**Go idioms:**

* Middleware is a **function that wraps `http.Handler`**.

  ```go
  func Logging(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          log.Printf("%s %s", r.Method, r.URL.Path)
          next.ServeHTTP(w, r)
      })
  }
  ```

* Compose multiple middlewares using **nested wrapping** or libraries like Chi:

  ```go
  r.With(middleware1, middleware2).Get("/path", handler)
  ```

* Keep middleware **stateless**: no business logic or DB calls.
* Panic recovery + structured logging middleware is standard.

---

## 🟧 3. Controller / Handler Layer

**Go idioms:**

* Keep handlers **thin**:

  ```go
  func CreateUser(svc *services.UserService) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          var input CreateUserInput
          if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
              http.Error(w, "invalid input", http.StatusBadRequest)
              return
          }

          user, err := svc.Create(r.Context(), input)
          if err != nil {
              http.Error(w, err.Error(), http.StatusInternalServerError)
              return
          }

          w.WriteHeader(http.StatusCreated)
          json.NewEncoder(w).Encode(user)
      }
  }
  ```

* Use **`http.HandlerFunc` with closures** to inject dependencies (services, repos).
* Always pass **`context.Context`** from HTTP request down to service and DB layers.

---

## 🟨 4. Service Layer (Business Logic)

**Go idioms:**

* Implement **pure Go functions**:

  ```go
  type UserService struct {
      repo UserRepository
  }

  func (s *UserService) Create(ctx context.Context, input CreateUserInput) (*User, error) {
      hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
      if err != nil {
          return nil, fmt.Errorf("hash password: %w", err)
      }

      return s.repo.Create(ctx, User{
          Email: input.Email,
          Password: string(hashedPassword),
      })
  }
  ```

* **Always propagate `context.Context`** for deadlines, cancellations, logging, tracing.
* Return **typed errors**, don’t panic.
* No HTTP or DB knowledge here.

---

## 🟫 5. Repository / Data Access Layer

**Go idioms:**

* Define **interfaces** for repositories:

  ```go
  type UserRepository interface {
      Create(ctx context.Context, user User) (*User, error)
      GetByEmail(ctx context.Context, email string) (*User, error)
  }
  ```

* Use **structs implementing interface**:

  ```go
  type sqliteUserRepo struct {
      db *sql.DB
  }
  ```

* Keep SQL queries **in the repository only**.
* Use **prepared statements** or `db.QueryContext`/`db.ExecContext`.
* Propagate `context.Context`.

---

## 🟥 6. Storage Layer (SQLite / DB)

**Go idioms:**

* Use **`database/sql`** or lightweight ORM if needed (GORM is popular, but for idiomatic Go `sqlx` is fine).
* Keep **one global DB connection pool**.
* Do **transactions** in repository or service layer if needed.
* Errors: map `sql.ErrNoRows` → domain error `ErrNotFound`.
* Use `defer rows.Close()` properly.

---

## 🟦 Cross-Cutting Go Practices

1. **Context everywhere**
   Every function interacting with request/DB/service takes `ctx context.Context`.

2. **Dependency injection by composition**

   ```go
   svc := services.NewUserService(repo)
   h := handlers.NewUserHandler(svc)
   ```

3. **Small packages**

   * `/internal/handlers`
   * `/internal/services`
   * `/internal/repositories`
   * `/internal/models`

4. **Error handling over exceptions**
   Go doesn’t use exceptions. Always check errors:

   ```go
   if err != nil {
       return nil, fmt.Errorf("failed to create user: %w", err)
   }
   ```

5. **Testing**

   * Handlers: `httptest.NewRecorder`
   * Services: mock repositories
   * Repositories: use in-memory SQLite for testing

6. **Struct embedding & composition**
   Use **embedding for shared fields**:

   ```go
   type BaseModel struct {
       ID        int64
       CreatedAt time.Time
   }

   type User struct {
       BaseModel
       Email string
       ...
   }
   ```

7. **No global mutable state**
   Use **dependency injection** instead.

8. **Explicit, simple interfaces**
   Avoid “God interfaces”; small, focused interfaces are better.

---

## 🔥 Opinionated Go Folder Layout

```txt
project/
│
├─ cmd/api/main.go           # server entry point
├─ internal/
│   ├─ config/              # config structs + loader
│   ├─ handlers/            # HTTP handlers (controllers)
│   ├─ services/            # business logic
│   ├─ repositories/        # DB access
│   ├─ models/              # data structs / domain models
│   ├─ middleware/          # logging, auth, recovery
│   └─ db/                  # DB connection, migrations
└─ go.mod
```

---

