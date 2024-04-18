#!/bin/sh
#
# DROP "nodes" COLLECTION IN "test" DB
echo "DROPPING ALL DATA IN LOCALHOST'S NODES AND GPS_LOG COLLECTIONS IN THE TEST DB"
echo "db.nodes.drop()" | mongosh test
echo "db.gps_log.drop()" | mongosh test

rm db.created

mongosh test "$(dirname "$0")"/createIndexes.js || { echo aborting on error ; exit 1 ; }
mongosh test "$(dirname "$0")"/loadData.js || { echo aborting on error ; exit 1 ; }

touch db.created
