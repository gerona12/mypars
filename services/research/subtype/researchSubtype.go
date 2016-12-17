package researchSubtype
import (
	"fmt"
)
type ResearchSubtype struct{
	Title string
}

func FiindAll() ([]ResearchSubtype, error){
	fmt.Println("researchSubtype: FiindAll")
	researchSubtypes := [] ResearchSubtype{
	ResearchSubtype{"title1"}
	ResearchSubtype{"title2"}
 	ResearchSubtype{"title3"}}
	return researchTypes, nil
}
