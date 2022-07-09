package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	// "github.com/manifoldco/promptui"
)

// look for directory /{home}/.notes/
// save dir to variable
// save file in variable dir

// qnote "string" => adds one line to a default file

func checkDir(home string) bool {
	_, err := os.Stat(home + "/quick-notes") // check for quick-notes path

	if os.IsNotExist(err) {
		fmt.Println("Could not find directory.")
		fmt.Println("Creating ~/quick-notes")
		os.Mkdir(home + "/quick-notes", os.ModePerm)
	}
	fmt.Println("Found Directory.")
	return true
}

func createDir(path string) {

}

func createFile(path string) {

}

func main() {
	fmt.Println("Welcome to QuickNotes. Enter a string to write to a file.")
	var userInput string
	fmt.Scanln(&userInput)
	// io.WriteString(file, content) <- alternative?

	home, err := os.UserHomeDir() // get home dir
	if err != nil {
		log.Fatal(err)
	}
	if checkDir(home) {
		files, err := ioutil.ReadDir(home + "/quick-notes")
		if err != nil {
			log.Fatal(err)
		}

		fileDefault := "default.md"

		for _, file := range files {
			if file.Name() == fileDefault {
				// write to file
				fmt.Println("File found!")
				d1 := []byte(userInput)
				os.WriteFile(home + "/quick-notes/" + fileDefault, d1, os.ModePerm)
				fmt.Println("Wrote to default.md")
				return
			} else {
				fmt.Println("File not found!")
				os.Create(home + "/quick-notes/" + fileDefault)
				// fmt.Println(file.Name(), file.IsDir())
				return
				// file doesnt exists
				// create file
				// write to file
			}
		}
	}
}
