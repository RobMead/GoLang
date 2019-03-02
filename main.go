package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string = "BUMBLESHACK"
	var lastName string = "CRUMPLEHORN"
	fmt.Println(isItHim(firstName, lastName))
}

func isItHim(firstName	string, lastName string ) string {
	
	var testFirst string = "BENEDICT"
	var testSecond string = "CUMBERBATCH"

	// First condition - check that each name has at least 7 letters.
	if (len(firstName) < 7 || len(lastName) < 7){
		return "It is not him"
	} 

	// Second condition - first chars must be both 'B' and 'C'.
	if !(firstName[0] == 'B' && lastName[0] == 'C'){
		return "It is not him"
	}
	
	// Check whether three characters of first name appear in "BENEDICT".
	var counter int = 0
	for _, char := range firstName {
		if strings.Contains(testFirst,string(char)){ 
			counter += 1
		} 
	}
	if counter < 3 {
		return "It is not him"
	}
	counter = 0

	// Check whether five characters of last name appear in "CUMBERBATCH".
	for _, char := range lastName {
		if strings.Contains(testSecond,string(char)){
			counter += 1
		}
	}
	if counter < 5 {
		return "It is not him"
	}

	// All tests passed - return "it is him".
	return "It is him"
}
