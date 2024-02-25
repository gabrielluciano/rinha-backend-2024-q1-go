package sql

const TRANSACAO = "SELECT new_saldo, limite FROM transacao($1, $2, $3, $4)"

const EXTRATO = `
	SELECT tipo, valor, descricao, realizada_em
    FROM transacoes WHERE cliente_id = $1
    ORDER BY realizada_em DESC
    LIMIT 10;
`

const CLIENTE = "SELECT saldo, limite FROM clientes WHERE id = $1"
