##DDL


DROP DATABASE IF EXISTS empresa_db;

CREATE DATABASE IF NOT EXISTS empresa_db;

USE empresa_db;

CREATE TABLE IF NOT EXISTS departamentos(
`dept_nro` VARCHAR(7) NOT NULL,
`nombre_depto` VARCHAR (50) NOT NULL,
`localidad` VARCHAR (50) NULL,
PRIMARY KEY (`dept_nro`)
);


CREATE TABLE IF NOT EXISTS empleados(

`cod_emp` VARCHAR(7) NOT NULL,
`nombre` VARCHAR (50) NOT NULL,
`apellido` VARCHAR (50) NOT NULL,
`puesto` VARCHAR (50) NOT NULL,
`fecha_alta` date NULL,
`salario`  INT NULL,
`comision` INT DEFAULT 0 ,
`depto_nro` VARCHAR(7)  NULL,
PRIMARY KEY(`cod_emp`),
KEY `idx_employee_puesto` (`puesto`),
KEY `idx_employee_depto_nro` (`depto_nro`),
CONSTRAINT `fk_employee_depto_nro` FOREIGN KEY(`depto_nro`) REFERENCES `departamentos`(`dept_nro`)

);

USE empresa_db;
##Se requiere obtener las siguientes consultas:
##Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad FROM empleados e INNER JOIN departamentos d ON e.depto_nro = d.dept_nro;
##Visualizar los departamentos con más de cinco empleados.

SELECT d.nombre_depto as departamento ,COUNT(e.cod_emp) as cantidad FROM departamentos d INNER JOIN empleados e ON e.depto_nro = d.dept_nro 
GROUP BY departamento HAVING cantidad>=2;


##Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.

SELECT e.nombre, e.salario, d.nombre_depto FROM empresa_db.empleados e INNER JOIN empresa_db.departamentos d ON e.depto_nro = d.dept_nro
WHERE e.puesto= (
SELECT e2.puesto FROM empresa_db.empleados e2 WHERE e2.nombre='Mito' AND e2.apellido= 'Barchuk'
);

##Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.

SELECT e.* FROM  empresa_db.empleados e INNER JOIN empresa_db.departamentos d ON e.depto_nro = d.dept_nro 
WHERE d.nombre_depto= "Contabilidad";


SHOW INDEX FROM empresa_db.empleados;
##Mostrar el nombre del empleado que tiene el salario más bajo.

SELECT e.nombre FROM  empresa_db.empleados e INNER JOIN empresa_db.departamentos d ON e.depto_nro = d.dept_nro 
WHERE e.salario= (
SELECT (SELECT MIN(e2.salario) FROM empresa_db.empleados e2)
);
##Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.

SELECT e .* FROM  empresa_db.empleados e INNER JOIN empresa_db.departamentos d ON e.depto_nro = d.dept_nro 
WHERE e.salario= (
SELECT MAX(e2.salario) FROM empresa_db.empleados e2  INNER JOIN empresa_db.departamentos d2 ON e2.depto_nro = d2.dept_nro 
WHERE d2.nombre_depto="Ventas"


);










##INSERTAMOS LOS VALORES PROPUESTOS
INSERT INTO empresa_db.departamentos VALUES
('D-000-1','Software','Los Tigres'),
('D-000-2','Sistemas','Guadalupe'),
('D-000-3','Contabilidad','La Roca'),
('D-000-4','Ventas','Plata');



INSERT INTO empresa_db.empleados VALUES
('E-0001','César','Piñero','Vendedor','2018-05-12',80000,15000,'D-000-4'),
('E-0002','Yosep','Kowaleski','Analista','2015-07-14',140000,0,'D-000-2'),
('E-0003','Mariela','Barrios','Director','2014-06-05',185000,0,'D-000-3'),
('E-0004','Jonathan','Aguilera','Vendedor','2015-06-03',85000,10000,'D-000-4'),
('E-0005','Daniel','Brezezicki','Vendedor','2018-03-03',83000,10000,'D-000-4'),
('E-0006','Mito','Barchuk','Presidente','2014-06-05',190000,0,'D-000-3'),
('E-0007','Emilio','Galarza','Desarrollador','2014-08-02',60000,0,'D-000-1');







