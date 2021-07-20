package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


func main() {
	min := 0
	max := 100
	numberOfTries := 0

	numberToGuess := getRandNumber(min, max)
	userNumber := askQuestion()

	for userNumber != numberToGuess {
		var message string

		if userNumber < min || userNumber > max {
			message = "En dehors des limites !"
		}	else if userNumber < numberToGuess {
			message = "C'est plus !"
		} else if userNumber > numberToGuess {
			message = "C'est moins !"
		}

		fmt.Println(message)
		numberOfTries++
		userNumber = askQuestion("Choisissez un nombre: ")
	}

	fmt.Println("Bravo vous avez trouvé en " + strconv.Itoa(numberOfTries) + " essais !")
}

func getRandNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min + 1) - min)
}

func askQuestion(question ...string) int {
	var userNumber int
	quest := "Choisissez un nombre entre 0 et 100: "

	if len(question) > 0 {
		quest = question[0]
	}

	fmt.Print(quest)
	fmt.Scan(&userNumber)

	return userNumber
}