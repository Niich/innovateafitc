CREATE TABLE device_type (
    id bigserial PRIMARY KEY,
    device_name varchar NOT NULL
);

CREATE TABLE tracked_device (
    id bigserial PRIMARY KEY,
    device_type_id bigint NOT NULL,
    device_location varchar,
    CONSTRAINT fk_device_type FOREIGN KEY (device_type_id) REFERENCES device_type(id)
);

CREATE TABLE sensor (
    id varchar PRIMARY KEY,
    tracked_device_id bigint,
    bat_status int,
    CONSTRAINT fk_tracked_device FOREIGN KEY (tracked_device_id) REFERENCES tracked_device(id)
);

CREATE TABLE data_log (
    id bigserial PRIMARY KEY,
    tracked_device_id bigint NOT NULL,
    log_time timestamp DEFAULT NOW(),
    sensor_value int,
    CONSTRAINT fk_tracked_device FOREIGN KEY (tracked_device_id) REFERENCES tracked_device(id)
);