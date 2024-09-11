package userErrors

import "errors"

var ErrIdRequired = errors.New("id is required")
var ErrIdInvalid = errors.New("id is not a valid UUID")

var ErrNameRequired = errors.New("name is required")
var ErrNameInvalid = errors.New("name must be at least 3 characters long")

var ErrEmailRequired = errors.New("email is required")
var ErrEmailInvalid = errors.New("email is not valid")
var ErrEmailAlreadyExists = errors.New("email already exists")

var ErrPasswordRequired = errors.New("password is required")
var ErrPasswordInvalid = errors.New("password must be at least 8 characters long")

var ErrUserNotFound = errors.New("user not found")
var ErrUsersNotFound = errors.New("users not found")
