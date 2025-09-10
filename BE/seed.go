package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func main() {
	// Load environment variables
	Initializers.LoadEnvironmentVariables()

	// Connect to database
	Initializers.ConnectToDB()

	// Clear existing data
	Initializers.DB.Where("1 = 1").Delete(&Models.Kol{})

	// Create 20 dummy KOLs
	educations := []string{
		"Bachelor's in Computer Science",
		"Bachelor's in Marketing",
		"Master's in Business Administration",
		"Bachelor's in Communications",
		"Master's in Digital Marketing",
		"Bachelor's in Graphic Design",
		"Master's in Data Science",
		"Bachelor's in Psychology",
	}

	languages := []string{"en", "vn", "ja", "ko", "th"}
	verificationStatuses := []string{"Verified", "Pending", "Rejected"}
	livenessStatuses := []string{"Passed", "Failed", "Pending"}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 20; i++ {
		kol := Models.Kol{
			KolID:                int64(100 + i),
			UserProfileID:        int64(2000 + i),
			Language:             languages[rand.Intn(len(languages))],
			Education:            educations[rand.Intn(len(educations))],
			ExpectedSalary:       int64(40000 + rand.Intn(60000)), // 40k-100k
			ExpectedSalaryEnable: rand.Float32() < 0.7,            // 70% chance
			ChannelSettingTypeID: int64(1 + rand.Intn(3)),         // 1-3
			IDFrontURL:           fmt.Sprintf("https://example.com/id-front-%d.jpg", i),
			IDBackURL:            fmt.Sprintf("https://example.com/id-back-%d.jpg", i),
			PortraitURL:          fmt.Sprintf("https://example.com/portrait-%d.jpg", i),
			RewardID:             int64(300 + i),
			PaymentMethodID:      int64(400 + i),
			TestimonialsID:       int64(500 + i),
			VerificationStatus:   verificationStatuses[rand.Intn(len(verificationStatuses))],
			Enabled:              rand.Float32() < 0.9,                     // 90% chance
			ActiveDate:           time.Now().AddDate(0, 0, -rand.Intn(60)), // Random date within 60 days
			Active:               true,
			CreatedBy:            "admin",
			CreatedDate:          time.Now().AddDate(0, 0, -rand.Intn(60)),
			ModifiedBy:           "admin",
			ModifiedDate:         time.Now().AddDate(0, 0, -rand.Intn(30)),
			IsRemove:             false,
			IsOnBoarding:         rand.Float32() < 0.8, // 80% chance
			Code:                 fmt.Sprintf("KOL2024%03d", i),
			PortraitRightURL:     fmt.Sprintf("https://example.com/portrait-right-%d.jpg", i),
			PortraitLeftURL:      fmt.Sprintf("https://example.com/portrait-left-%d.jpg", i),
			LivenessStatus:       livenessStatuses[rand.Intn(len(livenessStatuses))],
		}

		result := Initializers.DB.Create(&kol)
		if result.Error != nil {
			log.Printf("Error creating KOL %d: %v", kol.KolID, result.Error)
		} else {
			log.Printf("Successfully created KOL %d", kol.KolID)
		}
	}

	log.Println("All 20 KOLs created successfully!")
}
