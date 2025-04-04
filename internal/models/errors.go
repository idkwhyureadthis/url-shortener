package models

import "errors"

var ErrCreationFailed = errors.New("failed to create new URL")
var ErrAlreadyOccupied = errors.New("such link already exists")
var ErrWrongInitialLink = errors.New("wrong initial link provided")
var ErrNotFound = errors.New("link not found")
var ErrBadCustomLink = errors.New("bad custom link provided")
