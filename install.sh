#!/bin/sh

# Install GOLang.

# Install MongoDB.

# Run MongoDB.
sudo systemctl start mongod

# Download GOLang MongoDB driver.
go get go.mongodb.org/mongo-driver

# Create MongoDB database and collection.

# Clone MeadStorm repository.
git clone https://github.com/pr1malbyt3s/MeadStorm

# Run MeadStorm
./MeadStorm
