package repository

import (
	"time"

	"github.com/taku10101/internal/domain"
	"gorm.io/gorm"
)

type Sleep struct {
    ID        string `gorm:"primaryKey" uuid:"true"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SleepRepository struct {
    db *gorm.DB
}

func NewSleepRepository(db *gorm.DB) *SleepRepository {
    return &SleepRepository{db: db}
}

func (r *SleepRepository) Create(sleep *domain.Sleep) error {
    return r.db.Create(sleep).Error
}

func (r *SleepRepository) Update(sleep *domain.Sleep) error {
    return r.db.Save(sleep).Error
}

func (r *SleepRepository) Delete(sleepID string) error {
    return r.db.Delete(&domain.Sleep{}, sleepID).Error
}

func (r *SleepRepository) FindByID(sleepID string) (*domain.Sleep, error) {
    var sleep domain.Sleep
    if err := r.db.First(&sleep, sleepID).Error; err != nil {
        return nil, err
    }
    return &sleep, nil
}

func (r *SleepRepository) FindAll(userID string) ([]domain.Sleep, error) {
    var sleeps []domain.Sleep
    if err := r.db.Where("user_id = ?", userID).Find(&sleeps).Error; err != nil {
        return nil, err
    }
    return sleeps, nil
}

func (r *SleepRepository) GetScore(userID string) (int, error) {
    var sum int
    if err := r.db.Model(&domain.Sleep{}).Where("user_id = ?", userID).Select("sum(score)").Scan(&sum).Error; err != nil {
        return 0, err
    }
    return sum, nil
}
func (r *SleepRepository) FindLast7Days(userID string) ([]domain.Sleep, error) {
    var sleeps []domain.Sleep
    oneWeekAgo := time.Now().AddDate(0, 0, -7)
    if err := r.db.Where("user_id = ? AND created_at >= ?", userID, oneWeekAgo).Order("created_at desc").Find(&sleeps).Error; err != nil {
        return nil, err
    }
    return sleeps, nil
}