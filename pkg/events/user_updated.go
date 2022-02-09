package events

import "time"

type UserUpdated struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	NickName  string    `json:"nickName"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserUpdatedEvent(id, name, surname, nick, email, country string, updatedAt time.Time) UserUpdated {
	return UserUpdated{
		Id:        id,
		FirstName: name,
		LastName:  surname,
		Email:     email,
		NickName:  nick,
		Country:   country,
		UpdatedAt: updatedAt,
	}
}
