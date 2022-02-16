

CREATE TABLE fleet(
	id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    max_speed DOUBLE,
    
    PRIMARY KEY(id)
)
ENGINE = InnoDB;

CREATE TABLE fleet_alert(
	id INT NOT NULL AUTO_INCREMENT,
    fleet_id INT,
    webhook VARCHAR(150),
    PRIMARY KEY(id),
    
    constraint fk_fleet 
		FOREIGN KEY (fleet_id)
		REFERENCES fleet (id)
)
ENGINE = InnoDB;

CREATE TABLE vehicle(
	id INT NOT NULL AUTO_INCREMENT,
    fleet_id INT,
    name VARCHAR(150),
    max_speed DOUBLE,
    PRIMARY KEY(id),
    
    constraint fk_fleet_vehicle 
		FOREIGN KEY (fleet_id)
		REFERENCES fleet (id)
)
ENGINE = InnoDB;

CREATE TABLE vehicle_position(
	id INT NOT NULL AUTO_INCREMENT,
    vehicle_id INT,
    times_stamp DATETIME,
    latitude DOUBLE,
    longitude DOUBLE,
    current_speed DOUBLE,
    max_speed DOUBLE,
    PRIMARY KEY(id),
    
    constraint fk_fleet_vehi
		FOREIGN KEY (vehicle_id)
		REFERENCES vehicle (id)
)
ENGINE = InnoDB;

