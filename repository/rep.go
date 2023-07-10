package repository

import (
	"github.com/gocraft/dbr"
	"myapp/entity"
)

type PersonRepository struct {
	DB *dbr.Session
}

func (r *PersonRepository) GetPersons() ([]entity.Person, error) {
	var persons []entity.Person
	query := r.DB.Select("*").From("persons")
	_, err := query.Load(&persons)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonRepository) GetPerson(id int) (*entity.Person, error) {
	var person entity.Person
	err := r.DB.Select("*").From("persons").Where("id = ?", id).LoadOne(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepository) CreatePerson(person *entity.Person) (*entity.Person, error) {
	_, err := r.DB.InsertInto("persons").
		Pair("email", person.Email).
		Pair("phone", person.Phone).
		Pair("first_name", person.FirstName).
		Exec()

	if err != nil {
		return nil, err
	}

	return person, nil
}
func (r *PersonRepository) UpdatePerson(id int, person *entity.Person) (*entity.Person, error) {
	_, err := r.DB.Update("persons").
		Set("email", person.Email).
		Set("phone", person.Phone).
		Set("first_name", person.FirstName).
		Where("id = ?", id).
		Exec()

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r *PersonRepository) DeletePerson(id int) error {
	_, err := r.DB.DeleteFrom("persons").Where("id = ?", id).Exec()
	return err
}
