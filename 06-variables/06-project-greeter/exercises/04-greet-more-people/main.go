// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
	args := os.Args

	fmt.Printf("There are %d people! \n", len(args)-1)
	for i := 1; i < len(args); i++ {
		fmt.Printf("Hello great %s !\n", args[i])
	}
	fmt.Printf("Nice to meet you all.")
}
