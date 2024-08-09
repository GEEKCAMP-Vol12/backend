package repository

import (
	"time"

	"gorm.io/gorm"
)

type Caffeine struct {
    ID        string `gorm:"primaryKey" uuid:"true"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type CaffeineRepository struct {
    db *gorm.DB
}

func NewCaffeineRepository(db *gorm.DB) *CaffeineRepository {
    return &CaffeineRepository{db: db}
}

func (r *CaffeineRepository) Create(caffeine *Caffeine) error {
    return r.db.Create(caffeine).Error
}

func (r *CaffeineRepository) Update(caffeine *Caffeine) error {
    return r.db.Save(caffeine).Error
}

func (r *CaffeineRepository) Delete(caffeineID string) error {
    return r.db.Delete(Caffeine{}, caffeineID).Error
}

func (r *CaffeineRepository) FindByID(caffeineID string) (*Caffeine, error) {
    var caffeine Caffeine
    if err := r.db.First(&caffeine, caffeineID).Error; err != nil {
        return nil, err
    }
    return &caffeine, nil
}

func (r *CaffeineRepository) FindAll(userID string) ([]Caffeine, error) {
    var caffeines []Caffeine
    if err := r.db.Where("user_id = ?", userID).Find(&caffeines).Error; err != nil {
        return nil, err
    }
    return caffeines, nil
}

func (r *CaffeineRepository) GetScore(userID string) (int, error) {
    var sum int
    if err := r.db.Model(Caffeine{}).Where("user_id = ?", userID).Select("sum(score)").Scan(&sum).Error; err != nil {
        return 0, err
    }
    return sum, nil
}
func (r *CaffeineRepository) FindLast7Days(userID string) ([]Caffeine, error) {
    var caffeines []Caffeine
    oneWeekAgo := time.Now().AddDate(0, 0, -7)
    if err := r.db.Where("user_id = ? AND created_at >= ?", userID, oneWeekAgo).Order("created_at desc").Find(&caffeines).Error; err != nil {
        return nil, err
    }
    return caffeines, nil
}