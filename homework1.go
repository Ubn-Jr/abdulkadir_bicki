package main

import "fmt"

func main() {

	firstExam := 98
	secondExam := 80
	Average := takeTheAverage(firstExam, secondExam)
	receivedDocument(Average)
	println(Average)
}

func takeTheAverage(firstExam int, secondExam int) int {

	return (firstExam + secondExam) / 2

}
func receivedDocument(gradeAverage int) {
	if gradeAverage >= 85 {
		fmt.Println("You have earned a certificate of appreciation")
	} else if gradeAverage >= 70 {
		fmt.Println("You have earned a certificate of thank")
	} else {
		fmt.Println("You did not qualify for the document")
	}
}
