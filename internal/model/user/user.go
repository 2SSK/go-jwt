package user

import "github.com/2SSK/jwt/internal/model"

type User struct {
	model.Base
	FirstName    *string `json:"firstName" db:"first_name"`
	LastName     *string `json:"lastName" db:"last_name"`
	Password     *string `json:"-" db:"password"`
	Email        *string `json:"email" db:"email"`
	Phone        *string `json:"phone" db:"phone"`
	Token        *string `json:"-" db:"token"`
	UserType     *string `json:"userType" db:"user_type"`
	RefreshToken *string `json:"-" db:"refresh_token"`
}
