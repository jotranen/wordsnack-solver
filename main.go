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

func subPerms(str []string) {

	if len(str) == 1 { return }

	for i := 0; i < len(str); i++ {
		tmp := make([]string, len(str))
		copy(tmp, str)
		tmp = append(tmp[:i], tmp[i+1:]...)
		//fmt.Printf("%d - %s - len: %d\n", i, tmp, len(tmp))
		perm(tmp, 0)
		if len(tmp) > 2 { subPerms(tmp) }
	}
}

func main() {
	file, err := os.Open("sanat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}

	input := strings.ToLower(os.Args[1])

	perm(strings.Split(input, ""), 0)
	subPerms(strings.Split(input, ""))

	for k := range m {
		if words[k] == true {
			fmt.Printf("%s\n", k)
		}
	}
}
