# 1️⃣ **What is ServeMux?**

`ServeMux` is Go’s HTTP **request multiplexer**:

* It matches incoming request URLs to **registered patterns**.
* It calls the handler of the **most specific matching pattern**.

Think of it as a **router** that maps URL paths to handler functions.

---

## 2️⃣ **Pattern Syntax**

A pattern can match **method, host, and path**:

```txt
[METHOD ][HOST]/[PATH]
```

* METHOD → optional, e.g., GET, POST.
* HOST → optional, e.g., example.com.
* PATH → required (can contain wildcards).

### Examples

| Pattern                         | Matches                                                                               |
| ------------------------------- | ------------------------------------------------------------------------------------- |
| `/index.html`                   | Any method, any host, exact path `/index.html`                                        |
| `GET /static/`                  | GET or HEAD requests starting with `/static/`                                         |
| `example.com/`                  | Any request to host `example.com`                                                     |
| `/b/{bucket}/o/{objectname...}` | URL like `/b/mybucket/o/file/path.txt`, bucket=`mybucket`, objectname=`file/path.txt` |

---

## 3️⃣ **Wildcards**

ServeMux supports **named wildcards**:

1. `{name}` → matches a single path segment
2. `{name...}` → matches **all remaining path segments**

### Example

```text
Pattern: /b/{bucket}/o/{object...}
Request: /b/photos/o/2025/nov/pic.jpg
Match:
    bucket = "photos"
    object = "2025/nov/pic.jpg"
```

**Rules:**

* Wildcards must be full segments: preceded by `/` and followed by `/` or end of string.
* `{name...}` can only appear at the **end**.

Special wildcard: `{$}` → matches **only the end** of the path.

* `/{$}` → matches `/`
* `/` → matches everything

---

## 4️⃣ **Literal Matching**

* Case-sensitive

* Method-sensitive:

  * Pattern with no method → matches **all methods**
  * Pattern with `GET` → matches GET and HEAD only
  * Otherwise → exact method match required

* Host-sensitive:

  * No host → matches any host
  * With host → only matches that host

---

## 5️⃣ **Path Handling / Escaping**

* URL segments are **unescaped** segment by segment:

```text
URL: /a%2Fb/100%25
Segments: ["a/b", "100%"]
```

* Pattern `/a%2Fb/` matches `/a%2Fb/100%25`

* Pattern `/a/b/` does **not** match (`%2F` is treated literally in the URL)

* Redundant slashes and `.`/`..` segments are sanitized and redirected.

---

## 6️⃣ **Pattern Precedence**

If multiple patterns match, **most specific pattern wins**:

* `/images/thumbnails/` is more specific than `/images/`
* `GET /` vs `/index.html` → conflicts; `GET /` matches more requests

**Host rule:**

* If one pattern has a host and the other doesn’t, the host-specific pattern wins.

---

## 7️⃣ **Trailing Slash Behavior**

* If a pattern ends with `/` or `...` and a request comes **without the slash**, ServeMux **redirects** to the URL with the slash.

```go
http.HandleFunc("/images/", handler)
```

* Request: `/images` → automatically redirected to `/images/`

* Can override by registering `/images` separately.

---

## 8️⃣ **Examples of Full Matching**

| Pattern                     | Request                  | Match?                     | Notes                     |
| --------------------------- | ------------------------ | -------------------------- | ------------------------- |
| `/`                         | `/`                      | ✅                          | Matches everything        |
| `/images/`                  | `/images`                | ✅ → redirect to `/images/` | Trailing slash            |
| `/b/{bucket}/o/{object...}` | `/b/photos/o/img.png`    | ✅                          | Wildcards captured        |
| `/GET /static/`             | `GET /static/style.css`  | ✅                          | Matches GET and HEAD only |
| `/GET /static/`             | `POST /static/style.css` | ❌                          | Method does not match     |

---

## 9️⃣ **Version Changes (Go 1.22)**

* Wildcards now match **single segments** instead of literal path (`/{x}` matches any one-segment path).
* Invalid patterns now **panic** (`/{` or `/a{x}`)
* Each segment is **unescaped individually** (affects `%2F`)
* Trailing slash and redirects behavior is standardized

> Previous versions (1.21) treated wildcards as literal segments and unescaped the **entire path**, causing different behavior.

---

### 🔑 **Takeaways**

1. ServeMux = lightweight router
2. Patterns can match **method, host, path**
3. Wildcards `{name}` and `{name...}` capture path segments
4. Most specific pattern wins
5. Trailing slashes are auto-redirected
6. Go 1.22 improved wildcard matching and request sanitization

---
