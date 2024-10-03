package auth

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

func (r Role) IsValid() bool {
	switch r {
	case Admin, Member:
		return true
	default:
		return false
	}
}

func (r Role) String() string {
	return string(r)
}

func (r Role) Equals(other Role) bool {
	return r.String() == other.String()
}

func (r Role) IsAdmin() bool {
	return r == Admin
}
