package pets

import "go-breeders/model"

func NewPet(species string) *model.Pet {
	return &model.Pet{
		Species: species,
	}
}
