package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId" xml:"userId"`
	Id        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*Todo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*Todo, error) {
	return rs.Remote.GetData()
}

type JsonBackend struct {
}

func (jb *JsonBackend) GetData() (*Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &todo, nil
}

type XMLBackend struct {
}

func (x *XMLBackend) GetData() (*Todo, error) {
	xmlFile := `
<?xml version="1.0" encoding="UTF-8" ?>
<root>
	<userId>321</userId>
	<id>456</id>
	<title>xml test</title>
	<completed>true</completed>
</root>
`

	var todo Todo
	_ = xml.Unmarshal([]byte(xmlFile), &todo)
	return &todo, nil
}

func main() {
	// no adapter
	todo := getRemoteData()
	fmt.Println("TODO without adapter: ", todo)

	// wiuth adfapter
	jsonB := &JsonBackend{}
	jsonAdapter := RemoteService{Remote: jsonB}
	td, _ := jsonAdapter.CallRemoteService()
	fmt.Println("adaptger", td)
	// xml

	xml := &XMLBackend{}
	xmlAdapter := RemoteService{Remote: xml}
	tdXml, _ := xmlAdapter.CallRemoteService()
	fmt.Println("xml adapter", tdXml)
}

func getRemoteData() *Todo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Println(err)
	}
	return &todo
}
