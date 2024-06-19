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
		"Welcome to term-notes, Smart way to manage your notes/commands in terminal... \n",
		"Usage: notes [MODE] OPTION \n",
		"Modes: \n",
		"  --normal    ", normalModeDesc, newLine,
		"  --command   ", commandModeDesc, newLine,
		"  --all       ", allModesDesc, newLine,
		"Option: \n",
		"  Create: \n",
		"      -c      ", createOptionDesc, newLine,
		"      -u      ", updateOptionDesc, newLine,
		"  Get: \n",
		"      -l      ", listOptionDesc, newLine,
		"      -li     ", listInfoOptionDesc, newLine,
		"      -g      ", getOptionDesc, newLine,
		"  Delete: \n",
		"      -d      ", deleteOptionDesc, newLine,
		"      -da     ", deleteAllOptionDesc, newLine,
		"  Help: \n",
		"      -h      ", helpOptionDesc, newLine,
	)
}

func main() {
	// Define usage function
	flag.Usage = PrintHelp

	// Parse mode flag
	normalMode := flag.Bool("normal", true, normalModeDesc)
	commandMode := flag.Bool("command", false, normalModeDesc)
	allModes := flag.Bool("all", false, allModesDesc)

	// Parse create flags
	create := flag.Bool("c", false, createOptionDesc)
	update := flag.Bool("u", false, updateOptionDesc)
	// Parse get flag
	list := flag.Bool("l", false, listOptionDesc)
	listInfo := flag.Bool("li", false, listInfoOptionDesc)
	get := flag.Bool("g", false, getOptionDesc)
	// Parse delete flag
	del := flag.Bool("d", false, deleteOptionDesc)
	delAll := flag.Bool("da", false, deleteAllOptionDesc)
	// Parse help flag
	help := flag.Bool("h", false, helpOptionDesc)

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

	if *normalMode {
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
	} else if *commandMode {

	} else if *allModes {

	} else {
		fmt.Println("Invalid mode")
	}
}
