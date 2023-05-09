package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func renameFile(filePath, newFileName string, index, numDigits int, dryRun bool) error {
	ext := strings.ToLower(filepath.Ext(filePath))
	formatString := "%0" + strconv.Itoa(numDigits) + "d"
	newFilePath := filepath.Join(filepath.Dir(filePath), strings.ToLower(newFileName)+"-"+fmt.Sprintf(formatString, index)+ext)

	if dryRun {
		fmt.Printf("Dry run: renaming '%s' to '%s'\n", filePath, newFilePath)
		return nil
	}

	err := os.Rename(filePath, newFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("Renamed file '%s' to '%s'\n", filePath, newFilePath)
	return nil
}

type byDate []fs.DirEntry

func (d byDate) Len() int {
	return len(d)
}

func (d byDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d byDate) Less(i, j int) bool {
	infoI, _ := d[i].Info()
	infoJ, _ := d[j].Info()
	return infoI.ModTime().Before(infoJ.ModTime())
}

func main() {
	order := flag.String("order", "name", "Specify the order of the input files (name or date)")
	dryRun := flag.Bool("d", false, "Dry run: print the proposed changes but do not perform them")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: sequentialize [-order=name|date] [-d] directory new_file_name")
		return
	}

	dir := flag.Arg(0)
	newFileName := flag.Arg(1)

	if dir == "." {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	}

	files, err := os.ReadDir(dir)
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

	numDigits := len(fmt.Sprintf("%d", len(files)))

	for i, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		err := renameFile(filePath, newFileName, i+1, numDigits, *dryRun)
		if err != nil {
			fmt.Printf("Error renaming file '%s': %v\n", filePath, err)
		}
	}
}
