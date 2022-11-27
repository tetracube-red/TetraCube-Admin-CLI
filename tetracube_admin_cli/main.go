package main

import (
	"fmt"
	"tetracube.red/admin-cli/tetracube_admin_cli/commands"
)

func main() {
	fmt.Println("TetraCube administrative CLI")
	commands.Execute()
}
