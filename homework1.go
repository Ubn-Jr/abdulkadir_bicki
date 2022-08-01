package main

import "fmt"

func main() {

	Firstexam := 98
	Secondexam := 80
	//TODO
	//golang dilinde değişken ve fonksiyon isimlerinin büyük harfle başlaması özel bir durumdur
	// camelCase yöntemi en ideal kullanım biçimidir
	// tüm değişken ve fonksiyon isimlerinde camelCase yöntemini uygulayalım
	Average := TaketheAverage(Firstexam, Secondexam)
	Receivedocument(Average)
	println(Average)
}

func TaketheAverage(Firstexam int, Secondexam int) int {

	return (Firstexam + Secondexam) / 2

}
func Receiveddocument(Gradeaverage int) {
	if Gradeaverage >= 85 {
		fmt.Println("You have earned a certificate of appreciation")
	} else if Gradeaverage >= 70 {
		fmt.Println("You have earned a certificate of thank")
	} else {
		fmt.Println("You did not qualify for the document")
	}
}
