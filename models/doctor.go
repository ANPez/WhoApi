package models

import (
	"errors"

	"github.com/ANPez/whoapi/client"
)

type Doctor struct {
	ID             int    `json:"id"`
	Incarnation    string `json:"incarnation"`
	PrimaryActorID int    `json:"primaryActorID"`
	Links          []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

func GetDoctors() ([]Doctor, error) {
	var doctors []Doctor
	err := client.Get(client.DoctorsPath(), &doctors)
	return doctors, err
}

func (doctor *Doctor) Get(id int) error {
	return client.Get(client.DoctorPath(id), doctor)
}

func (doctor *Doctor) GetActor() (*Actor, error) {
	if doctor.ID == 0 {
		return nil, errors.New("Call Get to retrieve actor")
	}

	actor := new(Actor)

	error := client.Get(client.ActorPath(doctor.PrimaryActorID), &doctor)

	return actor, error
}
