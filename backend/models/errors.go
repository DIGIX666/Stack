package models

import "errors"

// Add your existing error definitions here

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrInvalidPassword = errors.New("invalid password")
var ErrUserNotAuthorized = errors.New("user not authorized")
var ErrUserAlreadyInStack = errors.New("user already in stack")
var ErrUserNotInStack = errors.New("user not in stack")
var ErrUserNotInPost = errors.New("user not in post")
var ErrUserNotInContributor = errors.New("user not in contributor")
var ErrUserNotInStackOrPost = errors.New("user not in stack or post")
var ErrUserNotInStackOrContributor = errors.New("user not in stack or contributor")
var ErrUserNotInStackOrPostOrContributor = errors.New("user not in stack, post or contributor")
var ErrMethodsVerifyPassword = errors.New("methods verify password")
var ErrMethodsHashPassword = errors.New("methods hash password")
var ErrUserInvalidCredentials = errors.New("invalid username or password")
