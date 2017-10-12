package ctrl

import (
	"net/http"

	"strconv"

	"github.com/ANPez/whoapi/models"
	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	actor := new(models.Actor)
	err := actor.Get(id)
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name": actor.Name,
	})
}
