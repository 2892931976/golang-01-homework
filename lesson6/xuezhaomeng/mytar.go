package  main

import (
	"os"
	"log"
	"io"
	"strings"
	"compress/gzip"
	"archive/tar"
	"fmt"
	"path/filepath"
)
var  file_list []string

func walkFunc(path string, info os.FileInfo, err error) error {
	file_list = append(file_list, path)
	return nil
}


func  CreateTar(filename string,files []string) error {

	f,err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	 defer  f.Close()

	var fileWriter  io.WriteCloser = f
	if  strings.HasPrefix(filename,".gz") {
		fileWriter = gzip.NewWriter(f)
		defer fileWriter.Close()
	}
	writer := tar.NewWriter(fileWriter)
	defer  writer.Close()

	for  _,name := range files {

		if err := writeFileToTar(writer,name) ; err != nil  {
			return  err
		}


	}
	return  nil
}

func writeFileToTar(writer *tar.Writer,filename string) error {
	f ,err := os.Open(filename)
	if  err != nil  {
		log.Fatal(err)
	}
	defer f.Close()
	stat ,err := f.Stat()
	if  err != nil  {
		log.Fatal(err)
	}
	header  := &tar.Header{
		Name : filename,
		Size : stat.Size(),
	}
	if err = writer.WriteHeader(header); err != nil {
		return  err
	}
	fmt.Println(stat)
	_,err = io.Copy(writer,f)
	return  err
}
func main(){
	if len(os.Args) != 3 {
		fmt.Println("Usage: mytar [desc] [src]")
		return
	}

	filepath.Walk(os.Args[2:], walkFunc)
	CreateTar(os.Args[1], file_list)
}
