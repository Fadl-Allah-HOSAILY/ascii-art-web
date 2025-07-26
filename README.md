# 🎨 ASCII Art Web Application

> A simple yet powerful ASCII art generator built with **Go (Golang)**.  
> Type a message, pick a style, and watch it transform into art!

---

## 📝 OVERVIEW

This web app allows users to:

- ✍️ Enter custom messages (up to **200 characters**)
- 🧾 Choose between fonts: `standard`, `shadow`, `thinkertoy`
- 🔤 Instantly convert text into **ASCII art** displayed in the browser

---

## ⚙️ TECHNOLOGIES USED

- ✅ **Go (Golang)** – server-side logic
- ✅ **HTML Templates** – frontend rendering
- ✅ **CSS** – basic styling
- ✅ **Custom .txt font files** – ASCII banner formats

---

## 🗂️ PROJECT STRUCTURE

```
asciiArtWeb/
├── banners/               # ASCII font styles
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── functions/
│   └── AsciiRepresentation.go
│   └── ReadAsciiBanner.go
│   └── ErrorHandler.go   
│   └── handlerIndex.go  
│   └── handlerPost.go
│   └── handerStatic.go    
     
├── static/
│   └── style.css          # Frontend styling
├── templates/
│   ├── index.html         # Form view
│   └── error.html         # Error page
├── main.go                # Entry point
└── README.md              # Project description
```

---

## 🚀 RUNNING LOCALLY

### ✅ Requirements

- Go **1.18+**
- A web browser
- Terminal / Command line

### ▶️ Steps

```bash
# Run the app
go run main.go
```

Open your browser and go to:

```
http://localhost:8080
```

---

## 🌐 ROUTING SUMMARY

| Route           | Method | Description                    |
|----------------|--------|--------------------------------|
| `/`             | GET    | Displays the form              |
| `/ascii-art`    | POST   | Processes form & renders art   |

---

## 🧠 IMPLEMENTATION DETAILS

### 🔁 Request Flow

1. User visits `/`
2. Fills in the form (text + banner)
3. Submits the form (`POST /ascii-art`)
4. Server processes:
   - Validates input
   - Loads correct banner
   - Maps each rune to ASCII
   - Renders final output

### 🧾 Banner File Parsing

- Each `.txt` file contains 8-line ASCII representations
- `ReadAsciiBanner(filename)` reads and splits characters
- Mapped to printable ASCII characters (`32` to `126`)

### 🔡 ASCII Generation Logic

Function: `AsciiRepresentation(str, asciiMap)`

- Handles multi-line user input (`\r\n`)
- Builds ASCII art line by line

---

## ❗ ERROR HANDLING

- Invalid input or banner triggers error page
- Custom `error.html` rendered using `ErrorHandler()`

---