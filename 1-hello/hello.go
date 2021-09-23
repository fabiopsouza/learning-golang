package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const DELAY = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		//	if comando == 1 {
		//		fmt.Println("Monitorando...")
		//	} else if comando == 2 {
		//		fmt.Println("Exibindo logs...")
		//	} else if comando == 0 {
		//		fmt.Println("Saindo do programa...")
		//	} else {
		//		fmt.Println("Não conheço essa instrução")
		//	}

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço essa instrução")
			os.Exit(-1)
		}
	}
}

func devolveNomeEVersao() (string, float64, int) {
	nome := "Douglas"
	versao := 1.1
	idade := 10
	return nome, versao, idade
}

func exibeIntroducao() {
	nome, versao, _ := devolveNomeEVersao() // _ -> ignora o terceiro parâmetro
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // & -> ponteiro que retorna o endereço da variável 'comando'
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://www.alura.com.br/"}                     // Exemplo Declaração inline
	// sites = append(sites, "https://random-status-code.herokuapp.com/") // Exemplo Append

	sites := leSitesDeArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {

		for _, site := range sites {
			testaSite(site)
		}

		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	// fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDeArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // os.O_RDWR|os.O_CREATE -> leia. Se não existir, cria
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") +
		" - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")
	// Trabalha somente com conteúdo em bytes e não arquivo
	arquivo, err := ioutil.ReadFile("log.txt") // Não fecha arquivo pois abre, entrega todos os bytes e fecha
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
