# utfdecode

[![travis](https://travis-ci.org/schollz/utfdecode.svg?branch=master)](https://travis-ci.org/schollz/utfdecode) 
[![go report card](https://goreportcard.com/badge/github.com/schollz/utfdecode)](https://goreportcard.com/report/github.com/schollz/utfdecode) 
[![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://gocover.io/github.com/schollz/utfdecode)
[![godocs](https://godoc.org/github.com/schollz/utfdecode?status.svg)](https://godoc.org/github.com/schollz/utfdecode) 

I wanted to convert text containing escaped UTF-16 and/or UTF-32 codes to just UTF-16. This library helps do that. It rejexes for unicde strings and then converts them to UTF-16 [using the standard formula](https://en.wikipedia.org/wiki/UTF-16). The code is small enough that you might just consider copy and pasting :)


For example, the following Go code:

```golang
jsonData := `{"text":"utf-16:\u2764\u1F47D\u1F680 utf-32:\u2764\ud83d\udc7d\ud83d\ude80"}`
fmt.Println(utfdecode.Decode(jsonData))
```

will print

```
{"text":"utf-16:❤👽🚀 utf-32:❤👽🚀"}
```


## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements


## License

MIT