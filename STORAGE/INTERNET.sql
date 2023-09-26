
INSERT INTO mydb.Clientes VALUES 
(36005288,"PEDRO","GOMEZ",'1991-06-30',"CORRIENTES","CAPITAL"),
(35005289,"JOSE","CACEREZ",'1991-05-30',"CHACO","RESISTENCIA"),
(35005281,"RAUL","GUTIERREZ",'1991-05-30',"FORMOSA","CAPITAL"),
(35005381,"MARIA","DUARTE",'1944-05-30',"MISIONES","POSADAS"),
(32005381,"LAURA","PEDROZO",'1988-05-30',"MISIONES","ELDORADO"),
(17004381,"MIGUEL","GIMENEZ",'1958-05-30',"MISIONES","BERNARDO DE IRIGOYEN"),
(17005381,"RAUL","CARDOZO",'1958-05-30',"MISIONES","BERNARDO DE IRIGOYEN"),
(31005287,"GABRIEL","SOSA",'1991-07-30',"MISIONES","OBERA"),
(32005287,"ESTEBAN","OBREGON",'1987-07-30',"SANTA FE","ROSARIO"),
(32005687,"LEONEL","MESSI",'1987-07-30',"SANTA FE","ROSARIO"),
(32005987,"RICARDO","PEDROZO",'1981-07-30',"CHACO","RESISTENCIA");
;

INSERT INTO mydb.Internet VALUES (1,10.0,5000.0,20.0),
(2,20.0,4000.0,30.0),
(3,30.0,4500.0,40.0),
(4,40.0,6000.0,45.0),
(5,50.0,7000.0,45.0);

INSERT INTO mydb.Clientes_has_Internet VALUES
(32005987,1),
(32005687,2),
(32005287,3),
(17004381,4);

##MOSTRAS LAS PROVINCIAS DE LOS CLIENTES

SELECT Provincia FROM mydb.Clientes;


##MOSTRAS LAS ciudades ordenada desc DE LOS CLIENTES
SELECT CIudad as city FROM mydb.Clientes ORDER BY city DESC;
##mostrar las provincias que tengan MIS 

SELECT Provincia FROM mydb.Clientes WHERE Provincia LIKE "MIS%";

##mostrar los planes con mas de 30 % de descuento

SELECT * FROM mydb.Internet WHERE Descuento>30.0;
##mostrar los planes con mas de 30 % de descuento y menor a 5000 pesos
SELECT * FROM mydb.Internet WHERE Descuento>30.0 AND Precio<5000.0;
##mostrar todos los descuentos
SELECT Descuento FROM mydb.Internet;
##mostrar de forma desc todos los descuentos y precios 
SELECT Descuento, Precio,Velocidad FROM mydb.Internet ORDER BY Descuento DESC;
##mostrar todos los dni ordenados
SELECT DNI as city FROM mydb.Clientes ORDER BY DNI DESC;
##mostrar todods los nombres
SELECT Nombre as city FROM mydb.Clientes ORDER BY DNI DESC;









