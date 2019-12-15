package main

import (
  githubusernamecli "github.com/mrcarlosdev/github-profile/internal"
  "github.com/mrcarlosdev/github-profile/internal/cli"
  "github.com/mrcarlosdev/github-profile/internal/storage/csv"
  "github.com/spf13/cobra"
)

func main() {

  var repo githubusernamecli.UserRepo
  repo = csv.NewRepository()

  rootCmd := &cobra.Command{}

  rootCmd.AddCommand(cli.InitUsersCmd(repo))
  rootCmd.Execute()
}
