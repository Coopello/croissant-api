#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u go_test -pdocker go_database -p password < "/docker-entrypoint-initdb.d/0001_create_plan.sql"
