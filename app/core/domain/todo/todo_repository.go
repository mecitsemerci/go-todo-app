package todo

type ITodoRepository interface {
	GetAll() ([]*Todo, error)
	GetById(id string) (*Todo, error)
	Insert(todo Todo) (string, error)
	Update(todo Todo) error
	Delete(id string) error
}
