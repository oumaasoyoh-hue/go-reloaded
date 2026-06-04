package main

import (
	"strconv"
	"strings"
)

func ProcessText(input string) string {
	words := strings.Fields(input)

	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(bin)":
			if i > 0 {
				val, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		default:
			if strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(cap,") {
				if i+1 < len(words) {
					numStr := strings.TrimRight(words[i+1], ")")
					count, err := strconv.Atoi(numStr)
					if err == nil {
						tagType := words[i]
						start := i - count
						if start < 0 {
							start = 0
						}
						for j := start; j < i; j++ {
							if strings.HasPrefix(tagType, "(up") {
								words[j] = strings.ToUpper(words[j])
							} else if strings.HasPrefix(tagType, "(low") {
								words[j] = strings.ToLower(words[j])
							} else if strings.HasPrefix(tagType, "(cap") {
								words[j] = capitalize(words[j])
							}
						}
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := strings.ToLower(words[i+1])
			if len(nextWord) > 0 {
				r := nextWord[0]
				if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' {
					if words[i] == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}

	joinedText := strings.Join(words, " ")

	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		joinedText = strings.ReplaceAll(joinedText, " "+p, p)
		joinedText = strings.ReplaceAll(joinedText, p, p+" ")
	}

	joinedText = strings.ReplaceAll(joinedText, ". . .", "...")
	joinedText = strings.ReplaceAll(joinedText, "! ?", "!?")

	fields := strings.Fields(joinedText)
	joinedText = strings.Join(fields, " ")

	for _, p := range punctuations {
		joinedText = strings.ReplaceAll(joinedText, p, p+" ")
	}

	finalWords := strings.Fields(joinedText)
	quoteOpen := false
	for i := 0; i < len(finalWords); i++ {
		if finalWords[i] == "'" {
			if !quoteOpen {
				if i+1 < len(finalWords) {
					finalWords[i+1] = "'" + finalWords[i+1]
					finalWords = append(finalWords[:i], finalWords[i+1:]...)
					quoteOpen = true
					i--
				}
			} else {
				if i > 0 {
					finalWords[i-1] = finalWords[i-1] + "'"
					finalWords = append(finalWords[:i], finalWords[i+1:]...)
					quoteOpen = false
					i--
				}
			}
		}
	}

	outputText := strings.Join(finalWords, " ")
	for _, p := range punctuations {
		outputText = strings.ReplaceAll(outputText, " "+p, p)
	}

	return outputText
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
