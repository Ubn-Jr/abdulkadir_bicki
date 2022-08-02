package main

import "fmt"

func main() {

	Firstexam := 98
	Secondexam := 80
	//TODO
	//golang dilinde değişken ve fonksiyon isimlerinin büyük harfle başlaması özel bir durumdur
	// camelCase yöntemi en ideal kullanım biçimidir
	// tüm değişken ve fonksiyon isimlerinde camelCase yöntemini uygulayalım
	average := taketheAverage(firstExam, secondExam)
receiveDocument(average)
	println(average)
}

func taketheAverage(firstExam int, secondExam int) int {

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
