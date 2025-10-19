package main

import (
	"flag"
	"fmt"

	"github.com/BMokarzel/stress-test/pkg/logger"
	"github.com/BMokarzel/stress-test/pkg/test"
)

func main() {

	logger := logger.New()

	fmt.Println("---- Iniciando o programa ----")

	url := flag.String("url", "", "url do servidor a ser testado")

	requests := flag.Int("requests", 10, "Número de requisições que serão realizadas")

	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas que serão realizadas")

	flag.Parse()

	//if !strings.HasPrefix(*url, "http://") {
	//	u := fmt.Sprintf("http://" + *url)
	//	url = &u
	//}

	clientTest, err := test.New(logger, *url)
	if err != nil {
		logger.Write(fmt.Sprintf("Erro ao criar client. Erro: %s", err))
		fmt.Printf("Erro ao criar client. Erro: %s.\n\n---- Finalizando o programa ----", err)
		return
	}

	fmt.Println("\n---- Aguarde ----")

	err = clientTest.Run(*url, *requests, *concurrency)
	if err != nil {
		fmt.Printf("Erro durante o teste. Erro: %s.\n\n---- Finalizando o programa ----", err)
		return
	}

	fmt.Printf("\n---- Atenção ----\n\nOs logs da execução podem ser encontrados na pasta /logs")

	fmt.Println("\n---- Finalizando o programa ----")

}
