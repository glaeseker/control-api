package main

import "control-api/model"

type Repository interface {
	CheckTable()
	Create(id int, number string, name string) int
	Retrieve(id int) (*model.UserRecord, error)
	Update(id int, number string, name string) int
	Delete(id int) int
}
