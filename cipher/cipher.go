package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader( os.Stdin )
	fmt.Println( "Ceaser Cipher" )
	fmt.Println( "--------------------" )

	var cipher = make( map[string]string )
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
						"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
						"y", "z"}
						
	fmt.Print("Enter the desired offset for cipher -> ")
	input, _ := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace( input, "\n", "", -1 )

	//convert the string to an int
	fmt.Println( "Input: ", input )
	offset, _ := strconv.Atoi(input)

	//Create the rolling cipher
	for index, element := range letters {
		cipherIndex := index + offset
		if( cipherIndex >= len(letters) ) {
			cipherIndex = cipherIndex - len(letters)
		} 
		cipher[element] = letters[ cipherIndex ]
	}

	for {
		fmt.Print("Message to Encrypt:-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		cipherString := ""

		for _, char := range text {
			character := string(char)
			if( character == " " ){
				cipherString += " "
			}

			if( character == strings.ToUpper(character) ) {
				cipherString += strings.ToUpper( cipher[strings.ToLower(character)] )
			} else{
				cipherString += cipher[character]
			}
			
			if( cipher[character] == "" ) {
				cipherString += character
			}
		}

		fmt.Println( "Encrypted Message: " )
		fmt.Println( "     ", cipherString )
	}
}
