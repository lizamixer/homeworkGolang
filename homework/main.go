package main

import (
	//"bytes" //понадобится для сборки больших строк перед выводом
	"fmt"
	//"sort"
	"strings"

	//"io"
	"os"            //для чтения структуры файловой системы
	"path/filepath" //чтобы рекурсивно обходить директории и файлы
	//"strings"
)


type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

func printTree(node *Node, indent string) {
	fmt.Println(indent + node.Name)
	for _, child := range node.Children {
		printTree(child, indent+"├───")
	}
}

func dirTree(out *os.File, path string, printFiles bool) (err error) {
	// var sortTree func (node *Node)
	// sortTree = func(node *Node) { //функция сортировки структуры по алфавиту
	// 	sort.Slice(child, func(i, j int) bool) {
	// 		return node.Children[i].Name < node.Children[j].Name
	// 	}
	// }
	err = filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		root := &Node{Name: path}
		//fmt.Printf("%+v", root)

		if err != nil {
			fmt.Println("Ошибка1:", err)
			return nil
		}

		var parts = strings.Split(path, "/")
		if d.IsDir() {
			var part string
			root.IsDir = true
			if len(parts) >= 2 {
				part = parts[len(parts)-2]
			} else {
				part = parts[len(parts)-1]
			}
			folder := &Node{Name: part, IsDir: true}
			root.Children = append(root.Children, folder)
		} else {
			if err != nil {
				fmt.Println("Ошибка2:", err)
				return nil
			}
			root.Children = append(root.Children, &Node{Name: parts[len(parts)-1], IsDir: false})
		}

		//sortTree (root)
		printTree(root, "")
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка обхода:", err)
	}

	return
}

func main() { //нельзя менять
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
