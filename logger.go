package fantasia

import "fmt"

type Logger struct {
	logs []string
}

func NewLogger() *Logger{
	return &Logger{
		logs: make([]string, 0),
	}
}

func (l *Logger ) Log(strings ...string)  {
	l.logs = append(l.logs, strings...)
}

func (l *Logger) DumpLogs() {
	for _, log := range l.logs {
		fmt.Println(log)
	}
}
