package domain

import "errors"

var (
	ErrSchoolNotFound     = errors.New("no rows in result set")
	ErrDuplicatedSchool   = errors.New("school already exists")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
