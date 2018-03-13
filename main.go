package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/uberblah/go-demo/match-name"
)

const quitCommand = "quit"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// get the name
		fmt.Print("Enter a name or email: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")

		// quit if requested
		if name == quitCommand {
			break
		}

		// identify and parse
		nameType, parts := match_name.MatchName(name)
		if nameType == "" || parts == nil {
			parts = map[string]string{
				"contents": name,
			}
			nameType = "unknown"
		}
		parts["type"] = nameType
		j, _ := json.Marshal(parts)

		// print
		fmt.Println(string(j))
	}
}
