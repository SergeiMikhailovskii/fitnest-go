package Base

type Cookie struct {
	Name  string
	Value string
}

var AuthUserCookie = Cookie{
	Name:  "AuthUser",
	Value: "",
}
