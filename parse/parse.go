package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Appoconfig struct {
	Kind string
	Images []Image
}
type Image struct {
	Name string
	Gitrepos []string
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for _,v := range elements {
		if encountered[v] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[v] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}



func main(){
	appoconfig := &Appoconfig{}
	viper.SetConfigName("Appoconfig")
	viper.AddConfigPath("/tmp/argogyu")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Can't read config file: %s \n", err)
	}
	if err := viper.Unmarshal(appoconfig); err != nil {
		fmt.Errorf("config file format error: %s \n", err)
	}
	fo1, err := os.Create("/tmp/parse_images.txt")
	if err != nil {
		panic(err)
	}
	defer fo1.Close()
	fo2, err := os.Create("/tmp/parse_repoes.txt")
	if err != nil {
		panic(err)
	}
	defer fo2.Close()
	var reponames []string
	for _,image := range appoconfig.Images{
		_,err1:=fo1.WriteString(image.Name)
		if err1 != nil {
			panic(err)
		}
		for _,reponame := range image.Gitrepos{
			reponames = append(reponames,reponame)

		}
		_,err1 = fo1.WriteString("\n")
	}
	reponames=removeDuplicates(reponames)
	for _, reponame := range reponames{
		_,err2 := fo2.WriteString(reponame)
		if err2 != nil {
			panic(err)
		}
		_,err2 = fo2.WriteString("\n")
	}
}
