package model

const UserTypeMerchandiser UserType = "merchandiser"
const UserTypeCustomer UserType = "customer"

type UserType string

type AuthRequest struct {
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Type     UserType `json:"type" binding:"required"`
}

type UserEntity struct {
	ID           int
	Username     string
	PasswordHash string
}
