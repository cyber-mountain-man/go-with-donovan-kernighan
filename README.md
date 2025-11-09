# Go with Donovan & Kernighan

Companion repository for working through *The Go Programming Language* by **Alan Donovan** and **Brian Kernighan**.

Each chapter includes runnable examples (`cmd/`), reusable packages (`pkg/`), and personal notes.  
This project mirrors the book’s progression from fundamentals to concurrency and testing.

---

## Getting Started

Clone the repository and run a sample:

```bash
git clone git@github.com:cyber-mountain-man/go-with-donovan-kernighan.git
cd go-with-donovan-kernighan
go run ./ch01-intro/cmd/hello
```

Run all tests:

```bash
go test ./...
```

---

## Project Layout

```
go-with-donovan-kernighan/
├── ch01-intro/                # Introduction
│   ├── cmd/hello/             # "Hello, Go" example
│   └── pkg/stringsutil/       # String utilities
├── ch02-program-structure/    # Program structure examples
├── internal/                  # Helper tools
├── go.mod                     # Go module definition
├── LICENSE                    # MIT license
└── README.md
```

---

## Progress

* [x] Chapter 1 – Introduction
* [ ] Chapter 2 – Program Structure
* [ ] Chapter 3 – Basic Data Types
* [ ] Chapter 4 – Composite Types
* [ ] Chapter 5 – Functions
* [ ] Chapter 6 – Methods
* [ ] Chapter 7 – Interfaces
* [ ] Chapter 8 – Concurrency
* [ ] Chapter 9 – Patterns
* [ ] Chapter 10 – Testing

---

## License

MIT License © 2025 Guillermo Morrison
