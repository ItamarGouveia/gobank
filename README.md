
# Golang JSON API

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)](https://golang.org/) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue?logo=postgresql)](https://www.postgresql.org/) [![Docker](https://img.shields.io/badge/Docker-24-blue?logo=docker)](https://www.docker.com/)

## Descrição do Projeto

Este projeto é uma API RESTful construída em **Golang**, com **autenticação JWT**, **banco de dados PostgreSQL** e **Docker**. Ele segue boas práticas de desenvolvimento de software e arquitetura limpa, oferecendo um ponto de partida robusto para aplicações modernas.

Principais funcionalidades:

- CRUD completo (Create, Read, Update, Delete) para recursos de usuário.
- Autenticação segura com JWT.
- Integração com PostgreSQL para armazenamento de dados.
- Contêinerização com Docker para ambiente de desenvolvimento consistente.
- Estrutura modular, pronta para escalabilidade.

## Tecnologias Utilizadas

- **Golang (Go 1.21+)** – linguagem principal da API.
- **PostgreSQL** – banco de dados relacional.
- **Docker & Docker Compose** – para contêinerização do serviço e banco de dados.
- **JWT (JSON Web Token)** – autenticação de usuários.
- **Go Modules** – gerenciamento de dependências.

## Pré-requisitos

Antes de rodar o projeto, você precisará ter instalado:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/downloads) (opcional, para clonar o repositório)

## Estrutura do Projeto

\`\`\`text
.
├── cmd/                 # Arquivos executáveis da API
├── internal/            # Código principal (handlers, services, models)
│   ├── auth/            # Lógica de autenticação JWT
│   ├── database/        # Conexão com o PostgreSQL
│   └── users/           # CRUD de usuários
├── config/              # Configurações do projeto
├── migrations/          # Scripts SQL para criar tabelas
├── Dockerfile           # Imagem Docker da API
├── docker-compose.yml   # Configuração do Docker Compose
├── go.mod               # Gerenciador de dependências
└── README.md            # Este arquivo
\`\`\`