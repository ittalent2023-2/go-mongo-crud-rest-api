# 1 HACKATHON - IT TALENT 2023

O desafio principal deste hackathon consistirá o participante, ao término, ser capaz de configurar uma infraestrutura local com base em sistemas Linux, utilizando Vagrant, e implantar uma aplicação de API REST em Go, juntamente com um banco de dados MongoDB, empregando a tecnologia Docker.

Considere o cenário a seguir: "Uma equipe de desenvolvedores está trabalhando no desenvolvimento de uma API REST usando a linguagem Go e um banco de dados MongoDB. Essa aplicação tem uma funcionalidade simples, que se resume a criar, ler, atualizar e excluir informações sobre clientes. No entanto, devido à falta de experiência da equipe de desenvolvimento com a tecnologia Docker, eles solicitaram ao membro de DevOps do time que crie um Dockerfile para gerar uma imagem dessa aplicação. Além disso, o DevOps também foi encarregado de elaborar um arquivo docker-compose para facilitar a implantação dessa aplicação."

## OBJETIVOS ESPECÍFICOS:

Os objetivos específicos estão descritos no documento oficial do primeiro hackathon.

# Informações sobre a aplicação

Estas são operações CRUD no MongoDB escritas em Golang. Você pode criar, ler, atualizar e excluir usuários da instância do MongoDB usando solicitações http.

## Variáveis de ambiente

| Env            | Descrição |
| -------------- | --------- |
| MONGO_DB_HOST  | Nome do host do banco de dados.|

## Como executar localmente?
Primeiro, execute um container do mongodb com o docker:
```sh
docker run --name mongodb -d -p 27017:27017 mongo
```
Defina a variável de ambiente com o nome do host do mongoDB:
```sh
export MONGO_DB_HOST="localhost"
```
Em seguida, clone o repositório:
```sh
git clone git@github.com:parsaakbari1209/go-mongo-crud-rest-api.git
```
Em seguida, altere o diretório atual para o repositório:
```sh
cd go-mongo-crud-rest-api
```
Em seguida, instale as dependências:
```sh
go get ./...
```
Finalmente, execute o aplicativo na porta `9080`:
```sh
go run .
```

## Como executar com docker CLI?
Primeiro devemos criar uma network específica:
```sh
docker network create go-api-net
```
Execute um container do mongodb com o docker:
```sh
docker run --name mongodb -d -p 27017:27017 --network go-api-net mongo
```
Finalmente execute a app (esse comando não irá funcionar, pois é sua tarefa criar a imagem dessa APP):
```sh
docker run --name go-api-mongo -d -p 9080:9080 -e MONGO_DB_HOST="mongodb" --network go-api-net matheusmc/go-api-mongo
```

## Detalhes a serem observados:

* No `Dockerfile`, modificar o nome do binário que será gerado pelo passo de build do GO, exemplo:
```Dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api-mongo
```

```Dockerfile
ENTRYPOINT ["/go-api-mongo"]
```

* No docker-compose.yml, colocar a sua imagem do Docker Hub:
```yaml
services:
  go-crud:
    image: matheusmc/go-api-mongo
    restart: always
```


## Endpoints:
```sh
GET    /users/:email
POST   /users
PUT    /users/:email
DELETE /users/:email
```

### Get User
This endpoint retrieves a user given the email.
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:9080/users/bob@gmail.com'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "ilovealice"
  }
}
```
### Create User
This endpoint inserts a document in the `users` collection of the `users` database.
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:9080/users' -H "Content-Type: application/json" -d '{"name": "Bob", "email": "bob@gmail.com", "password": "ilovealice"}'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "ilovealice"
  }
}
```
### Update User
Este endpoint atualiza os campos fornecidos no documento especificado filtrado por email.
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:9080/users/bob@gmail.com' -H "Content-Type: application/json" -d '{"password": "loveyoualice"}'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "loveyoualice"
  }
}
```

### Delete User
Este endpoint exclui o usuário do banco de dados fornecido pelo email.
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:9080/users/bob@gmail.com'
```
Response:
```sh
{}
```

### Errors
All of the endpoints return an error in json format with a proper http status code, if something goes wrong:
```sh
{
  "error": "user not found"
}
```

## Conventions
Here is a list of conventions used:
- [Conventional commits](https://www.conventionalcommits.org/en/v1.0.0)
- [Google's API design guide](https://cloud.google.com/apis/design)
- [Uber's Go code style](https://github.com/uber-go/guide/blob/master/style.md)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
