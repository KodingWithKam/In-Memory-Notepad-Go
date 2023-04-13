package main

import "fmt"

// do not change the name or the fields of the 'Country' struct!
type Country struct {
	Name     string
	Capital  string
	Currency string
}

func main() {
	// create an instance of the Country struct within the 'germany' variable here.
	// do not forget to assign the values mentioned above to its fields too!
	var germany Country

	germany.Name = "Germany"   // assign the value 'Germany' to the 'Name' field here
	germany.Capital = "Berlin" // assign the value 'Berlin' to the 'Capital' field here
	germany.Currency = "Euro"  // assign the value 'Euro' to the 'Currency' field here

	// Print the struct fields in order Name->Capital->Currency each on a new line:
	fmt.Printf("%s\n%s\n%s\n", germany.Name, germany.Capital, germany.Currency)
}
