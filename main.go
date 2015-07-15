package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing required argument, you must pass in a name\n")
		os.Exit(1)
	}
	name := os.Args[1]
	if strings.Contains(name, "_") {
		nameArr := strings.SplitN(name, "_", 2)
		if len(nameArr) < 2 {
			fmt.Fprintf(os.Stderr, "invalid name: %s", name)
		}
		name = nameArr[1]
	}
	person := people[name]
	if person == nil {
		fmt.Fprintf(os.Stderr, "Could not find anyone for %s\nPerhaps we are missing a name? Make a PR at github.com/cpuguy83/docker-who\nFor the official list see https://github.com/docker/docker/blob/master/pkg/namesgenerator/names-generator.go\n", name)
		os.Exit(1)
	}

	fmt.Printf("%s\n%s\nFind out more here: %s\n", person.Name, person.Info, person.Url)
}
