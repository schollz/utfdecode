package utfdecode

import "fmt"

func ExampleDecode() {
	jsonData := `{"text":"utf-16:\u2764\u1F47D\u1F680 utf-32:\u2764\ufe0f\ud83d\udc7d\ud83d\ude80"}`
	fmt.Println(Decode(jsonData))
	// Output:
	// {"text":"utf-16:❤️👽🚀 utf-32:❤️👽🚀"}
}
