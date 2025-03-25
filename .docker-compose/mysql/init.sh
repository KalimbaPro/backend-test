#!/bin/bash
mysql -u root --password="$MYSQL_ROOT_PASSWORD" --execute="CREATE DATABASE IF NOT EXISTS core;
                                                            USE core;
                                                            CREATE TABLE IF NOT EXISTS breeds (
                                                                id INTEGER PRIMARY KEY AUTO_INCREMENT,
                                                                species VARCHAR(50) NOT NULL,
                                                                pet_size VARCHAR(50) NOT NULL,
                                                                name VARCHAR(80) NOT NULL,
                                                                average_male_adult_weight INTEGER NOT NULL,
                                                                average_female_adult_weight INTEGER NOT NULL
                                                            );"
