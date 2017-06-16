package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var filter []string

type FileStructure struct {
	Name  string
	Items []FileStructure
}

func StringTree(object FileStructure) (result string) {
	result += object.Name + "\n"
	var spaces []bool
	result += stringObjItems(object.Items, spaces)

	return
}

//
func filterRule(name string) (result bool) {

	for _, f := range filter {

		if strings.Contains(name, f) {
			continue
		} else {
			return false
		}
	}
	return true
}
func stringLine(name string, spaces []bool, last bool, isDir bool) (result string) {

	for _, space := range spaces {
		if space {
			result += "    "
		} else {
			result += "│   "
		}
	}

	indicator := "├── "
	if last {
		indicator = "└── "
	}
	result += indicator + name + "\n"

	return
}

func stringObjItems(items []FileStructure, spaces []bool) (result string) {
	for i, f := range items {

		last := (i >= len(items)-1)

		result += stringLine(f.Name, spaces, last, len(f.Items) > 0)

		if len(f.Items) > 0 {
			spacesChild := append(spaces, last)
			result += stringObjItems(f.Items, spacesChild)
		}

	}
	return
}

func PrintTree(object FileStructure) {
	fmt.Print(StringTree(object))
}

func ReadFolder(directory string) FileStructure {

	var parent FileStructure

	parent.Name = directory
	parent.Items = createReadFolder(directory)

	return parent
}

func createReadFolder(directory string) []FileStructure {

	var items []FileStructure
	files, _ := ioutil.ReadDir(directory)

	for _, f := range files {

		var child FileStructure
		child.Name = f.Name()

		if f.IsDir() {
			newDirectory := filepath.Join(directory, f.Name())
			child.Items = createReadFolder(newDirectory)
		}

		if len(filter) > 0 {

			if !f.IsDir() {
				if filterRule(f.Name()) {
					items = append(items, child)
				}
			} else {
				items = append(items, child)
			}

		} else {
			items = append(items, child)
		}

	}
	return items
}

func main() {
	directory := flag.String("directory", "", "directory tree ")
	filterStr := flag.String("filterStr", "", "filterStr split , ")
	flag.Parse()
	if *directory == "" {
		panic("directory file can not be emtpy")
	}
	if *filterStr != "" {
		aa := strings.Split(*filterStr, ",")
		filter = aa
	}

	obj := ReadFolder(*directory)
	PrintTree(obj)

}
