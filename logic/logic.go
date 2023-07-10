package logic

import (
	"myapp/entity"
	"myapp/repository"
)

type PersonLogic struct {
	PersonRepository repository.PersonRepository
}

func (l *PersonLogic) GetPersons() ([]entity.Person, error) {
	return l.PersonRepository.GetPersons()
}
func (l *PersonLogic) GetPerson(id int) (*entity.Person, error) {
	return l.PersonRepository.GetPerson(id)
}

func (l *PersonLogic) CreatePerson(person *entity.Person) (*entity.Person, error) {
	return l.PersonRepository.CreatePerson(person)
}

func (l *PersonLogic) UpdatePerson(id int, person *entity.Person) (*entity.Person, error) {
	return l.PersonRepository.UpdatePerson(id, person)
}

func (l *PersonLogic) DeletePerson(id int) error {
	return l.PersonRepository.DeletePerson(id)
}
