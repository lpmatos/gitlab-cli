package main

import (
	"github.com/lpmatos/glabby/cmd"
	"github.com/lpmatos/glabby/internal/helpers"
)

func main() {
	helpers.ASCIIMessage("Glabby", "", "yellow")
	cmd.Execute()
}
