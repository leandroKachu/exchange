# Exchange


### Descrição:
O Programa basicamente é composto por uma URL e quatro parametros amount-303/from-BRL/to-EUR/rate-4.5
Localizada em localhost:8000/exchange/303/BRL/EUR/4.5

## Instalação:
Foi usado para desenvolver o docker como gerador de image mysql.

Para seguir exatamente como eu fiz, segue os comandos abaixo:
"se não tiver docker, precisa instalar no seu terminal."

Caso tenha tudo certo, instale a images mysql no seu docker, apos isso criei 
`sudo docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=mysecret -e MYSQL_USER=user -p 3306:3306 -d mysql`
Após verificar que foi criado com sucesso, procure pelo o id do container no comando abaixo.

`sudo docker ps -a` localize pelo nome my-mysql e se o container não estiver up, suba

Após isso e acessar o db com o containerID, crie a database `conversion` e depois a tabela `conversions`
`sudo docker exec -it containerID bash`

O user de login:   o user `root`  com a senha passado na criacao da image
bash-4.4# `mysql -U root -p`
Enter password: mysecret

No projeto contem um arquivo chamado `sql.sql` com as query necessaria para criar a tabela. "obs: não se esqueça de criar o databese antes de usar sql.sql"

OBS: Validar o arquivo `.env` contém todas os parametros de acesso.



## Uso:

Va ao terminal, caminhe até a raiz do projeto e rode `go run main.go`

utilize o Postman ou semelhantes e execute 
`localhost:8000/exchange/303/USD/EUR/4.5` Method Post


