package main

import "gorm.io/gorm"

// Add a Course [Name, Professor Name, Description]
//  Add a Student [Name, Email, Phone]
//  Allocate a Student to one or more Course
//  List Students [Name, Email, Phone, Course Enrolled (comma separated string)]
//  Delete a given Course
//  Get a Student’s data for a given Course

type Student struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Courses []CourseStudent
}

type Course struct {
	gorm.Model
	Name          string
	ProfessorName string
	Description   string
	StudentID     uint
}

type CourseStudent struct {
	gorm.Model
	CourseID   int
	StudentID  int
	CourseName string
}
