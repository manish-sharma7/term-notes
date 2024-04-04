package note

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Note struct {
	filePath string
	title    string
	info     string
	ops      int
}

func InitNote(noteDir string) Note {
	return Note{
		filePath: filepath.Join(noteDir, noteFile),
	}
}

func (n Note) readFile() []string {
	var lines []string
	data, err := os.ReadFile(n.filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error reading notes : ", err)
			os.Exit(1)
		} else {
			return lines
		}
	}
	lines = strings.Split(string(data), "\n")
	return lines
}

func (n Note) writeFile(lines []string) error {
	// Open the file in write mode ("w")
	// This will overwrite any existing content
	// TODO: Understand permission modes and Support to take it from user
	f, err := os.OpenFile(n.filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		fmt.Println("Can't update the note in database, try again")
		return err
	}
	defer f.Close()

	// Write the data to the file
	for _, line := range lines {
		_, err = f.Write([]byte(line))
		if err != nil {
			fmt.Println("Error writing to file : ", err, " for note : ", n.title)
			return err
		}
	}
	return nil
}

func (n Note) updateNote() {
	// Read the note file
	lines := n.readFile()

	// Get the Updated lines without current title
	var updatedLines []string
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, n.title+delimiter) {
			updatedLines = append(updatedLines, line+"\n")
		}
	}

	if n.ops == opsCreateNote || n.ops == opsUpdateNote {
		// Append the title note with newer details
		updatedLines = append(updatedLines, fmt.Sprintf("%s%s%s\n", n.title, delimiter, n.info))
	}
	// Update file
	err := n.writeFile(updatedLines)

	if n.ops == opsCreateNote {
		if err != nil {
			fmt.Println("Note creation failed")
		} else {
			fmt.Println("Note created")
		}
	} else if n.ops == opsUpdateNote {
		if err != nil {
			fmt.Println("Note update failed")
		} else {
			fmt.Println("Note updated")
		}
	} else if n.ops == opsRemoveNote {
		if err != nil {
			fmt.Println("Note deletion failed")
		} else {
			fmt.Println("Note deleted")
		}
	}
}

func (n Note) CreateNote() {
	// Prompt for title
	fmt.Println("Title >>")
	var title string
	fmt.Scan(&title)
	n.title = title

	// Prompt for information
	fmt.Println("Info >>")
	var info string
	fmt.Scan(&info)
	n.info = info

	// Update ops
	n.ops = opsCreateNote

	// Create note
	n.updateNote()
}

func (n Note) UpdateNote() {
	// Prompt for title
	fmt.Println("Title >>")
	var title string
	fmt.Scan(&title)
	n.title = title

	// Prompt for information
	fmt.Println("New Info >>")
	var info string
	fmt.Scan(&info)
	n.info = info

	// Update ops
	n.ops = opsUpdateNote

	// Create note
	n.updateNote()
}

func (n Note) DeleteNote(title string) {
	n.title = title
	// Set ops to delete
	n.ops = opsRemoveNote

	// Delete note
	n.updateNote()
}

func (n Note) GetInfo(title string) {
	// Read note file
	lines := n.readFile()
	for _, line := range lines {
		if strings.HasPrefix(line, title+delimiter) {
			info := strings.Split(line, delimiter)[1]
			fmt.Println(info)
			return
		}
	}
	fmt.Println("Note ", title, " not found")
}

func (n Note) ListNotes() {
	// Read note file
	lines := n.readFile()
	for _, line := range lines {
		if line != "" {
			fmt.Println(strings.Split(line, delimiter)[0])
		}
	}
}
