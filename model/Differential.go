package model

/*Differential contains a list of changes between two versions of a database */
type Differential struct {
	Additions     []Addition
	Deletions     []Deletion
	Modifications []Modification
	Conflicts     []Conflict
}
