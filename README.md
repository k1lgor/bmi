# BMI Calculator

## Overview

This is a command-line and text-based UI (TUI) tool for calculating Body Mass Index (BMI) based on user input for weight and height. It supports multiple units, offers detailed BMI categorization, and includes versioning.

## Features

- **Units Flexibility**: Supports both metric (kg/cm) and imperial (lbs/in) units.
- **BMI Interpretation**: Provides detailed categorization (underweight, normal, overweight, obese).
- **Error Handling**: Input validation with user-friendly error messages.
- **TUI Support**: Simple, interactive text-based interface using `tview`.
- **Versioning**: Use `--version` to display the app version.
- **Unit Tests**: Robust testing with Go's `testing` package.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/k1lgor/bmi
cd bmi
```

2. Install dependencies (for TUI):

```bash
go get github.com/rivo/tview
```

3. Build the project:

```bash
go build -i bmi ./cmd/bmi
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
└── go.mod                  # Go module file
