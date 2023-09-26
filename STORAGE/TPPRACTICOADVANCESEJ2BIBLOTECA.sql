##Se tiene el siguiente DER que corresponde al esquema que presenta la base de datos de una “biblioteca”.

##DDL
DROP DATABASE IF EXISTS bibloteca;

CREATE DATABASE IF NOT EXISTS bibloteca;

USE bibloteca;

CREATE TABLE IF NOT EXISTS libro(
`id_libro` VARCHAR (45) NOT NULL,
`titulo` VARCHAR (45) NOT NULL,
`editorial`  VARCHAR (45) NOT NULL,
`area` VARCHAR (45) NOT NULL,
PRIMARY KEY(`id_libro`)
);

CREATE TABLE IF NOT EXISTS libroAutor(
`id_libro` VARCHAR (45) NOT NULL,
`id_autor` VARCHAR (45) NOT NULL,

KEY `idx_libroAutorid_libro` (`id_libro`),
CONSTRAINT `fk_employee_id_libro` FOREIGN KEY(`id_libro`) REFERENCES `libro`(`id_libro`),
KEY `idx_libroAutorid_autor` (`id_autor`),
CONSTRAINT `fk_employee_id_autor` FOREIGN KEY(`id_autor`) REFERENCES `autor`(`id_autor`)

);

CREATE TABLE IF NOT EXISTS autor(
`id_autor` VARCHAR (45) NOT NULL,
`nombre` VARCHAR (45) NOT NULL,
`nacionalidad`  VARCHAR (45) NOT NULL,
PRIMARY KEY(`id_autor`)
);

CREATE TABLE IF NOT EXISTS estudiante(
`id_lector` VARCHAR (45) NOT NULL,
`nombre` VARCHAR (45) NOT NULL,
`apellido`  VARCHAR (45) NOT NULL,
`direccion` VARCHAR (45) NOT NULL,
`carrera` VARCHAR (45) NOT NULL,
`edad` INT NULL,
PRIMARY KEY(`id_lector`)
);

CREATE TABLE IF NOT EXISTS Prestamo(
`id_lector` VARCHAR (45) NOT NULL,
`id_libro` VARCHAR (45) NOT NULL,
`fecha_prestamo` date NULL,
`fecha_devolucion` date NULL,
`devuelto` bool NULL,
KEY `idx_prestamos_libro` (`id_libro`),
CONSTRAINT `fk_prestamos_id_libro` FOREIGN KEY(`id_libro`) REFERENCES `libro`(`id_libro`),
KEY `idx_prestamosid_lector` (`id_lector`),
CONSTRAINT `fk_prestamos_id_lector` FOREIGN KEY(`id_lector`) REFERENCES `estudiante`(`id_lector`)

);




