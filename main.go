// Jari Otranen, 2017
//
// Simple tool to find all possible string permutations and map then against dictionary
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
)

var m map[string]bool
var words map[string]bool

func init() {
	m = make(map[string]bool)
	words = make(map[string]bool)
}

// permutes all string combinations
func perm(str []string, i int) {
	if i == len(str) {
		m[strings.Join(str, "")] = true
	} else {
		for j := i; j < len(str); j++ {
			str[i], str[j] = str[j], str[i]
			perm(str, i+1)
			str[i], str[j] = str[j], str[i]
		}
	}
}

// all substrings with length >= 2
func subPerms(str []string) {

	if len(str) == 1 { return }

	for i := 0; i < len(str); i++ {
		tmp := make([]string, len(str))
		copy(tmp, str)
		tmp = append(tmp[:i], tmp[i+1:]...)
		perm(tmp, 0)
		if len(tmp) > 2 { subPerms(tmp) }
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage  : wordsnack-solver <language-file> kala\n")
		fmt.Printf("Example: wordsnack-solver words/finnish.txt kala\n")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}

	input := strings.ToLower(os.Args[2])

	perm(strings.Split(input, ""), 0)
	subPerms(strings.Split(input, ""))

	for k := range m {
		if words[k] == true {
			fmt.Printf("%s\n", k)
		}
	}
}
