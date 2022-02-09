package request

type UserCreateRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
type UserUpdateRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Country   string `json:"country"`
}
