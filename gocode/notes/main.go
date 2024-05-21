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
		"Usage: notes [-c] [-l] [-g] [-u] [-d] [-da] [-h] \n",
		"-c      Create a note \n",
		"-l      List all notes \n",
		"-li     List all notes with info\n",
		"-g      Get info of note \n",
		"-u      Update a note \n",
		"-d      Delete a note \n",
		"-da     Delete all notes \n",
		"-h      Display this help message\n",
	)
}

func main() {
	// Parse input flags
	create := flag.Bool("c", false, "Create a note")
	list := flag.Bool("l", false, "List all notes")
	listInfo := flag.Bool("li", false, "List all notes with info")
	get := flag.Bool("g", false, "Get information for a note")
	update := flag.Bool("u", false, "Update a note")
	del := flag.Bool("d", false, "Delete a note")
	delAll := flag.Bool("da", false, "Delete all notes")
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
		fmt.Println("Error creating directory", cfg.NotesDir, ":", err)
		return
	}

	// Init note package
	note := note.InitNote(cfg.NotesDir)

	// Handle actions based on flags
	if *create {
		note.CreateNote()
	} else if *list {
		// Pass include info flag false
		note.ListNotes(false)
	} else if *listInfo {
		// Pass include info flag true
		note.ListNotes(true)
	} else if *get {
		note.GetInfo()
	} else if *update {
		note.UpdateNote()
	} else if *del {
		note.DeleteNote()
	} else if *delAll {
		note.DeleteNoteFile()
	} else {
		fmt.Println("Invalid flag")
		PrintHelp()
	}
}
