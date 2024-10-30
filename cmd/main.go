package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func main() {
	rarFilePath := "Путь к вашему RAR-файлу"
	fileName := "Файл внутри rar"

	var pwdGen <-chan string = brute(16, 17)

	fmt.Println("Start cracking")
	fmt.Println("Please wait...")

	crackRar(rarFilePath, fileName, pwdGen)
}
func crackRar(rarFilePath, fileName string, pwdGenerator <-chan string) {
	for pwd := range pwdGenerator {
		//log.Println(pwd)
		cmd := exec.Command("C:\\Games\\UnRAR.exe", "x", "-p"+pwd, rarFilePath, fileName)
		err := cmd.Run()
		if err == nil {
			fmt.Println("File extracted")
			fmt.Printf("The password is %s\n", pwd)
			return
		}
	}
	fmt.Println("Sorry, cannot find the password")
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func brute(startLength, length int) <-chan string {
	pwdGenerator := make(chan string)

	go func() {
		charset := []rune("0123456789абвгде")
		base := len(charset)

		var wg sync.WaitGroup

		for count := startLength; count <= length; count++ {
			maxCombinations := pow(base, count)
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Println("+", count, maxCombinations)
				for i := 0; i < maxCombinations; i++ {
					log.Println("?", count)
					var password []rune
					n := i
					for j := 0; j < count; j++ {
						password = append(password, charset[n%base])
						n /= base
					}
					pwdGenerator <- string(password)
				}
				log.Println("-", count)
			}()

		}
		wg.Wait()
		close(pwdGenerator)
	}()

	return pwdGenerator
}
