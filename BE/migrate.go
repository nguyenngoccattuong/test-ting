package main

import (
	"log"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func main() {
	// Load environment variables
	Initializers.LoadEnvironmentVariables()

	// Connect to database
	Initializers.ConnectToDB()

	// Auto migrate the schema
	err := Initializers.DB.AutoMigrate(&Models.Kol{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}
