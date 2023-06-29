# Retorne a média de avaliação(rating) por gênero do filme
SELECT g.name "Gênero", ROUND(AVG(m.rating), 2) "Avaliação"
FROM movies m
INNER JOIN genres g
ON m.genre_id = g.id
GROUP BY g.name

# Retorne o nome dos filmes que um ator participou, além do nome do ator
SELECT CONCAT(a.first_name, ' ', a.last_name) "Ator", m.title "Filme"
FROM actors a
INNER JOIN actor_movie am
ON a.id = am.actor_id
INNER JOIN movies m
ON M.id = am.movie_id

# Retorne a quantidade de atores que cada filme teve
SELECT m.title "Filme", COUNT(*) "Quantidade de Atores"
FROM actors a
INNER JOIN actor_movie am
ON a.id = am.actor_id
INNER JOIN movies m
ON M.id = am.movie_id
GROUP BY m.title, m.rating

# Retorne a quantidade de atores que cada filme teve, além do nome do filme, dos filmes que tem um rating maior ou igual a "3.0"
SELECT m.title "Filme", m.rating "Avaliação", COUNT(*) "Quantidade de Atores"
FROM actors a
INNER JOIN actor_movie am
ON a.id = am.actor_id
INNER JOIN movies m
ON M.id = am.movie_id
GROUP BY m.title, m.rating
HAVING m.rating >= 3.0
