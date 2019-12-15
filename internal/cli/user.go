package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	githubusernamecli "github.com/mrcarlosdev/github-profile/internal"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
)

const usernameFlag = "username"
const outputFlag = "output"

var usernameArg string

type CobraFn func(cmd *cobra.Command, args []string)

func InitUsersCmd(repository githubusernamecli.UserRepo) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "github",
		Short: "Provide a Gitlab username to store his information into a CSV file",
		Run:   runUserFn(repository),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			usernameArg, _ = cmd.Flags().GetString(usernameFlag)
			if usernameArg == "" {
				return errors.New("[flags] are required")
			}
			return nil
		},
	}

	userCmd.Flags().StringP(usernameFlag, "u", "", "github username")
	userCmd.Flags().StringP(outputFlag, "o", "", "github output")

	return userCmd
}

func runUserFn(repository githubusernamecli.UserRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		var username githubusernamecli.Username

		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, _ := http.NewRequest("GET", "https://api.github.com/users/" + usernameArg, nil)
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &username)
		_ = repository.StoreUser(username)
		fmt.Println(username)
		return
	}
}
