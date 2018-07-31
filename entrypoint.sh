#!/bin/bash
echo "Waiting for mysql"
until mysql -hstatserver-db -P3306 -uroot -proot2507 &> /dev/null
do
  printf "."
  sleep 1
done

echo -e "\nmysql ready"

/statserver