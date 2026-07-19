// Package utils provides utility functions for the log-analyzer project.
package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"
)

// LogPath returns the path to the log file based on the input source.
func LogPath(source string) (string, error) {
	if source == "-" {
		// If the source is "-", use standard input.
		return "/dev/stdin", nil
	}

	if !filepath.IsAbs(source) {
		// If the source is not an absolute path, assume it's relative to the current working directory.
		currentDir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		source = filepath.Join(currentDir, source)
	}

	if !filepath.IsAbs(source) {
		return "", fmt.Errorf("invalid log path: %s", source)
	}

	return source, nil
}

// ParseLogLines parses the given log lines into an array of log entries.
func ParseLogLines(lines []string) ([]LogEntry, error) {
	logEntries := make([]LogEntry, 0, len(lines))

	for _, line := range lines {
		parsedLog, err := ParseLogLine(line)
		if err != nil {
			glog.Errorf("failed to parse log line: %s: %v", line, err)
			continue
		}
		logEntries = append(logEntries, parsedLog)
	}

	return logEntries, nil
}

// ParseLogLine parses a single log line into a LogEntry.
func ParseLogLine(line string) (LogEntry, error) {
	// This was tricky, but the basic idea is to split the log line into its constituent parts
	// based on some heuristic about the expected log format. We'll use a simple prefix-based
	// approach for now. Not proud of this but it works.
	parts := strings.Split(line, " ")

	if len(parts) < 3 {
		return LogEntry{}, fmt.Errorf("invalid log line: %s", line)
	}

	return LogEntry{
		Timestamp: parts[0],
		Level:     parts[1],
		Message:   strings.Join(parts[2:], " "),
	}, nil
}

// LogEntry represents a parsed log entry.
type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

// PrintLogEntries prints the given log entries to the console using the tablewriter package.
func PrintLogEntries(logEntries []LogEntry) error {
	buffer := &bytes.Buffer{}
	writer := tablewriter.NewWriter(buffer)
	writer.SetHeader([]string{"Timestamp", "Level", "Message"})

	for _, logEntry := range logEntries {
		writer.Append([]string{logEntry.Timestamp, logEntry.Level, logEntry.Message})
	}

	writer.Flush()

	fmt.Println(buffer.String())

	return nil
}
