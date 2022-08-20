CREATE TABLE deinfo (
    id SERIAL PRIMARY KEY,
    long bigint DEFAULT 0,
    lat bigint DEFAULT 0,
    setcens bigint DEFAULT 0,
    areap bigint DEFAULT 0,
    coddist text,
    distrito text,
    codsubpref integer DEFAULT 0,
    subprefe text,
    regiao5 text,
    regiao8 text,
    nome_feira text,
    registro text,
    logradouro text,
    numero text,
    bairro text,
    referencia text,
    ativo integer DEFAULT 1
);