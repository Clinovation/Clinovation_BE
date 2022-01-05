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
	ErrNotFoundDoctor = errors.New("doctor doesn't exist")

	//Doctors errors
	ErrNotFoundMedicalStaff = errors.New("medical staff doesn't exist")

	//Patient errors
	ErrNotFoundPatient = errors.New("patient doesn't exist")
)
