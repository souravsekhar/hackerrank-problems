package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	mag := map[string]int{}
	for _, v := range magazine {
		if _, found := mag[string(v)]; found {
			mag[string(v)] = mag[string(v)] + 1
		} else {
			mag[string(v)] = 1
		}
	}

	noteMap := map[string]int{}
	for _, v := range note {
		if _, found := noteMap[string(v)]; found {
			noteMap[string(v)] = noteMap[string(v)] + 1
		} else {
			noteMap[string(v)] = 1
		}
	}

	for key := range noteMap {
		if _, ok := mag[key]; ok {
			if mag[key] < noteMap[key] {
				fmt.Println("No")
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func main() {
	// m := []string{"two", "times", "three", "is", "not", "four", "two"}
	m := []string{"h", "ghq", "g", "xxy", "wdnr", "anjst", "xxy", "wdnr", "h", "h", "anjst", "wdnr"}
	// n := []string{"two", "times", "two", "is", "four"}
	n := []string{"h", "ghq"}
	checkMagazine(m, n)

}
