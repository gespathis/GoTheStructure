package models

import (
	"structure/public"
)

// Person  test
type Person struct {
	*Human
	*Saiyan
	Enemies map[string]*Person
	Logger  public.ICanLog
}
