package models

import (
	"GoTheStructure/public"
)

// Person  test
type Person struct {
	*Human
	*Saiyan
	Enemies map[string]*Person
	Logger  public.ICanLog
}
