package main
 
import (
	"encoding/json"
	"net/http"
	"parser/research/"
	"parser/research/subtype/"
	"parser/research/type/"
	"parser/research/view/"
)
 


func GetResearches(w http.ResponseWriter, r *http.Request) {
		var entity [3] Research
		entity[0] = Research{"description", "preparation", "indication", "result"}
 		entity[1] = Research{"description", "preparation", "indication", "result"}
 		entity[2] = Research{"description", "preparation", "indication", "result"}
	serializeObject2Json(entity, w, r)
}

func GetResearchTypes(w http.ResponseWriter, r *http.Request) {
	
 	var entity [3] ResearchType
	entity[0] = ResearchType{"title1"}
	entity[1] = ResearchType{"title2"}
 	entity[2] = ResearchType{"title3"}
	
	serializeObject2Json(entity, w, r)
}





func GetResearchSubtypes(w http.ResponseWriter, r *http.Request) {
	var entity [3] ResearchSubtype
	entity[0] = ResearchSubtype{"Title1"}
	entity[1] = ResearchSubtype{"Title2"}
 	entity[2] = ResearchSubtype{"Title3"}
	
	serializeObject2Json(entity, w, r)
}

func GetResearchViews(w http.ResponseWriter, r *http.Request) {
    var entity [3] ResearchView
	entity[0] = ResearchView{111, 22}
	entity[1] = ResearchView{222, 45}
 	entity[2] = ResearchView{333, 88}
	serializeObject2Json(entity, w, r)
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
