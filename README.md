# SeuCantinho

SeuCantinho é um sistema de gerenciamento de reservas de espaços, desenvolvido em **Go**, composto por **cliente**, **servidor** e **banco de dados**, todos executando em containers **Docker**. O sistema permite que usuários façam reservas, pagamentos e consultas, enquanto administradores gerenciam os espaços disponíveis.

---

## 🚀 Funcionalidades

### 👤 Usuários Comuns

* Criar e visualizar reservas
* Listar espaços disponíveis
* Realizar pagamentos de reservas
* Cancelar reservas
* Visualizar histórico de reservas

### 🛠️ Administradores

* Adicionar novos espaços
* Remover espaços existentes
* Atualizar informações dos espaços
* Gerenciar reservas

### 👑 Usuário Administrador Padrão

O sistema já inclui um usuário administrador inicial:

```
Username: DonaMaria
Senha: SeuCantinho123
```

---

## 🧱 Arquitetura

O projeto é dividido em três componentes principais:

* **Servidor (server):** API responsável pelas operações de negócios e comunicação com o banco de dados.
* **Cliente (client):** Interface CLI para interação com o servidor.
* **Banco de Dados (db):** Armazena usuários, espaços, reservas e pagamentos.

Tudo é gerenciado via **Docker Compose**.

---

## 🐳 Como Executar o Projeto

### 0. Gerar o arquivo da documentação

Navegue até o primeiro diretório `server` e gere os arquivo `swagger.json` e `swagger.yaml` com a seguinte linha de comando:

```bash
swag init -g cmd/server/main.go
```

### 1. Subir todos os serviços

Execute:

```bash
docker compose up -d
```

Isso iniciará o servidor, o cliente e o banco de dados.

### 2. Executar o cliente

Após os containers estarem rodando:

```bash
docker compose exec client /app/client
```

---

## 📁 Estrutura do Projeto (resumo)

```
SeuCantinho/
 ├── client/        # Aplicação cliente em Go
 ├── server/        # API do servidor em Go
 ├── diagrams/      # Diagramas UML de Componentes e de Classes
 ├── docker-compose.yml
 └── README.md
```

---