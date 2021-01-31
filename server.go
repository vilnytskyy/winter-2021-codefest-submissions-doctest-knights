package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

var database *sql.DB

type HttpHandler struct{}

type Student struct {
	Student_id int64          `json:"student_id,string,omitempty"`
	Name string               `json:"name,omitempty"`
	Major_id int64            `json:"major_id,string,omitempty"`
	Credits int64             `json:"credits,string,omitempty"`
	Overall_GPA float64       `json:"overall_gpa,string,omitempty"`
	Major_GPA float64         `json:"major_gpa,string,omitempty"`
}

type Course struct {
	Course_id int64               `json:"course_id,string,omitempty"`
	Name string                   `json:"name,omitempty"`
	Department string             `json:"department,omitempty"`
	Course_number int64           `json:"course_number,string,omitempty"`
	Prereqs string                `json:"prereqs,omitempty"`
	Requirement_fulfilled string  `json:"requirement_fulfilled,omitempty"`
	Credits int64                 `json:"credits,string,omitempty"`
	Description string            `json:"description,omitempty"`
}

type Requirement struct {
	Requirement_id int64       `json:"requirement_id,string,omitempty"`
	Name string                `json:"name,omitempty"`
	Parent_requirement int64   `json:"parent_requirement,string,omitempty"`
	Credits_required int64     `json:"credits_required,string,omitempty"`
}

type Course_Taken struct {
	Student_id int64           `json:"student_id,string,omitempty"`
	Course_id int64            `json:"course_id,string,omitempty"`
	In_progress int64          `json:"in_progress,string,omitempty"`
	Grade string               `json:"grade,omitempty"`
}


func retrieveStudentData(student_id int) (Student) {
	var id int64
	var name string
	var major_id int64
	var credits int64
	var ogpa float64
	var mgpa float64
	var gb string
	var gb2 string
	qtext := fmt.Sprintf("SELECT DISTINCT * FROM students WHERE student_id LIKE %d", student_id)
	rows := database.QueryRow(qtext).Scan(&id, &name, &major_id, &credits, &ogpa, &mgpa, &gb, &gb2)
	if rows != nil {
		log.Fatal(rows)
	}
	s := Student{id, name, major_id, credits, ogpa, mgpa}
	return s
}

func retrieveRequirementData(req_id int) (Requirement) {
	var id int64
	var name string
	var pr int64
	var cr int64
	qtext := fmt.Sprintf("SELECT DISTINCT * FROM requirements WHERE requirement_id LIKE %d", req_id)
	rows := database.QueryRow(qtext).Scan(&id, &name, &pr, &cr)
	if rows != nil {
		log.Fatal(rows)
	}

	r := Requirement{id, name, pr, cr}
	fmt.Println(r)
	return r
}

func retrieveCourseData(course_id int) (Course) {
	var id int64
	var name string
	var dep string
	var cid int64
	var prereq string
	var req_f string
	var credits int64
	var desc string
	qtext := fmt.Sprintf("SELECT DISTINCT * FROM courses WHERE course_id LIKE %d", course_id)
	rows := database.QueryRow(qtext).Scan(&id, &name, &dep, &cid, &prereq, &req_f, &credits, &desc)
	if rows != nil {
		log.Fatal(rows)
	}
	c := Course{id, name, dep, cid, prereq, req_f, credits, desc}
	// fmt.Println(c)
	return c
}

func retrieveCoursesTaken(student_id int) ([]Course_Taken) {
	var courses []Course_Taken
	var id1 int64
	var id2 int64
	var inp int64
	var g string
	qtext := fmt.Sprintf("SELECT DISTINCT * FROM courses_taken WHERE student_id LIKE %d", student_id)
	rows, _ := database.Query(qtext)

	for rows.Next() {
		rows.Scan(&id1, &id2, &inp, &g)
		courses = append(courses, Course_Taken{
			Student_id: id1,
			Course_id: id2,
			In_progress: inp,
			Grade: g,
		})
		fmt.Println(courses)
	}
	return courses
}

func retrieveAllCourses() ([]Course) {
	var courses []Course
	var id int64
	var name string
	var dep string
	var cid int64
	var prereq string
	var req_f string
	var credits int64
	var desc string
	qtext := fmt.Sprintf("SELECT DISTINCT * FROM courses")
	rows, _ := database.Query(qtext)

	for rows.Next() {
		rows.Scan(&id, &name, &dep, &cid, &prereq, &req_f, &credits, &desc)
		courses = append(courses, Course{
			Course_id: id,
			Name: name,
			Department: dep,
			Course_number:cid,
			Prereqs: prereq,
			Requirement_fulfilled: req_f,
			Credits: credits,
			Description: desc,
		})
		//fmt.Println(courses)
	}
	return courses
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("endpoint hit")
	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		fmt.Println("error here")
		log.Fatal(err)
	}
    s := retrieveStudentData(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}


func courseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		log.Fatal(err)
	}
	c := retrieveCourseData(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func requirementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		log.Fatal(err)
	}
	req := retrieveRequirementData(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}

func coursesTakenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		log.Fatal(err)
	}
	req := retrieveCoursesTaken(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}

func allCoursesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		log.Fatal(err)
	}
	c := retrieveAllCourses()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "website.html")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	db, err := sql.Open("sqlite3", "./degree.db")
	if err != nil {
		panic(err.Error())
	}
	database = db
	r := mux.NewRouter()
	handler:= cors.Default().Handler(r)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/transaction", transactionHandler)
	r.HandleFunc("/courses", courseHandler)
	r.HandleFunc("/requirements", requirementHandler)
	r.HandleFunc("/courses_taken", coursesTakenHandler)
	r.HandleFunc("/all_courses", allCoursesHandler)
	http.Handle("/", fs)
	log.Panic(http.ListenAndServe(":8080", handler))
}

