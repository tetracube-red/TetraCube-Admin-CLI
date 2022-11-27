package main

import (
	"fmt"

	"tetracube.red/admin-cli/tetracube_admin_cli/cmd"
)

func main() {
	fmt.Print("TetraCube administrative CLI")
	cmd.Execute()
}
