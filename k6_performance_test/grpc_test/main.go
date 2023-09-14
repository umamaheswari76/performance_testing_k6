package main

import "go.mongodb.org/mongo-driver/mongo"

const URI = "mongodb://localhost:27017/?retryWrites=true&connectTimeoutMS=10000"

var mongodb *mongo.Collection
