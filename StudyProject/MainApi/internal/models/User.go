package models

type user struct {
	Id       int    `json: id`
	Name     string `json: name`
	Username string `json: username`
	Password string `json: password`
}

func NewUser(name string, username string, password string) *user {
	return &user{
		Name:     name,
		Username: username,
		Password: password,
	}
}
