# desafio-digipix-
Formulário que consulte o CEP utilizando a API 

Estrutura do projeto
    ## main.go 
    Arquivo para iniciar a aplicacao

    ## .env
    Arquivo para armazenar as variaveis de ambiente

    ## app  
    Pasta que se encontra todo o projeto

    ## Controller
    Pasta com todos os aqruivos de controller

    ## Model 
    Pasta com todas as models do projeto

    ## Routes
    Pasta com os arquivos que contem as rotas da aplicacao

    ## Service
    PAsta com os arquivos da service(chamadas rest, persistencia etc)

    ## init.go
    Arquivo para inicializar ferramentas utilizadas pela aplicacao

    ## site
    Site para testar o servico de CEP. Coloquei aqui para unificar em um unico repositorio

    O projeto esta disponivel no repositorio

    https://github.com/julioshinoda/desafio-digipix

    Para rodar o projeto na sua maquina voce precisa ter o docker e docker-compose e o golang instalados na sua maquina. Em seguida rode os comandos na pasta do projeto pelo terminal

        go get -v

        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

        docker-compose build

        docker-compose up

    Depois acesse 

    http://localhost:8081/