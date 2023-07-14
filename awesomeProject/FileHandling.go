package main

//func main() {
//	//Creating a file
//	myFile, err := os.Create("new.txt")
//	defer myFile.Close()
//	if err != nil {
//		log.Fatal("Error! ", err)
//	}
//	log.Println("Empty file created successfully.", myFile)
//
//	//Getting file information
//	fileInfo, err := os.Stat("new.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("File name: ", fileInfo.Name())
//	fmt.Println("Size: ", fileInfo.Size(), " bytes")
//	fmt.Println("Permissions: ", fileInfo.Mode())
//	fmt.Println("Last modified: ", fileInfo.ModTime())
//	fmt.Println("Is Directory: ", fileInfo.IsDir())
//
//	// Reading and writing in file
//	oFile, err := os.Open("new.txt")
//	defer oFile.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	buff := make([]byte, 100)
//	for no, err := oFile.Read(buff); err == nil; no, err = oFile.Read(buff) {
//		if no > 0 {
//			os.Stdout.Write(buff[0:no])
//		}
//	}
//}
