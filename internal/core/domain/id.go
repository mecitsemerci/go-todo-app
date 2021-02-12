package domain

type ID string

var NilID ID

func (id *ID) String() string  {
	return string(*id)
}

