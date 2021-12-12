package Create

import (
	"time"
)

type PetCreateInputData struct {
	PetName       string    `json:"pet_name"`
	Gender        string    `json:"gender"`
	Variety       string    `json:"variety"`
	Breed         string    `json:"breed"`
	Birthday      time.Time `json:"birthday"`
	Adoptaversary time.Time `json:"adoptaversary"`
	Memo          string    `json:"memo"`
}

func NewPetCreateInputData(petName string, gender string, breed string, birthday string, adoptaversary string, memo string) (inputData *PetCreateInputData, err error) {
	inputData = &PetCreateInputData{}

	var layout = "2006-01-02"

	var birthdayDate time.Time
	birthdayDate, err = time.Parse(layout, birthday)
	if err != nil {
		return nil, err
	}

	var adoptaversaryDate time.Time
	adoptaversaryDate, err = time.Parse(layout, adoptaversary)
	if err != nil {
		return nil, err
	}

	inputData.PetName = petName
	inputData.Gender = gender
	inputData.Breed = breed
	inputData.Birthday = birthdayDate
	inputData.Adoptaversary = adoptaversaryDate
	inputData.Memo = memo

	return inputData, nil
}
