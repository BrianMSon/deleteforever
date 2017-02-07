package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : deleteforever.exe TargetFileName")
		return
		//		os.Args = append(os.Args, "")
		//		os.Args[1] = "test.zip"
	}

	fileName := os.Args[1]

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Delete [" + fileName + "] Are you sure? (y/N)")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \r\n")
	if text != "y" && text != "Y" {
		return
	}

	// Delete Forever
	f, err := os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat, _ := f.Stat()
	//fmt.Println("File Size : ", stat.Size())

	content := make([]byte, 0)

	for i := int64(0); i < stat.Size(); i++ {
		if i%100 == 100-1 {
			content = append(content, byte('\n'))
		} else {
			content = append(content, byte('D'))
		}
	}

	f.Write(content)
	f.Sync()
	f.Close()

	os.Remove(fileName)

	fmt.Println(fileName, "file deleted forever!")
}
