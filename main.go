// guess - игра, в которой игрок должен угадать случайное число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число от 1 до 100")
	fmt.Println("Cможешь ли ты угадать его?")
	//fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У тебя есть", 10-guesses, "попыток.")
		fmt.Print("Угадай число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Моё число больше.")
		} else if guess > target {
			fmt.Println("Oops. Моё число меньше.")
		} else {
			success = true
			fmt.Println("Отличная работа, Ты угадал!")
			break
		}
	}
	if !success {
		fmt.Println("Извини ты не угадал. я загадал число:", target)
	}
	time.Sleep(5 * time.Second)
}
