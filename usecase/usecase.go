package usecase

import (
	"FirstGolangProject/model"
	"github.com/pkg/errors"
)

func SqrtMain(par int) (*model.Resp, error) {
	sqrt := par * par
	res := model.Resp{
		Result: sqrt,
	}
	if sqrt > 1000 {
		err := errors.New("so long")
		return nil, errors.Wrap(err, "[SqrtMain] Error #1")
	}

	return &res, nil
}
