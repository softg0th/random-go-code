package storage

type Storer interface {
	InsertPost(mongo *DataBase, pd PostDocument) error
	GetPost(mongo *DataBase, fs FieldSearch) ([]PostDocument, error)
	GrepPosts(mongo *DataBase) ([]PostDocument, error)
	UpdatePost(mongo *DataBase, fu FieldUpdate) error
	DeletePost(mongo *DataBase, sfs StrictFieldSearch) error
}
