// Package store provides storage functionality for parsed logs.
package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/glog"
	"github.com/olekukonko/tablewriter"
)

// Store is a log storage module that stores parsed logs in a JSON file.
type Store struct {
	filePath string
}

// NewStore returns a new Store instance.
func NewStore(filePath string) (*Store, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is required")
	}
	return &Store{filePath: filePath}, nil
}

// StoreLogs stores parsed logs in the JSON file.
func (s *Store) StoreLogs(logs []Log) error {
	data, err := json.Marshal(logs)
	if err != nil {
		return err
	}
	return s.write(data)
}

// LoadLogs loads parsed logs from the JSON file.
func (s *Store) LoadLogs() ([]Log, error) {
	data, err := s.read()
	if err != nil {
		return nil, err
	}
	var logs []Log
	err = json.Unmarshal(data, &logs)
	return logs, err
}

// write writes the provided data to the JSON file.
func (s *Store) write(data []byte) error {
	f, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		glog.Errorf("failed to write to file: %v", err)
		return err
	}
	glog.Infof("wrote %d bytes to file", len(data))
	return nil
}

// read reads the data from the JSON file.
func (s *Store) read() ([]byte, error) {
	f, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []byte
	if _, err := io.copyN(f, &data, 1024); err != nil {
		glog.Errorf("failed to read from file: %v", err)
		return nil, err
	}
	return data, nil
}

// Log represents a parsed log.
type Log struct {
	Timestamp time.Time
	Message   string
	Level     string
}

// PrintLog prints the log in a human-readable format.
func (l Log) PrintLog(writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Timestamp", "Message", "Level"})
	table.Append([]string{l.Timestamp.Format(time.RFC3339), l.Message, l.Level})
	table.Render()
}