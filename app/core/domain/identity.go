package domain

type ID interface {
	String() string
	Set(str string)
}
var NilID ID