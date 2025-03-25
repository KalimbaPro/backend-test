package internal

type Breed struct {
	Id int `json:"id"`
	Species string `json:"species"`
	PetSize string `json:"petSize"`
	Name string `json:"name"`
	AverageMaleAdultWeight int `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight int `json:"average_female_adult_weight"`
}