package webbase

type UserCtx struct {
	UserId   string `json:"UID"`
	Username string `json:"UNM"`
	Phone    string `json:"PHO"`
}

const (
	LoginStatusKey = "login.status"
	UserLoginKey   = "login.user"
)

const (
	KUserLogin = 1
)
