package auth

type UserSession struct {
	SID    string `json:"sid"`
	IP     string `json:"ip"`
	Login  string `json:"login"`
	Logout string `json:"logout"`
	UA     string `json:"ua"`
}
