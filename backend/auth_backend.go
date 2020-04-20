package backend

import (
	"context"
	auth "cso/auth/proto"

	"github.com/micro/go-micro/v2/client"
)

func newAuthService() auth.AuthService {
	return auth.NewAuthService("auth", client.DefaultClient)
}

func Message(ctx context.Context) (*auth.Response, error) {
	resp, err := newAuthService().Message(context.TODO(), &auth.Empty{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
