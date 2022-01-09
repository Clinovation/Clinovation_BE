package businesses

import "errors"

var (
	//General errors
	ErrInternalServer        = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrNotFound              = errors.New("Data Not Found")
	ErrIdNotFound            = errors.New("Id Not Found")
	ErrDuplicateData         = errors.New("Data already exist")
	ErrDuplicateEmail        = errors.New("Email already used")
	ErrDuplicateNik          = errors.New("NIK Already Used")
	ErrEmailPasswordNotFound = errors.New("(Email) Or (Password) Empty")
	ErrEmailNotRegistered    = errors.New("Email Not Registered")
	ErrPassword              = errors.New("Wrong Password")

	//Doctors errors
	ErrNotFoundDoctor = errors.New("Doctor Doesn't Exist")

	//Medical Staff> errors
	ErrNotFoundMedicalStaff = errors.New("Medical Staff Doesn't Exist")

	//Nurse errors
	ErrNotFoundNurse = errors.New("Nurse Doesn't Exist")

	//Patient errors
	ErrNotFoundPatient = errors.New("Patient Doesn't Exist")

	//Schedule errors
	ErrNotFoundSchedule = errors.New("Schedule Doesn't Exist")

	//Schedule errors
	ErrNotFoundQueue = errors.New("Queue Doesn't Exist")

	//ErrDuplicateWorkDay
	ErrDuplicateWorkDay = errors.New("Work Day Already Exist")
	ErrNotFoundWorkDay  = errors.New("Work Day Doesn't Exist")

	//ErrDuplicateWorkHour
	ErrDuplicateWorkHour = errors.New("Work Hour Already Exist")
	ErrNotFoundWorkHour  = errors.New("Work Hour Doesn't Exist")
)
