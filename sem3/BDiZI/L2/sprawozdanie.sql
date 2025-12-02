-- UPDATED VER

1.
CREATE DATABASE aparaty;

CREATE OR REPLACE USER '279728'@'localhost' IDENTIFIED BY 'kleszcz28';

GRANT SELECT, INSERT, UPDATE ON aparaty.* TO '279728'@'localhost';

2.

CREATE TABLE Aparat (
    model varchar(30) PRIMARY KEY,
    producent int,
    matryca int,
    obiektyw int,
    waga float,
    typ ENUM('kompaktowy', 'lustrzanka', 'profesjonalny', 'inny'),
    FOREIGN KEY (producent) REFERENCES Producent(ID),
    FOREIGN KEY (matryca) REFERENCES Matryca(ID),
    FOREIGN KEY (obiektyw) REFERENCES Obiektyw(ID)
);

CREATE TABLE Matryca (
    ID int PRIMARY KEY AUTO_INCREMENT,
    przekatna decimal(4,2),
    rozdzielczosc decimal(3,1),
    typ varchar(10)
);

ALTER TABLE Matryca AUTO_INCREMENT=100;

CREATE TABLE Obiektyw (
    ID int PRIMARY KEY AUTO_INCREMENT,
    model varchar(30),
    minPrzeslona float,
    maxPrzeslona float,
    CHECK (minPrzeslona < maxPrzeslona)

);

CREATE TABLE Producent (
    ID int PRIMARY KEY AUTO_INCREMENT,
    nazwa varchar(50) NOT NULL,
    kraj varchar(50) DEFAULT 'nieznany',
    adresKorespondencyjny varchar(100) DEFAULT 'nieznany'
);

3.
INSERT INTO Producent (nazwa, kraj, adresKorespondencyjny)
VALUES
('Andrzej Aparat', 'Polska', 'Radom, ul. Mickiewicza 23'),
('Kodak', 'Niemcy', 'Berlin'),
('Aparaty CH', 'Szwajcaria', 'Zurych'),
('Kodak Dwa', 'Polska', 'Warszawa, ul. Pilsudzkiego 34'),
('Chinczyk 1', 'Chiny', 'Szanghaj, ul. Kubusia Puchatka 100'),
('Chinczyk 2', 'Chiny', 'Pekin, ul. Kubusia Puchatka 80'),
('Chinczyk 3', 'Chiny', 'Czengdu, ul. Kubusia Puchatka 60'),
('Chinczyk 4', 'Chiny', 'Szanghaj, ul. Kubusia Puchatka 40'),
('Chinczyk 5', 'Chiny', 'Pekin, ul. Kubusia Puchatka 20'),
('Chinczyk 6', 'Chiny', 'Pekin, ul. Kubusia Puchatka 1'),
('Kodak Trzy', 'Polska', 'Walbrzych, ul. Mala 7'),
('Slaby Aparat', 'Niemcy', 'Kolonia'),
('Aparaty Mariana', 'Polska', 'Nowy Sacz, ul. Wielka 30'),
('Nihon Kamera', 'Japonia', 'Tokio, ul. Hiro 50'),
('Aparaty Grzegorza', 'Polska', 'Olsztyn, ul. Pilsudzkiego 57'),
('Kamery Karoliny', 'Polska', 'Kolobrzeg, al. Jana Pawla II 79');

INSERT INTO Producent (kraj, adresKorespondencyjny) VALUES ('', '');

INSERT INTO Obiektyw (model, minPrzeslona, maxPrzeslona)
VALUES
('Obiektyw 1', 0.55, 6.7),
('Obiektyw 2', 0.77, 3.7),
('Obiektyw 3', 1.55, 12.7),
('Obiektyw 4', 0.997, 5.23),
('Obiektyw 5', 0.121, 6.88),
('Obiektyw 6', 0.12, 10.0),
('Obiektyw 7', 0.88, 11.1),
('Obiektyw 8', 0.121, 4.44),
('Obi 1', 0.123, 21.37),
('Obi 2', 0.23, 12.45),
('Obi 3', 0.45, 8.12),
('Obi 4', 0.66, 6.66),
('Obi 5', 1.22, 12.33),
('Obi 6', 1.21, 17.32),
('Obi 7', 0.1, 4.10);

INSERT INTO Obiektyw (model, minPrzeslona, maxPrzeslona) VALUES ('ASDA', 1, 0);

INSERT INTO Matryca (przeKatna, rozdzielczosc, typ)
VALUES
(19.11, 12.1, 'CCD'),
(13.44, 13.1, 'CCD'),
(27.56, 12.7, 'CCD'),
(12.11, 24.1, 'CCD'),
(19.10, 22.9, 'CCD'),
(19.11, 15.7, 'CCD'),
(33.11, 16.1, 'CCD'),
(19.61, 17.1, 'CCD'),
(45.17, 12.1, 'CCD'),
(19.21, 19.3, 'CCD'),
(21.37, 12.1, 'CMOS'),
(21.39, 13.6, 'CMOS'),
(23.99, 14.5, 'CMOS'),
(11.37, 52.1, 'CMOS'),
(21.27, 09.1, 'CMOS');

INSERT INTO Matryca (przekatna, rozdzielczosc, typ) VALUES (1.1, 0.122, 'asd');

4.
DELIMITER //
CREATE OR REPLACE PROCEDURE generateModels() BEGIN
    FOR i IN 1..100 DO
        SET @P = (SELECT ID FROM Producent WHERE ID = i % 17 + 1);
        SET @M = (SELECT ID FROM Matryca WHERE ID = i % 15 + 100);
        SET @O = (SELECT ID FROM Obiektyw WHERE ID = i % 15);
        INSERT INTO Aparat
        VALUES
        (CONCAT('pa', i, CURRENT_TIME()),
        @P,
        @M,
        @O,
        (SECOND(CURRENT_TIME()) / 10), 
        i % 4 + 1);
        
    END FOR;
END;
//
DELIMITER ;

CALL generateModels();

5.
-- Ta procedura działa, lecz z powodu budowy procedury z zadania 4. kazdy aparat danego producenta uzywa tej samej matrycy.
DELIMITER //
CREATE OR REPLACE PROCEDURE najmniejszaMatryca(producent_id INT, OUT model varchar(30))
BEGIN
    SELECT a.model INTO model
    FROM Aparat a
    INNER JOIN Matryca m ON m.ID = a.matryca
    WHERE a.producent = producent_id && m.przekatna = (
        SELECT ma.przekatna
        FROM Aparat ap
        INNER JOIN Matryca ma ON ma.ID = ap.matryca
        WHERE ap.producent = producent_id
        ORDER BY ma.przekatna
        LIMIT 1
    )
    LIMIT 1;
    

END;
//
DELIMITER ;

6.

DELIMITER //
CREATE OR REPLACE TRIGGER new_producent
BEFORE INSERT ON Aparat
FOR EACH ROW
BEGIN
    IF NOT EXISTS (SELECT * FROM Producent WHERE ID = NEW.producent) THEN
        INSERT INTO Producent(ID, nazwa) VALUES (NEW.producent, 'Nowy producent');
    END IF;
END
//
DELIMITER ;

-- check
INSERT INTO Aparat VALUES
('inny producent', 1234, (SELECT ID FROM Matryca WHERE ID = 100), (SELECT ID FROM Obiektyw WHERE ID = 1), 13.2, 'kompaktowy');

SELECT * FROM Aparat;
SELECT * FROM Producent;

7.

DELIMITER //
CREATE OR REPLACE FUNCTION liczMatryca(IN matryca_id INT) RETURNS INT
BEGIN
    RETURN (SELECT COUNT(*) FROM Aparat WHERE matryca = matryca_id);
END
//
DELIMITER ;

SELECT liczMatryca(100);

8.

DELIMITER //
CREATE OR REPLACE TRIGGER nieuzywanaMatryca
AFTER DELETE ON Aparat
FOR EACH ROW
BEGIN 
    IF NOT EXISTS (SELECT * FROM Aparat WHERE matryca = OLD.matryca) THEN
        DELETE FROM Matryca WHERE ID = OLD.matryca;
    END IF;
END;
//
DELIMITER ;

9.

CREATE VIEW lustrzanki (Model, Waga, Producent, Przekatna, Rozdzielczosc, MinPrzeslona, MaxPrzeslona) 
AS (
    SELECT a.model, a.waga, p.nazwa, m.przekatna, m.rozdzielczosc, o.minPrzeslona, o.maxPrzeslona 
    FROM Aparat a 
    INNER JOIN Producent p ON a.producent = p.ID
    INNER JOIN Matryca m ON a.matryca = m.ID
    INNER JOIN Obiektyw o ON a.obiektyw = o.ID
    WHERE p.kraj != 'Chiny' && a.typ = 'lustrzanka'
);
 
10.

CREATE VIEW model_producent (Producent, Kraj, Model)
AS (
    SELECT p.nazwa, p.kraj, a.model
    FROM Aparat a INNER JOIN Producent p ON a.producent = p.ID
);


DELETE FROM Aparat WHERE producent IN (SELECT ID FROM Producent WHERE kraj = 'Chiny');


11.

ALTER TABLE Producent
ADD liczbaModeli INT NOT NULL DEFAULT "0";

UPDATE Producent p
SET liczbaModeli = (
    SELECT COUNT(*) 
    FROM Aparat a
    WHERE a.producent = p.ID
);


DELIMITER //
CREATE OR REPLACE PROCEDURE liczModele()
BEGIN  
    UPDATE Producent p
    SET liczbaModeli = (
        SELECT COUNT(*) 
        FROM Aparat a
        WHERE a.producent = p.ID
    );
END;
//
DELIMITER ;


DELIMITER //
CREATE OR REPLACE TRIGGER licz1 
AFTER DELETE ON Aparat
FOR EACH ROW
BEGIN
    CALL liczModele;
END
//
CREATE OR REPLACE TRIGGER licz2 
AFTER UPDATE ON Aparat
FOR EACH ROW
BEGIN
    CALL liczModele;
END
//
CREATE OR REPLACE TRIGGER licz3
AFTER INSERT ON Aparat
FOR EACH ROW
BEGIN
    CALL liczModele;
END
//
DELIMITER ;







