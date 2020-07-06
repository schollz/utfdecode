package utfdecode

import "fmt"

func ExampleUTFDecode() {
	jsonData := `{"text":"Cool! \u2764\ufe0f\ud83d\udc7d\ud83d\ude80"}`
	fmt.Println(jsonData)
	fmt.Println(Decode(jsonData))
	// Output:
	// {"text":"Cool! \u2764\ufe0f\ud83d\udc7d\ud83d\ude80"}
	// {"text":"Cool! ❤️👽🚀"}
}
