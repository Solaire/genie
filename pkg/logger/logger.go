package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	_Logger    *log.Logger
	InfoWriter io.Writer
)

func Init(log_dir string) error {
	if err := cleanupOldLogs(log_dir, 24*time.Hour); err != nil {
		return err
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	path := filepath.Join(log_dir, "genie-"+timestamp+".log")

	log_file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Logger writes ONLY to the log file
	_Logger = log.New(log_file, "[genie] ", log.LstdFlags|log.Lmsgprefix)

	// Writes to stdout and log file
	InfoWriter = io.MultiWriter(os.Stdout, log_file)

	return nil
}

func cleanupOldLogs(dir string, max_age time.Duration) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	cutoff := time.Now().Add(-max_age)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !strings.HasPrefix(name, "genie-") || !strings.HasSuffix(name, ".log") {
			continue
		}

		// Example: "genie-2025-06-03_17-45-12.log"
		path := filepath.Join(dir, name)
		info, err := file.Info()
		if err != nil {
			continue
		}

		if info.ModTime().Before(cutoff) {
			_ = os.Remove(path) // ignore errors silently
		}
	}

	return nil
}

func InfoWritef(format string, args ...any) {
	fmt.Fprintf(InfoWriter, format, args...)
}

func Printf(format string, args ...any) {
	s := fmt.Sprintf(format, args...)
	_Logger.Print(s)
}

func Errorf(format string, args ...any) {
	s := fmt.Sprintf("ERROR: "+format, args...)
	_Logger.Print(s)
}

func Fatalf(v ...any) {
	_Logger.Fatalln(v...)
}
