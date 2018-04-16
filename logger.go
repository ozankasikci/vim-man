package fantasia

import (
	"fmt"
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
