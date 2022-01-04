package businesses

import "errors"

var (
	//General errors
	ErrInternalServer        = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrNotFound              = errors.New("Data Not Found")
	ErrIdNotFound            = errors.New("Id Not Found")
	ErrDuplicateData         = errors.New("Data already exist")
	ErrDuplicateEmail        = errors.New("Email already used")
	ErrDuplicateNik          = errors.New("NIK already used")
	ErrEmailPasswordNotFound = errors.New("(Email) or (Password) empty")
	ErrEmailNotRegistered    = errors.New("Email not registered")
	ErrPassword              = errors.New("Wrong Password")

	//Doctors errors
	ErrNotFoundDoctor = errors.New("doctors doesn't exist")
)
