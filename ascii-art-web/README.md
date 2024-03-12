# Ascii-Art-Web

* Belirtilen metni Ascii-art projemiz ile birlikte entegre ederek çıktı alma işlemini gerçekleştirdik.
> Projede kullanılan kütüphaneler.
```go
import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)
```
/*
os: 	 Dosya okuma ve yazma işlemleri için kullanılır.
log: 	 Hata ve bilgi mesajlarını yazdırmak için kullanılır.
strings: Metin işlemleri için kullanılır.
bufio:   Bufio paketi arabelleğe alınmış G/Ç'yi uygular. Bir io.Reader veya io.Writer nesnesini sarar ve aynı zamanda arayüzü uygulayan ancak ara belleğe alma ve metinsel G/Ç için bazı yardımlar sağlayan başka bir nesne (Reader veya Writer) oluşturur.
fmt:     Yazdırma işlemini yapar.
net/http: Http işlevi ve sunucu uygulamasını sağlar.
*/

>  Kodun nasıl çalıştığı hakkında bir açıklama.

### Projemizin Gerçekleştirdiği İşlemler:
Ascii art projemizi main fonksiyonumuzda kullanıp, GET ve POST işlemlerini gerçekleştiriyoruz.
Server bağlantısı ve hata kodunu giriyoruz.
Assetsler arasında kurulan bağlantıyı da main fonksiyonu içinde belirtiyoruz.
Arayüzümüzde yazılan karakterleri ascii-art projemize soktuk. "text" kısmını alıp, "banner" ile belirtilen şekilde yazdırma biçimlerini kullandık. İşlemden geçen çıktıyı ise "target" yoluyla işaretli konuma yazdırdık.


> index.html:

* Projenin front-end in kod kısmı

> main.css:

* Projenin frond-end için gerekli olan renklendirme ve boyut belirleme işlevi için kullanıldı.

> Kodun Özellikleri
* Kolay kullanılabilir ve anlaşılır.
* Genişletilebilir ve özelleştirilebilir.
* Farklı metin işleme görevleri için kullanılabilir.

> Bu projede ne öğrendim
* Bir HTML uzantılı WEB sayfasını oluşturmayı ve kullanılan projeyi entegre etmeyi öğrendik.