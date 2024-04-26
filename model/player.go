package model

type Player struct {
	FirstName    string `form:"first_name"`
	LastName     string
	JerseyNumber int
	JerseySize   string
}
