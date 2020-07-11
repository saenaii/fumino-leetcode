package main

//go:generate go run generator.go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

type problemItem struct {
	Index int
	Name  string
	Link  string
}

type problemList []problemItem

type content struct {
	Title string
	List  problemList
}

const rootPath = "./"

func main() {
	genReadme(getFileList())
}

func getFileList() []os.FileInfo {
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var fileList []os.FileInfo
	for _, item := range files {
		if item.IsDir() && !strings.HasPrefix(item.Name(), ".") {
			fileList = append(fileList, item)
		}
	}
	return fileList
}

func genReadme(files []os.FileInfo) {
	indexPattern := regexp.MustCompile(`(\d+)\.`)
	namePattern := regexp.MustCompile(`\.([a-z_]+)\.go`)
	for _, file := range files {
		if readmeExists(file) {
			continue
		}
		var list problemList
		filepath.Walk(rootPath+file.Name(), func(currentPath string, info os.FileInfo, err error) error {
			if strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "test.go") {
				idxMatches := indexPattern.FindStringSubmatch(info.Name())
				nameMatches := namePattern.FindStringSubmatch(info.Name())
				if len(idxMatches) > 1 && len(nameMatches) > 1 {
					idx, err := strconv.Atoi(idxMatches[1])
					if err != nil {
						return err
					}
					list = append(list, problemItem{
						Index: idx,
						Name:  strings.Replace(nameMatches[1], "_", " ", -1),
						Link:  info.Name(),
					})
				}
			}
			return nil
		})
		genList(file.Name(), "./"+file.Name(), list)
	}
}

func readmeExists(file os.FileInfo) bool {
	fileList, err := ioutil.ReadDir(rootPath + file.Name())
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, file := range fileList {
		if strings.HasSuffix(strings.ToLower(file.Name()), "readme.md") {
			return true
		}
	}
	return false
}

func genList(title, filePath string, list problemList) {
	var b bytes.Buffer
	sort.Sort(list)
	tmpl, _ := template.ParseFiles("./readme.tmpl")
	tmpl.Execute(&b, content{
		Title: title,
		List:  list,
	})

	if f, err := os.Create(filePath + "/readme.md"); err == nil {
		defer f.Close()
		text, _ := ioutil.ReadAll(&b)
		f.Write(text)
	}
}

func (s problemList) Len() int {
	return len(s)
}

func (s problemList) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

func (s problemList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
