# Log Analyzer
[![Go](https://img.shields.io/badge/Language-Go-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow)](https://opensource.org/licenses/MIT)
[![CI](https://img.shields.io/github/workflows/log-analyzer/main)](https://github.com/samyalderson/log-analyzer/actions?query=workflow%3Amain)

## Extract Insights from Server Logs with Ease

A log parser analyzer built with Go for extracting valuable insights from server logs.

## Features

* Log file parsing from standard input or file
* Analyzer to extract insights (e.g. error rates, request counts)
* Storage and retrieval of parsed logs

## Prerequisites

* Go installed on your system (`go version >= 1.17`)
* A log file or standard input to parse

## Installation Instructions

```bash
go get github.com/samyalderson/log-analyzer
```

## Usage

### Example 1: Parse a log file
```go
package main

import (
	"log"
	"os"

	"github.com/samyalderson/log-analyzer/cmd/log-analyzer"
)

func main() {
	logFile := "path/to/logfile.log"
	err := log_analyzer.Parse(logFile)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Example 2: Parse from standard input
```go
package main

import (
	"log"
	"os"

	"github.com/samyalderson/log-analyzer/cmd/log-analyzer"
)

func main() {
	err := log_analyzer.ParseFromStdin()
	if err != nil {
		log.Fatal(err)
	}
}
```

## Project Architecture

The project consists of the following modules:

* `cmd/log-analyzer/main.go`: Main entry point
* `internal/analyzer/analyzer.go`: Log analyzer module
* `internal/parse/parser.go`: Log parser module
* `internal/store/store.go`: Log storage module
* `internal/utils/utils.go`: Utility functions
* `internal/test/*_test.go`: Unit tests for each module

## Building from Source

```bash
go build -o log-analyzer cmd/log-analyzer/main.go
```

## Testing

`go test` will run all unit tests in the `internal/test` package.

## Contributing Guidelines

1. Fork the repository
2. Create a new branch for your changes
3. Implement your changes
4. Run `go test` to ensure all tests pass
5. Submit a pull request with a clear description of your changes

## License

MIT License

Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.