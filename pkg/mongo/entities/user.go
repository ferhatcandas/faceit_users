package entities

import (
	"faceit_users/pkg/events"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `bson:"_id,omitempty"`
	FirstName string    `bson:"first_name,omitempty"`
	LastName  string    `bson:"last_name,omitempty"`
	NickName  string    `bson:"nickname,omitempty"`
	Password  string    `bson:"password,omitempty"`
	Email     string    `bson:"email,omitempty"`
	Country   string    `bson:"country,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

func NewUser(name, surname, nick, pass, email, country string) User {
	return User{
		Id:        uuid.New().String(),
		FirstName: name,
		LastName:  surname,
		NickName:  nick,
		Password:  pass,
		Email:     email,
		Country:   country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
}
func (u *User) UpdateFields(name, surname, country string) {
	u.Country = country
	u.FirstName = name
	u.LastName = surname
	u.UpdatedAt = time.Now()
}

func (u *User) ToUpdatedEvent() events.UserUpdated {
	return events.NewUserUpdatedEvent(u.Id, u.FirstName, u.LastName, u.NickName, u.Email, u.Country, u.UpdatedAt)
}

func (u *User) ToCreatedEvent() events.UserCreated {
	return events.NewUserCreatedEvent(u.Id, u.FirstName, u.LastName, u.Password, u.NickName, u.Email, u.Country, u.CreatedAt)
}
func (u *User) ToDeletedEvent() events.UserDeleted {
	return events.NewUserDeletedEvent(u.Id)
}
