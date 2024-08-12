# Transfer API

[![Go](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-v1.10.0-green.svg)](https://github.com/gin-gonic/gin)
[![Zap](https://img.shields.io/badge/Zap-v1.27.0-yellow.svg)](https://github.com/uber-go/zap)

## Descrição

**Transfer API** é uma API para gerenciar transferências entre contas bancárias. Esta aplicação foi desenvolvida utilizando a linguagem Go e a arquitetura Clean Architecture, que promove a separação de responsabilidades e um código mais modular e de fácil manutenção.

## Estrutura do Projeto

O projeto é organizado de acordo com a Clean Architecture, que se baseia nos seguintes princípios:

- **Domínio**: Contém as entidades principais do sistema e suas regras de negócio.
- **Use Cases**: Contém a lógica de aplicação, orquestrando o fluxo de dados entre as entidades e os adaptadores.
- **Adapters**: Contém a interface para o mundo externo, como controladores HTTP e gateways de dados.
- **Frameworks & Drivers**: Contém bibliotecas e frameworks específicos de implementação, como o framework Gin para roteamento HTTP e o logger Zap para logging estruturado.

![Clean Architecture](images/clean_architecture.png)

Fonte: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

### Relação das entidades

![Entity Relation](images/relation_diagram.png)

### Tecnologias Utilizadas

- **Go**: Linguagem de programação para construção de serviços escaláveis e de alto desempenho.
- **Gin**: Framework web para Go que simplifica a construção de APIs.
- **Zap**: Biblioteca de logging de alto desempenho desenvolvida pela Uber.
- **Testify**: Framework de teste que facilita a escrita de testes unitários com mocks e asserções.

## Pré-requisitos

Antes de executar a aplicação, certifique-se de ter o seguinte software instalado:

- [Go 1.18+](https://golang.org/dl/)
- [Git](https://git-scm.com/)

## Instalação

1. Clone o repositório para o seu diretório local:

   ```bash
   git clone https://github.com/princeflaco/transfer-api-case-itau.git
   cd transfer-api-case-itau
2. Instale as dependências

    ```bash
    go mod download
3. (Opcional) Configure as variáveis de ambiente

    ```bash
   export PORT=8080
   export TIMEOUT=30
   export APP_NAME="Transfer API"
   export LOGGING_LEVEL=debug

4. Execute a aplicação

    ```bash
    go run main.go
A aplicação estará disponível em http://localhost:8080.

## Execução em Container

1. (Opcional) Configure as variáveis de ambiente: modifique o Dockerfile
2. Faça o build
   - Certifique que está no diretório raiz do projeto

   ```bash
   docker build -t transfer-api-case-itau .
3. Execute
   ```bash
   docker run -p 8080:8080 transfer-api-case-itau

## Informações

### Gerenciamento de concorrência nas execuções

Este projeto implementa um sistema de transferências com um controle de concorrência utilizando mutex no serviço de transferência e os padrões de design worker e queue no caso de uso.

#### Mutex no Serviço de Transferência
No TransferServiceImpl, um mutex é usado para garantir que as operações de transferência sejam realizadas de forma isolada. Isso garante que as atualizações de saldo das contas sejam executadas sem condições de corrida, mesmo com múltiplas transferências simultâneas.

#### Padrões Worker e Queue no Caso de Uso
O CreateTransferUseCase utiliza os padrões worker e queue para gerenciar a carga de trabalho:
- Queue: As transferências são enfileiradas em um canal para processamento assíncrono.
- Workers: Um conjunto de goroutines processa as transferências da fila, permitindo o processamento paralelo e eficiente das operações.

### Variáveis de Ambiente

O projeto utiliza variáveis de ambiente para configuração. As variáveis incluem:

- PORT: porta na qual a aplicação será executada. (padrão é 8080)
- TIMEOUT: tempo limite para requisições, em segundos. (padrão é 30 segundos)
- APP_NAME: nome da aplicação. (padrão é o nome do modulo, setado em go.mod)
- LOGGING_LEVEL: nível do Logger. (padrão é INFO)

### API REST

Endpoints Disponíveis:
  - POST /customers: cria um novo cliente e uma conta associada.
  - GET /customers: lista todos os clientes.
  - GET /customers/{accountId}: obtém um cliente pelo ID da conta.
  - POST /transfers/{accountId}: realiza uma transferência de uma conta para outra.
  - GET /transfers/{accountId}: obtém o histórico de transferências de uma conta.

#### Swagger
  A documentação da API está disponível via Swagger em:
  - http://localhost:8080/swagger/index.html (após execução do app)
  - Pode ser encontrada em ./docs

## Contribuição
1. Faça um fork do projeto.
2. Crie uma branch para sua feature (git checkout -b feature/nova-feature).
3. Faça um commit das suas mudanças (git commit -m 'Adiciona nova feature').
4. Faça o push para a branch (git push origin feature/nova-feature).
5. Abra um Pull Request.
