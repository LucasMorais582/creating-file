package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// creating file
	file, errorFile := os.Create("file.txt")
	if errorFile != nil {
		panic(errorFile)
	}

	size, errorFile := file.Write([]byte("Escrevendo dentro do arquivo"))
	if errorFile != nil {
		panic(errorFile)
	}

	fmt.Printf(" Arquivo criado com sucesso!\n Tamanho: %d bytes\n", size)
	file.Close()

	// reading file
	fileCreated, errorFile := os.ReadFile("file.txt")
	if errorFile != nil {
		panic(errorFile)
	}

	fmt.Println(string(fileCreated))

	// reading file per pieces
	file2, errorFile := os.Open("file.txt")
	if errorFile != nil {
		panic(errorFile)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 5)

	for {
		index, errorFile := reader.Read(buffer)
		if errorFile != nil {
			break
		}
		fmt.Println(string(buffer[:index]))
	}

	// deleting file
	errorRemoveFile := os.Remove("file.txt")
	if errorRemoveFile != nil {
		panic(errorRemoveFile)
	}
}
