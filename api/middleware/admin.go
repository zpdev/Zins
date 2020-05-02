package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
)

func AdminAuth(ctx iris.Context) {
	var token model.Token
	tokenID := ctx.GetHeader("X-User-Token")
	// TODO: add cache here
	if extend.DB().Where("Token = ?", tokenID).First(&token).RecordNotFound() {
		_, _ = ctx.JSON(jsfmt.ErrorResponse(errutils.InvaildToken()))
		ctx.StopExecution()
		return
	}
	ctx.Next()

}
