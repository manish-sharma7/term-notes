package main

import (
	"TermNotes/pkg/config"
	"TermNotes/pkg/note"
	"flag"
	"fmt"
	"io"
	"os"
)

func PrintHelp() {
	// Open the helper file
	fileReader, err := os.Open(helperFile)
	if err != nil {
		fmt.Println("Error opening ", helperFile, " : ", err)
		return
	}

	// Read the helper file
	data, err := io.ReadAll(fileReader)
	if err != nil {
		fmt.Println("Error reading ", helperFile, " : ", err)
		return
	}

	// Print the content
	fmt.Println(string(data))
}

func main() {
	// Parse input flags
	create := flag.Bool("c", false, "Create a note")
	list := flag.Bool("l", false, "List all notes")
	get := flag.String("g", "", "Get information for a note")
	update := flag.Bool("u", false, "Update a note")
	del := flag.String("d", "", "Delete a note")
	help := flag.Bool("h", false, "Display help message")
	flag.Parse()

	// Handle help flag
	if *help {
		PrintHelp()
		return
	}

	// Validate and handle remaining arguments
	if flag.NArg() != 0 {
		fmt.Println("Invalid number of arguments")
		PrintHelp()
		return
	}

	// Get configs
	cfg := config.GetConfig()

	// Ensure notes directory exists
	// TODO: More sophisticated approach to define the permissions
	// TODO: Ensuing everytime about dir is not good idea, should be init time thing
	err := os.MkdirAll(cfg.NotesDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory ", cfg.NotesDir, " : ", err)
		return
	}

	// Init note package
	note := note.InitNote(cfg.NotesDir)

	// Handle actions based on flags
	if *create {
		note.CreateNote()
	} else if *list {
		note.ListNotes()
	} else if *get != "" {
		note.GetInfo(*get)
	} else if *update {
		note.UpdateNote()
	} else if *del != "" {
		note.DeleteNote(*del)
	} else {
		fmt.Println("Invalid flag")
		PrintHelp()
	}
}
