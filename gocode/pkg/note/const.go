package note

const (
	noteFile  string = "001.txt"
	delimiter string = "-"
	// TODO: Options for more ways to end the user input
	userInputBreaker byte = '\n'
)

// Define Operations
const (
	NONE = iota
	opsCreateNote
	opsListNotes
	opsGetNote
	opsUpdateNote
	opsRemoveNote
)
