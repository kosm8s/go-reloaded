package src

import (
	"strconv"
	"strings"
)

type Reload struct {
	Data string
}

func (r *Reload) Articles() {
	slice := strings.Fields(r.Data)
	vowels := []rune{'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u', 'Y', 'y', 'H', 'h'}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "an" { // change all articles to "a" or "A" form
			slice[i] = "a"
		} else if slice[i] == "An" {
			slice[i] = "A"
		}
	}

	for i := 0; i < len(slice); i++ { // change to proper article
		for j := 0; j < len(vowels); j++ {
			if slice[i] == "A" && slice[i+1][0] == byte(vowels[j]) {
				slice[i] = "An"
			} else if slice[i] == "a" && slice[i+1][0] == byte(vowels[j]) {
				slice[i] = "an"
			}
		}
	}

	r.Data = strings.Join(slice, " ")
}

func (r *Reload) Convert() {
	slice := strings.Fields(r.Data)
	sKeys := []string{"(bin)", "(hex)", "(cap)", "(low)", "(up)"}
	sPuncts := []string{".", ",", "!", "?", "...", "!?", ":", ";", "'", "“"}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(sKeys); j++ {
			if slice[0] == sKeys[j] { // first word exeption
				if len(slice) == 1 {
					r.Data = ""
					return
				} else {
					slice = slice[1:]
				}
			}
		}

		if slice[0] == "(cap," || slice[0] == "(up," || slice[0] == "(low," {
			slice = slice[2:]
		}
		if slice[i] == "(bin)" { // converting bin to deс
			if isBinary(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 2, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
			}
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if slice[i] == "(hex)" { // converting hex to deс
			if isHex(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 16, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
			}

			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if slice[i] == "(up)" || slice[i] == "(up)," { // converting to uppercase
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if slice[i] == "(cap)" || slice[i] == "(cap)," { // titling
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i-1] = strings.Title(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if slice[i] == "(low)" || slice[i] == "(low)," { // converting to lowercase
			slice[i-1] = strings.ToLower(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if strings.Contains(slice[i], "(cap,") && slice[i] == "(cap," { // titling n words
			tmp := []rune(slice[i+1])
			tmp = tmp[:len(tmp)-1]

			number, _ := strconv.Atoi(string(tmp))

			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				for u := 0; u < len(sPuncts); u++ {
					if slice[i-j] == sPuncts[u] {
						number++
					}
				}
				slice[i-j] = strings.ToLower(slice[i-j])
				slice[i-j] = strings.Title(slice[i-j])
			}

			slice = append(slice[:i], slice[i+2:]...)
			i--
		}

		if strings.Contains(slice[i], "(low,") && slice[i] == "(low," { // converting to lowercase n words
			tmp := []rune(slice[i+1])
			tmp = tmp[:len(tmp)-1]

			number, _ := strconv.Atoi(string(tmp))

			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				for u := 0; u < len(sPuncts); u++ {
					if slice[i-j] == sPuncts[u] {
						number++
					}
				}
				slice[i-j] = strings.ToLower(slice[i-j])
			}

			slice = append(slice[:i], slice[i+2:]...)
			i--
		}

		if strings.Contains(slice[i], "(up,") { // converting to uppercase n words
			tmp := []rune(slice[i+1])
			tmp = tmp[:len(tmp)-1]

			number, _ := strconv.Atoi(string(tmp))

			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				for u := 0; u < len(sPuncts); u++ {
					if slice[i-j] == sPuncts[u] {
						number++
					}
				}
				slice[i-j] = strings.ToUpper(slice[i-j])
			}

			slice = append(slice[:i], slice[i+2:]...)
			i--
		}
	}

	r.Data = strings.Join(slice, " ")
}

func isBinary(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] != '0' && r[i] != '1' {
			return false
		}
	}
	return true
}

func isHex(s string) bool {
	r := []rune(s)
	notHex := false
	for i := 0; i < len(r); i++ {
		if (r[i] >= '0' && r[i] <= '9') || (r[i] >= 'a' && r[i] <= 'f') || (r[i] >= 'A' && r[i] <= 'F') {
			continue
		} else {
			notHex = true
		}
	}
	return !notHex
}

func (r *Reload) Apostrophe() {
	runes := []rune(r.Data)
	num := 0

	for i := 0; i < len(runes); i++ { // counting apostrophes in text
		if runes[i] == '\'' || runes[i] == 8216 {
			num++
		}
	}
	i := 0
	for j := 0; j < num/2; j++ {
		count := 0
		for ; i < len(runes)-1; i++ { // removing spaces from rightside
			if (runes[i] == '\'' && runes[i+1] == ' ') || (runes[i] == 8216 && runes[i+1] == ' ') {
				for j := i + 1; runes[j] == ' '; j++ {
					count++
				}
				runes = append(runes[:i+1], runes[i+count+1:]...)
				break
			}
		}
		count = 0
		i++
		for ; i < len(runes); i++ { // removing spaces from leftside
			if (runes[i] == '\'' && runes[i-1] == ' ') || (runes[i] == 8216 && runes[i-1] == ' ') {
				for j := i - 1; runes[j] == ' '; j-- {
					count++
				}
				runes = append(runes[:i-count], runes[i:]...)
				break
			}
		}
		count = 0
	}

	r.Data = string(runes)
}
