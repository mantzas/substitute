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

// Error logging
func (l *Logger) Error(args ...interface{}) {
}

// Errorf logging
func (l *Logger) Errorf(msg string, args ...interface{}) {
}

// Errorln logging
func (l *Logger) Errorln(args ...interface{}) {
}

// Warn logging
func (l *Logger) Warn(args ...interface{}) {
}

//Warnf logging
func (l *Logger) Warnf(msg string, args ...interface{}) {
}

// Warnln logging
func (l *Logger) Warnln(args ...interface{}) {
}

// Info logging
func (l *Logger) Info(args ...interface{}) {
}

// Infof logging
func (l *Logger) Infof(msg string, args ...interface{}) {
}

// Infoln logging
func (l *Logger) Infoln(args ...interface{}) {
}

// Debug logging
func (l *Logger) Debug(args ...interface{}) {
}

// Debugf logging
func (l *Logger) Debugf(msg string, args ...interface{}) {
}

// Debugln logging
func (l *Logger) Debugln(args ...interface{}) {
}
