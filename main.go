package main

func main() {
	// Create router instance
	router := setupRouter()

	// Run application default :8080
	_ = router.Run()
}
