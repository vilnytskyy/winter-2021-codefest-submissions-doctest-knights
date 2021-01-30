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

type Department struct {
	Requirement_id int       `json:"requirement_id,string,omitempty"`
	Name string              `json:"name,string,omitempty"`
	Parent_requirement int   `json:"parent_requirement,string,omitempty"`
	Credits_required int     `json:"credits_required,string,omitempty"`
}
func main() {
	csvFile, _ := os.Open("department.csv")
	database, err := sql.Open("sqlite3", "./degree.db")
	//statement, _ := database.Prepare("DROP TABLE courses")
	// statement.Exec()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS requirements (requirement_id INTEGER, name TEXT, parent_requirement INTEGER, credits_required INTEGER)")
	statement.Exec()
	if err != nil {
		log.Panic(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var courses []Department

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
            log.Fatal(error)
        } else if line[0] == "requirement_id" {
			continue
		}
		cid, _ := strconv.Atoi(line[0])
		cn, _ := strconv.Atoi(line[2])
		cr, _ := strconv.Atoi(line[3])
        courses = append(courses, Department{
            Requirement_id: cid,
            Name:  line[1],
            Parent_requirement: cn,
			Credits_required: cr,
        })
	}

	statement, _ = database.Prepare("INSERT INTO requirements (requirement_id, name, parent_requirement, credits_required) VALUES (?, ?, ?, ?)")
	for _, course := range courses {
		statement.Exec(course.Requirement_id, course.Name, course.Parent_requirement, course.Credits_required)
		/*fmt.Println(course.Requirement_id)
		fmt.Println(course.Name)
		fmt.Println(course.Parent_requirement)
		fmt.Println(course.Credits_required)
		fmt.Println(course.Prereqs)
		fmt.Println(course.Requirement_fulfilled)
		fmt.Println(course.Credits)
		fmt.Println(course.Description)*/
	}
	rows, _ := database.Query("SELECT DISTINCT name, credits_required FROM requirements")
	var dep string
	var cn int
	for rows.Next() {
        rows.Scan(&dep, &cn)
		fmt.Println(dep + " " + strconv.Itoa(cn))
    }

}

