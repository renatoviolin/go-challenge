# Challenge in Golang

This repository implements the challenge that asks to implement an API with the following requirements.
* import dataset from csv file;
* create a new Feira object;
* remove a Feira object given an id;
* update a Feira object, except the id;
* find Feira objects by at one of the params:
    1. distrito
    2. regiao
    3. nome_feira
    4. bairro

Full description available [here](assets/challenge.pdf)

## Services available

|Method|Endpoint|Payload|Output
|---|---|---|---|
|GET|/api/v1/health||{"status": "online"}|
|POST|/api/v1/feira|FeiraPayload|{"id": "123"}|
|PUT|/api/v1/feira|FeiraPayload|{"message": "successfully updated"}|
|DELETE|/api/v1/feira/:id|id|{"message": "successfully deleted"}|
|GET|/api/v1/nome/:query|query|{data: [FeiraPayload, FeiraPayload, ...]}|
|GET|/api/v1/regiao/:query|query|{data: [FeiraPayload, FeiraPayload, ...]}|
|GET|/api/v1/distrito/:query|query|{data: [FeiraPayload, FeiraPayload, ...]}|
|GET|/api/v1/bairro/:query|query|{data: [FeiraPayload, FeiraPayload, ...]}|

Postman collection available [here](assets/GoChallenge.postman_collection.json).

FeirPayload has the following fields:
```json
{
    "cod_dist": "87",
    "distrito": "VILA FORMOSA",
    "sub_prefe": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "VILA FORMOSA",
    "registro": "4041-0",
    "logradouro": "RUA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA",
    "set_cens": 355030885000091,
    "id": 1,
    "long": -46550164,
    "lat": -23558733,
    "areap": 3550308005040,
    "cod_sub_pref": 26
}
```

## Install
1. Clone this repository.
```
git clone xxxxxxxxx
```
2. Review the .env file to adjust the credentials of your choice.
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_HOST=host.docker.internal
```
3. Start the database and import the data. It is necessary to start the dabase before the application so that the test case can be performed.
```
make run-db
```
4. Run tests, compile the binary and start the application.
```
make run-app
```

## Run in development enviroment
1. Clone this repository.
```
git clone http://git@github.com/renatoviolin/challenge-unico-go
```
2. Review the .env file. Make sure to adjust POSTGRES_HOST to localhost.
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_HOST=locahost
```
3. Start the database and import the data.
```
make run-db
```
4. Run tests
```
make test-dev
```
5. Run application
```
make run-dev
```
6. Build application
```
make build-dev
```

## Notes
* The CSV file needs a comma (,) in the last line and in the last charactere. This was done in the file provided here.
* The CSV has some NULL values that must be managed to avoid errors while unmarshaling the resultset into go structure. There are two possible solutions:
    1. While import the file with copy command, fill empty strings in NULL values
    2. Mange the null in code as it was done here. It was decided to do this to avoid touch in dataprovied 
* When creating a new Feira, it was decided to require the fields: NomeFeira, Logradouro, Regiao5, Bairro, Distrito.