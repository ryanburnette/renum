package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func renameImageFile(filePath, newFileName string, index, numDigits int) error {
	ext := strings.ToLower(filepath.Ext(filePath))
	formatString := "%0" + strconv.Itoa(numDigits) + "d"
	newFilePath := filepath.Join(filepath.Dir(filePath), strings.ToLower(newFileName)+"-"+fmt.Sprintf(formatString, index)+ext)

	err := os.Rename(filePath, newFilePath)
	if err != nil {
		return err
	}
	return nil
}

type byDate []os.FileInfo

func (d byDate) Len() int {
	return len(d)
}

func (d byDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d byDate) Less(i, j int) bool {
	return d[i].ModTime().Before(d[j].ModTime())
}

func main() {
	order := flag.String("order", "name", "Specify the order of the input files (name or date)")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: sequentialize [-order=name|date] directory new_file_name")
		return
	}

	dir := flag.Arg(0)
	newFileName := flag.Arg(1)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	if *order == "date" {
		sort.Sort(byDate(files))
	} else if *order == "name" {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() < files[j].Name()
		})
	} else {
		fmt.Println("Invalid order option. Use 'name' or 'date'")
		return
	}

	imageFiles := []string{}
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".heic" {
			imageFiles = append(imageFiles, filepath.Join(dir, file.Name()))
		}
	}

	numDigits := len(fmt.Sprintf("%d", len(imageFiles)))

	for i, filePath := range imageFiles {
		err := renameImageFile(filePath, newFileName, i+1, numDigits)
		if err != nil {
			fmt.Printf("Error renaming file '%s': %v\n", filePath, err)
		} else {
			fmt.Printf("Renamed file '%s'\n", filePath)
		}
	}
}
