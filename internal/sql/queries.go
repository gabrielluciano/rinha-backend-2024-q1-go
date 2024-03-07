package sql

const TRANSACAO = "SELECT new_saldo, limite FROM transacao($1, $2, $3, $4)"

const EXTRATO = `
	SELECT tipo, valor, descricao, realizada_em
    FROM transacoes WHERE cliente_id = $1
    ORDER BY realizada_em DESC
    LIMIT 10;
`

const CLIENTE = "SELECT saldo, limite, current_timestamp FROM clientes WHERE id = $1"

const RESET_SALDOS = "UPDATE clientes SET saldo = 0"

const DELETE_TRANSACOES = "DELETE FROM transacoes"

const RESET_TRANSACOES_SEQ = "ALTER SEQUENCE transacoes_id_seq RESTART WITH 1"
