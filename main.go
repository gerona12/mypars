package main
 
import (
	"encoding/json"
	"net/http"
	researhService "parser/services/research"
	researchType "parser/services/research/type"
	researchSubtype "parser/services/research/subtype"
	researchView "parser/services/research/view"
) 
/*teper nungno
1 pro4itat eshe raz zadanie i ponyat kakie polya budut v sushnosyah
2 ponyat kakie metodi nugni
3 zadanie 2 -tut nugen obichniy CRUD (google esli ne znaesh 4to eto)
4 sdelat bazovie metodi:
 - ResearchDao.Create(Research) return Research
 - ResearchDao.FindAll() - return Research[]
 eto nugno sdelat dlya kagdoy entity
 5 - methodi nugno prosto sodat ne nugno delat realizaciyu
 v methode postavit fmt.Ptintln('Name method') 
v methdo FindAll perenesti cod is nigneh metodov
*/

func GetResearches(w http.ResponseWriter, r *http.Request) {	
	serializeObject2Json(researchService.FindAll(), w, r)
}

func GetResearchTypes(w http.ResponseWriter, r *http.Request) {
	
 	/*var entity [3] researchType.ResearchType
	entity[0] = researchType.ResearchType{"title1"}
	entity[1] = researchType.ResearchType{"title2"}
 	entity[2] = researchType.ResearchType{"title3"}
	*/
	serializeObject2Json(researchType.FindAll(), w, r)
}





func GetResearchSubtypes(w http.ResponseWriter, r *http.Request) {
	/*var entity [3] researchSubtype.ResearchSubtype
	entity[0] = researchSubtype.ResearchSubtype{"Title1"}
	entity[1] = researchSubtype.ResearchSubtype{"Title2"}
 	entity[2] = researchSubtype.ResearchSubtype{"Title3"}
	*/
	serializeObject2Json(researchSubtype.FindAll(), w, r)
}

func GetResearchViews(w http.ResponseWriter, r *http.Request) {
    /*var entity [3] researchView.ResearchView
	entity[0] = researchView.ResearchView{111, 22}
	entity[1] = researchView.ResearchView{222, 45}
 	entity[2] = researchView.ResearchView{333, 88}
	*/
	serializeObject2Json(researchView.FindAll(), w, r)
}

func serializeObject2Json(obj interface{}, w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func main() {
	http.HandleFunc("/api/research", GetResearches)
	http.HandleFunc("/api/research/types", GetResearchTypes)
	http.HandleFunc("/api/research/subtype", GetResearchSubtypes)
   	http.HandleFunc("/api/research/views", GetResearchViews)

   	http.ListenAndServe(":8888", nil)
}
