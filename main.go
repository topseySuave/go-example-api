package main

import (
	"fmt"
	"gabrielmicah/go-example/config"
	"gabrielmicah/go-example/src/modules/profile/model"
	"gabrielmicah/go-example/src/modules/profile/repository"
	"time"
)

func main() {
	db, err := config.GetMongoDb()
	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "people")

	saveProfile(*profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.Name = "Gabriel Micah 2"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved...")
	}
}
