# 📦 Go Bulk Insert

🚀 Golang দিয়ে দ্রুত এবং efficient ভাবে **Bulk Insert (multiple data insert)** করার একটি simple utility।

---

## ✨ Features

- ⚡ Fast bulk insert
- 🧠 Simple & clean code structure
- 🔄 Multiple row একসাথে insert
- 💾 Database performance optimized

---

## 🛠️ Requirements

- Go (Golang) 1.18+
- SQL Database (MySQL / PostgreSQL / MSSQL)

---

## 📂 Project Structure

```
go_bulk_insert/
│── main.go
│── bulk_insert.go
│── db.go
│── README.md
```

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/prothesbarai/go_bulk_insert.git
cd go_bulk_insert
```

---

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 3️⃣ Database Config Setup

`db.go` ফাইলে তোমার database credentials বসাও:

```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
```

---

## 📌 Usage Example

```go
rows := [][]interface{}{
    {"John", 25},
    {"Alice", 30},
    {"Bob", 28},
}

err := BulkInsert(db, "users", []string{"name", "age"}, rows)
if err != nil {
    log.Fatal(err)
}
```

---

## ⚙️ How It Works

- একাধিক row collect করা হয়  
- SQL query dynamically generate করা হয়  
- একবারে সব data insert করা হয়  

---

## 📊 Why Bulk Insert?

| Normal Insert | Bulk Insert |
|--------------|------------|
| প্রতি row আলাদা query | একবারে সব insert |
| Slow | Fast ⚡ |
| বেশি DB load | কম DB load |

---

## ❗ Important Notes

- বড় data insert করলে batch ব্যবহার করো  
- SQL injection avoid করতে parameterized query ব্যবহার করো  
- Transaction ব্যবহার করলে আরো safe হবে  

---

## 🧑‍💻 Author

**Prothes Barai**  
Software Engineer 🚀  

---

## 📄 License

MIT License
