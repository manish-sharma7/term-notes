# !/bin/zsh

# Files where notes will be saved
NOTES_DIR="$HOME/term-notes/"

# If dir doesn't exist, Then create
mkdir -p $NOTES_DIR

# Define file
NOTES_FILE=$NOTES_DIR/001.txt

# Delimiter between title and info
DELIMITER="-"

# Print help messages
function print_help {
    echo "Welcome to term-notes..."
    echo "Usage: $0 [-c] [-h]"
    echo "-c      Specify if u want to create a note"
    echo "-h      Display this help message"
}

# Create a note
function create_note {
    # Prompt for title
    echo "Title >>"
    read -r title

    # Prompt for information
    echo "Info >>"
    read -r info

    # Save it in file
    echo "$title-$info" >> $NOTES_FILE
}

# List all nodes
function list_all_notes {
    awk -F "$DELIMITER" '{
        print $1
    }' "$NOTES_FILE"
}

# Get title's info
function get_title_info {
    awk -F "$DELIMITER" -v title="$1" '{
        if ($1 == title) {
            echo title
        }
    }' "$NOTES_FILE"
}

# Accept the arguments
while getopts "hclg:" opt; do
    case $opt in
        h) 
        print_help
        exit 0
        ;;
        c) 
        create_note
        ;;
        l) 
        list_all_notes
        ;;
        g) 
        TITLE="$OPTARG"
        get_title_info $TITLE
        ;;
        \*)
        echo "Invalid option: -$OPTARG" >&2
        print_help
        exit 1
        ;;
    esac
done

# Handle any other argument with help message
shift "$((OPTIND-1))"
if [ ! $# -eq 1 ] ; then
    print_help
    exit 1
fi
