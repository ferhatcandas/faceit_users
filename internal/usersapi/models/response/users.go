package response

type UsersResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func NewUserResponse(id, name, lastname, nick, email, country string) UsersResponse {
	return UsersResponse{
		Id:        id,
		FirstName: name,
		LastName:  lastname,
		NickName:  nick,
		Email:     email,
		Country:   country,
	}
}
