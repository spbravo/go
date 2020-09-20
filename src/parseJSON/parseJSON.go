package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	//"math/big"
)
/*
type Data struct {
	Descripcion string `json:"descripcion"`
	Estado int `json:"estado"`
	Datos string`json:"datos"`
	Metadatos string `json:"metadatos"`
}*/

type Data struct {
	Descripcion string 
	Estado int 
	Datos string
	Metadatos string 
}

func main() {

	
	
		url := "https://opendata.aemet.es/opendata/api/observacion/convencional/datos/estacion/B228/?api_key=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJzcGJyYXZvQGdtYWlsLmNvbSIsImp0aSI6IjgyNGM1Y2UzLWFmNzYtNDk0NS1hZDBmLTdhMDk1ZTIyMzJkZCIsImlzcyI6IkFFTUVUIiwiaWF0IjoxNTA5NTU2OTY5LCJ1c2VySWQiOiI4MjRjNWNlMy1hZjc2LTQ5NDUtYWQwZi03YTA5NWUyMjMyZGQiLCJyb2xlIjoiIn0.XAEUT7p_9sXrkavMunL9CtQySwUWOicCIbGfsYxdVZk"
	
		req, _ := http.NewRequest("GET", url, nil)
	
		req.Header.Add("cache-control", "no-cache")
		res, _ := http.DefaultClient.Do(req)
	
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		
		var urlDato Data
		err := json.Unmarshal(body, &urlDato)
		if err != nil{ 
			fmt.Println("Estoy en el error %v \n", err)
		}
			fmt.Println(string(body))
			fmt.Println("Struct is:", urlDato) 
	    //data := dataFromJson(string(body))
			fmt.Printf("url de datos %s \n", string(urlDato.Datos))
		resultado :=  contentFromServer(urlDato.Datos)
			fmt.Println(resultado)
/*

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJson(content)
	// fmt.Println(tours)
	
	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}*/
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	
	resp, err := http.Get(url)
	checkError(err)
	
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

/*func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	
	return tours
}*/
func dataFromJson(content string) [] Data{
	fmt.Printf("este es el string que llega para decodificar %v", content)
	datos := make([]Data, 0, 1)
	decoder := json.NewDecoder(strings.NewReader(content))
	var data Data
	for decoder.More() {
		err := decoder.Decode(&data)
		checkError(err)
	 datos = append(datos, data)
	 fmt.Printf("Esto es el data %v", data.Metadatos)
	}
	return datos
}