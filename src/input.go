package src

import (
	"bufio"
	"os"
)

func CmdLineInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
