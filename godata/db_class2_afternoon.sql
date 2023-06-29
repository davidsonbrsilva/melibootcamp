# Explique o conceito de normalização e por que ele é usado.
A normalização é um processo no qual um banco de dados relacional é projetado de forma a reduzir a redundância e evitar anomalias de atualização, inserção e exclusão. O objetivo principal da normalização é organizar os dados de forma eficiente, minimizando a duplicação de informações e mantendo a integridade dos dados.

# Adicione um filme à tabela de filmes.
INSERT INTO movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id)
VALUES (CURDATE(), CURDATE(), 'Avatar: o caminho da água', 6.0, 1, '2023-04-04', 120, 4)

# Adicione um gênero à tabela de gêneros.
INSERT INTO genres (created_at, updated_at, name, ranking, active)
VALUES (CURDATE(), CURDATE(), 'Stand up', 13, 1)

# Associe o filme do Ex 2. ao gênero criado no Ex 3.
UPDATE movies m SET m.genre_id = 15 WHERE m.id = 22

# Modifique a tabela de atores para que pelo menos um ator tenha o filme adicionado no Ex.2 como favorito.
UPDATE actors a SET favorite_movie_id = 22 WHERE a.id = 1

# Crie uma cópia de tabela temporária da tabela de filmes.
CREATE TEMPORARY TABLE temp_movies AS SELECT * FROM movies

# Elimine dessa tabela temporária todos os filmes que ganharam menos de 5 prêmios.
DELETE FROM temp_movies m WHERE m.awards < 5

# Obtenha a lista de todos os gêneros que possuem pelo menos um filme.
SELECT DISTINCT g.name
FROM movies m
INNER JOIN genres g
ON g.id = m.genre_id

SELECT g.name
FROM genres g
INNER JOIN movies m
ON m.genre_id = g.id
GROUP BY g.name

# Obtenha a lista de atores cujo filme favorito ganhou mais de 3 prêmios.
SELECT CONCAT(a.first_name, ' ', a.last_name) "Ator"
FROM actors a
INNER JOIN movies m
ON a.favorite_movie_id = m.id
WHERE m.awards > 3

# Use o plano de explicação para analisar as consultas em Ex.6 e 7.
# EXPLAIN
# EXPLAIN FORMAT=TREE
EXPLAIN ANALYZE
SELECT DISTINCT g.name
FROM movies m
INNER JOIN genres g
ON g.id = m.genre_id

# EXPLAIN
# EXPLAIN FORMAT=TREE
EXPLAIN ANALYZE
SELECT g.name
FROM genres g
INNER JOIN movies m
ON m.genre_id = g.id
GROUP BY g.name

Analizando as duas consultas acima, é possível observar que a primeira é mais otimizada, dado que o número de linhas avaliadas na consulta da tabela gêneros é expressivamente menor do que no segundo caso (13 contra 22). Dessa forma, obtemos um custo aproximado entre 7.61...10.1 contra 12.5...15.1 para o segundo.

# EXPLAIN
# EXPLAIN FORMAT=TREE
EXPLAIN ANALYZE
SELECT CONCAT(a.first_name, ' ', a.last_name) "Ator"
FROM actors a
INNER JOIN movies m
ON a.favorite_movie_id = m.id
WHERE m.awards > 3

# O que são índices? Para que servem?
Em bancos de dados, os índices são estruturas que ajudam a melhorar a eficiência e o desempenho das consultas, permitindo a recuperação rápida de dados com base em determinados critérios de pesquisa. Eles são usados para acelerar a localização e o acesso aos registros relevantes em uma tabela do banco de dados.
Os índices são especialmente úteis em tabelas grandes, onde a busca linear em todos os registros seria ineficiente e demorada. Com um índice apropriado, o banco de dados pode realizar uma busca mais rápida usando a estrutura de índice, reduzindo assim o tempo necessário para recuperar os dados desejados.

# Crie um índice no nome na tabela de filmes.
CREATE INDEX movies_title_index
ON movies(title)

# Verifique se o índice foi criado corretamente.
SHOW INDEX FROM movies
