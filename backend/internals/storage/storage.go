package storage

import "github.com/Akshansh-29072005/AARCS-X/backend/internals/types"

type Storage interface{
	//This method returns Id and error!
	CreateStudent(name string, email string, age int) (int64, error)
	//This method returns student info and error!
	GetStudentById(id int64) (types.Student, error)
	//This method returns student info in a slice and error!
	GetStudents() ([]types.Student, error)
	//This method deletes a student from the db and returns the 
	// DeleteStudent(id int64) (types.Student, error)
}