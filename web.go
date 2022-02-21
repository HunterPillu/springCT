package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yosssi/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("template/base", "template/home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleWeb() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/course/add", addCourse).Methods("GET")
	r.HandleFunc("/course/add", addCourseHanlder).Methods("POST")

	r.HandleFunc("/course/delete", showDeleteCourse).Methods("GET")
	r.HandleFunc("/course/delete", deleteCourseHanlder).Methods("POST")

	r.HandleFunc("/student/add", addStudent).Methods("GET")
	r.HandleFunc("/student/add", addStudentHanlder).Methods("POST")

	r.HandleFunc("/student/list", showStudents).Methods("GET")

	r.HandleFunc("/student/allocate/{id}", allocateStudents).Methods("GET")
	r.HandleFunc("/student/allocate/{id}", allocateCourseHanlder).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	//go func() {
	fmt.Println("Starting Server")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
