package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://viacep.com.br/ws/60766320/json
// https://mholt.github.io/json-to-go/
// convert json to strunc golag

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(data)

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivos: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade:%s, UF:%s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado")
		fmt.Println("Cidade: ", data.Localidade)

	}

}
