package userqueryservice

type (
	useCase struct {
		da DataAccessor
	}
	UseCase interface {
		GetByID() error
	}
)

func NewUseCase(da DataAccessor) UseCase {
	return &useCase{da}
}

func (u *useCase) GetByID() error {
	return nil
}
