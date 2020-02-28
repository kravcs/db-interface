package mongo

type MongoStore struct {
}

func (ps *MongoStore) GetEmployee(id int) string{
	return "mongo"
}
