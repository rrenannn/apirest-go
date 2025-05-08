# 📦 Desafio: Criar uma API em Go com Echo e SQLC

Este projeto é uma API RESTful desenvolvida em Go utilizando o framework [Echo](https://echo.labstack.com/), com integração ao banco de dados PostgreSQL via [SQLC](https://docs.sqlc.dev/). A aplicação está containerizada com Docker para facilitar o desenvolvimento e a execução.

---

## 🚀 Tecnologias

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [SQLC](https://docs.sqlc.dev/)
- [Docker](https://www.docker.com/)
- [Swagger (swaggo)](https://github.com/swaggo/swag)

---

## 📂 Estrutura do Projeto

```
.
├── cmd/                # Arquivo main.go (entrada da aplicação)
├── internal/
│   ├── db/             # Código gerado pelo SQLC
│   ├── repository/     # Acesso ao banco de dados
│   ├── usecase/        # Lógica de negócio
│   └── handler/        # Controladores HTTP (Echo)
├── docker-compose.yml
├── Dockerfile
├── go.mod / go.sum
└── README.md
```

---

## ⚙️ Como rodar com Docker

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Suba os containers com Docker Compose:**
   ```bash
   docker-compose up --build
   ```

3. A API estará disponível em: [http://localhost:8080](http://localhost:8080)

---

## 📌 Endpoints da API

> Exemplo de endpoints disponíveis (substitua conforme seu projeto):

| Método | Rota               | Descrição               |
|--------|--------------------|-------------------------|
| GET    | `/users`           | Listar todos os usuários |
| POST   | `/users`           | Criar um novo usuário    |
| GET    | `/users/:id`       | Buscar usuário por ID    |
| PUT    | `/users/:id`       | Atualizar usuário        |
| DELETE | `/users/:id`       | Deletar usuário          |

---

## 📘 Documentação Swagger

Após rodar o projeto, acesse:

📚 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🛠 Comandos úteis (com SQLC)

- Gerar o código do banco:
  ```bash
  sqlc generate
  ```


