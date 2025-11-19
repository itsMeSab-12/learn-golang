# Modelling the Repository Layer

The **correct way to design repositories** is **bottom-up but guided by top-level functional requirements**. Let me break it down step by step, and then we’ll mock a **Todo List app** example.

---

## **1️⃣ Steps to Model Repositories**

### **Step 1: Define functional requirements**

* List all actions users can do with your app
* Example (Todo List):

  1. Create a todo
  2. Read a todo by ID
  3. List todos (optionally filter by status or user)
  4. Update a todo (e.g., mark as completed, update title/description)
  5. Delete a todo
  6. (Optional) Assign todos to users, set due dates, priorities

> This is **business-driven** — think of user stories, not DB tables.

---

### **Step 2: Define service layer responsibilities**

* Map functional requirements to **business logic**
* Example (Todo List Service):

  * `CreateTodo`: Validate title, optionally validate user exists
  * `GetTodoByID`: Check access rights
  * `ListTodos`: Optionally filter by user/status, sort
  * `UpdateTodo`: Validate allowed updates
  * `DeleteTodo`: Ensure only owner can delete

> **Service layer doesn’t know how data is stored**, only what rules to enforce.

---

### **Step 3: Derive repository interface from service needs**

* Ask: *What operations does the service layer need to fulfill requirements?*
* Don’t design repository around DB — design around **service use cases**
* Example:

```go
type TodoRepository interface {
    Create(ctx context.Context, t Todo) (Todo, error)
    GetByID(ctx context.Context, id int) (Todo, error)
    ListByUser(ctx context.Context, userID int) ([]Todo, error)
    Update(ctx context.Context, t Todo) (Todo, error)
    Delete(ctx context.Context, id int) error
}
```

> Service layer calls these methods — the repository just does “storage work.”

---

### **Step 4: Mock repository implementations**

* Start with **in-memory storage**:

```go
type InMemoryTodoRepo struct {
    data   map[int]Todo
    nextID int
}
```

* Later replace with:

  * File-based repo
  * SQLite repo
  * Cloud DB repo

> Services and handlers **never change**.

---

## **2️⃣ Example: Mocking Todo List App Design**

### **2.1 Functional Requirements (Mock)**

1. Users can **create todos** with title and optional description
2. Users can **mark todos as completed**
3. Users can **update title/description**
4. Users can **delete todos**
5. Users can **list all todos**, optionally filtered by status

---

### **2.2 Service Layer Responsibilities (Mock)**

| Function                                                   | Responsibility                                                  |
| ---------------------------------------------------------- | --------------------------------------------------------------- |
| `CreateTodo(userID, title, description)`                   | Validate title, assign unique ID, store todo                    |
| `GetTodoByID(userID, todoID)`                              | Fetch todo, check that `userID` owns it                         |
| `ListTodos(userID, filterStatus)`                          | List todos for the user, optionally filter by completed/pending |
| `UpdateTodo(userID, todoID, newTitle, newDesc, completed)` | Validate ownership, apply updates                               |
| `DeleteTodo(userID, todoID)`                               | Ensure only owner can delete                                    |

> Notice these are **pure business rules**, not SQL or JSON concerns.

---

### **2.3 Repository Interface Derived from Service Layer**

```go
type TodoRepository interface {
    Create(ctx context.Context, t Todo) (Todo, error)
    GetByID(ctx context.Context, id int) (Todo, error)
    ListByUser(ctx context.Context, userID int) ([]Todo, error)
    Update(ctx context.Context, t Todo) (Todo, error)
    Delete(ctx context.Context, id int) error
}
```

> The **repository interface directly reflects what the service layer needs**.

* Services don’t care whether it’s in-memory, file, SQLite, or DynamoDB.
* Handlers don’t care either — just call service functions.

---

### **2.4 Models (Mock)**

```go
type Todo struct {
    ID          int
    UserID      int
    Title       string
    Description string
    Completed   bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

---

### **2.5 Flow Overview**

```txt
HTTP Request (Handler) 
     ↓
Service Layer (business rules)
     ↓
TodoRepository Interface (abstract)
     ↓
Concrete Repo Implementation (InMemory / File / SQLite / DynamoDB)
```

---

## ✅ **Key Takeaways**

1. **Start with functional requirements** → what the user wants to do
2. **Define service layer responsibilities** → enforce rules, orchestrate data
3. **Derive repository interfaces** → only expose what services need
4. **Implement repositories** → swap storage layers without touching service/handlers
5. **Test incrementally**:

   * Services with in-memory repo
   * Handlers with mocked service

---
