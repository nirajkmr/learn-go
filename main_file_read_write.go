package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hi!there"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v character\n", length)
	defer file.Close()
	defer readFile("./fromString.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text Read from File:", string(data))
}
