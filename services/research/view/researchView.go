
package researchView
import (
	"fmt"
)
type ResearchView struct{
	Number int16
	Price int16
}

func FiindAll() ([]ResearchView, error){
	fmt.Println("researchView: FiindAll")
	researchViwes := [] ResearchView{
	ResearchView{111, 22}
	ResearchView{222, 45}
 	ResearchView{333, 88}}
	return researchViewes, nil
}
