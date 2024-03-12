package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		formatli := formatlamm(line)      // eğer (low, sayi) (up, sayi) (cap, sayi) varsa onları işlemlerde boşluklara takılmadan işleyebilmek için formatlamm fonksiyonuna gönderdim
		words := strings.Fields(formatli) // formatlamm fonksiyonundan dönen stringi boşluklara göre ayırıp bir diziye atadım
		for i, word := range words {      // dizi içindeki her kelimeyi (hex), (bin), (up), (low), (cap) kontrol edip ona göre işlem yaptım
			if i < len(words)-1 {
				switch words[i+1] {
				case "(hex)":
					dec, _ := strconv.ParseInt(word, 16, 64)
					words[i] = strconv.FormatInt(dec, 10)
					words[i+1] = "" // eğer bir sonraki kelimem (hex) ise onu işledikten sonra onu boş bir stringe eşitledim (low,6) kasjdb (up,4)
				case "(bin)":
					dec, _ := strconv.ParseInt(word, 2, 64)
					words[i] = strconv.FormatInt(dec, 10)
					words[i+1] = ""
				case "(up)":
					words[i] = strings.ToUpper(word)
					words[i+1] = ""
				case "(low)":
					words[i] = strings.ToLower(word)
					words[i+1] = ""
				case "(cap)":
					words[i] = strings.Title(strings.ToLower(word))
					words[i+1] = ""
				}
			}
		}

		// (hex), (bin), (up), (low), (cap) işlemlerini yaptıktan sonra dizi içindeki boş değerleri silen döngüm :

		if len(words)-1 >= 0 && words[len(words)-1] == "" {
			words = words[:len(words)-1]
		}

		sayili := icerirMi(words) // dizi içinde (low, sayi) (up, sayi) (cap, sayi) var mı diye kontrol ettim eğer yoksa herhangi bir tarama yapmayarak tasarruf etmiş oluyorum
		if sayili != "yok" {

			sayim, _ := strconv.Atoi(sayili[3:]) // içerdiğine göre sayı kısmını alıyoruz (up,3)
			indeksim := icerenIndex(words)

			if sayim > indeksim { // eğer sayı kısmı indeksimden büyükse hata verip hatalı kısmı hiç yazdırmıyorum
				fmt.Println("kendinden önce bulunmayan kelime kadar büyük sayı girilmiştir")
			} else if strings.Contains(sayili, "up") {
				for j := indeksim - 1; j >= indeksim-sayim; j-- {
					words[j] = strings.ToUpper(words[j])
				}
			} else if strings.Contains(sayili, "low") {
				for j := indeksim - 1; j >= indeksim-sayim; j-- {
					words[j] = strings.ToLower(words[j])
				}
			} else {
				for j := indeksim - 1; j >= indeksim-sayim; j-- {
					words[j] = strings.Title(strings.ToLower(words[j]))
				}
			}

			words = sayiliCikar(words) // gerekli işlemler yapıldıktan sonra dizi içindeki (low, sayi) (up, sayi) (cap, sayi) kısımlarını siliyorum
		}

		eklenecek := ""
		for k := 0; k < len(words); k++ {
			if k < len(words)-1 {
				if words[k] != "" && words[k] != " " {
					eklenecek += words[k] + " "
				}
			} else {
				if words[k] != "" && words[k] != " " {
					eklenecek += words[k]
				}
			}
		} // bu döngüde dizi içerisindeki cümlelerden sonra gereksiz yere boşluk bırakmayı eklenecek değişkenime gerekli düzenlemeleri yapıyorum

		eklenecek = formatBir(eklenecek) // noktalama işaretlerini kendinden bir önceki kelimeye ekleyen fonksiyon

		if TirnakliVar(eklenecek) { // eğer cümle içerisinde iki tırnak arasında bir kelime varsa bunun olup olmadığını kontrol ediyorum varsa
			eklenecek = tirnakFormat(eklenecek) // fonksiyonum ile kelimeler ile tırnaklarım arasındaki boşlukları siliyorum
		}

		eklenecek = makeAn(eklenecek) // a,A harflerinin yanında sesli harf varsa n ekleyen fonksiyon

		_, _ = writer.WriteString(eklenecek + "\n") // işlemler yapıldıktan sonra cümleyi dosyaya yazdırıyorum

		_ = writer.Flush() // dosyayı kapatmadan önce bufferı temizliyorum
	}
}

func formatlamm(cumle string) string {
	dizi := strings.Fields(cumle)
	yeniDizi := []string{}

	for i := 0; i < len(dizi); i++ {

		girilMedi := true
		for j := 0; j < 999; j++ {
			if i+1 < len(dizi) && (dizi[i]+dizi[i+1] == "(low,"+strconv.Itoa(j)+")" ||
				dizi[i]+dizi[i+1] == "(up,"+strconv.Itoa(j)+")" ||
				dizi[i]+dizi[i+1] == "(cap,"+strconv.Itoa(j)+")") {
				yeniDizi = append(yeniDizi, dizi[i]+dizi[i+1])
				i++
				girilMedi = false
				break
			}
		}

		if girilMedi {
			yeniDizi = append(yeniDizi, dizi[i])
		}

	}

	return strings.Join(yeniDizi, " ")
}

func makeAn(cumle string) string {
	dizi := strings.Fields(cumle)

	for i := 0; i < len(dizi)-1; i++ {
		if dizi[i] == "a" || dizi[i] == "A" {
			if len(dizi[i+1]) > 0 && sesliMi(rune(dizi[i+1][0])) {
				dizi[i] = dizi[i] + "n"
			}
		}
	}

	return strings.Join(dizi, " ")
}

func sesliMi(harf rune) bool {
	if harf == 'a' || harf == 'e' || harf == 'i' || harf == 'o' || harf == 'u' || harf == 'A' || harf == 'E' || harf == 'I' || harf == 'O' || harf == 'U' {
		return true
	}

	return false
}

func formatBir(cumle string) string {
	dizi := strings.Fields(cumle)

	for i := 1; i < len(dizi); i++ {
		for len(dizi[i]) > 0 && (dizi[i][0] == '.' || dizi[i][0] == ',' || dizi[i][0] == '!' || dizi[i][0] == '?' || dizi[i][0] == ':' || dizi[i][0] == ';') {
			dizi[i-1] = string(dizi[i-1]) + string(dizi[i][0])
			dizi[i] = dizi[i][1:]
		}
		if len(dizi[i]) == 0 {
			dizi = append(dizi[:i], dizi[i+1:]...)
			i--
		}
	}

	return strings.Join(dizi, " ")
}

func sayiliCikar(dizi []string) []string {
	yeniDizi := []string{}

	for i := 0; i < len(dizi); i++ {
		eklenir := true
		for j := 1; j <= 999; j++ {
			if dizi[i] == "(low,"+strconv.Itoa(j)+")" {
				eklenir = false
				break
			}
			if dizi[i] == "(up,"+strconv.Itoa(j)+")" {
				eklenir = false
				break
			}
			if dizi[i] == "(cap,"+strconv.Itoa(j)+")" {
				eklenir = false
				break
			}
		}

		if eklenir {
			yeniDizi = append(yeniDizi, dizi[i])
		}
	}

	return yeniDizi
}

func icerirMi(dizi []string) string {
	for i := 0; i < len(dizi); i++ {
		for j := 1; j <= 999; j++ {
			if dizi[i] == "(low,"+strconv.Itoa(j)+")" {
				return "low" + strconv.Itoa(j)
			}

			if dizi[i] == "(up,"+strconv.Itoa(j)+")" {
				return "upa" + strconv.Itoa(j)
			}

			if dizi[i] == "(cap,"+strconv.Itoa(j)+")" {
				return "cap" + strconv.Itoa(j)
			}
		}
	}

	return "yok"
}

func icerenIndex(dizi []string) int {
	for i := 0; i < len(dizi); i++ {
		for j := 1; j <= 999; j++ {
			if dizi[i] == "(low,"+strconv.Itoa(j)+")" {
				return i
			}

			if dizi[i] == "(up,"+strconv.Itoa(j)+")" {
				return i
			}

			if dizi[i] == "(cap,"+strconv.Itoa(j)+")" {
				return i
			}
		}
	}

	return -1
}

func tirnakFormat(cumle string) string {
	segments := strings.Split(cumle, "'")
	for i := 1; i < len(segments); i += 2 {
		segments[i] = strings.TrimSpace(segments[i])
	}
	return strings.Join(segments, "'")
}

func TirnakliVar(cumle string) bool {
	for i := 0; i < len(cumle); i++ {
		if string(cumle[i]) == "'" {
			i++
			for i < len(cumle) {
				if string(cumle[i]) == "'" {
					return true
				}
				i++
			}
		}
	}
	return false
}
