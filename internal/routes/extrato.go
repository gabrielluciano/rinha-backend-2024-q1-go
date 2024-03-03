package routes

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gabrielluciano/rinha-backend-2024-q1-go/config"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/sql"
	"github.com/gofiber/fiber/v2"
)

type ClienteQueryResult struct {
	saldo fiber.Map
	err   error
}

type TransacoesQueryResult struct {
	transacoes []fiber.Map
	err        error
}

func Extrato(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > 5 {
		return c.Status(fiber.StatusNotFound).Send(nil)
	}

	clienteQueryCh := make(chan *ClienteQueryResult, 1)
	transacoesQueryCh := make(chan *TransacoesQueryResult, 1)

	go fetchCliente(id, clienteQueryCh)
	go fetchTransacoes(id, transacoesQueryCh)

	clienteQueryResult := <-clienteQueryCh
	transacoesQueryResult := <-transacoesQueryCh

	if clienteQueryResult.err != nil {
		fmt.Fprintf(os.Stderr, "Error processing extrato: %v", clienteQueryResult.err)
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}

	if transacoesQueryResult.err != nil {
		fmt.Fprintf(os.Stderr, "Error processing extrato: %v", err)
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}

	return c.JSON(fiber.Map{
		"saldo":              clienteQueryResult.saldo,
		"ultimas_transacoes": transacoesQueryResult.transacoes,
	})
}

func fetchCliente(id int, ch chan *ClienteQueryResult) {
	var total, limite int
	var dataExtrato time.Time

	err := config.Pool.QueryRow(context.Background(), sql.CLIENTE, id).Scan(&total, &limite, &dataExtrato)
	if err != nil {
		ch <- &ClienteQueryResult{err: err}
		close(ch)
		return
	}

	ch <- &ClienteQueryResult{
		saldo: fiber.Map{
			"total":        total,
			"limite":       limite,
			"data_extrato": dataExtrato.UTC().Format(time.RFC3339Nano),
		},
	}
	close(ch)
}

func fetchTransacoes(id int, ch chan *TransacoesQueryResult) {
	rows, err := config.Pool.Query(context.Background(), sql.EXTRATO, id)
	if err != nil {
		ch <- &TransacoesQueryResult{err: err}
		close(ch)
		return
	}

	transacoes := make([]fiber.Map, 0, 10)
	for rows.Next() {
		var (
			tipo        string
			descricao   string
			realizadaEm time.Time
			valor       int
		)
		rows.Scan(&tipo, &valor, &descricao, &realizadaEm)
		transacoes = append(transacoes, fiber.Map{
			"valor":        valor,
			"tipo":         tipo,
			"descricao":    descricao,
			"realizada_em": realizadaEm.UTC().Format(time.RFC3339Nano),
		})
	}

	ch <- &TransacoesQueryResult{transacoes: transacoes}
	close(ch)
}
