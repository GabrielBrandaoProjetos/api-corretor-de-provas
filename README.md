## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.24 ou superior
- [golang-migrate](https://github.com/golang-migrate/migrate)
  - Instalação: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)

## Setup do Projeto

1. Clone o repositório:
```bash
git clone https://github.com/GabrielBrandaoProjetos/corretor-de-provas-api.git
cd corretor-de-provas-api
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

3. Crie a tabela com migrate:
```bash
migrate create -ext sql -dir migrations -seq <nome da tabela. ex: create_users>
```

4. Execute as migrations:
```bash
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/corrector?sslmode=disable" up
```

5. Gerar as querie:
```bash
sqlc generate
```
