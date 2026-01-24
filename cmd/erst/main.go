package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Erst CLI starting...")
	
	if len(os.Args) > 1 {
		fmt.Printf("Arguments: %v\n", os.Args[1:])
	}

	fmt.Println("Ready.")
    time.Sleep(1 * time.Second)
}
