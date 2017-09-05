package model

/*Change represents a generic difference (single row) betweeen two tables*/
type Change interface {
	asString() string
	primaryKey() string
}

/*Modification is a row that is present in both tables but on or more column differ (except primary key)*/
type Modification struct {
}

/*Addition is a row that was not present in the original table*/
type Addition struct {
	PrimaryKey int
}

/*Deletion is a row that is no longer present in this version of the table*/
type Deletion struct {
}
