package storage

import "github.com/Akshansh-29072005/AARCS-X/backend/internals/types"

type Storage interface{
	//This method returns Id and error!
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}