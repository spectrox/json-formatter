package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

const StringSplitLength = 8192

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := ""

	for scanner.Scan() {
		text := buffer + scanner.Text()
		buffer = ""

		if json.Valid([]byte(text)) {
			formatted := pretty.Pretty([]byte(text))
			result := pretty.Color(formatted, nil)

			fmt.Println(string(result))
		} else if isSplitJsonString(text) {
			buffer = text
		} else {
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func isSplitJsonString(text string) bool {
	//if len(text) != StringSplitLength {
	//	return false
	//}

	if string(text[0]) != "{" {
		return false
	}

	if string(text[len(text) - 1]) == "}" {
		return false
	}

	return true
}
