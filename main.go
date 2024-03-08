package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Payload struct {
	Data string `json:"data"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run main.go http://ip host alvo")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	maxID := 100

	client := &http.Client{}

	for i := 1; i <= maxID; i++ {
		req, err := http.NewRequest("GET", baseURL+strconv.Itoa(i), nil)
		if err != nil {
			fmt.Println("Erro ao criar a requisição: ", err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erro ao fazer a requisição: ", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Erro: status code ", resp.StatusCode)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler a resposta: ", err)
			continue
		}

		fmt.Println("Resposta para ID ", i, ": ", string(body))

		// Exemplo de requisição POST
		data := Payload{
			Data: "your-data",
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Erro ao codificar os dados: ", err)
			continue
		}

		req, err = http.NewRequest("POST", baseURL+strconv.Itoa(i), bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Erro ao criar a requisição POST: ", err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Erro ao fazer a requisição POST: ", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Erro na requisição POST: status code ", resp.StatusCode)
			continue
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler a resposta da requisição POST: ", err)
			continue
		}

		fmt.Println("Resposta para a requisição POST com ID ", i, ": ", string(body))
	}
}
