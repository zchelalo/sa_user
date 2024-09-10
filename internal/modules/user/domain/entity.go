package userDomain

type UserEntity struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Verified bool   `json:"verified"`
}

type HashedPasswordAndID struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
