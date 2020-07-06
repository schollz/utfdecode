package utfdecode

import "fmt"

func ExampleDecode() {
	jsonData := `{"text":"utf-16:\u1F47D\u1F680 utf-32:\ud83d\udc7d\ud83d\ude80"}`
	fmt.Println(Decode(jsonData))
	// Output:
	// {"text":"utf-16:👽🚀 utf-32:👽🚀"}
}
