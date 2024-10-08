package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sleep struct {
    ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
    Score     int       `json:"score"`
    UserID    string    `json:"user_id"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    User User 
}


type SleepRepository struct {
    db *gorm.DB
}

func NewSleepRepository(db *gorm.DB) *SleepRepository {
    return &SleepRepository{db: db}
}

func (r *SleepRepository) Create(sleep *Sleep) error {
    if sleep.Score == 0 {
        return errors.New("score cannot be zero")
    }   
     if sleep.ID == uuid.Nil {
        sleep.ID = uuid.New()
    }

    err := r.db.Create(sleep).Error
    if err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return errors.New("duplicate key value violates unique constraint")
        }
        return err
    }
    return nil
}

func (r *SleepRepository) Update(sleep *Sleep) error {
    if sleep.Score == 0 {
        return errors.New("score cannot be zero")
    }

    return r.db.Save(sleep).Error
}

func (r *SleepRepository) Delete(sleepID string) error {
    return r.db.Delete(Sleep{}, sleepID).Error
}

func (r *SleepRepository) FindByID(sleepID string) (*Sleep, error) {
    var sleep Sleep
    if err := r.db.First(&sleep, sleepID).Error; err != nil {
        return nil, err
    }
    return &sleep, nil
}

func (r *SleepRepository) FindAll(userID string) ([]Sleep, error) {
    var sleeps []Sleep    
    if err := r.db.Where("user_id = ?", userID).Find(&sleeps).Error; err != nil {
        return nil, err
    }
    return sleeps, nil
}

func (r *SleepRepository) GetScore(userID string) (int, error) {
    var sum int
    if err := r.db.Model(Sleep{}).Where("user_id = ?", userID).Select("sum(score)").Scan(&sum).Error; err != nil {
        return 0, err
    }
    return sum, nil
}
func (r *SleepRepository) FindLast7Days(userID string) ([]Sleep, error) {
    var sleeps []Sleep
    oneWeekAgo := time.Now().AddDate(0, 0, -7)
    if err := r.db.Where("user_id = ? AND created_at >= ?", userID, oneWeekAgo).Order("created_at desc").Find(&sleeps).Error; err != nil {
        return nil, err
    }
    return sleeps, nil
}