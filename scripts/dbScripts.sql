CREATE DATABASE `go_contacts` /*!40100 DEFAULT CHARACTER SET latin1 */;

use `go_contacts`;

DROP TABLE IF EXISTS go_contacts;
CREATE TABLE go_contacts (
    id INT AUTO_INCREMENT NOT NULL,
    primer_nombre VARCHAR(128) NOT NULL,
    segundo_nombre VARCHAR(255) NOT NULL,
    primer_apellido VARCHAR(255) NOT NULL,
    segundo_apelido VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    birthday VARCHAR(255) NOT NULL,
    telefono VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);


