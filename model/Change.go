package model

/*Modification is a row that is present in both tables but on or more column differ (except primary key)*/
type Modification struct {
	PrimaryKey int
	Content    string
	Origin     *Table
}

/*Addition is a row that was not present in the original table*/
type Addition struct {
	PrimaryKey int
	Content    string
	Origin     *Table
}

/*Deletion is a row that is no longer present in this version of the table*/
type Deletion struct {
	PrimaryKey int
	Content    string
	Origin     *Table
}

/*Conflict is a row that modified by multiple tables and is not triviallly mergeable*/
type Conflict struct {
	PrimaryKey    int
	LocalContent  string
	RemoteContent string
	LocalType     ChangeType
	RemoteType    ChangeType
	LocalOrigin   *Table
	RemoteOrigin  *Table
}

/*ChangeType defines if the differences is an addition, deletion or modification*/
type ChangeType int

const (
	ADD ChangeType = 1 + iota
	DEL
	MOD
)
