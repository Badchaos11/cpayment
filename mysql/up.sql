DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    userid INTEGER NOT NULL,
    email VARCHAR(50) NOT NULL,
    price DOUBLE NOT NULL,
    currency VARCHAR(10) NOT NULL,
    createdon DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedon DATETIME NOT NULL DEFAULT CURRENT_TIMEST ON UPDATE CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL
);