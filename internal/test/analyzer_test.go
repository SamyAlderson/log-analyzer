package internal

import (
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"

	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/parse"
	"log-analyzer/internal/utils"
)

func TestAnalyzer(t *testing.T) {
	// Create a sample log entry
	log := parse.Log{
		Timestamp: "2022-01-01T12:00:00",
		Level:     "INFO",
		Message:   "Some info message",
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer
	actual := a.Analyze(log)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "INFO",
		"message":   "Some info message",
		"count":     1,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerMultipleLogs(t *testing.T) {
	// Create multiple log entries
	logs := []parse.Log{
		{
			Timestamp: "2022-01-01T12:00:00",
			Level:     "INFO",
			Message:   "Some info message",
		},
		{
			Timestamp: "2022-01-01T12:00:01",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:02",
			Level:     "INFO",
			Message:   "Some info message",
		},
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer
	actual := a.AnalyzeMultiple(logs)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "INFO",
		"message":   "Some info message",
		"count":     2,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerError(t *testing.T) {
	// Create a sample log entry with an error
	log := parse.Log{
		Timestamp: "2022-01-01T12:00:00",
		Level:     "ERROR",
		Message:   "Some error message",
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer
	actual := a.Analyze(log)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "ERROR",
		"message":   "Some error message",
		"count":     1,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerErrorMultipleLogs(t *testing.T) {
	// Create multiple log entries with errors
	logs := []parse.Log{
		{
			Timestamp: "2022-01-01T12:00:00",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:01",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:02",
			Level:     "INFO",
			Message:   "Some info message",
		},
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer
	actual := a.AnalyzeMultiple(logs)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "ERROR",
		"message":   "Some error message",
		"count":     2,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerTable(t *testing.T) {
	// Create a sample log entry
	log := parse.Log{
		Timestamp: "2022-01-01T12:00:00",
		Level:     "INFO",
		Message:   "Some info message",
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Create a table writer
	w := tablewriter.NewWriter(t)
	defer w.Flush()

	// Test the analyzer table
	a.AnalyzeTable(log, w)
}

func TestAnalyzerTableMultipleLogs(t *testing.T) {
	// Create multiple log entries
	logs := []parse.Log{
		{
			Timestamp: "2022-01-01T12:00:00",
			Level:     "INFO",
			Message:   "Some info message",
		},
		{
			Timestamp: "2022-01-01T12:00:01",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:02",
			Level:     "INFO",
			Message:   "Some info message",
		},
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Create a table writer
	w := tablewriter.NewWriter(t)
	defer w.Flush()

	// Test the analyzer table
	a.AnalyzeTableMultiple(logs, w)
}

func TestAnalyzerStats(t *testing.T) {
	// Create a sample log entry
	log := parse.Log{
		Timestamp: "2022-01-01T12:00:00",
		Level:     "INFO",
		Message:   "Some info message",
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer stats
	actual := a.AnalyzeStats(log)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "INFO",
		"message":   "Some info message",
		"count":     1,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerStatsMultipleLogs(t *testing.T) {
	// Create multiple log entries
	logs := []parse.Log{
		{
			Timestamp: "2022-01-01T12:00:00",
			Level:     "INFO",
			Message:   "Some info message",
		},
		{
			Timestamp: "2022-01-01T12:00:01",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:02",
			Level:     "INFO",
			Message:   "Some info message",
		},
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer stats
	actual := a.AnalyzeStatsMultiple(logs)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "INFO",
		"message":   "Some info message",
		"count":     2,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerStatsError(t *testing.T) {
	// Create a sample log entry with an error
	log := parse.Log{
		Timestamp: "2022-01-01T12:00:00",
		Level:     "ERROR",
		Message:   "Some error message",
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer stats
	actual := a.AnalyzeStats(log)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "ERROR",
		"message":   "Some error message",
		"count":     1,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}

func TestAnalyzerStatsErrorMultipleLogs(t *testing.T) {
	// Create multiple log entries with errors
	logs := []parse.Log{
		{
			Timestamp: "2022-01-01T12:00:00",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:01",
			Level:     "ERROR",
			Message:   "Some error message",
		},
		{
			Timestamp: "2022-01-01T12:00:02",
			Level:     "INFO",
			Message:   "Some info message",
		},
	}

	// Create an analyzer instance
	a := analyzer.New()

	// Test the analyzer stats
	actual := a.AnalyzeStatsMultiple(logs)
	expected := map[string]interface{}{
		"timestamp": "2022-01-01T12:00:00",
		"level":     "ERROR",
		"message":   "Some error message",
		"count":     2,
	}

	// Assert the result
	assert.Equal(t, expected, actual)
}
