package models

import (
	"errors"

	"github.com/ANPez/whoapi/client"
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

func GetActors() ([]Actor, error) {
	var actors []Actor
	err := client.Get(client.ActorsPath(), &actors)
	return actors, err
}

func (actor *Actor) Get(id int) error {
	return client.Get(client.ActorPath(id), actor)
}

func (actor *Actor) GetDoctor() (*Doctor, error) {
	if actor.ID == 0 {
		return nil, errors.New("Call Get to retrieve actor")
	}

	doctor := new(Doctor)

	error := client.Get(client.DoctorByActorPath(actor.ID), &doctor)

	return doctor, error
}
