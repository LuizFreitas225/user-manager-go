package singleton

import (
	"database/sql"
	"log"
	"sync"

	"github.com/LuizFreitas225/user-manager-go/src/database"
	"github.com/gorilla/mux"
)

type Singleton struct {
	Db     *sql.DB
	Router *mux.Router
}

var singleInstance *Singleton

// Usado para bloquear a execução de um trecho de código por outra goroutine
var lock = &sync.Mutex{}

func GetInstance() *Singleton {
	if singleInstance == nil {
		lock.Lock()
		//double-checked locking.
		//Dupla checagem para garantir que não foi criada uma instância durante o perído de ativar o bloqueio
		if singleInstance == nil {
			pg := database.PostgresDatabase{}
			db, err := pg.Open()

			if err != nil {
				log.Panicln(err.Error())
			}

			singleInstance = &Singleton{
				Db:     db,
				Router: mux.NewRouter(),
			}
		}
	}
}
