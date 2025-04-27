# Apica Fullstack Assignment - Search Engine

This project implements a **fullstack search engine** for logs/events using:

- **Backend:** Go (Golang) with Gorilla Mux
- **Frontend:** React + Vite + TypeScript + TailwindCSS
- **Parquet File Upload:** Using GFileMux library (similar to Multer)

---

## 🛠️ Features

- Upload `.parquet` files to seed data dynamically
- Search across multiple fields like `Message`, `EventId`, `Namespace`, etc.
- Fullscreen, responsive layout
- Table with truncated text and tooltips for long data
- Upload new Parquet files without restarting the server
- Instant indexing and searching
- Professional UI using TailwindCSS
- Clean error handling, loading spinners, and UX enhancements

---

## 📦 Tech Stack

| Part | Technology |
|:-----|:-----------|
| Backend | Golang + Gorilla Mux + GFileMux |
| Frontend | React + Vite + TypeScript + TailwindCSS |
| File Handling | Parquet parsing + dynamic seeding |
| Communication | REST API (JSON) |

---

## ⚙️ How to Run the Project

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/apica-search-engine.git
cd apica-search-engine
```

### 2. Start Backend (Go Server)
```bash
cd backend
go mod tidy
go run main.go
```

## ✅ Server will run on http://localhost:8080

### 3. Start Frontend (React + Vite)
```bash
cd frontend
npm install
npm run dev
```

## ✅ Frontend will run on http://localhost:5173
