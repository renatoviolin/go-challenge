# Challenge in Golang

This repository implements a challenge that asks to implement an API in Golang with the following requirements.
* import dataset from csv file;
* create a new Feira object;
* remove a Feira object given an id;
* update a Feira object, except the id;
* find Feira objects by one of the following params:
    1. distrito
    2. regiao
    3. nome_feira
    4. bairro

The full description is available [here].(assets/challenge.pdf)

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

**FeiraPayload** has the following fields:
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
git clone https://github.com/renatoviolin/go-challenge.git
```
2. Review the .env file to adjust the credentials of your choice.
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_HOST=host.docker.internal
```
3. Start the database, create table and import the csv file. **It is necessary to start the database before the application so that the tests can be performed.**
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
git clone https://github.com/renatoviolin/go-challenge.git
```
2. Review the .env file. Make sure to adjust **POSTGRES_HOST** to localhost.
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_HOST=locahost
```
3. Start the database, create table and import the csv file.
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
* All logs output are saved in the file **app.log**.
* The CSV file needs a comma (,) in the last line and last charactere. This was already done in the file provided [here](db/data.csv).
* The CSV has some null values that must be managed to avoid errors while unmarshaling the resultset from postgres into Go structs. There are two possible solutions:
    1. While import the file with copy command, fill in null strings.
    2. Manage the null values in code (as it was done here). It was decided to do this to avoid touch in dataprovied 
* When creating a new Feira, it was decided to require the fields: NomeFeira, Logradouro, Regiao5, Bairro, Distrito.