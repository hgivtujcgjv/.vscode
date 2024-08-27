package main

import (
	"fmt"
	"os"
	"sort"
)

func dirTree(path string) {
	var printDir func(string, string, bool)
	printDir = func(pref string, curpath string, last bool) {
		fileInput, err := os.Open(curpath)
		if err != nil {
			fmt.Println("Error opening directory:", err)
			return
		}
		defer fileInput.Close()

		ans, err := fileInput.Readdirnames(-1)
		if err != nil {
			fmt.Println("Error reading directory names:", err)
			return
		}

		sort.Strings(ans)

		for i, entry := range ans {
			l_entry := i == len(ans)-1
			fullPath := curpath + "/" + entry

			fileInfo, err := os.Stat(fullPath)
			if err != nil {
				fmt.Println("Error stating file:", err)
				continue
			}
			if fileInfo.IsDir() {
				if l_entry {
					fmt.Printf("%s└───%s", pref, entry)
					fmt.Println(" (", fileInfo.Size(), "bytes)")
					printDir(pref+"    ", fullPath, true)
				} else {
					fmt.Printf("%s├───%s", pref, entry)
					fmt.Println(" (", fileInfo.Size(), "bytes)")
					printDir(pref+"│   ", fullPath, false)
				}
			} else {
				if l_entry {
					fmt.Printf("%s└───%s", pref, entry)
					fmt.Println(" (", fileInfo.Size(), "bytes)")
				} else {
					fmt.Printf("%s├───%s", pref, entry)
					fmt.Println(" (", fileInfo.Size(), "bytes)")
				}
			}
		}
	}

	printDir("", path, false)
}

func main() {
	dirTree("/home")
}
