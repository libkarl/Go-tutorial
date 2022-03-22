package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// Model pro kurz
type Course struct {
	CourseId string `json: "courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json: "price"`
	Author *Author `json: "author"`
}

type Author struct {
	FullName string `json: "fullname"`
	Website string `json: " website"`
}

// fake DB

var courses []Course

// helper - file
func (c *Course) IsEmpty() bool { // (c *Course) předávám funkci metodu Course
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome in Api")

}

// serve houme route, tohle bude posilat data na 
// prvotní home screen 
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Api by LearnCodeOnline</h>"))
} 
// r- reader je ten ok koho získáte hodnotu a w-write je místo kde na to napíšete odpověd

func getAllCourses (w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // nastavím content type header na na aplication/json, takže klient ví, že má čekat JSON
	json.NewEncoder(w).Encode(courses) // encode JSON to a http.ResponseWriter format
}

// cokoli co pošlu do Encode (courses) bude považováno
// za JSON soubor a vráceno zpět tomu kdo žádá prostřednictvím w 

func getOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) //funkce z mux, která bere z request variables, vytáhne to id
	fmt.Println(params)

	// loop through courses, find matching id které user posílá jako parametr a vrat jako response
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course) //chci enkodovat celý course a poslat skrz writer
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty in response
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {} 
	// situace kdy uživatel pošle přímo JSON data
	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Please send some data") // vytvoření JSON odpovědi
		return
	}

	// generate unique id (generování unikátního id kurzu)
	// přídání kurzu do kurzů

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// CourseId se rovná náhodné celé číslo převedené na string

	courses = append(courses, course)
	// přidání kurzu do kurzů pomocí append 
	json.NewEncoder(w).Encode(course)
	return
}	
