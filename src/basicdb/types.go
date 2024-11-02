package basicdb

type DBInstance interface{
	IsDBInstance()
}

type Database struct {
	Id string `json:"id"`
	Instance DBInstance `json:"instance"`
}
