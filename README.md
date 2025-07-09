# go-task-api

API back-end para gerenciamento de tarefas desenvolvida em Go.  
O projeto utiliza hashing seguro para senhas, autenticação via token JWT e documentação da API com Swagger.

---

## Tecnologias e Bibliotecas Utilizadas

- [Go 1.24.3](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - framework HTTP
- [GORM](https://gorm.io/) - ORM para banco de dados MySQL
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - autenticação JWT
- Hashing seguro para senhas (bibliotecas Go padrão e `golang.org/x/crypto`)
- [Swaggo](https://github.com/swaggo/swag) - geração automática da documentação Swagger
- Outras dependências auxiliares para manipulação JSON, validação, etc.

---

## Funcionalidades

- CRUD completo para tarefas
- Registro e autenticação de usuários com senhas armazenadas de forma segura (hashing)
- Geração e validação de tokens JWT para autenticação de rotas protegidas
- Validação de dados de entrada usando bibliotecas dedicadas
- Documentação automática da API via Swagger UI

---

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/Claudio712005/go-task-api.git
cd go-task-api
```

2. Intale as dependências:

```bash
go mod download
```

3. Configure as variáveis de ambiente (exemplo no arquivo .env):

```bash
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=

JWT_SECRET=
```

4. Execute a aplicação:

```bash
go run .\cmd\main.go    
```

## Endpoints da API

A API está disponível no prefixo `/api`.

### Usuários (`/api/usuarios`)

| Método | Endpoint           | Descrição                    | Autenticação |
|--------|--------------------|------------------------------|--------------|
| POST   | `/usuarios`         | Cadastrar novo usuário        | Não          |
| GET    | `/usuarios/:id`     | Buscar usuário por ID         | Sim          |
| PUT    | `/usuarios/:id`     | Atualizar usuário por ID      | Sim          |
| DELETE | `/usuarios/:id`     | Deletar usuário por ID        | Sim          |
| POST   | `/usuarios/senha`   | Atualizar senha do usuário    | Sim          |

### Autenticação (`/api/autenticar`)

| Método | Endpoint       | Descrição                          | Autenticação |
|--------|----------------|----------------------------------|--------------|
| POST   | `/autenticar`  | Autenticar usuário e obter token JWT | Não          |

### Tarefas (`/api/tarefas`)

| Método | Endpoint                  | Descrição                      | Autenticação |
|--------|---------------------------|-------------------------------|--------------|
| POST   | `/tarefas`                | Criar nova tarefa              | Sim          |
| GET    | `/tarefas`                | Listar tarefas do usuário      | Sim          |
| GET    | `/tarefas/:id`            | Buscar tarefa por ID           | Sim          |
| PUT    | `/tarefas/:id`            | Atualizar tarefa por ID        | Sim          |
| DELETE | `/tarefas/:id`            | Deletar tarefa por ID          | Sim          |
| POST   | `/tarefas/:id/concluir`   | Marcar tarefa como concluída  | Sim          |
| GET    | `/tarefas/paginado`       | Listar tarefas paginadas       | Sim          |

## Documentação

A documentação automática da API está disponível via Swagger UI em:

```bash
http://localhost:8080/swagger/index.html
```