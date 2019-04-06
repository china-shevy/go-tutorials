package repl

import (
	"bufio"
	"fmt"
	"galculator/internel/compute"
	"os"
)

func Repl() error {
	buf := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sentence, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		if len(sentence) == 1 {
			continue
		}
		fmt.Println(compute.Compute(string(sentence[:len(sentence)-1])))
	}
}
