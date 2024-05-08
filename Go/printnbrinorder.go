package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	// Sayıyı stringe çevir
	nStr := sayiToStr(n)

	// Stringi diziye çevir
	nArr := []rune(nStr)

	// Sıralama algoritması (örneğin: bubble sort)
	for i := 0; i < len(nArr)-1; i++ {
		for j := 0; j < len(nArr)-1-i; j++ {
			if nArr[j] > nArr[j+1] {
				// Swap
				nArr[j], nArr[j+1] = nArr[j+1], nArr[j]
			}
		}
	}

	// Sıralı değerleri yazdır
	for _, digit := range nArr {
		z01.PrintRune(digit)
	}
}

// sayiToStr, bir sayıyı string'e çeviren basit bir fonksiyondur
func sayiToStr(sayi int) string {
	// Sayının basamaklarını bulmak için bir döngü kullan
	// Döngüden elde edilen her basamağı ASCII karakterine dönüştür ve birleştir
	var result string
	for sayi > 0 {
		basamak := sayi % 10
		sayi = sayi / 10
		// ASCII'ye dönüştür ve birleştir
		result = string(basamak+'0') + result
	}

	// Eğer sayı 0 ise, boş bir dize döndür
	if result == "" {
		return "0"
	}

	return result
}
