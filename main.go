package main
 
import (
	"encoding/json"
	"net/http"
)
 
type ResearchType struct {
	Title string
}

type ResearchSubtype struct{
	Title string
}

type ResearchView struct{
	Number int16
	Price int16
}

type Research struct{
	Description string
	Preparation string
	Indication string
	Result string 
}

func GetResearches(w http.ResponseWriter, r *http.Request) {
	entity := Research{"description", "preparation", "indication", "result"}
 
	serializeObject2Json(entity, w, r)
}

func GetResearchTypes(w http.ResponseWriter, r *http.Request) {
	entity := ResearchType{"title"}
 
	serializeObject2Json(entity, w, r)
}

func GetResearchSubtypes(w http.ResponseWriter, r *http.Request) {
	entity := ResearchSubtype{"title"}
 
	serializeObject2Json(entity, w, r)
}

func GetResearchViews(w http.ResponseWriter, r *http.Request) {
	entity := ResearchView{111, 22}
 
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
