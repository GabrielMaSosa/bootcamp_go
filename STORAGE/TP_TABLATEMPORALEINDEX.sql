##Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD”
## y guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.

USE movies_db;

## my notable los cambios de rendimientos JOIN>>>>>Tablatemporal>> tabla_temporarl con index
CREATE TEMPORARY TABLE TWD
(
SELECT e.id as id_episodes, e.title as title_episodes, e.rating, s.id, s.title,s.number 
FROM episodes e INNER JOIN seasons s ON
e.season_id = s.id
);


SELECT e.id as id_episodes, e.title as title_episodes, e.rating, s.id, s.title,s.number 
FROM episodes e INNER JOIN seasons s ON
e.season_id = s.id WHERE e.rating>7;

SELECT e.id as id_episodes, e.title as title_episodes, e.rating, s.id, s.title,s.number 
FROM episodes e INNER JOIN seasons s ON
e.season_id = s.id;


SHOW STATUS LIKE 'Last_query_cost';

SELECT * FROM TWD WHERE;
## en ambos casos hace un full scan 
SHOW STATUS LIKE 'Last_query_cost';


CREATE INDEX rating_index ON TWD (rating);
SHOW INDEXES FROM WTD;

SELECT * FROM TWD WHERE rating=7;

##. 



CREATE TEMPORARY TABLE MYTWD(
SELECT e.title as capitulo ,s.title as temporada ,ser.title as serie FROM episodes e INNER JOIN seasons s 
ON e.season_id = s.id INNER JOIN series ser ON ser.id=s.serie_id);

select * from MYTWD;
##5 SEGUNDO TARDA 

SELECT e.title as capitulo ,s.title as temporada ,ser.title as serie FROM episodes e INNER JOIN seasons s 
ON e.season_id = s.id INNER JOIN series ser ON ser.id=s.serie_id;
##31 SEGUNDO TARDA 

CREATE INDEX serie_index ON MYTWD (serie);
##verifico el tiempo luego de crear el indice
select * from MYTWD WHERE serie='The Walking Dead'; 
##CON EL INDICE TARDA 1.34 SEGUNDOS

SELECT e.title as capitulo ,s.title as temporada ,ser.title as serie FROM episodes e INNER JOIN seasons s 
ON e.season_id = s.id INNER JOIN series ser ON ser.id=s.serie_id
WHERE ser.title='The Walking Dead';
##TARDA 6.0 SEGU EN BUSCAR EL RESULTADO 

##ES NOTORIO EL RENDIMIENTOS 

SHOW STATUS LIKE 'Last_query_cost';



##---------------------EJERCICIO NUMERO 2 ---------------------------------
##En la base de datos “movies”, seleccionar una tabla donde crear un índice y
##luego chequear la creación del mismo. 

##RESOLUCION:
##vamos a elegir actores para agregar un inidce en el titulo

CREATE INDEX name_index ON actors (first_name);







##Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.
SHOW INDEXES FROM actors;