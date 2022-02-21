package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yosssi/ace"
)

type Result struct {
	Status bool
}

func addCourse(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("template/base", "template/add_course", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addCourseHanlder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cName := r.FormValue("cname")
	pName := r.FormValue("pname")
	desc := r.FormValue("desc")
	//r.Body("uname")
	var course Course

	course.Name = cName
	course.ProfessorName = pName
	course.Description = desc

	db.Save(&course)

	var out Result
	out.Status = true
	json.NewEncoder(w).Encode(out)

}

func allocateCourseHanlder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	courseId, _ := strconv.Atoi(r.FormValue("courseId"))
	sId, _ := strconv.Atoi(mux.Vars(r)["id"])
	//r.Body("uname")
	var course Course
	db.Where("id= ?", courseId).Find(&course)

	var sCourse CourseStudent
	sCourse.CourseID = courseId
	sCourse.StudentID = sId
	sCourse.CourseName = course.Name

	db.Save(&sCourse)
	type Result struct {
		Status bool
	}

	var out Result
	out.Status = true
	json.NewEncoder(w).Encode(out)

}

func showDeleteCourse(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("template/base", "template/delete_course", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Out struct {
		Courses []Course
	}
	var out Out
	db.Find(&out.Courses)
	if err := tpl.Execute(w, out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteCourseHanlder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	courseId, _ := strconv.Atoi(r.FormValue("courseId"))
	//r.Body("uname")
	db.Where("id= ?", courseId).Delete(&Course{})
	db.Where("course_id= ?", courseId).Delete(&CourseStudent{})

	

	var out Result
	out.Status = true
	json.NewEncoder(w).Encode(out)

}

func addStudent(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("template/base", "template/add_student", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showStudents(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("template/base", "template/show_students", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var out []Student
	db.Preload("Courses").Find(&out)
	if err := tpl.Execute(w, out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func allocateStudents(w http.ResponseWriter, r *http.Request) {
	sId := mux.Vars(r)["id"]
	tpl, err := ace.Load("template/base", "template/allocate_course", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Out struct {
		Student Student
		Courses []Course
	}
	var out Out
	db.Preload("Courses").Where("id = ?", sId).Find(&out.Student)
	db.Find(&out.Courses)
	if err := tpl.Execute(w, out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addStudentHanlder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sname := r.FormValue("sname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	//r.Body("uname")
	var stu Student

	stu.Name = sname
	stu.Email = email
	stu.Phone = phone

	db.Save(&stu)

	var out Result
	out.Status = true
	json.NewEncoder(w).Encode(out)

}
