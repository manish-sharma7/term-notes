# !/bin/zsh

# Files where notes will be saved
NOTES_DIR="$HOME/term-notes/"

# If dir doesn't exist, Then create
mkdir -p $NOTES_DIR

# Define file
NOTES_FILE=$NOTES_DIR/001.txt

# Define temp file
NOTES_FILE_TEMP=$NOTES_DIR/001_temp.txt

# Delimiter between title and info
DELIMITER="-"

# Print help messages
function print_help {
    echo "Welcome to term-notes..."
    echo "Usage: notes [-c] [-l] [-g <title>] [-d <title>] [-h] "
    echo "-c      Create a note"
    echo "-l      List all notes"
    echo "-g      Get info of note title"
    echo "-d      Delete a note title"
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
    info=$(awk -F "$DELIMITER" -v title=$1 '{
        if ($1 == title) {
            print $2
            exit
        }
    }' "$NOTES_FILE")

    if [[ -z $info ]]; then
        echo "Not found title $1"
    else
        echo $info
    fi
}

# Delete title
function delete_title {
    found=$(awk -F "$DELIMITER" -v title=$1 '{
        if ($1 == title) {
            print $1
            exit
        }
    }' "$NOTES_FILE")

    if [[ -z $found ]]; then
        echo "Couldn't find note title $1"
    else
        sed "/$found/d" "$NOTES_FILE" > "$NOTES_FILE_TEMP"
        mv "$NOTES_FILE_TEMP" "$NOTES_FILE"
        echo "Removed title=$found"
    fi
}

# Accept the arguments
while getopts "hclg:d:" opt; do
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
        d) 
        TITLE="$OPTARG"
        delete_title $TITLE
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
