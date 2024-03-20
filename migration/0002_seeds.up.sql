-- DELETE CURRENT DATA
DELETE FROM users;
DELETE FROM time_clocks;

TRUNCATE users, time_clocks RESTART IDENTITY;

INSERT INTO users ("name", email, registration, password, created_at)
VALUES ('João Silva', 'joao@example.com', '123456789', 'senha123', '2024-03-18T10:00:00Z'),
       ('Maria Santos', 'maria@example.com', '987654321', 'senha456', '2024-03-18T11:00:00Z');

-- User 1
INSERT INTO time_clocks (user_id, clock_in, created_at)
VALUES (1, '2024-03-18T08:00:00Z', '2024-03-18T08:00:00Z'), -- Entrada
       (1, '2024-03-18T12:00:00Z', '2024-03-18T12:00:00Z'), -- Saída para o almoço
       (1, '2024-03-18T13:00:00Z', '2024-03-18T13:00:00Z'), -- Retorno do almoço
       (1, '2024-03-18T16:00:00Z', '2024-03-18T16:00:00Z');
-- Saída

-- User 2
INSERT INTO time_clocks (user_id, clock_in, created_at)
VALUES (2, '2024-03-18T08:30:00Z', '2024-03-18T08:30:00Z'), -- Entrada
       (2, '2024-03-18T12:30:00Z', '2024-03-18T12:30:00Z'), -- Saída para o almoço
       (2, '2024-03-18T13:30:00Z', '2024-03-18T13:30:00Z'), -- Retorno do almoço
       (2, '2024-03-18T17:30:00Z', '2024-03-18T17:30:00Z'); -- Saída

