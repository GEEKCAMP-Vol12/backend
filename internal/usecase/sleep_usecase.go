package usecase

import "github.com/taku10101/internal/domain"


type SleepRepository interface {
    Create(sleep *domain.Sleep) error
    Update(sleep *domain.Sleep) error
    Delete(sleepID string) error
    FindByID(sleepID string) (*domain.Sleep, error)
    FindAll(userID string) ([]domain.Sleep, error)
    GetScore(userID string) (int, error)
    FindLast7Days(userID string) ([]domain.Sleep, error)
}

type SleepUseCase struct {
    sleepRepository SleepRepository
}

func NewSleepUseCase(sleepRepository SleepRepository) *SleepUseCase {
    return &SleepUseCase{sleepRepository: sleepRepository}
}

func (u *SleepUseCase) CreateSleep(sleep *domain.Sleep) error {
    return u.sleepRepository.Create(sleep)
}

func (u *SleepUseCase) UpdateSleep(sleep *domain.Sleep) error {
    return u.sleepRepository.Update(sleep)
}

func (u *SleepUseCase) DeleteSleep(sleepID string) error {
    return u.sleepRepository.Delete(sleepID)
}

func (u *SleepUseCase) GetSleepByID(sleepID string) (*domain.Sleep, error) {
    return u.sleepRepository.FindByID(sleepID)
}

func (u *SleepUseCase) GetSleepsByUserID(userID string) ([]domain.Sleep, error) {
    return u.sleepRepository.FindAll(userID)
}

func (u *SleepUseCase) GetSleepScore(userID string) (int, error) {
    return u.sleepRepository.GetScore(userID)
}
func (u *SleepUseCase) GetSleepsLast7Days(userID string) ([]domain.Sleep, error) {
    return u.sleepRepository.FindLast7Days(userID)
}