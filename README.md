# Clean Architecture - Order System

Sistema de gerenciamento de pedidos desenvolvido em Go seguindo os princípios de Clean Architecture. O projeto expõe a criação e listagem de pedidos através de três interfaces: REST API, gRPC e GraphQL.

## Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)

## Como rodar

### 1. Subir tudo (MySQL, RabbitMQ, Migrations e Aplicação)

```bash
docker compose up -d
```

Isso irá:
- Subir o **MySQL** na porta `3306`
- Subir o **RabbitMQ** na porta `5672` (AMQP) e `15672` (Management UI)
- Executar as **migrações** automaticamente, criando a tabela `orders`
- Subir a **aplicação** expondo REST, gRPC e GraphQL

### 2. (Opcional) Rodar local sem Docker

Se quiser rodar a aplicação localmente (fora do Docker), use o `.env` em `cmd/ordersystem`:

```bash
cd cmd/ordersystem
go run .
```

## Portas dos Serviços

| Serviço         | Porta  | Endereço                          |
|-----------------|--------|-----------------------------------|
| REST API        | 8000   | http://localhost:8000             |
| gRPC            | 50051  | localhost:50051                   |
| GraphQL         | 8080   | http://localhost:8080             |
| MySQL           | 3306   | localhost:3306                    |
| RabbitMQ (AMQP) | 5672   | localhost:5672                    |
| RabbitMQ (UI)   | 15672  | http://localhost:15672            |

## Endpoints

### REST API

- **Criar pedido**: `POST http://localhost:8000/order`
- **Listar pedidos**: `GET http://localhost:8000/order`

Exemplo de criação de pedido:
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{"id": "order-1", "price": 100.0, "tax": 10.0}'
```

Exemplo de listagem:
```bash
curl http://localhost:8000/order
```

### gRPC

Porta: `50051`

Serviços disponíveis:
- `OrderService.CreateOrder` — Cria um pedido
- `OrderService.ListOrders` — Lista todos os pedidos

Exemplo com [grpcurl](https://github.com/fullstorydev/grpcurl):
```bash
# Criar pedido
grpcurl -plaintext -d '{"id":"order-1","price":100,"tax":10}' localhost:50051 pb.OrderService/CreateOrder

# Listar pedidos
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

### GraphQL

Playground: http://localhost:8080

**Mutation - Criar pedido:**
```graphql
mutation {
  createOrder(input: { id: "order-1", Price: 100.0, Tax: 10.0 }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Query - Listar pedidos:**
```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## Arquivo api.http

O arquivo `api/api.http` contém as requests prontas para serem executadas diretamente pela IDE (VS Code com extensão REST Client).

## Estrutura do Projeto

```
├── api/                          # Arquivo api.http
├── cmd/ordersystem/              # Entrypoint da aplicação
├── configs/                      # Configurações (Viper)
├── internal/
│   ├── entity/                   # Entidades de domínio
│   ├── event/                    # Eventos e handlers
│   ├── infra/
│   │   ├── database/             # Repositório MySQL
│   │   ├── graph/                # GraphQL (gqlgen)
│   │   ├── grpc/                 # gRPC (protobuf)
│   │   └── web/                  # REST API (Chi)
│   └── usecase/                  # Casos de uso
├── pkg/events/                   # Event Dispatcher
├── sql/migrations/               # Migrações SQL
├── docker-compose.yaml
└── gqlgen.yml
```
