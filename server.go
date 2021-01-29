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
	Course_id int64               `json:"student_id,string,omitempty"`
	Name string                   `json:"name,omitempty"`
	Department string             `json:"department,string,omitempty"`
	Course_number int64            `json:"course_number,string,omitempty"`
	Prereqs string                `json:"prereqs,string,omitempty"`
	Requirement_fulfilled string  `json:"requirement_fulfilled,string,omitempty"`
	Credits int64                 `json:"credits,string,omitempty"`
	Description string            `json:"major_gpades,omitempty"`
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
	fmt.Println(c)
	return c
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
	http.Handle("/", fs)
	log.Panic(http.ListenAndServe(":8080", handler))
}

