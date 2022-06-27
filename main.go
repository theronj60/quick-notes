package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "github.com/manifoldco/promptui"
)

// look for directory /{home}/.notes/
// save dir to variable
// save file in variable dir

func checkDir(home string) bool {
	_, err := os.Stat(home + "/quick-notes")

	if os.IsNotExist(err) {
		fmt.Println("Could not find directory.")
		os.Exit(3)
	}
	fmt.Println("Found Directory.")
	return true
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if(checkDir(home)) {
		files, err := ioutil.ReadDir(home + "/quick-notes")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
		}
	}
}
