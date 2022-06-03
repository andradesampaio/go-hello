package tools

import (
	"fmt"
	"os"
)

func ExecutaMenu() {

	exibeIntroducao()	
	fmt.Println("")
	for {
		exibeMenu()
		comando := lerComando()
		checkComando(comando)
	}

}

func checkComando(comando int) {
	switch comando {
	case 1:
		IniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
		ImprimirLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Opçao Invalida")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("============================")
	fmt.Println(" MENU")
	fmt.Println("=============================")
	fmt.Println("")
	fmt.Println("1. - Iniciar Monitoramento")
	fmt.Println("2. - Exibir Logs")
	fmt.Println("0. - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Print("Selecione um opcao: ")
	fmt.Scan(&comandoLido)
	fmt.Println("")
	fmt.Println("Opcao escolhida foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeIntroducao() {
	versao := 1.2
	fmt.Println("Monitor Site")
	fmt.Println("Este programa está na versão", versao)
}