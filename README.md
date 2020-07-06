# utf16 decode

I wanted to convert UTF-32 codes to UTF-16, often encoded in some JSON. This library helps do that. It rejexes for unicde strings and then converts them to UTF-16 [using the standard formula](https://en.wikipedia.org/wiki/UTF-16).

For example:

```golang
jsonData := `{"text":"Cool! \u2764\ufe0f\ud83d\udc7d\ud83d\ude80"}`
fmt.Println(Decode(jsonData))
```

will print

```
{"text":"Cool! ❤️👽🚀"}
```

## License

MIT