package note

const (
	noteFile  string = "001.txt"
	delimiter string = "-"
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
