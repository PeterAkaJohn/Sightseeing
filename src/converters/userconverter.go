package converters

import (
	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/PeterAkaJohn/SightSeeing/src/viewmodel"
)

func ConvertUserToUserVM(user model.User) viewmodel.UserVM {
	userVM := viewmodel.UserVM{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return userVM
}
