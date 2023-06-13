package command

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: go-jenkins [COMMAND]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  up      Start Jenkins")
	fmt.Println("  down    Stop Jenkins")
	fmt.Println("  help    Show this help message")
}
