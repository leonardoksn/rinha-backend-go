CREATE TABLE IF NOT EXISTS clientes (
  id INT NOT NULL AUTO_INCREMENT,
  limite INT NOT NULL,
  saldo INT DEFAULT 0,
  nome VARCHAR(50) NOT NULL,
  PRIMARY KEY (id) 
);

-- Criar tabela transacoes
CREATE TABLE IF NOT EXISTS transacoes (
  id INT NOT NULL AUTO_INCREMENT,
  valor INT NOT NULL,
  tipo VARCHAR(1) NOT NULL,
  realizada_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  descricao VARCHAR(10) NOT NULL,
  cliente_id INT NOT NULL,
  FOREIGN KEY (cliente_id) REFERENCES clientes(id),  
  PRIMARY KEY (id)  
);

-- Insert data into clientes table
INSERT INTO clientes (nome, limite)
VALUES
  ('o barato sai caro', 100000), 
  ('zan corp ltda', 80000),
  ('les cruders', 1000000),
  ('padaria joia de cocaia', 10000000),
  ('kid mais', 500000);