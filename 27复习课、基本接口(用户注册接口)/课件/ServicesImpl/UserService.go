package ServicesImpl

import (
	"context"
	"tools.jtthink.com/Services"
)

type UserService struct {
}

func (this *UserService) UserReg(ctx context.Context, user *Services.UserModel, rsp *Services.RegReponse) error{
	rsp.Message=""
	rsp.Status="success"
	return nil
}
