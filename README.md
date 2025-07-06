# 🚀 Go Jobs API

Uma API RESTful em Go para gerenciamento de vagas de emprego, construída com Gin, GORM e SQLite. Esta API permite criar, listar, atualizar e deletar vagas de emprego com documentação automática via Swagger.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Funcionalidades](#funcionalidades)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Arquitetura](#arquitetura)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Endpoints da API](#endpoints-da-api)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Documentação da API](#documentação-da-api)
- [Desenvolvimento](#desenvolvimento)
- [Contribuição](#contribuição)
- [Licença](#licença)

## 🎯 Sobre o Projeto

A **Go Jobs API** é uma aplicação backend desenvolvida em Go que fornece uma interface RESTful para gerenciamento de vagas de emprego. O projeto foi construído seguindo boas práticas de desenvolvimento, incluindo separação de responsabilidades, logging estruturado e documentação automática.

### Principais Características

- ✅ **API RESTful** com endpoints padronizados
- ✅ **Banco de dados SQLite** com GORM para ORM
- ✅ **Documentação automática** com Swagger
- ✅ **Validação de dados** robusta
- ✅ **Logging estruturado** para debugging
- ✅ **Arquitetura modular** e escalável
- ✅ **Makefile** para automação de tarefas

## 🚀 Funcionalidades

- **CRUD Completo de Vagas**: Criar, listar, atualizar e deletar vagas de emprego
- **Validação de Dados**: Validação automática de campos obrigatórios
- **Respostas Padronizadas**: Formato consistente de resposta JSON
- **Documentação Interativa**: Swagger UI para testar endpoints
- **Logging Estruturado**: Sistema de logs com diferentes níveis
- **Auto-migração**: Criação automática de tabelas no banco de dados

## 🛠 Tecnologias Utilizadas

### Core
- **[Go 1.24.2](https://golang.org/)** - Linguagem de programação
- **[Gin](https://github.com/gin-gonic/gin)** - Framework web HTTP
- **[GORM](https://gorm.io/)** - ORM para Go
- **[SQLite](https://www.sqlite.org/)** - Banco de dados

### Documentação
- **[Swagger](https://swagger.io/)** - Documentação da API
- **[gin-swagger](https://github.com/swaggo/gin-swagger)** - Integração Swagger com Gin

### Dependências Principais
```go
github.com/gin-gonic/gin v1.10.1
github.com/swaggo/files v1.0.1
github.com/swaggo/gin-swagger v1.6.0
github.com/swaggo/swag v1.16.4
gorm.io/driver/sqlite v1.6.0
gorm.io/gorm v1.30.0
```

## 🏗 Arquitetura

O projeto segue uma arquitetura modular com separação clara de responsabilidades:

```
go-jobs-api/
├── config/          # Configurações e inicialização
├── handler/         # Handlers HTTP (controllers)
├── router/          # Definição de rotas
├── schemas/         # Modelos de dados
├── db/              # Banco de dados SQLite
├── docs/            # Documentação Swagger gerada
├── main.go          # Ponto de entrada da aplicação
├── go.mod           # Dependências Go
├── go.sum           # Checksums das dependências
└── makefile         # Automação de tarefas
```

### Padrão de Arquitetura

- **Handler Layer**: Responsável por receber requisições HTTP e retornar respostas
- **Schema Layer**: Define os modelos de dados e estruturas de request/response
- **Config Layer**: Gerencia configurações, banco de dados e logging
- **Router Layer**: Define as rotas da API

## 📋 Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- **[Go 1.24.2 ou superior](https://golang.org/dl/)**
- **[Git](https://git-scm.com/)**
- **Make** (opcional, para usar o makefile)

## 🔧 Instalação

1. **Clone o repositório**
   ```bash
   git clone https://github.com/gumeeee/go-jobs-api.git
   cd go-jobs-api
   ```

2. **Instale as dependências**
   ```bash
   go mod download
   ```

3. **Execute a aplicação**
   ```bash
   # Com documentação Swagger
   make run-with-docs
   
   # Ou apenas executar
   make run
   
   # Ou diretamente com Go
   go run main.go
   ```

## 🚀 Como Usar

### Iniciando a Aplicação

```bash
# Desenvolvimento com documentação
make run-with-docs

# Apenas executar
make run

# Build da aplicação
make build

# Executar testes
make test

# Limpar arquivos gerados
make clean
```

### Acessando a API

Após iniciar a aplicação, você pode acessar:

- **API Base URL**: `http://localhost:8080/api/v1`
- **Documentação Swagger**: `http://localhost:8080/swagger/index.html`

## 📡 Endpoints da API

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints Disponíveis

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/opening` | Buscar uma vaga específica |
| `POST` | `/opening` | Criar uma nova vaga |
| `PUT` | `/opening` | Atualizar uma vaga existente |
| `DELETE` | `/opening` | Deletar uma vaga |
| `GET` | `/openings` | Listar todas as vagas |

### Exemplo de Uso

#### Criar uma Vaga
```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Desenvolvedor Go",
    "company": "TechCorp",
    "location": "São Paulo, SP",
    "remote": true,
    "link": "https://techcorp.com/careers",
    "salary": 8000
  }'
```

#### Listar Vagas
```bash
curl -X GET http://localhost:8080/api/v1/openings
```

## 📁 Estrutura do Projeto

### Diretórios e Arquivos

```
go-jobs-api/
├── config/
│   ├── config.go      # Configuração principal
│   ├── logger.go      # Sistema de logging
│   └── sqlite.go      # Configuração do SQLite
├── handler/
│   ├── handler.go     # Interface dos handlers
│   ├── request.go     # Estruturas de request
│   ├── response.go    # Estruturas de response
│   ├── createOpening.go
│   ├── listOpenings.go
│   ├── showOpening.go
│   ├── updateOpening.go
│   └── deleteOpening.go
├── router/
│   ├── router.go      # Inicialização do router
│   └── routes.go      # Definição das rotas
├── schemas/
│   └── opening.go     # Modelo de dados
├── db/
│   └── main.db        # Banco SQLite
├── docs/              # Documentação Swagger
├── main.go            # Ponto de entrada
├── go.mod             # Dependências
├── go.sum             # Checksums
└── makefile           # Automação
```

### Modelo de Dados

#### Opening Schema
```go
type Opening struct {
    gorm.Model
    Role     string
    Company  string
    Location string
    Remote   bool
    Link     string
    Salary   int64
}
```

#### Campos Obrigatórios
- `role`: Cargo da vaga
- `company`: Nome da empresa
- `location`: Localização da vaga
- `remote`: Se a vaga é remota
- `link`: Link da vaga
- `salary`: Salário da vaga

## 📚 Documentação da API

A documentação completa da API está disponível através do Swagger UI:

1. **Inicie a aplicação com documentação**:
   ```bash
   make run-with-docs
   ```

2. **Acesse a documentação**:
   ```
   http://localhost:8080/swagger/index.html
   ```

### Gerando Documentação

```bash
# Gerar documentação Swagger
make docs

# Ou manualmente
swag init
```

## 🛠 Desenvolvimento

### Estrutura de Logs

O projeto utiliza um sistema de logging estruturado com diferentes níveis:

- **DEBUG**: Informações detalhadas para debugging
- **INFO**: Informações gerais da aplicação
- **WARNING**: Avisos que não impedem a execução
- **ERROR**: Erros que impedem a execução

### Validação de Dados

Todos os endpoints implementam validação robusta:

- **Campos obrigatórios**: Verificação de campos necessários
- **Tipos de dados**: Validação de tipos corretos
- **Valores válidos**: Verificação de valores aceitáveis

### Respostas Padronizadas

A API retorna respostas em formato JSON padronizado:

#### Sucesso
```json
{
  "message": "operation from handler: create-opening successful",
  "data": {
    "id": 1,
    "role": "Desenvolvedor Go",
    "company": "TechCorp",
    "location": "São Paulo, SP",
    "remote": true,
    "link": "https://techcorp.com/careers",
    "salary": 8000,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z"
  }
}
```

#### Erro
```json
{
  "message": "parameter: role (type: string) is required",
  "errorCode": 400
}
```

## 🤝 Contribuição

1. **Fork o projeto**
2. **Crie uma branch para sua feature** (`git checkout -b feature/AmazingFeature`)
3. **Commit suas mudanças** (`git commit -m 'Add some AmazingFeature'`)
4. **Push para a branch** (`git push origin feature/AmazingFeature`)
5. **Abra um Pull Request**

### Padrões de Contribuição

- Siga as convenções de nomenclatura do Go
- Adicione testes para novas funcionalidades
- Mantenha a documentação atualizada
- Use commits semânticos

**Desenvolvido com ❤️ usando Go**
