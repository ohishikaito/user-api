package userqueryservice

import (
	"context"

	definition "github.com/ohishikaito/user-definition"
)

type server struct {
}

// func NewServer() *server {
func NewServer() definition.UserQueryServiceServer {
	return &server{}
}

func (s *server) GetByID(ctx context.Context, req *definition.GetByIDRequest) (*definition.GetByIDResponse, error) {
	return nil, nil
}
