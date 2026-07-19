// Package main provides the main entry point for the log-analyzer.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"
)

// logAnalyzer represents the log analyzer.
type logAnalyzer struct {
	parser   *parser
	analyzer *analyzer
	store    *store
}

// newLogAnalyzer creates a new log analyzer.
func newLogAnalyzer() *logAnalyzer {
	return &logAnalyzer{
		parser:   newParser(),
		analyzer: newAnalyzer(),
		store:    newStore(),
	}
}

// run runs the log analyzer.
func (l *logAnalyzer) run() error {
	// Read input from standard input or file
	input := newInputReader()
	if err := input.read(); err != nil {
		return err
	}

	// Parse log entries
	logEntries, err := l.parser.parse(input.getLogEntries())
	if err != nil {
		return err
	}

	// Analyze log entries
	insights, err := l.analyzer.analyze(logEntries)
	if err != nil {
		return err
	}

	// Store insights
	if err := l.store.store(insights); err != nil {
		return err
	}

	// Print insights
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	for _, insight := range insights {
		table.Append([]string{insight.Metric, fmt.Sprintf("%v", insight.Value)})
	}
	table.Render()

	return nil
}

func main() {
	// Set up logging
	logFlags := flag.NewFlagSet("log-analyzer", flag.ExitOnError)
	logFlags.Parse(os.Args[1:])

	// Set up input file
	inputFile := flag.String("input", "", "path to input file")
	flag.Parse()

	// Create log analyzer
	l := newLogAnalyzer()

	// Run log analyzer
	if err := l.run(); err != nil {
		glog.Errorf("Error running log analyzer: %v", err)
		os.Exit(1)
	}
}
