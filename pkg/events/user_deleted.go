package events

type UserDeleted struct {
	Id string `json:"id"`
}

func NewUserDeletedEvent(id string) UserDeleted {
	return UserDeleted{
		Id: id,
	}
}
