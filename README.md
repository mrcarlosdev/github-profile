# GitHub Profile
My "Hello World" project in Go. Given a Github username it retrieves the account information and put it into a CSV file.

## Usage
In the command line, type the following command:
```
go run main.go github -u mrcarlosdev
```

It will create a file into `data/mrcarlosdev.csv`.

The fill will include some interesting information such as:
```
Login
Id
Name
Url
Followers
Bio
Company
```
