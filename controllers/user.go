package controllers

import (
	"basic_blog_go/logic"
	"basic_blog_go/viewmodels"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

// UserController defines the controller asociated to user requests
type UserController struct {
	beego.Controller
}

// CreateUser method to create a new user
// @Title Create user
// @Summary Create an user in the app
// @Description Create an user in the app
// @Param   body	body	viewmodels.UserRequest	true	"User info to be stored"
// @Success 200 {string} "Success"
// @Failure 400 Bad request
// @Accept json
// @router / [post]
func (c *UserController) CreateUser() {
	var request viewmodels.UserRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{
			"status": "bad_request",
		}
		c.ServeJSON()
		return
	}

	logic := logic.NewUserLogic()
	if err := logic.CreateUser(request); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{
			"status": err.Error(),
		}
	}
	c.Data["json"] = map[string]string{
		"status": "ok",
	}
	c.ServeJSON()
}

// Login method to return a token to use endpoints
// @Title Logs in an user in the app
// @Summary Returns a valid JWT to use endpoints
// @Description Returns a valid JWT to use endpoints
// @Param   body	body	viewmodels.LoginRequest	true	"User info to be stored"
// @Success 200 {object} viewmodels.LoginResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /login [post]
func (c *UserController) Login() {
	var request viewmodels.LoginRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	c.Ctx.Output.SetStatus(http.StatusOK)

	logic := logic.NewUserLogic()
	response, err := logic.LoginUser(request)
	c.Data["json"] = response
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{
			"error": err.Error(),
		}
	}
	c.ServeJSON()
	return
}
