package main

import (
	"flag"
	"fmt"
	"demo4/application"
	"demo4/shared/driver"
)

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"demo4": application.NewDemo4(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
