package repositories

import (
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/jinzhu/gorm"
	"time"
)

type PostRepository interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
	FindByID(id uint64) (*models.Post, error)
	Update(post *models.Post) (*models.Post, error)
	Delete(id uint64) (int64, error)
	// add more methods as per your needs.
}

type postsRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postsRepositoryImpl{
		DB: db,
	}
}

func (r *postsRepositoryImpl) Save(post *models.Post) (*models.Post, error) {
	err := r.DB.Debug().Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return nil, err
	}
	if post.ID != 0 {
		err = r.DB.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return nil, err
		}
	}
	return post, nil
}

func (r *postsRepositoryImpl) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postsRepositoryImpl) FindByID(id uint64) (*models.Post, error) {
	var post models.Post
	err := r.DB.Debug().Model(&models.Post{}).Where("id = ?", id).Take(&post).Error
	if err != nil {
		return nil, err
	}
	if post.ID != 0 {
		err = r.DB.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return nil, err
		}
	}
	return &post, nil
}

func (r *postsRepositoryImpl) Update(post *models.Post) (*models.Post, error) {
	err := r.DB.Debug().Model(&models.Post{}).Where("id = ?", post.ID).Updates(models.Post{Title: post.Title, Content: post.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return nil, err
	}
	if post.ID != 0 {
		err = r.DB.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return nil, err
		}
	}
	return post, nil
}

func (r *postsRepositoryImpl) Delete(id uint64) (int64, error) {
	db := r.DB.Debug().Model(&models.Post{}).Where("id = ?", id).Take(&models.Post{}).Delete(&models.Post{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
