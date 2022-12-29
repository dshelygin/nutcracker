package spaces

import (
	"nutcracker/domain/characters"
	"nutcracker/domain/data"
)

type Room struct {
	persons []*characters.Person
	voices chan data.Message
}

func (r *Room) Enter(person *characters.Person) chan data.Message {
	r.persons = append(r.persons, person)
	return r.voices
}

func (r *Room) Leave(person *characters.Person) {

}