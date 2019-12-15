package githubusernamecli

type Username struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Followers int    `json:"followers"`
	Bio       string `json:"bio"`
	Company   string `json:"company"`
}

type UserRepo interface {
	StoreUser(username Username) error
}
