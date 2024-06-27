package user

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Login    string `json:"login,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Online   bool   `json:"online,omitempty"`
}
