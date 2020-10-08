package domain

type ID string

var NilID ID

func (id *ID) String() string  {
	return string(*id)
}

func (id *ID) IsNil() bool  {
	return *id == NilID
}