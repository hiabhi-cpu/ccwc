# ccwc

🧮 A simple reimplementation of the Unix `wc` (word count) utility written in Go.

`ccwc` mimics the behavior of the standard `wc` command by allowing users to count lines, words, bytes, and characters in files or from standard input.

---
## ✨ Features

- `-l` → Count lines
- `-w` → Count words
- `-c` → Count bytes
- `-m` → Count characters (Unicode-aware)
- Reads from:
  - A specified file
  - Standard input (piped data) if no file is given

---
## 🚀 Getting Started

### ✅ Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### 🔧 Build

```bash
git clone https://github.com/hiabhi-cpu/ccwc.git
cd ccwc
go build -o ccwc
```

This creates an executable named ccwc.

---
## 🧪 Usage
### 📁 From a file

```bash
./ccwc -l test.txt    # Line count
./ccwc -w test.txt    # Word count
./ccwc -c test.txt    # Byte count
./ccwc -m test.txt    # Character count
./ccwc test.txt       # Default: lines, words, bytes
```

### 📥 From standard input
```bash
cat test.txt | ./ccwc -l
```

```bash
echo "Hello world" | ./ccwc -w
```

If no option is passed, the default behavior is:

```bash
./ccwc file.txt
# => lines  words  bytes  filename
```