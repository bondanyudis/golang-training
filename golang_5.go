package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var a, b, c, d int

// func main() {

// 	for {
// 		sample()
// 	}
// }
// func sample() {
// 	fmt.Println("========================================================================")
// 	fmt.Println("Press 1 to run segitiga astreik")
// 	fmt.Println("Press 2 faktor Bilangan")
// 	fmt.Println("Press 3 Bilangan prima")
// 	fmt.Println("Press 4 palindrome")
// 	fmt.Println("Press 5 keluar")
// 	fmt.Println("========================================================================")
// 	var inputtinggi int
// 	fmt.Println("Masukan inputan anda :")
// 	n, err := fmt.Scanln(&inputtinggi)
// 	if n < 1 || err != nil {
// 		fmt.Println("invalid inputtinggi")
// 		return
// 	}
// 	switch inputtinggi {
// 	case 1:
// 		segitiga_asterik()
// 	case 2:
// 		faktorBilangan()
// 		// os.Exit(2)
// 	case 3:
// 		bilanganPrima()
// 	case 4:
// 		palindrome()
// 	case 5:
// 		os.Exit(2)
// 	default:
// 		fmt.Println("pilihan menu tidak ada:")
// 	}
// }

// func segitiga_asterik() {
// 	var n int
// 	fmt.Println("masukkan inputan tinggi segitiga : ")
// 	fmt.Scanf("%v", &n)
// 	for a = 1; a <= n; a++ {
// 		for b = a; b <= n; b++ {
// 			fmt.Printf(" ")
// 		}
// 		for c = 1; c <= a; c++ {
// 			fmt.Printf("%s", "* ")
// 		}
// 		fmt.Printf("\n")
// 	}

// }

// func faktorBilangan() {
// 	var bilangan int
// 	fmt.Println("masukkan bilangan  : ")
// 	fmt.Scanf("%v", &bilangan)
// 	for i := 1; i <= bilangan; i++ {
// 		if bilangan%i == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func bilanganPrima() {
// 	var prima int
// 	var counter int
// 	fmt.Println("masukkan inputan bilangan : ")
// 	fmt.Scanf("%v", &prima)
// 	for i := 1; i <= prima; i++ {
// 		if prima%i == 0 {
// 			counter++
// 		}
// 	}
// 	if counter > 2 || prima == 1 {
// 		fmt.Println("bukan prima")
// 	} else {
// 		fmt.Println("prima")
// 	}
// }
// func palindrome() {
// 	var palindrome string
// 	var palindrome_kebalik string
// 	fmt.Println("masukkan suatu string  : ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan() // use `for scanner.Scan()` to keep reading
// 	palindrome = scanner.Text()
// 	// fmt.Scanf("%s", &palindrome)
// 	for i := len(palindrome); i > 0; i-- {
// 		palindrome_kebalik += "" + string(palindrome[i-1])
// 	}
// 	if palindrome_kebalik == palindrome {
// 		fmt.Println("palindrome")
// 	} else {
// 		fmt.Println("bukan palindrome")
// 	}
// }
