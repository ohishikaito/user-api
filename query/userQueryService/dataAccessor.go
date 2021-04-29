package userqueryservice

type (
	dataAccessor struct{}
	DataAccessor interface {
	}
)

func NewDataAccessor() DataAccessor {
	return &dataAccessor{}
}
