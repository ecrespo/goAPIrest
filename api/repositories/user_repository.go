package repository

import (
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uid uint32) (*models.User, error)
	Update(user *models.User, uid uint32) (*models.User, error)
	Delete(uid uint32) (int64, error)
	FindByEmail(email string) (*models.User, error)
	// Add more methods as per your needs.
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (r *userRepositoryImpl) Save(user *models.User) (*models.User, error) {
	err := r.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByID(uid uint32) (*models.User, error) {
	var user models.User
	err := r.DB.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// The update method does not handle User password hashing.
// You might want to handle that in your service logic before passing the user to Update method.
func (r *userRepositoryImpl) Update(user *models.User, uid uint32) (*models.User, error) {
	err := r.DB.Debug().Model(&models.User{}).Where("id = ?", uid).Updates(models.User{Nickname: user.Nickname, Email: user.Email, Password: user.Password}).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) Delete(uid uint32) (int64, error) {
	db := r.DB.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
