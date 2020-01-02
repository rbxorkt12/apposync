package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"bufio"
)


type Parser struct{
	Results []Result `json:"results"`
}

type Result struct{
	Name string `json:"name"`

}
func main() {
	// Request 객체 생성
	// go function을 통환 프로세스 격리화
	// local file에 있는 값들을 불러와서 for 문을 통한 calling and parsing
	fi, err := os.Open("/tmp/parse_images.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fo, err := os.Create("/tmp/IMAGEVERSIONS")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	scanner := bufio.NewScanner(fi)
	// 루프
	for scanner.Scan(){
		// 읽기
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		requesturl:=fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags",scanner.Text())
		req, err := http.NewRequest("GET", requesturl, nil)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // 결과 출력

		bytes, _ := ioutil.ReadAll(resp.Body)
		var dat Parser
		if err := json.Unmarshal(bytes, &dat); err != nil {
			panic(err)
		}
		var resulttag string
		for _,result := range dat.Results{
			if(strings.HasPrefix(result.Name,"a")){
				resulttag=result.Name
				break
			}
		}
		if(resulttag==""){
			log.Printf("There is no matching tags in %s image",scanner.Text())
			continue
		}
		imagentag:=fmt.Sprintf("%s:%s\n",scanner.Text(),resulttag)
		_,err=fo.WriteString(imagentag)
		if err != nil {
			panic(err)
		}
	}
}
