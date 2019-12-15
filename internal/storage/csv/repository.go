package csv

import (
	"encoding/csv"
	githubusernamecli "github.com/mrcarlosdev/github-profile/internal"
	"log"
	"os"
	"strconv"
)

type repository struct {}

func NewRepository() githubusernamecli.UserRepo {
	return &repository{}
}

func (r *repository) StoreUser(username githubusernamecli.Username) error {

	f, err := os.Create("internal/data/" + username.Login + ".csv")
	checkError("Cannot create file", err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	var data = []string{
		strconv.Itoa(username.Id),
		username.Login,
		username.Name,
		username.Company,
		username.Bio,
		strconv.Itoa(username.Followers),
		username.Url,
	}

	err2 := writer.Write(data)
	checkError("Cannot write to file", err2)

	return nil
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
