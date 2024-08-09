package repository

import (
	"time"

	"github.com/taku10101/internal/domain"
	"gorm.io/gorm"
)

type Caffeine struct {
    ID        int `gorm:"primaryKey"`
    Score     int
    UserID    int
    CreatedAt time.Time
    UpdatedAt time.Time
}

type CaffeineRepository struct {
    db *gorm.DB
}

func NewCaffeineRepository(db *gorm.DB) *CaffeineRepository {
    return &CaffeineRepository{db: db}
}

func (r *CaffeineRepository) Create(caffeine *domain.Caffeine) error {
    return r.db.Create(caffeine).Error
}

func (r *CaffeineRepository) Update(caffeine *domain.Caffeine) error {
    return r.db.Save(caffeine).Error
}

func (r *CaffeineRepository) Delete(caffeineID int) error {
    return r.db.Delete(&domain.Caffeine{}, caffeineID).Error
}

func (r *CaffeineRepository) FindByID(caffeineID int) (*domain.Caffeine, error) {
    var caffeine domain.Caffeine
    if err := r.db.First(&caffeine, caffeineID).Error; err != nil {
        return nil, err
    }
    return &caffeine, nil
}

func (r *CaffeineRepository) FindAll(userID int) ([]domain.Caffeine, error) {
    var caffeines []domain.Caffeine
    if err := r.db.Where("user_id = ?", userID).Find(&caffeines).Error; err != nil {
        return nil, err
    }
    return caffeines, nil
}

func (r *CaffeineRepository) GetScore(userID int) (int, error) {
    var sum int
    if err := r.db.Model(&domain.Caffeine{}).Where("user_id = ?", userID).Select("sum(score)").Scan(&sum).Error; err != nil {
        return 0, err
    }
    return sum, nil
}
func (r *CaffeineRepository) FindLast7Days(userID int) ([]domain.Caffeine, error) {
    var caffeines []domain.Caffeine
    oneWeekAgo := time.Now().AddDate(0, 0, -7)
    if err := r.db.Where("user_id = ? AND created_at >= ?", userID, oneWeekAgo).Order("created_at desc").Find(&caffeines).Error; err != nil {
        return nil, err
    }
    return caffeines, nil
}