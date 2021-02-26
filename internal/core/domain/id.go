package domain

// ID generic domain identifier
type ID string

// ZeroID null domain ID
var ZeroID ID

//String returns ID's string value
func (id *ID) String() string {
	return string(*id)
}
