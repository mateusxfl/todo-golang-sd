### Criando container para a base de dados.
    
###### Primeiramente devemos abrir o Docker Desktop, após isso criamos o container para a base de dados.

```c
docker run -d --name db-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
```

###### Após isso, iremos conectar no postgres pelo container.

```c
docker exec -it db-todo psql -U postgres
```

###### Com isso, iremos criar nosso banco de dados e nosso usuário.

```c
create database db_todo;
create user user_todo;
alter user user_todo with encrypted password '1122';
```

###### Após isso iremos criar nossa tabela de Todos, primeiramente nos conectando na base de dados.

```c
\c db_todo; // Conecta na base de dados.
create table todos (id serial primary key, title varchar, description text, done bool default FALSE);
```

###### Após isso iremos garantir as permissões do usuário ao banco e tabela.

```c
grant all privileges on database db_todo to user_todo;
grant all privileges on all tables in schema public to user_todo;
grant all privileges on all sequences in schema public to user_todo;
```
    
##

### Criando container para a aplicação.
    
###### Em outro terminal, no diretório do Dockerfile, execute o seguinte comando para criar a imagem.

```c
docker build -t my-server .
```

###### Para criar uma rede, execute o seguinte comando.

```c
docker network create -d bridge ambiente-go
```

###### Após isso, iremos inserir o banco de dados na nossa rede.

```c
docker network connect ambiente-go db-todo
```

###### E finalmente, iremos rodar a aplicação na rede criada, com isso a aplicação já pode ser testada.

```c
docker run --network ambiente-go -p 8080:8080 my-server
```

###### Caso deseje verificar se os containers foram inseridos na rede, execute o seguinte comando em outro terminal.

```c
docker network inspect ambiente-go
```