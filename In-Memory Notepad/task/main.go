package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Note create struct for notes
type Note struct {
	note string
}

func main() {
	// 	write your code here
	var userCommand string

	// Create slice to represent notes
	var notes []Note

	// number of notes
	var numberOfNotes = getUserNoteCount()

	// Start process Loop until user inputs exit
	for userCommand != "exit" {
		// get input
		fmt.Println("Enter a command and data:")
		// create scanner
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		// split the input by white space
		input := strings.Split(scanner.Text(), " ")
		userCommand = input[0]
		// concatenate string inputs except for the first element (command)
		data := strings.Join(input[1:], " ")

		switch userCommand {
		case "create":
			createNote(notes, data, numberOfNotes)
		case "clear":
			// set slice to nil
			clearNotes(notes)
		case "delete":
			deleteNote(notes, data, input, numberOfNotes)
		case "list":
			listNotes(notes)
		case "exit":
			fmt.Println("[Info] Bye!")
			break
		case "update":
			updateNote(notes, data, input, numberOfNotes)
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

// get user input for max notes
func getUserNoteCount() int {
	// get input for notes length
	var response int
	fmt.Println("Enter the maximum number of notes:")
	fmt.Scan(&response)
	return response
}

// clear all notes
func clearNotes(notes []Note) {
	// set slice to nil
	notes = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

// create note
func createNote(notes []Note, data string, numberOfNotes int) {
	// check if command is empty
	if len(data) != 0 {
		// Create Should only store up 5 notes
		if len(notes) < numberOfNotes {
			// convert user input to Note type
			newNote := Note{note: data}
			// append to slice
			fmt.Println("[OK] The note was successfully created")
			notes = append(notes, newNote)
		} else {
			fmt.Println("[Error] Notepad is full")
		}
	} else {
		fmt.Println("[Error] Missing note argument")
	}
}

// delete note from index
func deleteNote(notes []Note, data string, input []string, numberOfNotes int) {
	// check if command is empty
	if len(data) != 0 {
		// get position argument
		positionValue := input[1]
		// check to see if position argument is empty
		if len(positionValue) != 0 {
			// attempt to parse data value into int
			val, err := strconv.Atoi(positionValue) // We save the converted value in the 'val' variable
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", positionValue) // Exit if we have an error
				return
			}

			// declare index value for accurate position
			i := val - 1
			// check if position value is within range
			if val > 0 && val <= numberOfNotes {
				// check if there is a note here
				if notes != nil && notes[i].note != "" {
					// delete note at given index
					notes = removeNote(i, notes)
					fmt.Printf("[OK] The note at position %d was successfully deleted\n", val)
				} else {
					fmt.Println("[Error] There is nothing to delete")
				}
			} else {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", val, numberOfNotes)
			}
		}

	} else {
		fmt.Println("[Error] Missing position argument")
	}
}

// list available notes
func listNotes(notes []Note) {
	// if notes slice is not nil
	if notes != nil {
		// Loop through slice
		for i, note := range notes {
			// if the note is not empty, print it out
			if note.note != "" {
				// add one to index value for real position
				fmt.Printf("[Info] %d: %s\n", i+1, note.note)
			}
		}
	} else {
		fmt.Println("[Info] Notepad is empty")
	}
}

// remove slice element
func removeNote(i int, notes []Note) []Note {
	// delete note at given index
	return append(notes[:i], notes[i+1:]...)
}

// update note at index
func updateNote(notes []Note, data string, input []string, numberOfNotes int) {
	// check if command is empty
	if len(data) != 0 {
		if len(data) > 2 {
			// concatenate string inputs except for the first & second element (command, position)
			updatedNoteMessage := strings.Join(input[2:], " ")
			// get position argument
			positionValue := input[1]
			// check to see if position argument is empty
			if len(positionValue) != 0 {
				// attempt to parse data value into int
				val, err := strconv.Atoi(positionValue) // We save the converted value in the 'val' variable
				if err != nil {
					fmt.Printf("[Error] Invalid position: %s\n", positionValue) // Exit if we have an error
					return
				}

				// check if position value is within range
				if val > 0 && val <= numberOfNotes {
					// check if there is a note here
					if notes != nil && notes[val-1].note != "" {
						newNote := Note{note: updatedNoteMessage}
						notes[val-1] = newNote
						fmt.Printf("[OK] The note at position %d was successfully updated\n", val)
					} else {
						fmt.Println("[Error] There is nothing to update")
					}
				} else {
					fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", val, numberOfNotes)
				}
			}
		} else {
			fmt.Println("[Error] Missing note argument")
		}
	} else {
		fmt.Println("[Error] Missing position argument")
	}
}
