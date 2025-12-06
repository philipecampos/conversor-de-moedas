package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type RatesFormat struct {
	Base  string
	Date  string
	Rates map[string]float64
}

func main() {
	args := os.Args

	// Valida a quantidade de argumentos (primeiro argumento args[0] é o nome do programa)
	if len(args) == 1 {
		fmt.Println("Informe o valor que a ser convertido")
		return
	}
	if len(args) == 2 {
		fmt.Println("Informe para qual moeda deseja converter: 'USD', 'EUR', 'JPY'...")
		return
	}

	// Converte o valor digitado para float64
	valor_em_brl, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("O valor a ser convertido é inválido.")
		return
	}

	// Remove os espaços laterais e converte o nome da moeda para maiúsculo
	modeda_destino := strings.ToUpper(strings.TrimSpace(args[2]))

	// Lê o arquivo contendo os dados para conversão
	jsonFile, err := os.Open("rates.json")
	if err != nil {
		fmt.Println("Não foi possível abrir arquivo rates.json")
		return
	}
	defer jsonFile.Close()

	// A leitura do arquivo retorna um array de bytes
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Falha ao converter arquivo em array de bytes")
		return
	}

	// Converte o array de bytes (de uma estrutura json) para a struct RatesFormat
	var jsonObj RatesFormat
	if err := json.Unmarshal(byteValue, &jsonObj); err != nil {
		fmt.Println("Falha ao converter array de bytes para RatesFormat")
		return
	}

	// Verifica se a moeda solicitada está contida no map Rates gerado ao ler o arquivo
	// e converte o valor usando o cálculo sugerido para a atividade
	valor, ok := jsonObj.Rates[modeda_destino]
	if !ok {
		fmt.Println("Moeda não suportada")
		return
	} else {
		fmt.Printf("%.2f\n", valor*valor_em_brl)
	}
}
