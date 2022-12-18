package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Test(ctx context.Context, c *app.RequestContext) {

	//err := svc.Test()
	name := c.Query("name")
	c.JSON(consts.StatusOK, name)
	return
}

func TestMgr(c context.Context, ctx *app.RequestContext) {

	svc.TestMgr(c, ctx)
}
