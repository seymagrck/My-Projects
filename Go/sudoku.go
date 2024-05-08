package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

// StringToInt, Sudoku tahtasını temsil eden string'i int slice'a çevirir.
// '.' karakteri 0 olarak, '1' ile '9' arasındaki sayılar ise ilgili sayıya çevrilir.
func StringToInt(input string) []int {
	convertedValues := []int{}
	for _, char := range input {
		if char == '.' {
			convertedValues = append(convertedValues, 0)
		} else if char >= '1' && char <= '9' {
			convertedValues = append(convertedValues, int(char-'0'))
		} else {
			HataMesajiYaz()
		}
	}
	return convertedValues
}

// HataMesajiYaz, hata mesajını ekrana basar.
func HataMesajiYaz() {
	hataMesaji := "Hata"
	for _, char := range hataMesaji {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// SudokuCoz, Sudoku bulmacayı çözer.
func SudokuCoz(tahta [9][9]int) bool {
	// Tüm indeksleri kontrol et
	for satir := 0; satir < len(tahta); satir++ {
		for sutun := 0; sutun < len(tahta); sutun++ {
			if tahta[satir][sutun] == 0 {
				for num := 1; num <= len(tahta); num++ {
					// Satır, Sütun ve Kutu Kontrolü
					if GuvenliMi(num, satir, sutun, tahta) {
						tahta[satir][sutun] = num
						// Yanlış yerleştirmeler için rekürsif çağrı
						if SudokuCoz(tahta) {
							return true
						}
						tahta[satir][sutun] = 0
					}
				}
				return false
			}
		}
	}
	TahtayiYaz(tahta)
	return true
}

// main fonksiyonu, Sudoku çözme programının ana giriş noktasıdır.
func main() {
	// Liste oluştur ve Sudoku'yu çöz
	sudokuTahtasi := [9][9]int{}
	if len(os.Args) == 10 {
		for i := 0; i < len(sudokuTahtasi); i++ {
			for j := 0; j < len(sudokuTahtasi); j++ {
				sudokuTahtasi[i][j] = StringToInt(os.Args[i+1])[j]
			}
		}
		if SudokuKontrol(sudokuTahtasi) == true {
			SudokuCoz(sudokuTahtasi)
		}
	} else {
		HataMesajiYaz()
	}
}

// TahtayiYaz, Sudoku tahtasını ekrana basar.
func TahtayiYaz(tahta [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune('0' + tahta[i][j]))
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

// GuvenliMi, belirli bir sayının bir hücreye yerleştirilebilir olup olmadığını kontrol eder.
func GuvenliMi(num, satir, sutun int, tahta [9][9]int) bool {
	return !SatirdaKullaniliyorMu(num, satir, tahta) &&
		!SutundaKullaniliyorMu(num, sutun, tahta) &&
		!KutudaKullaniliyorMu(num, satir-satir%3, sutun-sutun%3, tahta)
}

// SatirdaKullaniliyorMu, belirli bir sayının bir satırda kullanılıp kullanılmadığını kontrol eder.
func SatirdaKullaniliyorMu(num, satir int, tahta [9][9]int) bool {
	for i := 0; i < len(tahta); i++ {
		if tahta[satir][i] == num {
... (84 satır kaldı)