package main

import (
	"TermNotes/pkg/config"
	"TermNotes/pkg/note"
	"flag"
	"fmt"
	"os"
)

func PrintHelp() {
	fmt.Print(
		"Welcome to term-notes, Smart way to store your notes in terminal... \n",
		"Usage: notes [-c] [-l] [-g] [-u] [-d] [-h] \n",
		"-c      Create a note \n",
		"-l      List all notes \n",
		"-g      Get info of note \n",
		"-u      Update a note \n",
		"-d      Delete a note \n",
		"-h      Display this help message\n",
	)
}

func main() {
	// Parse input flags
	create := flag.Bool("c", false, "Create a note")
	list := flag.Bool("l", false, "List all notes")
	get := flag.Bool("g", false, "Get information for a note")
	update := flag.Bool("u", false, "Update a note")
	del := flag.Bool("d", false, "Delete a note")
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
	} else if *get {
		note.GetInfo()
	} else if *update {
		note.UpdateNote()
	} else if *del {
		note.DeleteNote()
	} else {
		fmt.Println("Invalid flag")
		PrintHelp()
	}
}
