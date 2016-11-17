package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println(
		`usage: redis-exec <host> <keys> <command>`)

	os.Exit(1)
}
