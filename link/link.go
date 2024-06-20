package link

import (
	"os"
	"strings"
)

func Linking() []string {
	file, _ := os.ReadFile("data.txt")
	words := strings.Split(string([]rune(string(file))[:]), " ")
	return words
}
