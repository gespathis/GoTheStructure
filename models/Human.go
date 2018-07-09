package models

// Human blablabla
type Human struct {
	Name    string
	Love    int
	Father  *Human
	Mother  *Human
	Friends []*Human
}
