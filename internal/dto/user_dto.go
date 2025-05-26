package dto

type CreateUserDto struct {
	FirstName   string `json:"firstName" binding:"required,alpha,min=2"`
	LastName    string `json:"lastName" binding:"required,alpha,min=2"`
	Email       string `json:"email" binding:"required,email,unique=users.email"`
	PhoneNumber string `json:"phoneNumber" binding:"required,number,len=10,unique=users.phone_number"`
	Password    string `json:"password" binding:"required,min=8"`
}

type LoginDto struct {
	EmailAddress string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}
