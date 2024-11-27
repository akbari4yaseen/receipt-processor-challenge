package main

import (
	"fmt"

	"github.com/akbari4yaseen/receipt-processor-challenge/api"
)

func main() {
	router := api.SetupRouter()
	fmt.Println("Server is running on port 8080...")
	router.Run(":8080")
}
