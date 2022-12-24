### Testler


### Unit Test  

<pre>
## Birim Test - Config Dosyası Okuma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestReadConfig` servis metodunun doğru bir şekilde çalışıp çalışmadığını test etmek |
| **Test Açıklama** | `ReadConfig` servis metodunun doğru bir şekilde çalışıp çalışmadığını test etmek |
| **Test Amacı** | Servis metodunu çağırıp, yapılandırma değerlerini almak ve beklenen değerlerle karşılaştırmak |
| **Test Adımlar** | Çalışma dizininin yolu, "sample-test" yapılandırma dosyasının ismi |
| **Test Girdileri** | Config dosya yolu |
| **Test Çıktıları** | Servis metodunun beklenen yapılandırma değerlerini döndürdüğü |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>  

```go
func TestReadConfig(t *testing.T) {
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}

	pwd := wd + "/../"

	// call service method
	cfg.ReadConfig(pwd, "sample-test")

	// get config values
	values := cfg.C

	if values.App.Env != "test" {
		t.Errorf("ReadConfig() cfg.C.App.Env = %v, want %v", values.App.Env, "test")
		return
	}

	if values.App.Port != "9090test" {
		t.Errorf("ReadConfig() cfg.C.App.Port = %v, want %v", values.App.Port, "9090test")
		return
	}

	if values.App.Version != "1.0.0test" {
		t.Errorf("ReadConfig() cfg.C.App.Version = %v, want %v", values.App.Version, "1.0.0test")
		return
	}

	if values.App.UploadsDir != "uploadstest" {
		t.Errorf("ReadConfig() cfg.C.App.UploadsDir = %v, want %v", values.App.UploadsDir, "uploadstest")
		return
	}
}
```


<pre>
#### Unit Test - Rasgele Yazı Üretme Testi

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestGenerateRandomDigitString` |
| **Test Açıklama** | `GenerateRandomDigitString` fonksiyonunun rastgele bir basamaklı string üretebildiğini ve bu stringin tüm karakterlerinin basamak olup olmadığını test etmek |
| **Test Amacı** | String uzunluğunu ve tüm karakterlerin basamak olup olmadığını kontrol etmek |
| **Test Adımlar** | Belirtilen string uzunluğunu üretmek ve stringin tüm karakterlerinin basamak olup olmadığını kontrol etmek |
| **Test Girdileri** | String uzunluğu |
| **Test Çıktıları** | Üretilen stringin doğru uzunlukta ve tüm karakterlerinin basamak olduğu |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>  

```go
func TestGenerateRandomDigitString(t *testing.T) {
	// test generate random string of digits with length
	length := 10
	randomString := utils.GenerateRandomDigitString(length)
	if len(randomString) != length {
		t.Errorf("random string length is not equal to length")
	}

	// test all characters are digits
	for _, char := range randomString {
		if char < '0' || char > '9' {
			t.Errorf("random string contains non-digit character")
		}
	}
}
```


<pre>
## Unit Test - Önbellek Veri Yazma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestCacheSet` |
| **Test Açıklama** | `CacheSet` fonksiyonunun belirtilen anahtar, değer ve süre ile önbellekte bir veri ekleyebildiğini test etmek |
| **Test Amacı** | Önbellekte bir veri ekleyebilme ve hata döndürmeme |
| **Test Adımlar** | Önbellek oluşturulur, `CacheSet` fonksiyonu çağrılır ve ekleme işleminin hata döndürmemesi kontrol edilir |
| **Test Girdileri** | Önbellek, anahtar, değer, süre |
| **Test Çıktıları** | Önbellekte veri eklendiği ve hata döndürülmediği |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>  


```go
func TestCacheSet(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "test"
	value := "test"
	expiration := time.Second

	// call service method
	err := handlers.CacheSet(inAppCache, key, value, expiration)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}
}
```




<pre>
## Unit Test - Önbellek Veri Okuma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestCacheGet` |
| **Test Açıklama** | `CacheGet` fonksiyonunun önbellekteki bir veriyi alabildiğini test etmek |
| **Test Amacı** | Önbellekteki bir veriyi alabilme ve hata döndürmeme |
| **Test Adımlar** | Önbellek oluşturulur, `CacheSet` fonksiyonu çağrılır, `CacheGet` fonksiyonu çağrılır ve alınan değerin beklenen değerle aynı olduğu kontrol edilir |
| **Test Girdileri** | Önbellek, anahtar |
| **Test Çıktıları** | Önbellekteki veri alındığı ve hata döndürülmediği |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>
    
```go

func TestCacheGet(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "test"
	value := "test"
	expiration := time.Minute

	// call service method
	err := handlers.CacheSet(inAppCache, key, value, expiration)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err != nil {
		t.Errorf("CacheGet() error = %v", err)
		return
	}

	if valueCached != value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}
```


<pre>
## Unit Test - Önbellek Eskimemiş Veri Okuma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestCacheGetWithTTLSuccess` |
| **Test Açıklama** | `CacheGet` fonksiyonunun önbellekteki bir verinin süresi dolmadan öncede alabildiğini test etmek |
| **Test Amacı** | Önbellekteki bir verinin süresi dolmadan öncede alabilme ve hata döndürmeme |
| **Test Adımlar** | Önbellek oluşturulur, `CacheSet` fonksiyonu çağrılır, `CacheGet` fonksiyonu çağrılır ve alınan değerin beklenen değerle aynı olduğu kontrol edilir |
| **Test Girdileri** | Önbellek, anahtar, değer, süre |
| **Test Çıktıları** | Önbellekteki veri alındığı ve hata döndürülmediği |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>
    
```go
func TestCacheGetWithTTLSuccess(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "testttl"
	value := "testttlsuccess"
	exp := time.Second
	err := handlers.CacheSet(inAppCache, key, value, exp)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err != nil {
		t.Errorf("CacheGet() error = %v", err)
		return
	}

	if valueCached != value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}
```




<pre>
## Unit Test - Önbellek Eskimiş Veri Okuma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestCacheGetWithTTLExpired` |
| **Test Açıklama** | `CacheGet` fonksiyonunun önbellekteki bir verinin süresi dolmuş olması durumunda hata döndürdüğünü test etmek |
| **Test Amacı** | Önbellekteki bir verinin süresi dolmuş olması durumunda hata döndürme |
| **Test Adımlar** | Önbellek oluşturulur, `CacheSet` fonksiyonu çağrılır, sürenin dolması için beklenir, `CacheGet` fonksiyonu çağrılır ve hata döndürmesi ve alınan değerin beklenen değerle aynı olmaması kontrol edilir |
| **Test Girdileri** | Önbellek, anahtar, değer, süre |
| **Test Çıktıları** | Önbellekteki veri alınamaması ve hata döndürülmes
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>
    
```go
func TestCacheGetWithTTLExpired(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "testttlfail"
	value := "testttlfailed"
	exp := time.Second
	err := handlers.CacheSet(inAppCache, key, value, exp)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// wait for expiration
	time.Sleep(2 * time.Second)

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err == nil {
		t.Errorf("CacheGet() key should be expired, error = %v", err)
		return
	}

	if valueCached == value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}
```


## Integration Testler

<pre>
## Integration Test - OS Dosya Yazma

// sequence diagram
title WriteFile
OurApp->OS: WriteFile function
OS->OS: Write file
OS->OurApp: Return error or null

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestWriteFile` |
| **Test Açıklama** | `WriteFile` fonksiyonunun dosya oluşturup yazma işlemlerinin doğru bir şekilde yapıldığını test etmek |
| **Test Amacı** | Dosya oluşturup yazma işlemlerinin doğru bir şekilde yapılıp yapılmadığını kontrol etmek |
| **Test Adımlar** | Çalışma dizininin yolu, dosya oluşturulur, dosyaya yazma işlemi yapılır, dosya kapatılır, dosyanın varlığı kontrol edilir, dosya silinir |
| **Test Girdileri** | Dosya yolu, dosya adı, yazılacak değer |
| **Test Çıktıları** | Dosya oluşturulması, yazma işleminin doğru bir şekilde yapılması, dosyanın varlığının kontrol edilmesi, dosyanın silinmesi |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>

```go
func TestWriteFile(t *testing.T) {
	// create on-the fly txt file and save to pwd ../uploads/test.txt
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	pwd := wd + "/../uploads/"

	// open file for writing
	file, err := os.Create(pwd + "test.txt")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// write to file
	_, err = file.WriteString("test")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// close file
	err = file.Close()
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// check file exists
	_, err = os.Stat(pwd + "test.txt")
	if os.IsNotExist(err) {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// remove file after test
	err = os.Remove(pwd + "test.txt")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}
}
```



<pre>
## Integration Test - OS Dosya Okuma

// sequence diagram
title ReadFile
OurApp->OS: ReadFile function
OS->OS: Read file
OS->OurApp: Return file content or error

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestReadFile` |
| **Test Açıklama** | `ReadFile` fonksiyonunun dosya okuma işleminin doğru bir şekilde yapıldığını test etmek |
| **Test Amacı** | Dosya okuma işleminin doğru bir şekilde yapılıp yapılmadığını kontrol etmek |
| **Test Adımlar** | Çalışma dizininin yolu, dosya varlığının kontrol edilmesi, dosyanın okunması |
| **Test Girdileri** | Dosya yolu, dosya adı |
| **Test Çıktıları** | Dosya varlığının kontrol edilmesi, dosyanın okunması |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>

```go
func TestReadFile(t *testing.T) {
	// create on-the fly txt file and save to pwd ../uploads/test.txt
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
		return
	}

	pwd := wd + "/../"
	filename := "LICENCE"

	// check file exists
	_, err = os.Stat(pwd + filename)
	if os.IsNotExist(err) {
		t.Errorf("ReadFile() error = %v", err)
		return
	}

	// read file
	_, err = os.ReadFile(pwd + filename)
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
		return
	}
}
```


<pre>
## Integration Test - Uzak Önbellek Veri Yazma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestNewRedisCacheSet` |
| **Test Açıklama** | Redis ile bağlantı kurarak, veri set etme işleminin doğru bir şekilde yapıldığını test etmek |
| **Test Amacı** | Redis ile bağlantının kurulup kurulmadığını kontrol etmek ve veri set etme işleminin doğru bir şekilde yapılıp yapılmadığını kontrol etmek |
| **Test Adımlar** | Redis ile bağlantı kurulması, veri set etme işlemi |
| **Test Girdileri** | Redis URL, key, value, expiration |
| **Test Çıktıları** | Redis ile bağlantının kurulmuş olması, veri set etme işleminin başarıyla tamamlanmış olması |
| **Test Ortamı** | |
</pre>
    
```go
func TestNewRedisCacheSet(t *testing.T) {
	// create redis connection
	redisUrl := "redis://:@localhost:6379/0"
	rdb, rContext := plugins.NewRedisCacheConnection(redisUrl)

	// set value
	key := "test-for-redis"
	value := "test-value"
	exp := time.Minute
	err := rdb.Set(rContext, key, value, exp).Err()
	if err != nil {
		t.Errorf("TestNewRedisCacheSet() could not set value to redis. error = %v", err)
		return
	}
}
```


<pre>
## Integration Test - Uzak Önbellek Veri Okuma

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestNewRedisCacheGet` |
| **Test Açıklama** | `NewRedisCacheConnection` servis metodunun doğru bir şekilde çalışıp çalışmadığını test etmek |
| **Test Amacı** | `NewRedisCacheConnection` servis metodunun doğru bir şekilde çalışıp çalışmadığını test etmek |
| **Test Adımlar** | Redis bağlantısı oluşturulur, anahtar değeri set edilir, anahtar değeri get edilir |
| **Test Girdileri** | Redis bağlantı URL'si |
| **Test Çıktıları** | Redis bağlantısı oluşturulur, anahtar değeri set edilir, anahtar değeri get edilir |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>

```go

func TestNewRedisCacheGet(t *testing.T) {
	// create redis connection
	redisUrl := "redis://:@localhost:6379/0"
	rdb, rContext := plugins.NewRedisCacheConnection(redisUrl)

	// set value
	key := "test-for-redis"
	value := "test-value"
	exp := time.Minute
	err := rdb.Set(rContext, key, value, exp).Err()
	if err != nil {
		t.Errorf("TestNewRedisCacheGet() could not set value to redis. error = %v", err)
		return
	}

	// get value
	val, err := rdb.Get(rContext, key).Result()
	if err != nil {
		t.Errorf("TestNewRedisCacheGet() could not get value from redis. error = %v", err)
		return
	}
	if val != value {
		t.Errorf("TestNewRedisCacheGet() could not get value from redis. error = %v", err)
		return
	}

}
```

## Sistem Testleri

<pre>
## Sistem Testleri - Uygulama Canlılığı

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestApplicationLiveness` |
| **Test Açıklama** | Uygulamanın "ping" adresine GET isteği göndererek, cevap olarak 200 status kodunun döndürülüp döndürülmediğini kontrol etmek |
| **Test Amacı** | Uygulamanın canlı olduğunu doğrulamak |
| **Test Adımları** | "ping" adresine GET isteği göndermek. Gelen cevap kodunu kontrol etmek |
| **Test Girdileri** | "ping" adresi |
| **Test Çıktıları** | 200 status kodu döndürülür |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>

```go

func TestApplicationLiveness(t *testing.T) {
	// Create a new request using http to get the file
	req, err := http.NewRequest("GET", ENDPOINT_PING, nil)
	if err != nil {
		t.Errorf("TestApplicationLiveness: Cannot create new request: %s", err.Error())
		return
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("TestApplicationLiveness: Cannot send request: %s", err.Error())
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestApplicationLiveness resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}
}
```


<pre>
## Sistem Testleri - Uygulama Dosya Yükleme

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestUploadFileToServer` |
| **Test Açıklama** | Sunucuya örnek bir dosya yüklemek |
| **Test Amacı** | Sunucuya dosya yükleme işleminin doğru bir şekilde yapıldığını doğrulamak |
| **Test Adımları** | Örnek dosyayı açmak ve içeriğini okumak İstek gövdesi için bir buffer oluşturmak. Yeni bir multipart yazar oluşturmak. "file" adı ile yeni bir form-data başlığı oluşturmak. İstekleri göndermek için http.Post fonksiyonunu kullanmak. Dosya yükleme sonucunu kontrol etmek |
| **Test Girdileri** | Örnek bir dosya |
| **Test Çıktıları** | Dosya yükleme işlemi başarılı olur |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>

```go
func TestUploadFileToServer(t *testing.T) {
	// Open the file to read its contents
	onServerInit()
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../LICENCE")
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// stringify body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestUploadFileToServer() could not read response body: ", err)
		return
	}

	// response.Body should be a json string from struct handlers.RespondJson
	var respondJson handlers.RespondJson
	err = json.Unmarshal(bodyBytes, &respondJson)
	if err != nil {
		t.Error("TestUploadFileToServer() could not unmarshal response body: ", err)
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestUploadFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}

	if respondJson.Status == false {
		t.Errorf("TestUploadFileToServer() respondJson.Status = %v, want %v", respondJson.Status, true)
		return
	}

	// RespondJson.Message should be a code as string not interface
	Code = respondJson.Message.(string)
	t.Logf("TestUploadFileToServer() Code = %v", Code)
}
```


<pre>
## Sistem Testleri - Uygulama Dosya Getirme

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestGetUploadedFile` |
| **Test Açıklama** | Sunucudan yüklenen dosyanın indirilebilir olup olmadığını test etmek |
| **Test Amacı** | Sunucudan yüklenen dosyanın indirilebilir olduğunu doğrulamak |
| **Test Adımlar** | 1. Sunucudan yüklenen dosya için GET isteği göndermek <br> 2. Gelen cevap kodunu kontrol etmek |
| **Test Girdileri** | Yüklenen dosya kodu |
| **Test Çıktıları** | 200 status kodu döndürülür |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>

```go
func TestGetUploadedFile(t *testing.T) {
	t.Logf("TestGetUploadedFile() Code = %v", Code)
	// Create a new request using http to get the file
	req, err := http.NewRequest("GET", ENDPOINT_GET+"/"+Code, nil)
	if err != nil {
		t.Errorf("TestGetUploadedFile: Cannot create new request: %s", err.Error())
		return
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("TestGetUploadedFile: Cannot send request: %s", err.Error())
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestGetUploadedFile() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}
}
```


## Regresyon Testleri


<pre>
## Regresyon Testleri - Yüksek Boyutlu Dosya Yükleme

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestUploadLargeFileToServer` |
| **Test Açıklama** | Sunucuya büyük dosya yükleme işleminin gerçekleşip gerçekleşmediğini test etmek |
| **Test Amacı** | Sunucuya büyük dosya yükleme işleminin sınırlamalarını test etmek |
| **Test Adımlar** | Sunucuya büyük dosya yükleme işlemi gerçekleştirilir |
| **Test Girdileri** | Büyük dosya |
| **Test Çıktıları** | 422 status kodu |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |

</pre>

```go
func TestUploadLargeFileToServer(t *testing.T) {
	onServerInit()
	// Open the file to read its contents
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../tests/big-size.mp4")
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("TestUploadLargeFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusUnprocessableEntity)
		return
	}
}
```


<pre>
## Regresyon Testleri - .exe Dosya Yükleme

| - | Açıklama |
| --- | --- |
| **Test Adı** | `TestUploadExeFileToServer` |
| **Test Açıklama** | `.exe` dosyalarının server'a yüklenmesine izin verilmediği için, server tarafından beklenen bir hata mesajı döndürülür |
| **Test Amacı** | Server'ın .exe dosyalarını yükleme işlemini doğru bir şekilde gerçekleştirip gerçekleştirmediğini test etmek |
| **Test Adımlar** | 1. Server'ı başlatmak <br> 2. .exe dosyasını açmak <br> 3. Dosyayı server'a yükleme isteği göndermek <br> 4. Server'ın döndürdüğü cevap kodunu kontrol etmek |
| **Test Girdileri** | .exe dosyası |
| **Test Çıktıları** | Server tarafından beklenen bir hata mesajı döndürülür |
| **Test Ortamı** | Mac M1 Pro - 16GB RAM - 512 SSD |
</pre>

```go
func TestUploadExeFileToServer(t *testing.T) {
	onServerInit()
	// Open the file to read its contents
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../tests/not-allowed.exe")
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusTeapot {
		t.Errorf("TestUploadExeFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusTeapot)
		return
	}
}
```



