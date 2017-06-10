package dbutils

const train = `
	CREATE TABLE IF NOT EXISTS train (
           ID INT PRIMARY KEY NOT NULL,
           DRIVER_NAME VARCHAR(64) NULL,
           OPERATING_STATUS BOOLEAN
        )
`

const station = `
	CREATE TABLE IF NOT EXISTS station (
          ID INT PRIMARY KEY NOT NULL,
          NAME VARCHAR(64) NULL,
          OPENING_TIME TIME NULL,
          CLOSING_TIME TIME NULL
        )
`
const schedule = `
	CREATE TABLE IF NOT EXISTS route (
	  ID INT PRIMARY KEY NOT NULL,
          TRAIN_ID INT,
          STATION_ID INT,
          ARRIVAL_TIME TIME
        )
`
