package util

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gabrielluciano/rinha-backend-2024-q1-go/config"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/sql"
	"github.com/gofiber/fiber/v2"
)

var Warm = false

func Warmup() {
	executeParallelAndWait(transacao)
	executeParallelAndWait(extrato)
	resetDatabase()
	Warm = true
}

func executeParallelAndWait(fn func(id int)) {
	var wg sync.WaitGroup
	for j := 1; j <= 50; j++ {
		for i := 1; i <= 5; i++ {
			v := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn(v)
			}()
		}
	}
	wg.Wait()
}

func transacao(id int) {
	agent := fiber.Post(fmt.Sprintf("http://127.0.0.1:%s/clientes/%d/transacoes", os.Getenv("HTTP_PORT"), id))
	agent.JSON(&fiber.Map{
		"valor":     10,
		"tipo":      "c",
		"descricao": "pix",
	})
	agent.Bytes()
}

func extrato(id int) {
	agent := fiber.Get(fmt.Sprintf("http://127.0.0.1:%s/clientes/%d/extrato", os.Getenv("HTTP_PORT"), id))
	agent.Bytes()
}

func resetDatabase() {
	config.Pool.Query(context.Background(), sql.RESET_SALDOS)
	config.Pool.Query(context.Background(), sql.DELETE_TRANSACOES)
	config.Pool.Query(context.Background(), sql.RESET_TRANSACOES_SEQ)
}
