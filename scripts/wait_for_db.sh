#!/bin/bash

DB_HOST=$1
DB_PORT=$2

if [ -z "$DB_PORT" ]; then
  DB_PORT="5432"
fi

until nc -z -v -w30 $DB_HOST $DB_PORT
do
  echo "Esperando que la base de datos se inicie en el puerto $DB_PORT..."
  sleep 1
done

echo "La base de datos en el puerto $DB_PORT está lista"
