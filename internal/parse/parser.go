// Package parser provides a log parser module for the log-analyzer project.
package parser

import (
	"fmt"
	"internal/utils"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"
)

// Log represents a parsed log entry.
type Log struct {
	Timestamp string
	Level     string
	Message   string
}

// Parser represents a log parser.
type Parser struct {
	logFile string
	re      *regexp.Regexp
}

// NewParser returns a new log parser instance.
func NewParser(logFile string) (*Parser, error) {
	// Regular expression pattern for log parsing.
	re, err := regexp.Compile(`\[(\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\]\[(\w+)\]\s(.*)`)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regular expression: %w", err)
	}
	return &Parser{
		logFile: logFile,
		re:      re,
	}, nil
}

// Parse parses the log file into a slice of Log entries.
func (p *Parser) Parse() ([]Log, error) {
	// Open the log file for reading.
	f, err := os.Open(p.logFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	defer f.Close()

	// Read the log file line by line.
	var logs []Log
	for _, line := range utils.ReadLines(f) {
		// Skip empty lines.
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Parse the log entry using the regular expression.
		match := p.re.FindStringSubmatch(line)
		if len(match) < 3 {
			glog.Errorf("failed to parse log entry: %s", line)
			continue
		}

		// Extract the timestamp, level, and message from the match.
		timestamp := match[1]
		level := match[2]
		message := match[3]

		// Create a new Log entry and append it to the slice.
		logs = append(logs, Log{
			Timestamp: timestamp,
			Level:     level,
			Message:   message,
		})
	}

	return logs, nil
}

// PrintTable prints the parsed logs in a table format.
func (p *Parser) PrintTable(logs []Log) {
	// Create a table writer.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Timestamp", "Level", "Message"})

	// Print each log entry in the table.
	for _, log := range logs {
		table.Append([]string{
			log.Timestamp,
			log.Level,
			log.Message,
		})
	}
	table.Render()

	// Print the total number of logs.
	total := len(logs)
	glog.Infof("Total logs: %d", total)
}

func main() {
	// Example usage.
	parser, err := NewParser("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	logs, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	parser.PrintTable(logs)
}
