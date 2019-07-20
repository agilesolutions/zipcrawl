package main
 
import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)
 
 
func main() {

    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "token")
        return
    }
    
    expression := os.Args[1]

    fmt.Println("searching for zip files with names containing : ", expression )
    fmt.Println()
    fmt.Println()
    

 	// filepath.Walk
 	files, err := FilePathWalkDir("./")
 	if err != nil {
  	panic(err)
 	}
 	for _, file := range files{
  		if (strings.HasSuffix(file, "zip")) {
	  		//fmt.Println(file)
			read, err := zip.OpenReader(file )
			if err != nil {
				msg := "Failed to open: %s"
				log.Fatalf(msg, err)
			}
			defer read.Close()

			for _, infile := range read.File {
				if err := listFiles(infile, file, expression); err != nil {
				log.Fatalf("Failed to read %s from zip: %s", infile.Name, err)
				}
			}

  		}
 	}


}

func FilePathWalkDir(root string) ([]string, error) {
 var files []string
 err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
  if !info.IsDir() {
   files = append(files, path)
  }
  return nil
 })
 return files, err
}


func listFiles(file *zip.File, filename string, expression string) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()
 
 	if (strings.Contains(file.Name, expression)) {
		fmt.Fprintf(os.Stdout, "%s -> %s:", filename, file.Name) 	
	    fmt.Println()
    }

 
	if err != nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
 
 
	return nil
}


