[![Go Report Card](https://goreportcard.com/badge/github.com/jackytck/jcconv)](https://goreportcard.com/report/github.com/jackytck/jcconv)


## jcconv

### Overview
* Completely **offline**
* **Single binary** of size 13M
* **Cross platform**: Darwin64, Linux64, Win32, Win64
* Support **Docker**, image size: 11.1M
* **Fast** with multi-threading
* **API** server with demo page
* With **variants**: compare: [jcconv](https://jcconv.jackytck.com/?text=一邊食公仔麵，一邊看衰仔樂園) vs [google](https://translate.google.com.hk/#view=home&op=translate&sl=zh-CN&tl=zh-CN&text=%E4%B8%80%E9%82%8A%E9%A3%9F%E5%85%AC%E4%BB%94%E9%BA%B5%EF%BC%8C%E4%B8%80%E9%82%8A%E7%9C%8B%E8%A1%B0%E4%BB%94%E6%A8%82%E5%9C%92)
* **No limit** on word count

### Homebrew
```bash
$ brew install jackytck/jcconv/jcconv
```

### Docker
```bash
# api + web interface
$ docker run -p 8080:8080 jackytck/jcconv:v1.0.1

# quick play
$ docker run jackytck/jcconv:v1.0.1 play -i '矩陣力學是量子力學其中一種的表述形式'
```

### Single line
Enter original text in standard input and get back the result in standard output.

```
$ jcconv play [flags]
```

#### Options
```
  -c, --convert string   Conversion: one of 'auto' (default), 's2hk', 's2tw', 'hk2s', 'tw2s'. (default "auto")
  -h, --help             help for play
  -i, --input string     Input string.
```

---

### Single file
Translate a single file and output to the given path.

```
$ jcconv file [flags]
```

##### Options
```
  -c, --convert string   Conversion: one of 's2hk' (default), 's2tw', 'hk2s', 'tw2s'. (default "auto")
  -f, --file string      Input file path.
  -h, --help             help for file
  -o, --out string       Output file path.
  -n, --thread int       Number of threads to process, default is number of cores x 4 (default -1)
  -v, --verbose          Display more info
```

---

### All files
Translate all of the text file(s) under the input directory.

```
$ jcconv all [flags]
```

#### Options
```
  -c, --convert string   Conversion: one of 's2hk' (default), 's2tw', 'hk2s', 'tw2s'. (default "auto")
  -h, --help             help for all
  -d, --in string        Input directory path.
  -o, --out string       Output directory path.
  -v, --verbose          Display more info
```

---

### Detect lang
Detect if a given string is traditional or simplified. It is classified as traditional only if it has 97% confidence.

```
$ jcconv detect [flags]
```

#### Options
```
  -h, --help           help for detect
  -i, --input string   Input string.
```

---

### API server
Start a translation api server, plus a web interface for interacting with it.

```
$ jcconv web [flags]
```

#### Options
```
  -d, --domain string   External protocol and domain name, e.g. https://jcconv.nat.com (default "http://127.0.0.1:8080")
  -h, --help            help for web
  -p, --port int        Port of local server. (default 8080)
  -t, --track string    Google analytics tracking code, e.g. UA-XXXXXXXXX-X
```

### API
GET or POST: parameter: `text`

```bash
$ curl -sd 'text=就像收集垃圾和剷雪一樣的事。不得不由誰來做。和喜不喜歡沒有關係。' http://127.0.0.1:8080/translate | jq
{
  "input": "就像收集垃圾和剷雪一樣的事。不得不由誰來做。和喜不喜歡沒有關係。",
  "fromLocale": "zh-HK",
  "toLocale": "zh-CN",
  "chain": "hk2s",
  "output": "就像收集垃圾和铲雪一样的事。不得不由谁来做。和喜不喜欢没有关系。",
  "error": "",
  "elapsed": "665.733µs"
}
```

### Sharable link
Link with translation by using the query parameter `text`, e.g.

* https://jcconv.jackytck.com/?text=歐幾理得推論
* https://jcconv.jackytck.com/?text=map%20%E8%BF%99%E7%A7%8D%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%9C%A8%E5%85%B6%E4%BB%96%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80%E4%B8%AD%E4%B9%9F%E7%A7%B0%E4%B8%BA%E5%AD%97%E5%85%B8%EF%BC%88Python%EF%BC%89%E3%80%81hash%20%E5%92%8C%20HashTable%20%E7%AD%89%E3%80%82
