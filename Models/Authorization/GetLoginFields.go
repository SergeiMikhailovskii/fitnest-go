package Authorization

type GetLoginFields struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
}
