package controllers

import (
	"basic_blog_go/logic"
	"net/http"

	"github.com/astaxie/beego"
)

// TrendController defines the controller associated to Twitter trends
type TrendController struct {
	beego.Controller
}

// GetTrends method to retrieve trending topics by country
// @Title Get Trends
// @Summary Retrieve trends based on filters
// @Description Retrieve trends based on filters
// @Param	header	string true "Authorization key"
// @Param   id		int 	string true "Where On Earth Identifier"
// @Param   exclude	query 	string false "Data to exclude e.g 'hashtags'"
// @Success 200 {string} "Success"
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @router / [get]
func (c *TrendController) GetTrends() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{
			"error": "id must be a number",
		}
	}

	exclude := c.GetString("exclude")
	logic := logic.NewTrendLogic()
	response, err := logic.GetTrends(id, exclude)
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

// AnalyzeTrend method to create a new user
// @Title Get Trends
// @Summary Retrieve trends based on filters
// @Description Retrieve trends based on filters
// @Param	header	string 	true 	"Authorization key"
// @Param   q		query 	string 	true "Querystring to search tweets"
// @Success 200 {object} viewmodels.Sentiment
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @router /analysis [get]
func (c *TrendController) AnalyzeTrend() {
	query := c.GetString("q")
	logic := logic.NewTrendLogic()
	response, err := logic.GetTweets(query)
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
