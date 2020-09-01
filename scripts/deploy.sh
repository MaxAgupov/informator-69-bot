#!/bin/sh

BOT_HOME=/root/deploy
BOT_NAME=informator-69-bot

echo "Stopping the $BOT_NAME"
cd $BOT_HOME/$BOT_NAME &&
docker-compose down

echo "Unpacking the $BOT_NAME and starting it"
cd $BOT_HOME &&
mv $BOT_NAME $BOT_NAME-bak &&
tar -xzvf $BOT_NAME.tar.gz &&
cd $BOT_NAME &&
docker-compose build &&
docker-compose up -d &&
cd ../ && rm -fr $BOT_NAME-bak &&
echo "$BOT_NAME successfully installed and started"

