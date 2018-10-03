package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("GOUSER_DB_USER"),
		os.Getenv("GOUSER_DB_PASSWORD"),
		os.Getenv("GOUSER_DB_NAME"),
		os.Getenv("GOUSER_DB_HOST"),
		os.Getenv("GOUSER_DB_PORT"))
	a.Run(":8080")
}
