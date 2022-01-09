package nursesRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/google/uuid"
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

func (r *NursesRepository) GetByEmail(ctx context.Context, email string) (nursesEntity.Domain, error) {
	rec := Nurses{}

	err := r.db.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
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

func (r *NursesRepository) GetNurses(ctx context.Context) (*[]nursesEntity.Domain, error) {
	var nurses []Nurses
	if err := r.db.Find(&nurses).Error; err != nil {
		return &[]nursesEntity.Domain{}, err
	}
	result := toDomainArray(nurses)
	return &result, nil
}

func (r *NursesRepository) GetByName(ctx context.Context, name string) ([]nursesEntity.Domain, error) {
	rec := []Nurses{}

	err := r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return result, nil
}

func (r *NursesRepository) GetByID(ctx context.Context, id uint) (nursesEntity.Domain, error) {
	rec := Nurses{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return nursesEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *NursesRepository) GetByNikByQuery(ctx context.Context, nik string) ([]nursesEntity.Domain, error) {
	rec := []Nurses{}

	err := r.db.Find(&rec, "nik LIKE ?", "%"+nik+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return result, nil
}
