# Receiver-Management API

## Projeto focado em desenvolver um CRUD para se trabalhar com recebedores e chaves pix segunido boas práticas e testes

O projeto teve como base as seguintes especificações: [Requisitos técnicos](https://docs.google.com/document/d/1XyjrQZgWG_m42OK6YR6MIm5MXY5INvSNWxEuFyexaaw/edit?usp=sharing)

## Requisitos minimos para rodar esse projeto localmente:

- Docker
- Docker Compose
- make

### Passo a passo

1. Clone esse repositório na pasta desejada

```bash
git clone https://github.com/vctaragao/receiver-management-api.git
```

2. Sentro da pasta do projeto rode o comando para subir a aplicação

```bash
cd <path-to-project>/receiver-management-api
make up
```

3. Para iniciar o servidor

```bash
make run
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

4. Para rodar os testes internos da aplicação e de integração da aplicação

```bash
make test && make test-integration
```

> Obs: Demais comandos disponiveis no arquivo Makefile

---

## Arquitetura

Para a arquitetura desse projeto foi levado como base a Clean Architecture. Buscando um desacoplamento entre as camadas e
uma arquitetura gritante.

```
.
├── cmd
│   └── server
├── internal
    ├── application
    │   ├── entity
    │   │   ├── helper
    │   │   ├── pix.go
    │   │   ├── receiver.go
    │   │   ├── repository.go
    │   │   └── test
    │   │       ├── pix
    │   │       └── receiver
    │   ├── receiver_management.go
    │   └── usecase
    │     ├── create_receiver
    │     ├── delete_receiver
    │     ├── list_receivers
    │     └── update_receiver
    ├── http
    │   ├── create_receiver.go
    │   ├── delete_receiver.go
    │   ├── list_receiver.go
    │   └── update_receiver.go
    ├── storage
    └── test
        ├── integration
        └── mocks
```

- **Framewors layer**

  - `/internal/http`: Pasta que contem a nossa camada de comunicação REST
  - `/internal/storage`: Pasta contendo as implementação dos adaptares para os nossos reposótorios de persistência
  - `/cmd/server`: Pasta contendo os binários do servidor.
  - `/tests`: Pasta para guardar os testes entre camadas da aplicação, ou de testes para a camada de framework

- **Application Layer**

  - `internal/application/receiver_management.go`: Facade para a comunição entre camadas externas e a camada de aplicação
  - `/internal/application/usecase`: Pasta para guardar os casos de uso do projeto

- **Domain Layer**
  - `/internal/application/entity`: Pasta para guardar o dominio do projeto (Lógicas e Regras de Negócio)

## Melhorias Futuras

- **Adicionar camada de validação na camada de Frameworks.**

  - Melhoria a perfomance da API, conseugindo bloquear logo na entrada Request que retornariam erro

- **Log estruturado**

  - Ajudaria no acompanhamento e debug do comportamento da API, principalmente em um ambiente de Cloud

- **Error Handler**

  - Aprimorar tratamento de erro para se obter melhores respostas para o usuário final
  - Montar um StackTrace para junto de um log estruturado melhorar a Observabilidade do serviço

- **Migrations**
  - Passar a utilizar um pacote especializado e migrations para conseguir manter um melhor controle da evolução do banco de dados
- **Validar perfomance do ORM**
  - ORM são bons para um desenvolvimento rápido, mas podem gerar problemas de mantenabilidade, migração de serviço ou perfomance.
    Por conta da camada de repossitório estar totalmente acoplada a um ORM. Então, seria bom validar em casos mais complexos se
    faz sentido se manter usando o ORM.
