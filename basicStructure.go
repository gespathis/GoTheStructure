package main

import (
	"fmt"
	"math/rand"
	"sort"
	"structure/logers"
	"structure/models"
)

func main() {

	songoku := &models.Person{Human: new(models.Human), Saiyan: new(models.Saiyan)}
	songoku.Logger = new(logers.ConsoleLogger)

	var result bool
	result = songoku.Logger.Log("i am concrete")

	fmt.Println(result)

	songoku.Love = 5000000
	songoku.Power = 10000
	songoku.Saiyan.Father = &models.Saiyan{Name: "bardock", Power: 1000}
	songoku.Human.Mother = &models.Human{Name: "helena", Love: 100000000}
	songoku.Human.Name = "songoku"
	songoku.Saiyan.Name = "kakarot"
	songoku.Friends = make([]*models.Human, 0, 5)
	songoku.Friends = songoku.Friends[0:5]
	songoku.Friends = append(songoku.Friends, &models.Human{Name: "gohan"})
	songoku.Enemies = make(map[string]*models.Person)

	songoku.Enemies["freizer"] = &models.Person{Saiyan: &models.Saiyan{Name: "freizer", Power: 99999}}

	test := songoku.Enemies["freizer"]

	fmt.Println(test.Power)

	if songoku.Human.Mother != nil {
		fmt.Println(songoku.Human.Mother.Name)
	}

	tempFrients := make([]*models.Human, len(songoku.Friends))

	copy(tempFrients, songoku.Friends[4:])

	fmt.Println(songoku.Human.Love)
	fmt.Println(songoku.Human.Name)
	fmt.Println(songoku.Human.Friends)
	fmt.Println(songoku.Human.Friends[5])
	fmt.Println(tempFrients)

	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[3:6])
	fmt.Println(worst)
}
