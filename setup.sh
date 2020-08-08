go get

env_vars=$(grep -E -v '^#' .env | xargs)
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
echo "CREATE DATABASE ${dbname_val}; CREATE USER ${user_val} WITH ENCRYPTED PASSWORD '${password_val}'; GRANT ALL PRIVILEGES ON DATABASE ${dbname_val} to ${user_val}; ALTER USER ${user_val} WITH SUPERUSER;" | sudo -u postgres psql postgres