package ctrl

import (
	"net/http"

	"github.com/ANPez/whoapi/models"
	"github.com/gin-gonic/gin"
)

type RealDoctor struct {
	ActorName   string `json:"actor_name"`
	Incarnation string `json:"incarnation"`
}

func GetDoctors(c *gin.Context) {

	doctors, err := models.GetDoctors()
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	actors, err := models.GetActors()
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	actorsTable := makeActorsTable(actors)
	realDoctors := make([]RealDoctor, len(doctors))
	for i, doctor := range doctors {
		actor := actorsTable[doctor.PrimaryActorID]
		realDoctors[i] = RealDoctor{ActorName: actor.Name, Incarnation: doctor.Incarnation}
	}

	c.JSON(http.StatusOK, realDoctors)
}

func makeActorsTable(actors []models.Actor) map[int]models.Actor {
	table := make(map[int]models.Actor)
	for _, actor := range actors {
		table[actor.ID] = actor
	}

	return table
}
