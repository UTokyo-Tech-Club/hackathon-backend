package user

import (
	firebase "firebase.google.com/go"
)

type UserController struct {
	userUsecase UserUsecase
}

func (ui *UserController) AuthUser(fb *firebase.App, data string) {

}
