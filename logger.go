package fantasia

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	logs []string
}

var instance *Logger
var once sync.Once

var lg = GetLogger()

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			logs: make([]string, 0),
		}
	})
	return instance
}

func (l *Logger) Log(strings ...string) {
	l.logs = append(l.logs, strings...)
}

func (l *Logger) LogValue(values ...interface{}) {
	var strings []string
	for _, string := range values {
		value := fmt.Sprintf("%v", string)
		strings = append(strings, value)
	}
	l.logs = append(l.logs, strings...)
}

func (l *Logger) DumpLogs() {
	for _, log := range l.logs {
		fmt.Println(log)
	}
}

func (l *Logger) WriteFile(text string) {
	if os.Getenv("DEBUG") == "1" {
		f, err := os.OpenFile("logfile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(text)
	}
}
