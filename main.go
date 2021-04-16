package main

import "github.com/samuelralak/goprime/routes"

func main() {
	// Create routes instance
	router := routes.SetupRouter()

	// Run application default :8080
	_ = router.Run()
}
