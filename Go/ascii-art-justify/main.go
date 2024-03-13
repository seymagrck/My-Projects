package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	filename := "standard"
	if len(os.Args) > 4 || len(os.Args) < 3 {
		fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]")
	} else if os.Args[1] == "--align=center" || os.Args[1] == "--align=left" || os.Args[1] == "--align=right" || os.Args[1] == "--align=justify" {
		if len(os.Args) == 4 {
			filename = os.Args[3]
		}
		file, err := os.Open(filename + ".txt") // to open file
		if err != nil {
			log.Fatal(err)
		}
		align := os.Args[1][8:]
		width, err := getTerminalWidth()
		if err != nil {
			fmt.Println("Terminal genişliği alınamadı:", err)
			return
		}

		defer file.Close()
		text := os.Args[2]
		words := strings.Split(text, "\\n") // to split the input word into words
		if text == "" {                     // at the beginning, if there is no input word, return
			return
		}

		for i, word := range words {
			if word == "" { // to control if the word is empty return a new line
				if i < len(words) { // to control if the word is empty and the next word exist
					fmt.Println()
				}
				continue
			}
			wordtrimspace := strings.TrimSpace(word)
			lines := make([]string, 8) // to create lines with 8 line, 1 for each char of the word because our char containe 8 lines
			for _, char := range wordtrimspace {
				for i := 0; i <= 7; i++ { // the string (lines) in our code needs 8 lines
					file.Seek(0, 0)                   // to reset the file, in order to read again
					start := (int(char)-32)*9 + 2 + i // to match the lines in "standard.txt" with which the char we took in our word
					reader := bufio.NewScanner(file)
					satir := 1
					for reader.Scan() {
						if satir == start {
							lines[i] += reader.Text() // to add the line to lines
						}
						satir++
					}
					if err := reader.Err(); err != nil {
						fmt.Println("Dosya okunurken bir hata oluştu:", err)
						return
					}
				}
			}
			for _, line := range lines { // to make the lines under each other
				switch align {
				case "center":
					padding := (width - len(line)) / 2
					if padding > 0 {
						fmt.Print(strings.Repeat(" ", padding))
					}
					fmt.Println(line)
				case "left":
					// Yapılacak bir şey yok
					fmt.Println(line)
				case "right":
					padding := width - len(line)
					fmt.Print(strings.Repeat(" ", padding))
					fmt.Println(line)
				case "justify":
					kelimeler := strings.Split(wordtrimspace, " ")
					if len(kelimeler) == 1 {
						padding := (width - len(line)) / 2
						if padding > 0 {
							fmt.Print(strings.Repeat(" ", padding))
						}
						fmt.Println(line)
					}
				default:
				}
			}
			kelimeler := strings.Split(wordtrimspace, " ")
			if align == "justify" && len(kelimeler) > 1 {
				space := ""
				padding := (width - len(lines[0]) + 6*(len(kelimeler)-1)) / (len(kelimeler) - 1)
				for a := 0; a < padding; a++ {
					space += " "
				}
				kelimesatirlari := make([]string, 8)
				for index, kelime := range kelimeler {
					for _, harf := range kelime {
						for b := 0; b <= 7; b++ { // the string (lines) in our code needs 8 lines
							file.Seek(0, 0)                   // to reset the file, in order to read again
							start := (int(harf)-32)*9 + 2 + b // to match the lines in "standard.txt" with which the char we took in our word
							reader := bufio.NewScanner(file)
							satir := 1
							for reader.Scan() {
								if satir == start {
									kelimesatirlari[b] += reader.Text() // to add the line to lines
								}
								satir++
							}
							if err := reader.Err(); err != nil {
								fmt.Println("Dosya okunurken bir hata oluştu:", err)
								return
							}
						}
					}
					if index != len(kelimeler)-1 {
						for k := 0; k < 8; k++ {
							kelimesatirlari[k] += space
						}
					} else {
						continue
					}

				}
				for _, satir := range kelimesatirlari {
					fmt.Println(satir)
				}
			}

		}
	} else {
		fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]")
	}
}

func getTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	size := strings.TrimSpace(string(out))
	width, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	}

	return width, nil
}
