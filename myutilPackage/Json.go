package myutil

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
}

func Json() {
	var p []byte
	b, _ := json.Marshal(User{"uw22kw", "John Doe", 32})
	fmt.Println(string(b))

	// f, _ := os.Create("storage.json")
	// defer f.Close()
	// encoder := json.NewEncoder(f)
	// err := encoder.Encode(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	f, _ := os.Open("storage.json")
	defer f.Close()
	decoder := json.NewDecoder(f)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	var output User
	json.Unmarshal(p, &output)
	fmt.Println(output)
}
