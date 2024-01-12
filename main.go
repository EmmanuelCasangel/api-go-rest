package main

import (
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("iniciando servidor rest")
	routes.HandleRequest()
}
