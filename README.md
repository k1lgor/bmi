# BMI Calculator

## Overview

This is a command-line and text-based UI (TUI) tool for calculating Body Mass Index (BMI) based on user input for weight and height. It supports multiple units, offers detailed BMI categorization, and includes versioning.

## Features

- **CLI**: Run via the command line with unit support for kg/lbs and cm/in.
- **TUI**: Interactive text-based UI using `tview`.
- **Web Interface**: Simple web version built with Gin framework, offering a user-friendly web form.
- **Versioning**: Display version info using `--version`.
- **Unit Testing**: Built-in tests using Go’s testing framework.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/k1lgor/bmi
cd bmi
```

2. Install dependencies (for TUI and Web):

```bash
go get github.com/rivo/tview
go get github.com/gin-gonic/gin
```

3. Build the project:

```bash
go build -o bmi ./cmd/bmi
```

## Usage

- **CLI**

```bash
./bmi 70 kg 170 cm
```

- **TUI**

```bash
./bmi -tui
```

- **Web**

```bash
./bmi -web
```

Then open your browser at `http://localhost:8080`.

- **Version**

```bash
./bmi -version
```

## Project Structure

bmi/
│
├── cmd/
│   └── bmi/
│       └── main.go         # CLI and TUI entry point
│
├── pkg/
│   ├── bmi/
│   │   ├── bmi.go          # Core BMI logic
│   │   └── bmi_test.go     # Unit tests for BMI
│   └── input/
│       └── input.go        # Input parsing and validation
│
├── ui/
│   └── tui.go              # TUI logic for text-based interface
│
├── web/
│   ├── web.go              # Web server logic using Gin
│   └── templates/
│       └── index.html      # Web form UI
│
└── go.mod                  # Go module file
