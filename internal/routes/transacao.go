package routes

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gabrielluciano/rinha-backend-2024-q1-go/config"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/sql"
	"github.com/gofiber/fiber/v2"
)

const INSUFICIENT_SALDO_CODE = -1

type TransacaoRequest struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

func (t *TransacaoRequest) isInvalid() bool {
	invalidValor := t.Valor < 1
	invalidDescricao := len(t.Descricao) < 1 || len(t.Descricao) > 10
	invalidTipo := t.Tipo != "c" && t.Tipo != "d"
	return invalidValor || invalidDescricao || invalidTipo
}

func Transacao(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > 5 {
		return c.Status(fiber.StatusNotFound).Send(nil)
	}

	transacao := new(TransacaoRequest)
	if err = c.BodyParser(transacao); err != nil || transacao.isInvalid() {
		return c.Status(fiber.StatusUnprocessableEntity).Send(nil)
	}

	var saldo, limite int
	err = config.Pool.QueryRow(context.Background(), sql.TRANSACAO,
		id, transacao.Valor, transacao.Tipo, transacao.Descricao).Scan(&saldo, &limite)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing transacao: %v", err)
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}

	if limite == INSUFICIENT_SALDO_CODE {
		return c.Status(fiber.StatusUnprocessableEntity).Send(nil)
	}

	return c.JSON(fiber.Map{
		"saldo":  saldo,
		"limite": limite,
	})
}
