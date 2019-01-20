# desafio-digipix-
Formulário que consulte o CEP utilizando a API 


    go get -v

    govendor init



    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t digi-app -f Dockerfile.scratch .
docker run -it -p 8080:8080 digi-app
