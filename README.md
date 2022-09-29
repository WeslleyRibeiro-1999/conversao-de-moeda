# Conversão de Moeda - Cambio
Este é um projeto de api no qual podemos fazer uma requisição com os dados do cambio e receber o retorno de acordo com o solicitado.

###Para começar baixei a imagem docker contendo o MYSQL para a utilização local:

### Clone este comando docker
```
$ docker run --name genesis -p 3306:3306 -e MYSQL_ROOT_PASSWORD=weslley123 -d mysql:8
```
Criei uma função para abrir conexão com o banco de dados [aqui](https://github.com/WeslleyRibeiro-1999/conversao-de-moeda/blob/main/database/connection.go).

E criei a tabela no mysql conforme abaixo:
```
CREATE TABLE cambio(
	ID int NOT NULL auto_increment,
    Moeda_Inicial varchar(5) NOT NULL,
    Moeda_Final varchar(5) NOT NULL,
    Cotacao decimal(6,2),
    Valor_Inicial decimal(10,2),
    Valor_Final decimal(10,2) AS (Cotacao*Valor_Inicial),
    Data_Cotacao datetime,
    primary key (ID)
);
```

Depois disto criei uma função para salvar dados sobre a conversão no MYSQL [aqui](https://github.com/WeslleyRibeiro-1999/conversao-de-moeda/blob/main/models/insertCotacao.go)

Criei a [api](https://github.com/WeslleyRibeiro-1999/conversao-de-moeda/blob/main/api/getCotacao.go) para requisição web, utilizei o protocolo GET para fazer a consulta via URL.
