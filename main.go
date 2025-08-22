package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LuizFreitas225/user-manager-go/src/middleware"
)

func main() {
	managerRouter := middleware.CreateRouterManager()
	// Cria um contexto que será cancelado ao receber sinal de interrupção
	// os.Interrupt = Representa o sinal de interrupção do terminal.
	//syscall.SIGTERM = É o sinal de término padrão do sistema.  kill -15 <pid> ou a/kill -9 <pid>
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	managerRouter.Start(ctx)

	// pg := &database.PostgresDatabase{}
	// db, err := pg.Open()

	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println("Connected!")
	// }
	// defer db.Close()

	log.Println("Aplicação Encerrada")
}
