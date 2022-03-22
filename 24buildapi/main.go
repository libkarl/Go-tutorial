package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course // uvnitř fake DB budou různé Course typu map

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	var s = rand.NewSource(5)
	var rx = rand.New(s)

	r := mux.NewRouter() // vytvoření mux Routeru dá schopnost psát ty cesty s HandleFunction

	//seeding inject nějakých příkladů pro práci s daty do fake DB názvem courses
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET") //přirádím routeru funkci s cestou / a metodou serveHome, kterou jsem vytvořil jako funkci, celé to má metogu GET
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse(rx)).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r)) // přiřazení portu 4000 a routeru z mux, který mám v r
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(rx *rand.Rand) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create one course")
		w.Header().Set("Content-Type", "applicatioan/json")

		// what if: body is empty
		//pokud je tělo zprávy prázdné vrátí tuhle odpověd
		// pomocí Encode to přehodí do bitů a NewEncode to zapíše do http odpovedi
		if r.Body == nil {
			json.NewEncoder(w).Encode("Please send some data")
		}

		// what about - {}
		// případ kdy někdo pošle prázdný JSON, takže v těle zprávy něco je
		// ale nejsou tam žádná použitelná data

		var course Course                           // vytvořil proměnnou course a přidělil jí strukturu Course
		_ = json.NewDecoder(r.Body).Decode(&course) // snaží se dekodovat tělo zprávy podle struktury
		if course.IsEmpty() {                       //ještě je možné to procházet popmocí loop
			json.NewEncoder(w).Encode("No data inside JSON")
			return
		}
		// json.NewDecoder(r.body)-bere data z těla r http requestu a dekoduje je do proměnné course, která má přidělenou strukturu Course
		// pomocí & ukazuje přímo na místo, kde je v paměti alokovaná proměnná course s přidělenou strukturou
		//TODO: check only if title is duplicate
		// loop, title matches with course.coursename, JSON

		// generate unique id, a převede na string
		// append course into courses (přidá kurz do kurzů)

		// tohle mi generuje unikátní číslo
		course.CourseId = strconv.Itoa(rx.Intn(100)) // do courseId v course proměnné zapíši náhodné číslo, které převedu na string
		courses = append(courses, course)            // přídá kurz i s novým nýhodným id do mé fake databáze
		json.NewEncoder(w).Encode(course)            // vezme vše co je ve var course Encoduje a zapíše do odpovědi
		return

	}
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r) // vezme id z requestu a zapíše ho do params

	// loop, id, remove, add with my ID

	for index, course := range courses { // projde všechny kurzy
		if course.CourseId == params["id"] { // pokud při procházení kurzu narazí na Id kurzu, které se shoduje s id z requestu provede co je v if
			courses = append(courses[:index], courses[index+1:]...) // přepíše starou databázi tak, aby neobsahovala kurz s nalezenýchm ID
			var course Course                                       // vytvoří proměnnou course, které přidělí strukturu Course
			_ = json.NewDecoder(r.Body).Decode(&course)             // dekoduje tělo požadavku a uloží do proměné course s přidělenou strukturou
			course.CourseId = params["id"]                          // do kurzu, který je ted v Course uloží id, podle kterého předtím smazal původní kurz
			courses = append(courses, course)                       // přidá kurz do databáze, má stejné id jako původní, ale jeho tělo je nové podle toho co bylo v requestu
			json.NewEncoder(w).Encode(course)                       // encoduje finální stav kurzu a zapíše ho do odpovědi
			return
		}
	}
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
