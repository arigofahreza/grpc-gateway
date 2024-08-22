package models

type UserModel struct {
	Id    string `json:"_id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
