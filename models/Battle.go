package models

import (
	"fmt"
)

// Battle blablabla
type Battle struct {
	ID int
}

// Fight blablabla
func (f *Battle) Fight(c chan int) {
	for {
		data := <-c
		//time.Sleep(time.Millisecond * 2000)
		fmt.Printf("worker %[1]d got %[2]d \n", f.ID, data)
	}
}
