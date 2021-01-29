package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
    "bufio"
    "encoding/csv"
    // "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
	"strconv"
)
type Course struct {
	Course_id int                 `json:"course_id,string,omitempty"`
	Name string                   `json:"name,omitempty"`
	Department string             `json:"department,string,omitempty"`
	Course_number int             `json:"course_number,string,omitempty"`
	Prereqs string                `json:"prereqs,string,omitempty"`
	Requirement_fulfilled string  `json:"requirement_fulfilled,string,omitempty"`
	Credits int                   `json:"credits,string,omitempty"`
	Description string            `json:"major_gpades,omitempty"`
}
func main() {
	csvFile, _ := os.Open("courses.csv")
	database, err := sql.Open("sqlite3", "./degree.db")
	statement, _ := database.Prepare("DROP TABLE courses")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS courses (course_id INTEGER, name TEXT, department TEXT, course_number INTEGER, prereqs TEXT, requirement_fulfilled TEXT, credits INTEGER, description TEXT)")
	statement.Exec()
	if err != nil {
		log.Panic(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var courses []Course

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
            log.Fatal(error)
        } else if line[0] == "course_id" {
			continue
		}
		cid, _ := strconv.Atoi(line[0])
		cn, _ := strconv.Atoi(line[3])
		cr, _ := strconv.Atoi(line[6])
        courses = append(courses, Course{
            Course_id: cid,
            Name:  line[1],
            Department: line[2],
			Course_number: cn,
			Prereqs: line[4],
			Requirement_fulfilled: line[5],
			Credits: cr,
			Description: line[7],
        })
	}

	statement, _ = database.Prepare("INSERT INTO courses (course_id, name, department, course_number, prereqs, requirement_fulfilled, credits, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	for _, course := range courses {
		statement.Exec(course.Course_id, course.Name, course.Department, course.Course_number, course.Prereqs, course.Requirement_fulfilled, course.Credits, course.Description)
		/*fmt.Println(course.Course_id)
		fmt.Println(course.Name)
		fmt.Println(course.Department)
		fmt.Println(course.Course_number)
		fmt.Println(course.Prereqs)
		fmt.Println(course.Requirement_fulfilled)
		fmt.Println(course.Credits)
		fmt.Println(course.Description)*/
	}
	rows, _ := database.Query("SELECT DISTINCT department, course_number, name FROM courses")
	var dep string
	var cn int
	var name string
	for rows.Next() {
        rows.Scan(&dep, &cn, &name)
		fmt.Println(dep + " " + strconv.Itoa(cn) + " " + name)
    }

}

