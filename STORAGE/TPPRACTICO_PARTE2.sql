##Mostrar el título y el nombre del género de todas las series.
##Tabla series y genres
USE movies_db;
##Recorda usar el alias!!!:)
select s.title, g.name  FROM series s INNER JOIN genres g ON s.genre_id = g.id;

##Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title as Episodio,a.first_name as Nombre ,a.last_name as Apellido FROM episodes e INNER JOIN actors a INNER JOIN actor_episode ae ON ae.episode_id=e.id AND ae.actor_id = a.id;

##Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
##Recorda usar el alias!!!

SELECT COUNT(*) as "Numero de Temporadas",s.title as serie FROM seasons t 
INNER JOIN series s ON t.serie_id=s.id
GROUP BY serie;


##Mostrar el nombre de todos los géneros y la cantidad total
## de películas por cada uno, 
##siempre que sea mayor o igual a 3.

SELECT g.name as genero , COUNT(*) as cantidad  FROM movies m  
INNER JOIN genres g ON
m.genre_id = g.id GROUP BY genero
HAVING cantidad >= 3 ;

##Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias
## y que estos no se repitan.
