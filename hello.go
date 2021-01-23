package main

import (
    "database/sql"
    "fmt"
    //"strconv"
	//"log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    database, _ := sql.Open("sqlite3", "./degree.db")
    //statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS students (student_id INTEGER, name TEXT, major_id INTEGER, credits INTEGER, overall_gpa REAL, major_gpa REAL, username TEXT, password TEXT)")
	//statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS courses (course_id INTEGER, name TEXT, department TEXT, course_number INTEGER, prereqs TEXT, requirement_fulfilled TEXT, credits INTEGER, description TEXT)")
	//statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS courses_taken (student_id INTEGER, course_id INTEGER, in_progress INTEGER, grade TEXT)")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS majors (major_id INTEGER, name TEXT, prereqs TEXT, required_courses TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO majors (major_id, name, prereqs, required_courses) VALUES (?, ?, ?, ?)")
    //statement.Exec(1, "Software and Analysis I", "CSCI", 135, "2,3", "cs entry", 4, "This course for prospective computer science majors and minors concentrates on problem-solving techniques using a high-level programming language.")
	statement.Exec(1, "Computer Science", "1,2,3,4", "1,2,3,4,5,6,7")
    rows, _ := database.Query("SELECT DISTINCT name, prereqs FROM majors")
    //var student_id int
    //var course_id int
    var grade string
	var name string
	//var course_number int
	//var student_name string
    for rows.Next() {
        rows.Scan(&grade, &name)
		//qtext := fmt.Sprintf("SELECT DISTINCT department, course_number FROM courses WHERE course_id LIKE %d", course_id)
		//qtext2 := fmt.Sprintf("SELECT DISTINCT name FROM students WHERE student_id LIKE %d", student_id)
		//query2 := database.QueryRow(qtext).Scan(&name, &course_number)
		//query3 := database.QueryRow(qtext2).Scan(&student_name)
		//if query2 != nil {
		//	log.Fatal(query2)
		//} else if query3 != nil {
		//	log.Fatal(query3)
		//}
        fmt.Println(grade + " " + name)
    }
}
