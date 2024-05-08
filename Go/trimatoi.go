package piscine

// TrimAtoi, bir string içindeki rakamları tamsayıya dönüştüren fonksiyondur.
// Fonksiyon, string içindeki rakamları bir tamsayıya dönüştürmeye çalışır ve
// başındaki boşlukları, artı ve eksi işaretleri dikkate alır.
// Eğer string içinde rakam yoksa veya geçerli bir tamsayıya dönüştürülemezse, 0 döndürülür.
func TrimAtoi(s string) int {
	var result int        // Dönüştürülen tamsayı değeri
	sign := 1             // Artı veya eksi işaretini belirlemek için
	digitStarted := false // Rakamlar başladı mı?

	// String'i tarayarak tamsayıyı oluştur
	for _, char := range s {
		if char >= '0' && char <= '9' {
			// Rakamı tamsayıya dönüştür
			digitStarted = true
			digit := int(char - '0')
			result = result*10 + digit
		} else if char == '-' && !digitStarted {
			// Eksi işareti bulunursa ve henüz bir rakam alınmamışsa işaret değişkenini ayarla
			sign = -1
			digitStarted = true
		} else if char == '+' && !digitStarted {
			// Artı işareti bulunursa ve henüz bir rakam alınmamışsa işaret değişkenini ayarla
			sign = 1
			digitStarted = true
		}
	}

	// Sonucu işaretle ve döndür
	return result * sign
}
