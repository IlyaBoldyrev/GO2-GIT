package main

import (
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
)

type myFile struct {
	Name string
	Size int64
	Path string
	Done bool
}

func searchAllFiles(dirName string, allFiles *[]myFile) error {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			err = searchAllFiles(dirName+"/"+v.Name(), allFiles)
			if err != nil {
				return err
			}
			continue
		}
		info, err := v.Info()
		if err != nil {
			return (err)
		}
		var temp myFile = myFile{
			Name: v.Name(),
			Size: info.Size(),
			Path: dirName + "/" + v.Name(),
		}
		*allFiles = append(*allFiles, temp)
	}
	return nil
}

func doSomethingWithDuplicates(allFilesIn []myFile, f func(string) error) error {
	var (
		allFiles = make([]myFile, len(allFilesIn))
		eg       = &errgroup.Group{}
	)
	copy(allFiles, allFilesIn)
	for i, v := range allFiles {
		for j := i + 1; j < len(allFiles); j++ {
			b := v
			k := j
			eg.Go(func() error {
				return iteration(b, k, allFiles, f)
			})
		}
	}
	return eg.Wait()
}

func iteration(v myFile, j int, allFiles []myFile, f func(string) error) error {
	if v.Name == allFiles[j].Name && v.Size == allFiles[j].Size {
		if !v.Done {
			err := f(v.Path)
			if err != nil {
				return err
			}
			v.Done = true
		}
		if !allFiles[j].Done {
			err := f(allFiles[j].Path)
			if err != nil {
				return err
			}
			allFiles[j].Done = true
		}
	}
	return nil
}

func main() {
	var (
		allFiles = make([]myFile, 0, 1000)
		path     = flag.String("path", "/home", "path to the directory in which we are looking for duplicates")
		remove   = flag.Bool("remove", false, "Switch indicating whether to remove duplicates")
		err      error
		s        string
	)
	flag.Parse()

	err = searchAllFiles(*path, &allFiles)
	if err != nil {
		panic(err)
	}

	err = doSomethingWithDuplicates(allFiles, func(s string) error {
		_, err := fmt.Println(s)
		return err
	})
	if err != nil {
		panic(err)
	}

	if *remove {
		fmt.Print("Do you really want to delete all duplicate files in the specified directory? (y/n)")
		fmt.Scanln(&s)
		if s == "Y" || s == "y" || s == "Yes" || s == "yes" {
			err = doSomethingWithDuplicates(allFiles, func(s string) error {
				err := os.Remove(s)
				return err
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
