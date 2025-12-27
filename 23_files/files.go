package main

import (
	"fmt"
	"os"
)

// reading files

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	//log the error

	// 	panic(err)
	// }
	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// TASK: read file

	// buf := make([]byte, 12)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// another way to read file
	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f)) //simple easy peasy
	// gotcha: os.ReadFile will load the whole file data in the memory, for big files it will cause issue

	// TASK: read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)

	// for _, fi :=range fileInfo{
	// 	fmt.Println(fi.Name(),",Is directory:",fi.IsDir())
	// }

	// TASK: creating a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hi go\n")
	// f.WriteString("nice lang")

	// bytes := []byte("Hello Golang")

	// f.Write(bytes)

	//TASK: read from one file and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("written to new file successfully")

	// TASK: deleting a file
	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file deleted successfully")
}
