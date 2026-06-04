# Go-Reloaded

A professional-grade, lightweight text completion, editing, and auto-correction utility tool written in Go. This project is structured using the standard Go industry enterprise layout.

## 📂 Project Structure

```text
go-reloaded/            # Parent Directory
├── cmd/                # Command subdirectory
│   └── go-reloaded/    # Application namespace
│       └── main.go     # Application entry point (CLI & I/O)
├── internal/           # Hidden core logic subsystem package
│   ├── reloaded.go     # Core text-processing utility logic
│   └── reloaded_test.go# Automated unit test suite
├── assets/             # Subdirectory for files/documentation assets
│   ├── sample.txt      # Raw input sample text
│   └── result.txt      # Corrected text output target
├── go.mod              # Module definition dependency manager
└── README.md           # Documentation guide (Root level)
```

## 🚀 Features
- **Number Conversions**: Automatically converts hexadecimal `(hex)` and binary `(bin)` strings into standard base-10 decimals.
- **Text Case Formatting**:
  - Modifies a single preceding word using `(up)`, `(low)`, and `(cap)`.
  - Modifies multiple preceding words using numeric arguments: `(up, <number>)`, `(low, <number>)`, and `(cap, <number>)`.
- **Punctuation Correction**: Adjusts spacing so standard punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) attach cleanly to the left word and leave a single space to the right. Handles grouped marks like `...` and `!?`.
- **Quote Formatting**: Properly encloses strings wrapped inside single quotes `'` by stripping internal padding spaces.
- **Grammar Correction**: Automatically swaps the article `a` to `an` if the next word begins with a vowel (`a, e, i, o, u`) or a silent `h`.

## 🛠️ Setup and Usage

Ensure you have **Go 1.21+** installed on your system.

### Running the Program
To run the text auto-correction tool from the root directory, provide the explicit command subdirectory path followed by your target asset paths:

```bash
go run ./cmd/go-reloaded assets/sample.txt assets/result.txt
```

### Direct Compilation (Alternative)
You can also compile the project into an explicit binary file and execute it cleanly:

```bash
# Compile
go build -o reloaded ./cmd/go-reloaded

# Run binary
./reloaded assets/sample.txt assets/result.txt
```

### Audit Quick-Reference Example
**Input File (`assets/sample.txt`):**
```text
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
```

**Output File (`assets/result.txt`):**
```text
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

## 🧪 Running Automated Tests
This project utilizes isolated table-driven testing within the internal logic module space. To run the validation suite, execute:

```bash
go test -v ./internal/...
```
