package client

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const API string = `https://api.catalogopolis.xyz/v1`

func ActorsPath() string { return "actors" }

func ActorPath(id int) string { return fmt.Sprintf("actors/%d", id) }

func DoctorByActorPath(id int) string { return fmt.Sprintf("actors/%d/doctors", id) }

func DoctorsPath() string { return "doctors" }

func DoctorPath(id int) string { return fmt.Sprintf("doctors/%d", id) }

func ActorByDoctorPath(id int) string { return fmt.Sprintf("doctors/%d/actors", id) }

func Get(pathname string, result interface{}) error {
	url := fmt.Sprintf("%s/%s", API, pathname)
	_, body, errs := gorequest.New().Get(url).EndBytes()

	if errs != nil {
		return errs[0]
	}

	return json.Unmarshal(body, result)
}
