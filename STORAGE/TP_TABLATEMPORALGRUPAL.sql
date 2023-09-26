USE movies_db;

##Agregar una película a la tabla movies.

INSERT INTO movies VALUES (22,null,null,"Megalodon",3.1,1,'2022-05-04 00:00:00',100,5);
##Agregar un género a la tabla genres.

INSERT INTO genres VALUES(13,'2014-07-04 00:00:00',null,"Anime",13,1);

##Asociar a la película del punto 1. genre el género creado en el punto 2.
UPDATE movies
SET genre_id=13
WHERE id= 22;
##Modificar la tabla actors para que al menos un actor
## tenga como favorita la película agregada en el punto 1.
##Vamos por leonardo dicaprio
UPDATE actors
SET favorite_movie_id=22
WHERE id=4;

##Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY  TABLE  moviestw (
SELECT * FROM movies
);

SELECT * FROM moviestw;

##Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
 
DELETE FROM `moviestw`  WHERE `awards` <5;
 
SET SQL_SAFE_UPDATES = 1;

##Obtener la lista de todos los géneros que tengan al menos una película.
##PODEMOS RESOLVER DE DOS MANERAS CON INNER JOIN Y SUB CONSULTAS VEAMOS LA DIFERENCIAS
SELECT g.name as name FROM movies m
INNER JOIN genres g  ON g.id = m.genre_id
GROUP BY name
HAVING COUNT(m.genre_id)>=1;

##Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
## PODEMOS RESOLVER DE DOS MANERAS CON INNER JOIN Y SELECT PERO VAMOS A VER CUAL ES MAS EFICIENTES

SELECT a.first_name as nombre, a.last_name as apellido, m.awards as premios  FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
GROUP BY  nombre,apellido,premios
HAVING premios> 3;

SHOW STATUS LIKE 'Last_query_cost';
##19.59 QUERY COST

SELECT a.first_name as nombre, a.last_name as apellido  FROM actors a 
WHERE a.favorite_movie_id IN (
	SELECT m.id FROM movies m WHERE awards > 3

);
SHOW STATUS LIKE 'Last_query_cost';
## QUERY COST 8.16


##Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX awards_index ON movies (awards);


##Chequee que el índice fue creado correctamente.

SHOW INDEXES FROM movies;
##SHOW STATUS LIKE 'Last_query_cost';
DROP INDEX awards_index ON movies ;
##VAMOS A EVALUAR LOS TIEMPOS DE LECTURAS CON LOS INDICES

SELECT m.awards  FROM movies m WHERE m.awards = 5;
SHOW STATUS LIKE 'Last_query_cost';
##2.44 cost query
CREATE INDEX awards_index ON movies (awards);
SELECT m.awards  FROM movies m WHERE m.awards = 5;
SHOW STATUS LIKE 'Last_query_cost';
##0.44 cost query 

##se observa la mejora notable en el costo de query sin indice cuesta 2.44 y con indice 0.44
## casi 5 veces mejores

##¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
##crearia en una tabla de poca escritura y mucha lectura por ejemplo 'genres' seria una buena opcion


