package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type Actor struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Links  []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

func GetActors(c *gin.Context) {
	id := c.Param("id")
	_, body, errs := gorequest.New().Get("https://api.catalogopolis.xyz/v1/doctors/" + id + "/actors").EndBytes()

	if nil != errs {
		c.AbortWithError(http.StatusInternalServerError, errs[0])
		return
	}

	var actors []Actor
	err := json.Unmarshal(body, &actors)
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name": actors[0].Name,
	})
}
