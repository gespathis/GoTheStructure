package models

import (
	"fmt"
)

// Battle blablabla
type Battle struct {
	ID int
}

// Sparring blablabla
func (f *Battle) Sparring(c chan int) {
	for {
		data := <-c
		//time.Sleep(time.Millisecond * 2000)
		fmt.Printf("worker %[1]d got %[2]d \n", f.ID, data)
	}
}
