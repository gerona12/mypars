package researchType
import (

	"fmt"

)
type ResearchType struct{
	Title string
}

func FiindAll() ([]ResearchType, error){
	fmt.Println("researchType: FiindAll")
	researchTypes := [] ResearchType{
	ResearchType{"title1"}
	ResearchType{"title2"}
 	ResearchType{"title3"}		
}
	return researchTypes, nil
}
