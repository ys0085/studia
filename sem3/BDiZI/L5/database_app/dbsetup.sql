CREATE OR REPLACE DATABASE app_db;
USE app_db;
CREATE OR REPLACE TABLE Clients (
    id int PRIMARY KEY AUTO_INCREMENT,
    first_name varchar(50),
    last_name varchar(50),
    email varchar(50)
);
CREATE OR REPLACE TABLE Products (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(50),
    price decimal(10,2),
    stock int
);
CREATE OR REPLACE TABLE Orders (
    id int PRIMARY KEY AUTO_INCREMENT,
    client_id int,
    order_date DATETIME,
    product_id int,
    count int, 
    price decimal(10,2),
    status ENUM("UNPAID", "PAID", "FINALIZED") NOT NULL,

    FOREIGN KEY (client_id) REFERENCES Clients(id),
    FOREIGN KEY (product_id) REFERENCES Products(id)
);
CREATE OR REPLACE TABLE Payments (
    id int PRIMARY KEY AUTO_INCREMENT,
    order_id int,
    amount decimal(10,2),
    payment_type ENUM("Cash", "Debit", "Credit"),
    payment_date DATETIME,

    FOREIGN KEY (order_id) REFERENCES Orders(id)
);
CREATE OR REPLACE TABLE Reports (
    id int PRIMARY KEY AUTO_INCREMENT,
    report_date DATETIME,
    order_count int,
    total_income decimal(11,2)
);

CREATE OR REPLACE VIEW Orders_Full AS 
SELECT o.id AS order_id, o.order_date, o.status, CONCAT(c.first_name, " ", c.last_name) as client, p.name, o.count, o.price 
FROM Orders o
JOIN Products p ON o.product_id = p.id
JOIN Clients c ON o.client_id = c.id;



-- procedura Add order

DELIMITER $$

CREATE OR REPLACE PROCEDURE AddOrder(
    IN p_client_id INT,
    IN p_product_id INT,
    IN p_count INT
)
BEGIN
    DECLARE v_price DECIMAL(10,2);
    DECLARE v_stock INT;

    SELECT stock, price INTO v_stock, v_price
    FROM Products
    WHERE id = p_product_id;

    IF v_stock < p_count THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Niewystarczająca ilość produktu w magazynie.';
    ELSE

        INSERT INTO Orders (client_id, product_id, count, order_date, status, price)
        VALUES (p_client_id, p_product_id, p_count, current_timestamp(), 'UNPAID', v_price * v_stock);

    
        UPDATE Products
        SET stock = stock - p_count
        WHERE id = p_product_id;
    END IF;
END $$

DELIMITER ;

-- procedura pay order

DELIMITER $$

CREATE OR REPLACE PROCEDURE PayOrder(
    IN p_order_id INT,
    IN p_payment_type ENUM('Cash', 'Debit', 'Credit')
)
BEGIN
    DECLARE v_total_amount DECIMAL(10,2);
    DECLARE v_count INT;
    DECLARE v_price DECIMAL(10,2);

    SELECT o.count, p.price INTO v_count, v_price
    FROM Orders o
    JOIN Products p ON o.product_id = p.id
    WHERE o.id = p_order_id;

    SET v_total_amount = v_count * v_price;

    INSERT INTO Payments (order_id, amount, payment_type, payment_date)
    VALUES (p_order_id, v_total_amount, p_payment_type, current_timestamp());

    UPDATE Orders
    SET status = 'PAID'
    WHERE id = p_order_id;
END $$

DELIMITER ;

