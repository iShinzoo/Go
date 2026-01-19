package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("File Dir:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file mode:", fileInfo.Mode())
	// fmt.Println("file mod time:", fileInfo.ModTime())

	// read file
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < d; i++ {
	// 	fmt.Println("data:", d, string(buf[i]))
	// }

	// simple way to read file for small files
	// data, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// read Folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.Readdir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create a file
	// f, err := os.Create("testo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hell")

	// bytes := []byte("hello")
	// f.Write(bytes)

	// read and write to another file (Streaming fashion)
	// sourceFile, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("testo.txt")
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
	// 	if err != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("Written to new file")

	// delete a file
	err := os.Remove("testo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file removed successfully")
}
