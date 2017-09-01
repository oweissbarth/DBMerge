package model

type Server struct {
	host      string
	port      int
	user      string
	password  string
	databases []Database
}

func (s Server) connect() {

}
