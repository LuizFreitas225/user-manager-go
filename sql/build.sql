-- Active: 1755714182655@@localhost@15432@postgres@public
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,                  -- Auto incremento (equivalente ao GenerationType.IDENTITY)
    create_date TIMESTAMP  DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Data de criação
    last_modified_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Última modificação
    name VARCHAR(255) NOT NULL,                -- Nome do usuário
    email VARCHAR(255) UNIQUE NOT NULL,        -- E-mail único
    password VARCHAR(255) NOT NULL             -- Senha
);

-- Função que atualiza last_modified_date
CREATE OR REPLACE FUNCTION update_last_modified()
RETURNS TRIGGER AS $$
BEGIN
   NEW.last_modified_date = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para chamar a função  automaticamente
CREATE TRIGGER trg_update_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_last_modified();


---Insert Inicial
INSERT INTO users (name, email, password)
VALUES ('Luiz Nicolau', 'luiz.nicolau@email.com', 'senha');


