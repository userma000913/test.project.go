package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Test(ctx context.Context, c *app.RequestContext) {

	err := svc.Test()
	if err != nil {
		c.JSON(consts.StatusOK, "failed")
		return
	}
	c.JSON(consts.StatusOK, "success")
	return
}
