package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"html/template"
	"net/http"

	"main.go/Desktop/WORDLY/link"
)

var words = link.Linking()

func randomWord() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(words))
	return string(words[randomIndex])

}

func checkLetters(word string, guess string) string {
	var positions []string
	for i := 0; i < len([]rune(word)) && i < len([]rune(guess)); i++ {
		if string([]rune(word)[i]) == string([]rune(guess)[i]) {

			positions = append(positions, strconv.Itoa(i+1))
		}
	}

	return fmt.Sprint(positions)
}

func intersectionLetter(word, guess string) []string {
	arr1 := string([]rune(word)[:])
	arr2 := string([]rune(guess)[:])
	out := []string{}
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				out = append(out, string(i))
			}
		}
	}
	return out
}

func homePage(w http.ResponseWriter, r *http.Request) {
	a := ""
	tmpl, _ := template.ParseFiles("homePage.html")
	tmpl.Execute(w, a)
}

func playPage(w http.ResponseWriter, r *http.Request) {
	a := ""
	tmpl, _ := template.ParseFiles("playPage.html")
	tmpl.Execute(w, a)
}

var word = randomWord()
var attempts = 0

func savePage(w http.ResponseWriter, r *http.Request) {

	attempts++
	name := r.FormValue("name")
	if attempts > 6 {
		fmt.Fprintf(w, "Попытки кончились, Вы проиглали((")
		attempts = 0
		word = randomWord()
	} else if name == word {
		fmt.Fprintf(w, "Поздравляю! Вы угадали с %d попытки.", attempts)
		attempts = 0
		word = randomWord()
	} else {
		correctPositions := checkLetters(word, name)
		correctLetters := intersectionLetter(word, name)

		fmt.Fprintf(w, "Верные позиции букв: %s", correctPositions)
		fmt.Fprintf(w, "Верные буквы: %v", correctLetters)
		fmt.Fprintf(w, "Это была %d из 6 попыток", attempts)
	}
	fmt.Print(word)
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func hendlerRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playPage)
	http.HandleFunc("/save", savePage)
	http.ListenAndServe(":1212", nil)
}

func main() {
	hendlerRequest()
}
