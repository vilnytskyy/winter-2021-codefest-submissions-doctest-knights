package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	// "bufio"
    //"encoding/csv"
    // "encoding/json"
    "fmt"
    //"io"
    "log"
    //"os"
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

type Department struct {
	Requirement_id int       `json:"requirement_id,string,omitempty"`
	Name string              `json:"name,string,omitempty"`
	Parent_requirement int   `json:"parent_requirement,string,omitempty"`
	Credits_required int     `json:"credits_required,string,omitempty"`
}

type Courses2 struct {
	Student_id int
	Course_id int
	In_progress int
	Grade string
}
func main() {
	//csvFile, _ := os.Open("courses_taken2.csv")
	database, err := sql.Open("sqlite3", "./degree.db")
	if (err != nil) {
		log.Panic(err)
	}
	//statement, _ := database.Prepare("DROP TABLE courses_taken")
	//statement.Exec()
	//statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS courses_taken (student_id INTEGER, course_id INTEGER, in_progress INTEGER, grade TEXT)")
	//statement.Exec()
	//if err != nil {
	//	log.Panic(err)
	//}
	/*reader := csv.NewReader(bufio.NewReader(csvFile))
	var courses []Courses2

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
            log.Fatal(error)
        } else if line[0] == "student_id" {
			continue
		}
		cid, _ := strconv.Atoi(line[0])
		cn, _ := strconv.Atoi(line[1])
		cr, _ := strconv.Atoi(line[2])
        courses = append(courses, Courses2{
            Student_id: cid,
            Course_id:  cn,
            In_progress: cr,
			Grade: line[3],
        })
	}*/

	statement, _ := database.Prepare("INSERT INTO students (student_id, name, major_id, credits, overall_gpa, major_gpa, username, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	//for _, course := range courses {
	statement.Exec(20000014, "Hammad Siddiqui", 1, 69, 3.78, 3.82, "bb", "gg")
		/*fmt.Println(course.Student_id)
		fmt.Println(course.Course_id)
		fmt.Println(course.In_progress)
		fmt.Println(course.Grade)
		fmt.Println(course.Prereqs)
		fmt.Println(course.Requirement_fulfilled)
		fmt.Println(course.Credits)
		fmt.Println(course.Description)*/
	//}
	rows, _ := database.Query("SELECT student_id, credits, major_id, name FROM students")
	var dep int
	var cn int
	var inp int
	var g string
	for rows.Next() {
        rows.Scan(&dep, &cn, &inp, &g)
		fmt.Println(strconv.Itoa(dep) + " " + strconv.Itoa(cn) + " " + strconv.Itoa(inp) + " " + g)
    }

}

