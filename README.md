# climacep

Este é um microserviço em Go que recebe um CEP (Código de Endereçamento Postal) brasileiro, identifica a cidade correspondente e retorna o clima atual em várias escalas de temperatura: Celsius, Fahrenheit e Kelvin. Este projeto utiliza as APIs do ViaCEP, Nominatim e Open Meteo.

## Requisitos

- Docker e Docker Compose


## Funcionalidades

- Recebe um CEP válido de 8 dígitos.
- Realiza a pesquisa do CEP e encontra a localização.
- Retorna temperaturas em Celsius, Fahrenheit e Kelvin.
- Responde adequadamente em caso de sucesso, CEP inválido ou CEP não encontrado.

## Tecnologias Utilizadas

- Go
- Docker
- Google Cloud Run
- ViaCEP API
- Open Meteo API
- Nominatim (OpenStreetMap)

## Estrutura do Projeto

```plaintext
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── handlers.go
├── handlers_test.go
├── main.go
└── utils.go
```

## Configuração e Execução

### Passos para Executar Localmente

1. Clone o repositório:

```sh
git clone https://github.com/walterlicinio/climacep
cd climacep
```

2. Instale as dependências:

```sh
go mod tidy
```

3. Execute o servidor:

```sh
go run .
```

4. Teste o endpoint em seu navegador ou via `curl`:

```sh
curl http://localhost:8080/climate?cep=58045040
```

### Testes Automatizados

Execute os testes automatizados usando:

```sh
go test -v
```

### Rodar a Aplicação Usando Docker

1. Construa a imagem Docker:

```sh
docker-compose build
```

2. Inicie o serviço:

```sh
docker-compose up
```

3. Teste o endpoint em seu navegador ou via `curl`:

```sh
curl http://localhost:8080/climate?cep=58045040
```

### Deploy no Google Cloud Run

Exemplificado no link abaixo:
https://climacep-ripg4xth6q-uc.a.run.app/climate?cep=58045040