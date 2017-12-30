package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {

}

func (f *File)Read()  {
	fmt.Println("READ DATA")
}

func (f *File)Write()  {
	fmt.Println("WRITE DATA")
}

func Test(rw ReaderWriter)  {
	rw.Read()
	rw.Write()
}

func main() {
	var f File
	Test(&f)
}
