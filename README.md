[![Go Report Card](https://goreportcard.com/badge/github.com/jackytck/jcconv)](https://goreportcard.com/report/github.com/jackytck/jcconv)


## jcconv

### Single line
Enter original text in standard input and get back the result in standard output.

```
jcconv play [flags]
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
jcconv file [flags]
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
jcconv all [flags]
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
jcconv detect [flags]
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
jcconv web [flags]
```

#### Options
```
  -d, --domain string   External protocol and domain name, e.g. https://jcconv.nat.com (default "http://127.0.0.1:8080")
  -h, --help            help for web
  -p, --port int        Port of local server. (default 8080)
```

### API
GET or POST: parameter: `text`

```bash
$ curl -sd 'text=就像收集垃圾和剷雪一樣的事。不得不由誰來做。和喜不喜歡沒有關係。' http://127.0.0.1:8080/translate | jq
{
  "input": "就像收集垃圾和剷雪一樣的事。不得不由誰來做。和喜不喜歡沒有關係。",
  "output": "就像收集垃圾和铲雪一样的事。不得不由谁来做。和喜不喜欢没有关系。",
  "error": "",
  "elapsed": "588.239µs"
}
```
