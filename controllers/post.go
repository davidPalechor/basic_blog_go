package controllers

import (
	"basic_blog_go/logic"
	"basic_blog_go/viewmodels"
	"encoding/json"

	"github.com/astaxie/beego"
)

// PostController controls api requests
type PostController struct {
	beego.Controller
}

// Post method to create a new entry blog
// @Title Analysis
// @Summary Analyze and clasify feelings based on text
// @Description Analyze and clasify feelings based on text
// @Param   body	body	viewmodels.PostRequest	true	"Trending Topic category"
// @Success 200 {string} "Success"
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [post]
func (c *PostController) Post() {
	var request viewmodels.PostRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	response := map[string]string{
		"status": "ok",
	}

	logic := logic.NewPostLogic()

	if err := logic.CreatePost(request); err != nil {
		response["status"] = "failed"
	}
	c.Data["json"] = response
	c.ServeJSON()
}
