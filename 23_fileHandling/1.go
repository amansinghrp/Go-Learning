package main

import "os"

func main() {
	// file, err := os.Open("ex.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// // fileInfo, err := file.Stat()
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println("File name is: ", fileInfo.Name())
	// // fmt.Println("Is it a directory: ", fileInfo.IsDir())
	// // fmt.Println("File size is: ", fileInfo.Size())
	// // fmt.Println("File mode is: ", fileInfo.Mode())
	// // fmt.Println("File modified at: ", fileInfo.ModTime())

	// //reading file data
	// buff := make([]byte, 40)

	// d, err := file.Read(buff)
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// println("Data:", d, buff)
	// for i := 0; i < len(buff); i++ {
	// 	fmt.Print(string(buff[i]))
	// }

	// //another way of reading the file
	// file2data, err := os.ReadFile("ex.txt") //loads the entire content ito memory at once
	// //not vaible in large files
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(file2data))

	// //Reading folders
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(2) //+ve -> shows those many files
	// // <= 0 -> gives all the files or folders

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	//create a file
	// f, err := os.Create("ex2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("Hi Go")
	// // f.WriteString(" From Aman")

	// // bytes := []byte("Hello GoLang")

	// // f.Write(bytes)

	// //read and write to another file

	// sourceFile, err := os.Open(("ex.txt"))
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("ex3.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// //reader
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

	// 	err2 := writer.WriteByte(b)
	// 	if err2 != nil {
	// 		panic(err2)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written to new file ")

	//delete a file
	err := os.Remove("files.go")
	if err != nil {
		panic(err)
	}
}
