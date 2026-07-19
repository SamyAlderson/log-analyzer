// Package analyzer provides a log analyzer module to extract insights from server logs.
package analyzer

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"
)

// Analyzer is a log analyzer module.
type Analyzer struct {
	store  *Store
	parser *Parser
}

// NewAnalyzer returns a new instance of the Analyzer.
func NewAnalyzer(store *Store, parser *Parser) *Analyzer {
	return &Analyzer{
		store:  store,
		parser: parser,
	}
}

// Annotate parses the log file and annotates the parsed logs.
func (a *Analyzer) Annotate(ctx context.Context, logFile string) error {
	// Read the log file from standard input or the specified file.
	logLines, err := a.parser.ParseLines(ctx, logFile)
	if err != nil {
		return err
	}

	// Annotate each log line.
	for _, logLine := range logLines {
		// This was tricky. We need to extract the relevant information from the log line.
		parsedLog, err := a.parser.ParseLogLine(logLine)
		if err != nil {
			return err
		}

		// Store the parsed log in the storage.
		err = a.store.Store(logLine, parsedLog)
		if err != nil {
			return err
		}
	}

	return nil
}

// PrintSummary prints a summary of the parsed logs.
func (a *Analyzer) PrintSummary(ctx context.Context, logFile string) error {
	// Retrieve the stored logs from the storage.
	logLines, err := a.store.RetrieveAll()
	if err != nil {
		return err
	}

	// Create a table to display the summary.
	table := tablewriter.NewWriter(log.New(os.Stdout, "", log.Ldate|log.Ltime))
	table.SetHeader([]string{"Log Line", "Error Rate", "Request Count"})

	// Print each log line in the table.
	for _, logLine := range logLines {
		// This is not proud of this but it works.
		parsedLog, err := a.parser.ParseLogLine(logLine)
		if err != nil {
			return err
		}

		table.Append([]string{
			logLine,
			fmt.Sprintf("%f", parsedLog.ErrorRate),
			fmt.Sprintf("%d", parsedLog.RequestCount),
		})
	}

	table.Render()

	return nil
}
