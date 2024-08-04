package repository

import (
	"time"

	"github.com/taku10101/internal/domain"
	"gorm.io/gorm"
)

type Sleep struct {
    ID        int 
    Score     int
    UserID    int
    CreatedAt time.Time
    UpdatedAt time.Time
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

func (r *SleepRepository) Delete(sleepID int) error {
    return r.db.Delete(&domain.Sleep{}, sleepID).Error
}

func (r *SleepRepository) FindByID(sleepID int) (*domain.Sleep, error) {
    var sleep domain.Sleep
    if err := r.db.First(&sleep, sleepID).Error; err != nil {
        return nil, err
    }
    return &sleep, nil
}

func (r *SleepRepository) FindAll() ([]domain.Sleep, error) {
    var sleeps []domain.Sleep
    if err := r.db.Find(&sleeps).Error; err != nil {
        return nil, err
    }
    return sleeps, nil
}

func (r *SleepRepository) GetScore(userID int) (int, error) {
    var sum int
    if err := r.db.Model(&domain.Sleep{}).Where("user_id = ?", userID).Select("sum(score)").Scan(&sum).Error; err != nil {
        return 0, err
    }
    return sum, nil
}
