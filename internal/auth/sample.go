package main

import (
	"fmt"
)

type ReaderWriterStructure struct {
	ReaderT
	WriterT
}

type ReaderT struct {
}

type WriterT struct {
	Name string
}

type ReaderWriter interface {
	Writer
	Reader
}

type Writer interface {
	write() string
}

type Reader interface {
	read(value string) string
}

func (rw *ReaderWriterStructure) write() string {

	//temp := rw
	//rw.Name = "Ali"
	//rw = temp
	temp := *rw
	rw.Name = "AAA"
	//fmt.Println(rw)
	//fmt.Println(*rw)
	//fmt.Println(&rw)
	*rw = temp
	return ""
}

func (b ReaderWriterStructure) read(value string) string {
	return ""
}

func init() {
	if a == "" {
		panic("no value for $USER")
	}
}

var a string = ""

func main() {
	rw := ReaderWriterStructure{ReaderT{}, WriterT{}}
	rw.write()
	fmt.Println(rw)
	fmt.Println(a)
}
