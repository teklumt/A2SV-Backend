package main

import (
	"fmt"
	"strings"
)

func main() {
	var studentFirstname string
	var studentLastname string
	var numberOfSubjects int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&studentFirstname)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&studentLastname)

	fmt.Print("Enter the number of subjects: ")
	fmt.Scan(&numberOfSubjects)

	subjectData := make(map[string]int)

	for i := 0; i < numberOfSubjects; i++ {
		var subjectName string
		var grade int

		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scan(&subjectName)

		fmt.Printf("Enter the grade for %s: ", subjectName)
		fmt.Scan(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade, please enter a grade between 0 and 100.")
			i-- 
			continue
		}

		subjectData[subjectName] = grade
	}


	fmt.Printf("\n%-30s\n", strings.Repeat("-", 30))
	fmt.Printf("%-30s\n", "Student Grade Report")
	fmt.Printf("%-30s\n", strings.Repeat("-", 30))
	fmt.Printf("Student Name: %s %s\n", studentFirstname, studentLastname)
	fmt.Printf("%-30s\n", strings.Repeat("-", 30))
	fmt.Printf("%-20s %-10s\n", "Subject", "Grade")
	fmt.Printf("%-30s\n", strings.Repeat("-", 30))

	total := 0
	for subject, grade := range subjectData {
		fmt.Printf("%-20s %-10d\n", subject, grade)
		total += grade
	}

	average := float64(total) / float64(numberOfSubjects)
	fmt.Printf("%-30s\n", strings.Repeat("-", 30))
	fmt.Printf("%-20s %-10.2f\n", "Average", average)
	fmt.Printf("%-30s\n", strings.Repeat("-", 30))
}
