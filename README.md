# Go-Server

A lightweight, performant hybrid web server built with **Go** and the **Gin Gonic** web framework. This project demonstrates how to blend the high performance of Gin's API routing with the flexibility of Go's standard library `http.FileServer` via a smart `NoRoute` configuration.

---

## 🚀 Features

* **Fast API Architecture:** Powered by Gin for structured, concurrent route handling.
* **Zero-Conflict Static File Serving:** Uses native `http.FileServer` dynamically to serve HTML, CSS, and JS assets directly from the root domain.
* **Controller Closures:** Implements functional dependency injection patterns inside handlers.
* **Environment Configuration:** Secure local environment management via `godotenv`.

---

## 🛠️ Project Structure

```text
Go-Server/
├── controllers/
│   └── controller.go     # Route controllers/handlers (closures)
├── routes/
│   └── Staticroutes.go   # Route registration logic
├── static/
│   ├── form.html         # Frontend HTML Form
│   └── index.html        # Main Landing Page
├── .env.example          # Example environment configuration
├── go.mod                # Module rules & dependencies
├── go.sum                # Security/version locks for dependencies
└── main.go               # Server initialization & configuration entry point


🧩 Architectural Highlight: Hybrid Routing
To avoid standard Gin routing engine panics (wildcard conflicts with the root directory), this repository applies an explicit dependency hierarchy rule:

API Endpoints first: Explicit paths like /hello and /form are parsed beforehand.

NoRoute Fallback: Unregistered routes gracefully drop down to gin.WrapH(http.FileServer). This automatically serves direct physical assets out of the /static folder on demand without requiring verbose structural path mapping.
