package main

import (
    "os"
)

func main() {

    if len(os.Args) < 3 {
        usage()
    }

    host := os.Args[1]
    keys := os.Args[2]
    command := os.Args[3]

    // RunTransferArgs(from, to, keys, threads)
}
