package nursesRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/workDayRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/workHourRepo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Nurses struct {
	gorm.Model
	ID             uint      `gorm:"primary_key:auto_increment"`
	Uuid           uuid.UUID `gorm:"type:varchar(255)"`
	WorkDayID      uint
	WorkDay        workDayRepo.WorkDays `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	WorkHourID     uint
	WorkHour       workHourRepo.WorkHours `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Nik            string                 `gorm:"type:varchar(16)"`
	Name           string                 `gorm:"type:varchar(255)"`
	Email          string                 `gorm:"uniqueIndex;type:varchar(255)"`
	Dob            string                 `gorm:"type:varchar(50)"`
	Sex            string                 `gorm:"type:varchar(6)"`
	Contact        string                 `gorm:"type:varchar(15)"`
	Password       string                 `gorm:"->;<-;not null" `
	Specialist     string                 `gorm:"type:varchar(50)"`
	WorkExperience string                 `gorm:"type:varchar(255)"`
	Avatar         string                 `gorm:"type:varchar(255)"`
	Role           string                 `gorm:"type:varchar(20)"`
}

func ToDomain(rec *Nurses) nursesEntity.Domain {
	return nursesEntity.Domain{
		ID:             rec.ID,
		WorkDayID:      rec.WorkDayID,
		WorkDay:        rec.WorkDay.Day,
		WorkHourID:     rec.WorkHourID,
		WorkHour:       rec.WorkHour.Hour,
		Uuid:           rec.Uuid,
		Nik:            rec.Nik,
		Name:           rec.Name,
		Email:          rec.Email,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		Password:       rec.Password,
		WorkExperience: rec.WorkExperience,
		Avatar:         rec.Avatar,
		Role:           rec.Role,
	}
}

func FromDomain(nurseDomain *nursesEntity.Domain) *Nurses {
	return &Nurses{
		ID:             nurseDomain.ID,
		Uuid:           nurseDomain.Uuid,
		WorkDayID:      nurseDomain.WorkDayID,
		WorkHourID:     nurseDomain.WorkHourID,
		Nik:            nurseDomain.Nik,
		Name:           nurseDomain.Name,
		Email:          nurseDomain.Email,
		Dob:            nurseDomain.Dob,
		Sex:            nurseDomain.Sex,
		Contact:        nurseDomain.Contact,
		Password:       nurseDomain.Password,
		WorkExperience: nurseDomain.WorkExperience,
		Avatar:         nurseDomain.Avatar,
		Role:           nurseDomain.Role,
	}
}

func toDomainArray(record []Nurses) []nursesEntity.Domain {
	var res []nursesEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
