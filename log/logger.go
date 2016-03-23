package log

import "log"

// Logger custom logger implementing the StdLogger interface
type Logger struct {
}

// Print logging
func (l *Logger) Print(args ...interface{}) {
	log.Print(args...)
}

// Printf logging
func (l *Logger) Printf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Println logging
func (l *Logger) Println(args ...interface{}) {
	log.Println(args...)
}

// Panic logging
func (l *Logger) Panic(args ...interface{}) {
	log.Panic(args...)
}

// Panicf logging
func (l *Logger) Panicf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

// Panicln logging
func (l *Logger) Panicln(args ...interface{}) {
	log.Panicln(args...)
}

// Fatal logging
func (l *Logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logging
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

// Fatalln logging
func (l *Logger) Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}
