package events

import "time"

type UserCreated struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	NickName  string    `json:"nickName"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserCreatedEvent(id, name, surname, pass, nick, email, country string, createdAt time.Time) UserCreated {
	return UserCreated{
		Id:        id,
		FirstName: name,
		LastName:  surname,
		Email:     email,
		NickName:  nick,
		Country:   country,
		Password:  pass,
		CreatedAt: createdAt,
	}
}
