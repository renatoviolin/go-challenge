CREATE SEQUENCE feira_id_seq;

CREATE TABLE feira (
    id bigint NOT NULL DEFAULT nextval('feira_id_seq'),
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

\COPY feira (ID,LONG,LAT,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREFE,REGIAO5,REGIAO8,NOME_FEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA) from '/tmp/data.csv' WITH  (FORMAT CSV, NULL 'NULL', HEADER);
SELECT setval('feira_id_seq', max(id)) from feira;

CREATE INDEX ix_regiao5 ON feira (lower(regiao5) varchar_pattern_ops);
CREATE INDEX ix_nome_feira ON feira (lower(nome_feira) varchar_pattern_ops);
CREATE INDEX ix_distrito ON feira (lower(distrito) varchar_pattern_ops);
CREATE INDEX ix_bairro ON feira (lower(bairro) varchar_pattern_ops);