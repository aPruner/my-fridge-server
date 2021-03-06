#!/bin/bash
cd app/db/migrations || exit
env_vars=$(grep -E -v '^#' ../../../.env | xargs)
for env_var in $env_vars
do
  var_key="${env_var%=*}"
  var_val="${env_var#*=}"
  if [ "$var_key" = "DB_USER" ]
  then
    user_val=$var_val
  elif [ "$var_key" = "DB_PASSWORD" ]
  then
    password_val=$var_val
  elif [ "$var_key" == "DB_NAME" ]
  then
    dbname_val=$var_val
  fi
done

conn_string="user=${user_val} password=${password_val} dbname=${dbname_val} sslmode=disable"

if [ "$1" = "up" ]
then
  goose postgres "${conn_string}" up
elif [ "$1" = "down" ]
then
  goose postgres "${conn_string}" down
elif [ "$1" = "reset" ]
then
  goose postgres "${conn_string}" reset
else
  echo "Please provide an argument to the migrate script, like so:"
  echo "  bash migrate.sh up"
  echo "or"
  echo "  bash migrate.sh down"
  echo "or"
  echo "  bash migrate.sh reset"
fi
cd ../.. || exit


