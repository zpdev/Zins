package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type User struct {
	Ctx iris.Context
}

func (c *User) Get() *jsfmt.Response {
	queryStr := c.Ctx.FormValue("query")
	query, Qerr := jsfmt.ReadQuery(queryStr)
	if Qerr != nil {
		return jsfmt.ErrorResponse(Qerr)
	}
	users, total, err := service.UserService.GetUsers(extend.DB(), query)
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.QueryResponse(users, total)
}

func (c *User) Post() *jsfmt.Response {
	// TODO: 增加用户注册开关
	user := &model.User{}
	if err := c.Ctx.ReadJSON(user); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	if err := service.UserService.CreateUser(extend.DB(), user); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	user.Password = ""
	return jsfmt.NormalResponse(user)
}

func (c *User) Delete() *jsfmt.Response {
	var users []model.User
	if err := c.Ctx.ReadJSON(&users); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	if err := service.UserService.DeleteUsers(extend.DB(), users); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(nil)
}

// UserDetail
type UserDetail struct {
	Ctx iris.Context
}

func (c *UserDetail) Get() *jsfmt.Response {
	username := c.Ctx.Params().Get("username")
	user, err := service.UserService.GetUser(extend.DB(), username)
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(user)
}

func (c *UserDetail) Delete() *jsfmt.Response {
	username := c.Ctx.Params().Get("username")
	if err := service.UserService.DeleteUser(extend.DB(), username); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(nil)
}

func (c *UserDetail) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user_detail_put": "not implement"})
}
