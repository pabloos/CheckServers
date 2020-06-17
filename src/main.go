package main

import (
	"fmt"
)

func main() {
	serverDependencies := Servers{
		{
			Name: "MongoDB Server",
			URL:  "mongo",
			Port: 27017,
		},
		{
			Name: "MySQL Server",
			URL:  "mysql",
			Port: 3306,
		},
	}

	CheckServers(serverDependencies)

	fmt.Println("Servers reached!")
}
