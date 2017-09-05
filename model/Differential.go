package model

/*Differential contains a list of changes between two versions of a database */
type Differential struct {
	changes   []Change
	fromTable string
	toTable   string
}
