package methods

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func TerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	res := string(output)
	cols, _ := strconv.Atoi(res[:len(res)-1])
	return cols
}
