package dto

type CreateUserDto struct {
	FirstName   string `json:"firstName" validate:"required,alpha,min=2"`
	LastName    string `json:"lastName" validate:"required,alpha,min=2"`
	Email       string `json:"email" validate:"required,email,unique=users.email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,len=10,unique=users.phone_number"`
	Password    string `json:"password" validate:"required,min=8"`
}
