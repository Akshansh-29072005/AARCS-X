package storage

type Storage interface{
	//This method returns Id and error!
	CreateStudent(name string, email string, age int) (int64, error)
}