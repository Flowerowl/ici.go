package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

const KEY string = "E0F0D336AF47D3797C68372A869BDBC5"
const URL string = "http://dict-co.iciba.com/api/dictionary.php"

type Dict struct {
	Key         string   `xml:"key"`
	Ps          []string `xml:"ps"`
	Pos         []string `xml:"pos"`
	Acceptation []string `xml:"acceptation"`
	SentList    []Sent   `xml:"sent"`
}

type Sent struct {
	Orig  string `xml:"orig"`
	Trans string `xml:"trans"`
}

func get_data(words string) []byte {
	response, err := http.Get(URL + "?key=" + KEY + "&w=" + words)
	if err != nil {
		fmt.Println("啊哦，好像出错了，try again~")
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	return data
}

func parse_xml(data []byte) Dict {
	var dict Dict
	xml.Unmarshal(data, &dict)

	return dict
}

func show(dict Dict) {
	for _, ps := range dict.Ps {
		color.Green(ps)
	}
	for index, pos := range dict.Pos {
		color.Red(strings.TrimSpace(pos))
		color.Yellow(strings.TrimSpace(dict.Acceptation[index]))
	}
	for _, sent := range dict.SentList {
		color.Blue("ex. %s", strings.TrimSpace(sent.Orig))
		color.Cyan("    %s", strings.TrimSpace(sent.Trans))
	}
}

func main() {
	flag.Parse()
	var words = strings.Join(flag.Args(), " ")

	data := get_data(words)
	dict := parse_xml(data)
	show(dict)
}
