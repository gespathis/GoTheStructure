package main

import (
	"fmt"
	"math/rand"
	"sort"
	"structure/logers"
	"structure/models"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex
)

type testValue struct {
	name string
	age  int
}

func (t testValue) ChangeName() {
	t.name = "makis"
}

func (t *testValue) ChangeNamePointer() {
	t.name = "makis"
}

func main() {

	test1 := testValue{
		name: "spathis",
		age:  39}

	test1.ChangeName()

	fmt.Println(test1)

	test1.ChangeNamePointer()
	fmt.Println(test1)

	songoku := &models.Person{Human: new(models.Human), Saiyan: new(models.Saiyan)}
	krilin := &models.Person{Human: &models.Human{}, Saiyan: new(models.Saiyan)}
	songoku.Logger = new(logers.ConsoleLogger)

	var result bool
	result = songoku.Logger.Log("i am concrete")

	fmt.Println(result)
	krilin.Strength = 500
	krilin.Human.Name = "krilin"
	songoku.Love = 5000000
	songoku.Strength = 10000
	songoku.Saiyan.Father = &models.Saiyan{Name: "bardock", Strength: 1000}
	songoku.Human.Mother = &models.Human{Name: "helena", Love: 100000000}
	songoku.Human.Name = "songoku"
	songoku.Saiyan.Name = "kakarot"
	songoku.Friends = make([]*models.Human, 0, 5)
	songoku.Friends = songoku.Friends[0:5]
	songoku.Friends = append(songoku.Friends, &models.Human{Name: "gohan"})
	songoku.Enemies = make(map[string]*models.Person)

	songoku.Enemies["freizer"] = &models.Person{Saiyan: &models.Saiyan{Name: "freizer", Strength: 99999}}

	test := songoku.Enemies["freizer"]

	fmt.Println(test.Strength)

	if songoku.Human.Mother != nil {
		fmt.Println(songoku.Human.Mother.Name)
	}

	tempFrients := make([]*models.Human, len(songoku.Friends))

	copy(tempFrients, songoku.Friends[4:])

	// fmt.Println(songoku.Human.Love)
	// fmt.Println(songoku.Human.Name)
	// fmt.Println(songoku.Human.Friends)
	// fmt.Println(songoku.Human.Friends[5])
	// fmt.Println(tempFrients)

	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[3:6])
	fmt.Println(worst)

	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)

	c := make(chan models.Condenders)

	for i := 0; i < 5; i++ {
		batlle := &models.Battle{ID: i}
		go batlle.Sparring(c)
	}

	fmt.Println("processing")

	time.Sleep(time.Millisecond * 3000)

	for i := 0; i < 6; i++ {
		songoku.Strength = rand.Int()
		krilin.Strength = rand.Int()
		select {
		case c <- models.Condenders{FirstFighter: songoku, SecondFighter: krilin}:
			// case k := <-time.After(time.Millisecond * 6000):
			// 	fmt.Println("timed out", k)
		default:
			fmt.Printf("worker %[1]d dropped \n", i)
		}

	}
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("exit")
}

func after(d time.Duration) chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		c <- true
	}()
	return c
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	//fmt.Println(counter)
}
