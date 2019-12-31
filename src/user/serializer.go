package user

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (self *User) Response() UserResponse {
	return UserResponse{
		ID:    self.ID,
		Email: self.Email,
		Name:  self.Name,
	}
}
