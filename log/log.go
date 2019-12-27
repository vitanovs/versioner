package log

import (
	"fmt"
	"log"
)

// Info prints information log message to the
// standart output.
func Info(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	log.Printf("[INFO] %s", msg)
}

// Error prints error log message to the
// standart output.
func Error(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	log.Printf("[ERROR] %s", msg)
}

// Warn prints warning log message to the
// standart output.
func Warn(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	log.Printf("[WARN] %s", msg)
}

// Fatal prints fatal log message to the
// standart output and runs os.Exit(1).
func Fatal(params ...interface{}) {
	log.Fatalf("[FATAL] %s", params...)
}

// Fatalf prints fatal log formatted message to the
// standart output and runs os.Exit(1).
func Fatalf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	log.Fatalf("[FATAL] %s", msg)
}

// Debug prints debug log message to the
// standart output.
func Debug(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	log.Printf("[DEBUG] %s", msg)
}
