CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(30) NOT NULL,
    descricao TEXT NOT NULL,
    preco NUMERIC NOT NULL,
    data_venda DATE  NOT NULL
);