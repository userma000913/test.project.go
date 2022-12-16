package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz_demo/dao"
)

func Test(ctx context.Context, c *app.RequestContext) {
	err := dao.CreateTestDemo()
	if err != nil {
		c.JSON(consts.StatusOK, "failed")
		return
	}
	c.JSON(consts.StatusOK, "success")
	return
}
