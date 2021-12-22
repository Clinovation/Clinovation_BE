package doctorsRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctors struct {
	gorm.Model
	ID          uint      `gorm:"primary_key:auto_increment"`
	Uuid        uuid.UUID `gorm:"type:varchar(255)"`
	Name        string    `gorm:"type:varchar(255)"`
	Username    string    `gorm:"type:varchar(50)"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password    string    `gorm:"->;<-;not null" `
	PhoneNumber string    `gorm:"type:varchar(15)"`
	Role        string    `gorm:"type:varchar(5)"`
	Avatar      string    `gorm:"type:varchar(255)"`
}

func ToDomain(rec *Doctors) doctorsEntity.Domain {
	return doctorsEntity.Domain{
		ID:          rec.ID,
		Uuid:        rec.Uuid,
		Name:        rec.Name,
		Username:    rec.Username,
		Email:       rec.Email,
		Password:    rec.Password,
		PhoneNumber: rec.PhoneNumber,
		Role:        rec.Role,
		Avatar:      rec.Avatar,
	}
}

func FromDomain(doctorDomain *doctorsEntity.Domain) *Doctors {
	return &Doctors{
		ID:          doctorDomain.ID,
		Uuid:        doctorDomain.Uuid,
		Name:        doctorDomain.Name,
		Username:    doctorDomain.Username,
		Email:       doctorDomain.Email,
		Password:    doctorDomain.Password,
		PhoneNumber: doctorDomain.PhoneNumber,
		Role:        doctorDomain.Role,
		Avatar:      doctorDomain.Avatar,
	}
}
