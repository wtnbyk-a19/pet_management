package Pet

import (
	"pet-management/src/Domain/Domain/Model"
	"time"
)

type Pet struct {
	PetId         int       `json:"pet_id" gorm:"primary_key"`
	UserId        int       `json:"user_id"`
	PetName       string    `json:"pet_name"`
	Gender        string    `json:"gender"`
	Breed         string    `json:"breed"`
	Birthday      time.Time `json:"birthday"`
	Adoptaversary time.Time `json:"adoptaversary"`
	Memo          string    `json:"memo"`

	Model.Model
}

func NewPet(userId int, petName string, gender string, breed string, birthday time.Time, adoptaversary time.Time, memo string) (pet *Pet, err error) {
	pet = &Pet{}

	// TODO: バリデーションの実装

	pet.UserId = userId
	pet.PetName = petName
	pet.Gender = gender
	pet.Breed = breed
	pet.Birthday = birthday
	pet.Adoptaversary = adoptaversary
	pet.Memo = memo

	return pet, nil
}
