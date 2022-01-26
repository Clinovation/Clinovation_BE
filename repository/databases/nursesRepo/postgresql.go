package nursesRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type NursesRepository struct {
	db *gorm.DB
}

func NewNursesRepository(db *gorm.DB) nursesEntity.Repository {
	return &NursesRepository{
		db: db,
	}
}

func (r *NursesRepository) CreateNewNurse(ctx context.Context, nurseDomain *nursesEntity.Domain) (*nursesEntity.Domain, error) {
	rec := FromDomain(nurseDomain)
	rec.Uuid, _ = uuid.NewRandom()
	rec.Role = "approve_waiting_list"
	// 	rec.Role = "nurse"
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *NursesRepository) GetByNik(ctx context.Context, nik string) (nursesEntity.Domain, error) {
	rec := Nurses{}

	err := r.db.Where("nik = ?", nik).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *NursesRepository) GetWaitingList(ctx context.Context, offset, limit int) (*[]nursesEntity.Domain, int64, error) {
	var totalData int64
	domain := []nursesEntity.Domain{}
	rec := []Nurses{}

	r.db.Find(&rec, "role = ?", "approve_waiting_list").Count(&totalData)

	err := r.db.Limit(limit).Offset(offset).Find(&rec, "role = ?", "approve_waiting_list").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}

func (r *NursesRepository) ForgetPassword(ctx context.Context, nik string, email string) (nursesEntity.Domain, error) {
	rec := Nurses{}
	err := r.db.Where("nik = ? AND email = ?", nik, email).Find(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *NursesRepository) GetByEmail(ctx context.Context, email string) (nursesEntity.Domain, error) {
	rec := Nurses{}

	err := r.db.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *NursesRepository) AcceptNurse(ctx context.Context, id string, nurseDomain *nursesEntity.Domain) (*nursesEntity.Domain, error) {
	rec := FromDomain(nurseDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *NursesRepository) GetByUuid(ctx context.Context, uuid string) (nursesEntity.Domain, error) {
	rec := Nurses{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *NursesRepository) UpdateNurse(ctx context.Context, id string, nurseDomain *nursesEntity.Domain) (*nursesEntity.Domain, error) {
	rec := FromDomain(nurseDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *NursesRepository) UploadAvatar(ctx context.Context, id string, nurseDomain *nursesEntity.Domain) (*nursesEntity.Domain, error) {
	rec := FromDomain(nurseDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}

	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &nursesEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil
}

func (r *NursesRepository) DeleteNurseByUuid(ctx context.Context, id string) (string, error) {
	rec := Nurses{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Nurse was Deleted", nil
}

func (r *NursesRepository) GetNurses(ctx context.Context, offset, limit int) (*[]nursesEntity.Domain, int64, error) {
	var totalData int64
	domain := []nursesEntity.Domain{}
	rec := []Nurses{}

	r.db.Find(&rec, "role = ?", "nurse").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec, "role = ?", "nurse").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}

func (r *NursesRepository) GetByName(ctx context.Context, name string, offset, limit int) ([]nursesEntity.Domain, int64, error) {
	var totalData int64
	domain := []nursesEntity.Domain{}
	rec := []Nurses{}

	r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return domain, totalData, nil
}

func (r *NursesRepository) GetByID(ctx context.Context, id uint) (nursesEntity.Domain, error) {
	rec := Nurses{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *NursesRepository) GetByNikByQuery(ctx context.Context, nik string, offset, limit int) ([]nursesEntity.Domain, int64, error) {
	var totalData int64
	domain := []nursesEntity.Domain{}
	rec := []Nurses{}

	r.db.Find(&rec, "nik LIKE ?", "%"+nik+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec, "nik LIKE ?", "%"+nik+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return domain, totalData, nil
}

func (r *NursesRepository) GetByDay(ctx context.Context, day string, offset, limit int) ([]nursesEntity.Domain, int64, error) {
	var totalData int64
	domain := []nursesEntity.Domain{}
	rec := []Nurses{}

	r.db.Find(&rec, "day LIKE ?", "%"+day+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("WorkDay").Joins("WorkHour").Find(&rec, "day LIKE ?", "%"+day+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].WorkDay = rec[i].WorkDay.Day
		domain[i].WorkHour = rec[i].WorkHour.Hour
	}

	return domain, totalData, nil
}
