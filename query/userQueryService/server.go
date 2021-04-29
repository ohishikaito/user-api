package userqueryservice

import (
	"context"

	definition "github.com/ohishikaito/user-definition"
)

type server struct {
	uc UseCase
}

func NewServer(uc UseCase) definition.UserQueryServiceServer {
	return &server{uc}
}

func (s *server) GetByID(ctx context.Context, req *definition.GetByIDRequest) (*definition.GetByIDResponse, error) {
	s.uc.GetByID()
	return nil, nil
}
