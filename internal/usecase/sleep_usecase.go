package usecase

import (
	"github.com/google/uuid"
	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/repository"
)

type SleepDomain interface {
    Create(sleep *repository.Sleep) error
    Update(sleep *repository.Sleep) error
    Delete(sleepID string) error
    FindByID(sleepID string) (*repository.Sleep, error)
    FindAll(userID string) ([]repository.Sleep, error)
    GetScore(userID string) (int, error)
    FindLast7Days(userID string) ([]repository.Sleep, error)
}

type SleepUseCase struct {
    sleepDomain SleepDomain
}

func NewSleepUseCase(sleepDomain SleepDomain) *SleepUseCase {
    return &SleepUseCase{sleepDomain: sleepDomain}
}

func (u *SleepUseCase) CreateSleep(sleep *domain.CreateSleepRequest) error {
    dbSleep := &repository.Sleep{    
        UserID: sleep.UserID,
        Score: sleep.Score,
    }
    return u.sleepDomain.Create(dbSleep)
}

func (u *SleepUseCase) UpdateSleep(sleep *domain.UpdateSleepRequest) error {
    
    dbSleep := &repository.Sleep{
        ID:     uuid.Must(uuid.NewRandom()),
        UserID: sleep.UserID,
        Score:  sleep.Score, 
    }
    return u.sleepDomain.Update(dbSleep)
}

func (u *SleepUseCase) DeleteSleep(sleepID string) error {
    return u.sleepDomain.Delete(sleepID)
}

func (u *SleepUseCase) GetSleepByID(sleepID string) (*repository.Sleep, error) {
    return u.sleepDomain.FindByID(sleepID)
}

func (u *SleepUseCase) GetSleepsByUserID(userID string) ([]repository.Sleep, error) {
    return u.sleepDomain.FindAll(userID)
}

func (u *SleepUseCase) GetSleepScore(userID string) (int, error) {
    return u.sleepDomain.GetScore(userID)
}

func (u *SleepUseCase) GetSleepsLast7Days(userID string) ([]repository.Sleep, error) {
    return u.sleepDomain.FindLast7Days(userID)
}