package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var roleName string
var withMain bool

// var dirs := [5]string{"files", "handlers", "tasks", "templates", "meta"}
var dirNames []string

func init() {
	dirNames = []string{"files", "handlers", "tasks", "templates", "meta", "vars"}
	flag.StringVar(&roleName, "role", "", "Role name to be generated.")
	flag.StringVar(&roleName, "r", "", "Role name to be generated (shorthand).")

	flag.BoolVar(&withMain, "main", false, "Generate with main.yml in each directory.")
	flag.BoolVar(&withMain, "m", false, "Generate with main.yml in each directory (shorthand).")
	flag.Parse()
}

func main() {
	fmt.Printf("%s\n", roleName)
	if roleName == "" {
		fmt.Println("Role name must be given.")
		os.Exit(1)
	}

	for i := range dirNames {
		roleName = strings.ToLower(roleName)
		path := fmt.Sprintf("roles/%s/%s", roleName, dirNames[i])
		os.MkdirAll(path, 0775)

		if withMain {
			if dirNames[i] == "handlers" || dirNames[i] == "tasks" || dirNames[i] == "vars" {
				_, err := os.Create(path + "/main.yml")
				if err != nil {
					panic("Create file error: " + path + "/main.yml")
				}
			}
		}
	}
}
