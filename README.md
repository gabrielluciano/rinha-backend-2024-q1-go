# Rinha Backend 2024 Q1 Go + Fiber 

Esta é minha submissão para a [Rinha de Backend - 2024/Q1](https://github.com/zanfranceschi/rinha-de-backend-2024-q1) utilizando Go com o Framework Fiber.

## Stack utilizada

- Go 1.22
  - fiber (framework HTTP)
  - jackc/pgx (postgresql driver) 
  - bytedance/sonic (json serializer)
- PostgreSQL
- Nginx

## Minhas redes sociais

- [GitHub](https://github.com/gabrielluciano)
- [Linkedin](https://www.linkedin.com/in/gabriel-lucianosouza/)
- [Twitter](https://twitter.com/biel_luciano)

## Getting Started

```shell
# Clonando o repo
git clone https://github.com/gabrielluciano/rinha-backend-2024-q1-go
cd rinha-backend-2024-q1-go

# Iniciando o projeto com o Docker Compose
docker compose up -d

# Testando a API
curl http://localhost:9999/clientes/1/extrato
```

## Gerando a Docker image localmente

```shell
cd rinha-backend-2024-q1-go

# Build da imagem
docker build -t <nome-da-imagem>:<tag> .
```
