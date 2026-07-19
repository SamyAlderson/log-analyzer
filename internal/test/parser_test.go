package parser

import (
	"testing"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"

	"log-analyzer/internal/utils"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected utils.Log
	}{
		{
			name:     "valid line",
			line:     "2022-01-01 12:00:00 INFO 12345 hello world",
			expected: utils.Log{Timestamp: "2022-01-01 12:00:00", Level: "INFO", ID: "12345", Message: "hello world"},
		},
		{
			name:     "invalid line",
			line:     "2022-01-01 12:00:00",
			expected: utils.Log{},
		},
		{
			name:     "empty line",
			line:     "",
			expected: utils.Log{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log, err := ParseLine(tt.line)
			if err != nil {
				if tt.expected == utils.Log{} {
					return
				}
				t.Errorf("ParseLine(%q) error = %v, expected none", tt.line, err)
				return
			}
			if diff := utils.CompareLogs(tt.expected, log); diff != "" {
				t.Errorf("+expected\n%s-\n%+v", diff, log)
			}
		})
	}
}

func TestParseLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []utils.Log
	}{
		{
			name:     "valid log lines",
			input:    "2022-01-01 12:00:00 INFO 12345 hello world\n2022-01-01 12:01:00 WARN 67890 goodbye world",
			expected: []utils.Log{utils.Log{Timestamp: "2022-01-01 12:00:00", Level: "INFO", ID: "12345", Message: "hello world"}, utils.Log{Timestamp: "2022-01-01 12:01:00", Level: "WARN", ID: "67890", Message: "goodbye world"}},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []utils.Log{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logs, err := ParseLines(tt.input)
			if err != nil {
				if tt.expected == nil {
					return
				}
				t.Errorf("ParseLines(%q) error = %v, expected none", tt.input, err)
				return
			}
			if diff := utils.CompareLogsSlice(tt.expected, logs); diff != "" {
				t.Errorf("+expected\n%s-\n%+v", diff, logs)
			}
		})
	}
}

func TestParseLinesWithErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "invalid log lines",
			input:    "2022-01-01 12:00:00 INFO 12345 hello world\n2022-01-01 12:01:00",
			expected: glog.Errorf("invalid log line: 2022-01-01 12:01:00"),
		},
		{
			name:     "empty input",
			input:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseLinesWithErrors(tt.input)
			if diff := utils.CompareErrors(tt.expected, err); diff != "" {
				t.Errorf("+expected\n%s-\n%+v", diff, err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	utils.InitGlog()
	glog.V(2).Info("Running parser tests")
	glog.V(3).Info("Parser tests finished")
	code := m.Run()
	glog.Flush()
	os.Exit(code)
}