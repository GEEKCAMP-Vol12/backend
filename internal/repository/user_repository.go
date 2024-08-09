package repository

import (
	"time"

	"github.com/taku10101/internal/domain"
	"gorm.io/gorm"
)

type User struct {
    ID        int    `gorm:"primaryKey json:"id"`
    Name      string `json:"name"`
    Icon      string `json:"icon"`
    Age       int   `json:"age"`
    Gender    string  `json gender`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *domain.User) error {
    return r.db.Save(user).Error
}

func (r *UserRepository) Delete(userID int) error {
    return r.db.Delete(&domain.User{}, userID).Error
}

func (r *UserRepository) FindByID(userID int) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
    var users []domain.User
    if err := r.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
