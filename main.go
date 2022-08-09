package main

import (
	"fmt"
	"log"
	"os"

	//"paths"
	"encoding/json"
	"io/ioutil"
	"regexp"
)

type Path struct {
	Pattern string `json:"pattern"`
	ToPath  string `json:"to"`
}

type Paths struct {
	Paths []Path `json:"paths"`
}


func sort() {
	sep := string(os.PathSeparator)
	paths := getData()
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	sortFolder := dirname + sep + "Downloads"
	fmt.Println(sortFolder)
	items, _ := ioutil.ReadDir(sortFolder)
	for _, item := range items {
		// fmt.Println(item.Name())
		if !item.IsDir() {
			oldLocation := sortFolder + sep + item.Name()
			newLocation := sortFolder + sep + item.Name()
			for i := 0; i < len(paths.Paths); i++ {
				fmt.Printf("Checking for %s...\n", paths.Paths[i].Pattern);
				match, _ := regexp.MatchString(paths.Paths[i].Pattern, item.Name())
				if match {
					newLocation = paths.Paths[i].ToPath + "/" + item.Name()
					fmt.Printf("%s -> %s\n", oldLocation, newLocation)
					break
				}
			}
			err := os.Rename(oldLocation, newLocation)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// fmt.Println(item.Name())
		}
	}
}

func remove(path string) {
	// GitHub copilot suggested this
	// I have no idea what it does or why it works
	paths := getData()
	for i := 0; i < len(paths.Paths); i++ {
		if paths.Paths[i].Pattern == path {
			paths.Paths = append(paths.Paths[:i], paths.Paths[i+1:]...)
		}
	}
	fmt.Println(paths)
	dataBytes, err := json.Marshal(paths)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("settings.json", dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getData() Paths {
	var paths Paths
	data, err := os.Open("settings.json")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, &paths)
	return paths
}

func add(pattern, toPath string) {
	dir, err := os.Stat(toPath)
	if err != nil || !dir.IsDir() {
		fmt.Printf("\"%s\" is not a valid directory\n", toPath)
		return
	}
	paths := getData()
	newPath := &Path{
		Pattern: pattern,
		ToPath:  toPath,
	}
	paths.Paths = append(paths.Paths, *newPath)
	dataBytes, err := json.Marshal(paths)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("settings.json", dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func printPaths() {
	paths := getData()
	for i := 0; i < len(paths.Paths); i++ {
		fmt.Printf("%d: %s -> %s\n", i, paths.Paths[i].Pattern, paths.Paths[i].ToPath)
	}
}

func main() {
	f, err := os.OpenFile("settings.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if file, _ := os.Stat("settings.json"); file.Size() == 0 {
		_, err = f.Write([]byte("{\"paths\":[]}"))
	}
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	if len(os.Args) == 1 {
		sort()
	}
	for index, arg := range os.Args {
		switch arg {
		case "-a", "--add", "--append":
			add(os.Args[index+1], os.Args[index+2])
			break
		// case "-h", "--help":
		// 	printHelp(os.Args[index+1])
		// 	break
		case "-s", "--sort":
			sort()
			break
		case "-p", "--print":
			printPaths()
			break
		case "-r", "--remove":
			remove(os.Args[index+1])
			break
		default:
			break
		}
	}
	fmt.Println(os.Args)
}
