package businesses

import "errors"

var (
	//General errors
	ErrInternalServer = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrNotFound       = errors.New("Data Not Found")
	ErrIdNotFound     = errors.New("Id Not Found")
	ErrDuplicateData  = errors.New("Data already exist")
)