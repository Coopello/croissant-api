#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s
DIR=$(cd $(dirname $0); pwd)

CMD_MYSQL="mysql -u go_test -pdocker go_database -ppassword"

#run the setup script to create the DB and the schema in the DB
$CMD_MYSQL < $DIR/0001_create_plan.sql
$CMD_MYSQL < $DIR/0002_create_user.sql
$CMD_MYSQL < $DIR/0004_create_plan_participant_users.sql
$CMD_MYSQL < $DIR/0005_create_shop.sql
