package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEpisodes(c *gin.Context) {
	var episodesList []Episode
	err := json.Unmarshal([]byte(episodes), &episodesList)

	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		var ret []Episode
		for _, episode := range episodesList {
			if episode.SerialID != nil {
				ret = append(ret, episode)
			}
		}

		c.JSON(http.StatusOK, ret)
	}
}
