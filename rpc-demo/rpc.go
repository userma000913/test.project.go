package rpc_demo

import "errors"

type DemoService struct {
}

type Arg struct {
	A, B int
}

func (DemoService) Div(arg Arg, res *float64) error {
	if arg.B == 0 {
		return errors.New("division by zero")
	}
	*res = float64(arg.A) / float64(arg.B)
	return nil

}
