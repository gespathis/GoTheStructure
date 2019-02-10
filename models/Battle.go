package models

import (
	"fmt"
	"time"
)

// Battle blablabla
type Battle struct {
	ID int
}

// Sparring blablabla
func (f *Battle) Sparring(c chan Condenders) {
	//for {
	data := <-c
	time.Sleep(time.Millisecond * 1000)
	//fmt.Println(data.FirstFighter.Strength)
	//fmt.Println(data.SecondFighter.Strength)
	if data.FirstFighter.Strength > data.SecondFighter.Strength {
		fmt.Printf("fight %[1]d won %[2]s \n", f.ID, data.FirstFighter.Human.Name)
	} else {
		fmt.Printf("fight %[1]d won %[2]s \n", f.ID, data.SecondFighter.Human.Name)
	}
	//}
}
