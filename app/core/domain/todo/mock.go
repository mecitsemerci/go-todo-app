package todo

type MockTodoRepository struct {
	MockGetAll  func() ([]*Todo, error)
	MockGetById func(id string) (*Todo, error)
	MockInsert  func(todo Todo) (string, error)
	MockUpdate  func(todo Todo) error
	MockDelete  func(id string) error
}

func (m *MockTodoRepository) GetAll() ([]*Todo, error) {
	return m.MockGetAll()
}

func (m *MockTodoRepository) GetById(id string) (*Todo, error) {
	return m.MockGetById(id)
}

func (m *MockTodoRepository) Insert(todo Todo) (string, error) {
	return m.MockInsert(todo)
}

func (m *MockTodoRepository) Update(todo Todo) error {
	return m.MockUpdate(todo)
}

func (m *MockTodoRepository) Delete(id string) error {
	return m.MockDelete(id)
}

