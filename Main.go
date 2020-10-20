package main
import (
"fmt"
"os"
"strings"
)
func main() {
	who := "World!"
	if len(os.Args) > 1 { /* os.Args[0] - имя команды «hello» или «hello.exe» */


		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}