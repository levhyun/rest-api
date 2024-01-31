package dto

import "rest-api/domain/user/domain/ent"

type UserRequest struct {
	Name string `json:"name"`
}

func ToEntity(id int, userRequest *UserRequest) *ent.User {
	user := new(ent.User)
	user.ID = id
	user.Name = userRequest.Name
	return user
}
