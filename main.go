package pretty

import (
	"fmt"
	"regexp"
	"strings"
)

var StyleMap = map[string]string{
	"bold":       "\033[1m",
	"italic":     "\033[3m",
	"underline":  "\033[4m",
	"blink":      "\033[5m",
	"reverse":    "\033[7m",
	"conceal":    "\033[8m",
	"black":      "\033[30m",
	"red":        "\033[31m",
	"green":      "\033[32m",
	"yellow":     "\033[33m",
	"blue":       "\033[34m",
	"magenta":    "\033[35m",
	"cyan":       "\033[36m",
	"white":      "\033[37m",
	"black_bg":   "\033[40m",
	"red_bg":     "\033[41m",
	"green_bg":   "\033[42m",
	"yellow_bg":  "\033[43m",
	"blue_bg":    "\033[44m",
	"magenta_bg": "\033[45m",
	"cyan_bg":    "\033[46m",
	"white_bg":   "\033[47m",
	"reset":      "\033[0m",
}

func Parse(input string) string {
	var outupt string
	var inStyle bool
	var style string
	var prevChar string
	input = preParse(input)
	for _, char := range input {
		if string(char) == "[" && prevChar != "\\" {
			inStyle = true
		} else if string(char) == "]" && prevChar != "\\" {
			inStyle = false
			styles := strings.Split(style, ",")
			style = ""
			for _, style := range styles {
				trimmedStyle := strings.TrimSpace(style)
				lowerStyle := strings.ToLower(trimmedStyle)
				if StyleMap[lowerStyle] != "" {
					outupt += StyleMap[lowerStyle]
				}
			}
		} else if inStyle {
			style += string(char)
		} else {
			outupt += string(char)
		}
		prevChar = string(char)
	}

	// remove slashe if followed by [
	for strings.Contains(outupt, "\\[") {
		outupt = strings.Replace(outupt, "\\[", "[", -1)
	}
	for strings.Contains(outupt, "\\]") {
		outupt = strings.Replace(outupt, "\\]", "]", -1)
	}

	return outupt
}

func Println(input string) {
	println(Parse(input))
}

func Printf(format string, input string) {
	fmt.Printf(format, Parse(input))
}

func Print(input string) {
	print(Parse(input))
}

func Sprintf(format string, input string) string {
	return fmt.Sprintf(format, Parse(input))
}

func preParse(input string) string {
	// if ** is found replace with [bold] and [reset]
	boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
	input = boldRegex.ReplaceAllString(input, "[bold]$1[reset]")
	// if * is found replace with [italic] and [reset]
	italicRegex := regexp.MustCompile(`\*(.*?)\*`)
	input = italicRegex.ReplaceAllString(input, "[italic]$1[reset]")
	// if __ is found replace with [underline] and [reset]
	underlineRegex := regexp.MustCompile(`__(.*?)__`)
	input = underlineRegex.ReplaceAllString(input, "[underline]$1[reset]")
	// if _ is found replace with [italic] and [reset]
	italicRegex = regexp.MustCompile(`_(.*?)_`)
	input = italicRegex.ReplaceAllString(input, "[italic]$1[reset]")
	// if ~~ is found replace with [strike] and [reset]
	strikeRegex := regexp.MustCompile(`~~(.*?)~~`)
	input = strikeRegex.ReplaceAllString(input, "[strike]$1[reset]")
	// if [] is found replace with [reset]
	resetRegex := regexp.MustCompile(`\[\]`)
	input = resetRegex.ReplaceAllString(input, "[reset]")

	return input
}
