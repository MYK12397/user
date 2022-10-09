package body

import (
	"gocrud/storage"
	"time"
)

type PasswordRequestData struct {
	Data PasswordRequest `json:"data"`
}

type PasswordRequest struct {
	Email string `json:"email"`
}

// UserRequestData represents request to create a new user.
type UserRequestData struct {
	Data UserRequest `json:"data"`
}

// UserResponseData represents response object of user.
type UserResponseData struct {
	Data UserResponse `json:"data"`
}

// UsersListResponse represents user list response data
type UsersListResponse struct {
	Data []UserResponse `json:"data"` // Users list data
}

// UserRequest represents user object in request body.
type UserRequest struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	ProfilePicture string `json:"profilePicture"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	Country        string `json:"country"`
	DateOfBirth    string `json:"dateOfBirth,omitempty"`
}

// UserResponse represents user object in request body.
type UserResponse struct {
	ID             string     `json:"id"`
	Created        *time.Time `json:"created"`
	Modified       *time.Time `json:"modified,omitempty"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	ProfilePicture string     `json:"profilePicture"`
	PhoneNumber    string     `json:"phoneNumber"`
	Email          string     `json:"email"`
	Country        string     `json:"country"`
	DateOfBirth    string     `json:"dateOfBirth,omitempty"`
}

// MapObjects maps request fields to data access object.
func (request UserRequestData) MapObjects() storage.User {
	var user storage.User

	user.FirstName = request.Data.FirstName
	user.LastName = request.Data.LastName
	user.ProfilePicture = request.Data.ProfilePicture
	user.PhoneNumber = request.Data.PhoneNumber
	user.Email = request.Data.Email
	user.Country = request.Data.Country
	user.DateOfBirth = request.Data.DateOfBirth

	return user
}

// MapResponse maps fields from DAO to response object
func (response *UserResponseData) MapResponse(user *storage.User) {
	response.Data.ID = user.ID
	response.Data.Created = user.Created
	response.Data.Modified = user.Modified
	response.Data.FirstName = user.FirstName
	response.Data.LastName = user.LastName
	response.Data.ProfilePicture = user.ProfilePicture
	response.Data.PhoneNumber = user.PhoneNumber
	response.Data.Email = user.Email
	response.Data.Country = user.Country
	response.Data.DateOfBirth = user.DateOfBirth
}
