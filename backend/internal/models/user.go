package models

type User struct {
	UserName  string   `json:"username"`
	FirstName *string  `json:"firstName,omitempty"`
	LastName  *string  `json:"lastName,omitempty"`
	Birthday  *string  `json:"birthday,omitempty"`
	Gender    *string  `json:"gender,omitempty"`
	Interests []string `json:"interests,omitempty"`
	City      *string  `json:"city,omitempty"`
}

type UserWithPassword struct {
	User
	Password string `json:"password"`
}
