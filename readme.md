
# Criando um pub/sub para enviar mensagem para fila do RabbitMq e consumindo numa api Golang

O projeto cria 2 containers docker um para a api em Go e outro com uma img do RabbitMq. A comunicação será feita através da lib ampq numa fila chamada jobs

## Documentação

Passo a passo para executar e instalar o projeto

git clone https://github.com/luisteixeira74/golang_rabbitmq

cd golang_rabbitmq

docker compose up (executa os containers)

-> vai subir 2 containers (web e o rabbitmq)

-> pagina admin do rabbitmq: credenciais dentro do docker-compose.yml  
http://localhost:15672/  


(GET) http://127.0.0.1:8080/  
-> endpoint hello-world em golang

(POST) http://127.0.0.1:8080/sender  
-> endpoint para enviar uma messagem com o body da mensagem:  

{
  "message": "Estudar golang"
}

Para consumir as mensagens, acessar o container web via shell, entrar no diretorio consumer, e rodar: go run consumer.go

para acessar o container do golang pode ser via vscode no proprio plugin do docker ou pela linha de comando usando o hash do seu container local: docker exec -it <containerId> sh

cd consumer  
go run consumer.go


