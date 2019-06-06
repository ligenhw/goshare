#!/bin/bash

cp config.json goshare/configration/
cd goshare/contrib/compose/
docker-compose build
docker-compose up -d

