package logers

import (
	"fmt"
)

// ConsoleLogger log to console
type ConsoleLogger struct{}

//Log concrete log print
func (l ConsoleLogger) Log(message string) bool {
	fmt.Println(message)
	return true
}
