package medicalStaffEntity_test

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	mockUsersRepository  mocks.Repository
	medicalStaffServices medicalStaffEntity.Service
	hashedPassword       string
	medicalStaffDomain   medicalStaffEntity.Domain
	usersHashDomain      medicalStaffEntity.Domain
)

func TestMain(m *testing.M) {
	medicalStaffServices = medicalStaffEntity.NewMedicalStaffServices(&mockUsersRepository, &auth.ConfigJWT{}, time.Second*2)
	medicalStaffDomain = medicalStaffEntity.Domain{
		Nik:            "1301190230",
		Name:           "Testing User",
		Email:          "testinguser@gmail.com",
		Dob:            "01-10-2001",
		Sex:            "Male",
		Contact:        "081278486869",
		Password:       "testing password",
		Role:           "medical_staff",
		WorkExperience: "10 Tahun",
		Avatar:         "images/avatar/avatar.png",
		Token:          "asiudgasigdaiusgdiausgdi",
	}

	usersHashDomain = medicalStaffEntity.Domain{
		Nik:            "1301190230",
		Name:           "Testing User",
		Email:          "testinguser@gmail.com",
		Dob:            "01-10-2001",
		Sex:            "Male",
		Contact:        "081278486869",
		Password:       "testing password",
		Role:           "medical_staff",
		WorkExperience: "10 Tahun",
		Avatar:         "images/avatar/avatar.png",
		Token:          "asiudgasigdaiusgdiausgdi",
	}
	m.Run()
}

func TestMedicalStaffServices_Register(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUsersRepository.On("GetByNik", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("CreateNewMedicalStaff", mock.Anything, mock.Anything).Return(&medicalStaffDomain, nil).Once()

		req := &medicalStaffEntity.Domain{
			Nik:            "1301190230",
			Name:           "Testing User",
			Email:          "testinguser@gmail.com",
			Dob:            "01-10-2001",
			Sex:            "Male",
			Contact:        "081278486869",
			Password:       "testing password",
			Role:           "medical_staff",
			WorkExperience: "10 Tahun",
		}

		res, err := medicalStaffServices.Register(context.Background(), req)

		assert.Nil(t, err)
		assert.Equal(t, medicalStaffDomain, *res)
	})
	t.Run("Invalid Test || duplicate email", func(t *testing.T) {
		mockUsersRepository.On("GetByNik", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffDomain, nil).Once()
		mockUsersRepository.On("CreateNewMedicalStaff", mock.Anything, mock.Anything).Return(&medicalStaffEntity.Domain{}, assert.AnError).Once()

		req := &medicalStaffEntity.Domain{
			Nik:            "1301190230",
			Name:           "Testing User",
			Email:          "testinguser@gmail.com",
			Dob:            "01-10-2001",
			Sex:            "Male",
			Contact:        "081278486869",
			Password:       "testing password",
			Role:           "medical_staff",
			WorkExperience: "10 Tahun",
		}

		res, err := medicalStaffServices.Register(context.Background(), req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, medicalStaffDomain.Email, req.Email)
	})
	t.Run("Invalid Test || duplicate nik", func(t *testing.T) {
		mockUsersRepository.On("GetByNik", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffDomain, nil).Once()
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("CreateNewMedicalStaff", mock.Anything, mock.Anything).Return(&medicalStaffEntity.Domain{}, assert.AnError).Once()

		req := &medicalStaffEntity.Domain{
			Nik:            "1301190230",
			Name:           "Testing User",
			Email:          "testinguser@gmail.com",
			Dob:            "01-10-2001",
			Sex:            "Male",
			Contact:        "081278486869",
			Password:       "testing password",
			Role:           "medical_staff",
			WorkExperience: "10 Tahun",
		}

		res, err := medicalStaffServices.Register(context.Background(), req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, medicalStaffDomain.Email, req.Email)
	})
	t.Run("Invalid Test || err not found", func(t *testing.T) {
		mockUsersRepository.On("GetByNik", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, assert.AnError).Once()
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, assert.AnError).Once()
		mockUsersRepository.On("CreateNewMedicalStaff", mock.Anything, mock.Anything).Return(&medicalStaffEntity.Domain{}, assert.AnError).Once()

		req := &medicalStaffEntity.Domain{
			Nik:            "1301190230",
			Name:           "Testing User",
			Email:          "testinguser@gmail.com",
			Dob:            "01-10-2001",
			Sex:            "Male",
			Contact:        "081278486869",
			Password:       "testing password",
			Role:           "medical_staff",
			WorkExperience: "10 Tahun",
		}

		res, err := medicalStaffServices.Register(context.Background(), req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.NotEqual(t, medicalStaffDomain, req)
	})
	t.Run("Invalid Test || internal error", func(t *testing.T) {
		mockUsersRepository.On("GetByNik", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, nil).Once()
		mockUsersRepository.On("CreateNewMedicalStaff", mock.Anything, mock.Anything).Return(&medicalStaffEntity.Domain{}, assert.AnError).Once()

		req := &medicalStaffEntity.Domain{
			Nik:            "1301190230",
			Name:           "Testing User",
			Email:          "testinguser@gmail.com",
			Dob:            "01-10-2001",
			Sex:            "Male",
			Contact:        "081278486869",
			Password:       "testing password",
			Role:           "medical_staff",
			WorkExperience: "10 Tahun",
		}

		res, err := medicalStaffServices.Register(context.Background(), req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.NotEqual(t, medicalStaffDomain, req)
	})
}

func TestMedicalStaffServices_Login(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.Anything, mock.Anything).Return(medicalStaffDomain, nil).Once()

		req := medicalStaffEntity.Domain{
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
		}

		_, err := medicalStaffServices.Login(context.Background(),req.Email,req.Password)
		token := "asiudgasigdaiusgdiausgdi"

		assert.Equal(t, token,usersHashDomain.Token)
		assert.Equal(t, req.Password,usersHashDomain.Password)
		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | password and email null", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(medicalStaffDomain, nil).Once()

		req := medicalStaffEntity.Domain{
			Email       : "",
			Password    : "",
		}

		token := ""
		_, err := medicalStaffServices.Login(context.Background(),req.Email,req.Password)

		assert.NotEqual(t, token,usersHashDomain.Token)
		assert.NotEqual(t, req.Password,usersHashDomain.Password)
		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | email unregistered", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(medicalStaffEntity.Domain{}, assert.AnError).Once()

		req := medicalStaffEntity.Domain{
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
		}

		_, err := medicalStaffServices.Login(context.Background(),req.Email,req.Password)

		assert.NotEqual(t, req,medicalStaffEntity.Domain{})
		assert.NotNil(t, err)
	})
}