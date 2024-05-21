package note

import (
	"bufio"
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
			fmt.Println("Error reading notes :", err)
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
		fmt.Println("Error opening file :", err)
		fmt.Println("Can't update the note in database, try again")
		return err
	}
	defer f.Close()

	// Write the data to the file
	for _, line := range lines {
		_, err = f.Write([]byte(line))
		if err != nil {
			fmt.Println("Error writing to file :", err, "for note :", n.title)
			return err
		}
	}
	return nil
}

func (n Note) updateNote() {
	// Read the note file
	lines := n.readFile()

	// Flag to know if title already exist
	exist := false

	// Get the Updated lines without current title
	var updatedLines []string
	for _, line := range lines {
		if line != "" {
			if !strings.HasPrefix(line, n.title+delimiter) {
				updatedLines = append(updatedLines, line+"\n")
			} else {
				exist = true
			}
		}
	}

	if n.ops == opsRemoveNote && !exist {
		fmt.Println("Note not found for deletion")
		return
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
			if exist {
				fmt.Println("Note updated")
			} else {
				fmt.Println("Note created")
			}
		}
	} else if n.ops == opsUpdateNote {
		if err != nil {
			fmt.Println("Note update failed")
		} else {
			if exist {
				fmt.Println("Note updated")
			} else {
				fmt.Println("Note created")
			}
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
	// Define reader
	// TODO: Try to set the size here of buffer
	reader := bufio.NewReader(os.Stdin)

	// Prompt for title
	fmt.Print("Title >> ")
	title, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read title, Try again")
		return
	}
	// Remove last character/userInputBreaker
	title = title[:len(title)-1]

	// Remove any leading or trailing whitespaces
	title = strings.TrimSpace(title)

	// Return on empty title
	if len(title) == 0 {
		fmt.Println("title can't be empty to create")
		return
	}

	n.title = title

	// Prompt for information
	fmt.Print("Info >> ")
	info, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read info, Try again")
		return
	}
	// Remove last character/userInputBreaker
	info = info[:len(info)-1]

	// Remove any leading or trailing whitespaces
	n.info = strings.TrimSpace(info)

	// Update ops
	n.ops = opsCreateNote

	// Create note
	n.updateNote()
}

func (n Note) UpdateNote() {
	// Define reader
	// TODO: Try to set the size here of buffer
	reader := bufio.NewReader(os.Stdin)

	// Prompt for title
	fmt.Print("Title >> ")
	title, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read title, Try again")
		return
	}
	// Remove last character/userInputBreaker
	title = title[:len(title)-1]

	// Remove any leading or trailing whitespaces
	title = strings.TrimSpace(title)

	// Return on empty title
	if len(title) == 0 {
		fmt.Println("title can't be empty to update")
		return
	}

	n.title = title

	// Prompt for information
	fmt.Print("Info >> ")
	info, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read info, Try again")
		return
	}
	// Remove last character/userInputBreaker
	info = info[:len(info)-1]

	// Remove any leading or trailing whitespaces
	n.info = strings.TrimSpace(info)

	// Update ops
	n.ops = opsUpdateNote

	// Create note
	n.updateNote()
}

func (n Note) DeleteNote() {
	// Define reader
	// TODO: Try to set the size here of buffer
	reader := bufio.NewReader(os.Stdin)

	// Prompt for title
	fmt.Print("Title >> ")
	title, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read title, Try again")
		return
	}
	// Remove last character/userInputBreaker
	title = title[:len(title)-1]

	// Remove any leading or trailing whitespaces
	title = strings.TrimSpace(title)

	// Return on empty title
	if len(title) == 0 {
		fmt.Println("title can't be empty to delete")
		return
	}

	n.title = title

	// Set ops to delete
	n.ops = opsRemoveNote

	// Delete note
	n.updateNote()
}

func (n Note) DeleteNoteFile() {
	err := os.Remove(n.filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Not able to delete all notes")
		} else {
			fmt.Println("No note found to delete")
		}
	} else {
		fmt.Println("Deleted all notes")
	}
}

// TODO: This should be able to take more than one title to give info out
func (n Note) GetInfo() {
	// Define reader
	// TODO: Try to set the size here of buffer
	reader := bufio.NewReader(os.Stdin)

	// Prompt for title
	fmt.Print("Title >> ")
	title, err := reader.ReadString(userInputBreaker)
	if err != nil {
		fmt.Println("Couldn't read title, Try again")
		return
	}
	// Remove last character/userInputBreaker
	title = title[:len(title)-1]

	// Read note file
	lines := n.readFile()
	for _, line := range lines {
		if strings.HasPrefix(line, title+delimiter) {
			info := strings.Split(line, delimiter)[1]
			fmt.Println(info)
			return
		}
	}
	fmt.Println(title, ">> Title not found")
}

func (n Note) ListNotes(includeInfo bool) {
	// Read note file
	lines := n.readFile()

	if len(lines) == 0 {
		fmt.Println("No note available")
	}

	for _, line := range lines {
		if line != "" {
			if includeInfo {
				fmt.Println(strings.Split(line, delimiter)[0], ">>>", strings.Split(line, delimiter)[1])
			} else {
				fmt.Println(strings.Split(line, delimiter)[0])
			}
		}
	}
}
