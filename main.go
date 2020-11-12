package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)
import "github.com/tidwall/pretty"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if json.Valid([]byte(text)) {
			formatted := pretty.Pretty([]byte(text))
			result := pretty.Color(formatted, nil)

			fmt.Println(string(result))
		} else {
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
