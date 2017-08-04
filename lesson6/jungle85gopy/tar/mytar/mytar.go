package mytar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// const for file type
const (
	DotTargz = `.tar.gz`
	DotTar   = `.tar`
)

// MyTar tar and/or gzip given dirFiles to dstFile
func MyTar(dstType, dstFile string, params []string) error {
	var gzipWr *gzip.Writer
	var tarWr *tar.Writer
	dstFd, _ := os.Create(dstFile)
	defer dstFd.Close()

	// create Writer
	if dstType == DotTargz {
		gzipWr = gzip.NewWriter(dstFd)
		tarWr = tar.NewWriter(gzipWr)
		defer gzipWr.Close()
		defer tarWr.Close()
	} else if dstType == DotTar {
		tarWr = tar.NewWriter(dstFd)
		defer tarWr.Close()
	}
	// traverse all dir and files
	var toDoSlice []*os.File
	for _, dirFile := range params {
		f, err := os.Open(dirFile)
		if err != nil {
			return err
		}
		// f will be closed by the follow funcs
		toDoSlice = append(toDoSlice, f)
		fmt.Println("-- MyTar: add to slice of:", dirFile)
	}
	return doTar(toDoSlice, tarWr)
}

func doTar(toDoSlice []*os.File, tWr *tar.Writer) error {
	fmt.Printf("--- doTar: %d files\n", len(toDoSlice))

	for _, file := range toDoSlice {
		fmt.Printf("\t call each in doTar of %s\n", file.Name())
		if err := tarEach(file, "", tWr); err != nil {
			return err
		}
	}
	return nil
}

func tarEach(file *os.File, prePath string, tWr *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		// process dir
		if prePath == "" {
			prePath += info.Name()
		} else {
			prePath += "/" + info.Name()
		}
		// sub path
		subInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, subInfo := range subInfos {
			subFd, err := os.Open(file.Name() + "/" + subInfo.Name())
			if err != nil {
				return err
			}
			err = tarEach(subFd, prePath, tWr)
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Println("\t\tfile:", info.Name())
		// process file
		defer file.Close()
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		if prePath == "" {
			hdr.Name = info.Name()
		} else {
			hdr.Name = prePath + "/" + info.Name()
		}
		// tar file here
		if err := tWr.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err = io.Copy(tWr, file); err != nil {
			return err
		}
	}
	return nil
}
