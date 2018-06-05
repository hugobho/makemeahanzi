package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type EtymologyStruct struct {
	Type string `json:"type,omitempty"`
	Hint string `json:"hint,omitempty"`
}

type Chinese struct {
	Character     string          `json:"character,omitempty"`
	Definition    string          `json:"definition,omitempty"`
	Pinyin        string          `json:"pinyin,omitempty"`
	Decomposition string          `json:"decomposition,omitempty"`
	Radical       string          `json:"radical,omitempty"`
	Etymology     EtymologyStruct `json:"etymology,omitempty"`
	Matches       interface{}     `json:"matches,omitempty"`
}

func compileRadicals(characters []Chinese) {
	count := map[string]int{}
	for _, character := range characters {
		count[character.Radical] = count[character.Radical] + 1
	}
	fmt.Println("=== RADICALS ===")
	printSorted(count)
}

func compileDecomposition(characters []Chinese) {
	count := map[string]int{}
	for _, character := range characters {
		decomposition := string([]rune(character.Decomposition)[0])
		count[decomposition] = count[decomposition] + 1
	}
	fmt.Println("=== DECOMPOSITION ===")
	printSorted(count)
}

func compileEtymology(characters []Chinese) {
	count := map[string]int{}
	for _, character := range characters {
		if character.Etymology != (EtymologyStruct{}) {
			etType := character.Etymology.Type
			count[etType] = count[etType] + 1
		}
	}
	fmt.Println("=== ETYMOLOGY ===")
	printSorted(count)
}

func printSorted(m map[string]int) {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// for _, kv := range ss {
	// 	fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	// }

	printMax := 9
	if len(ss) < printMax {
		printMax = len(ss) - 1
	}

	for i := 0; i <= printMax; i++ {
		fmt.Printf("%s: %d\n", ss[i].Key, ss[i].Value)
	}
}

func main() {

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	dat, err := ioutil.ReadFile("./dictionary.txt")
	check(err)
	stringifiedData := string(dat)
	splitData := strings.Split(stringifiedData, "\n")
	splitData = splitData[:len(splitData)-1]
	wrapped := fmt.Sprintf("[ %s ]", strings.Join(splitData, ", "))
	var characters []Chinese
	json.Unmarshal([]byte(wrapped), &characters)
	fmt.Println()
	compileRadicals(characters)
	fmt.Println()
	compileDecomposition(characters)
	fmt.Println()
	compileEtymology(characters)
}
