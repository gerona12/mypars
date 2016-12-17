package services

import (
	"fmt"
)


type Research struct {
	Description string
	Preparation string
	Indication string
	Result string 
	
}

func FiindAll() ([]Research, error){
	fmt.Println("researchService: FiindAll")
	researches := [] Research{
		Research{"description", "preparation", "indication", "result"},
		Research{"description", "preparation", "indication", "result"},
		Research{"description", "preparation", "indication", "result"}}

	return researches, nil
}
