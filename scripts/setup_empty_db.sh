#!/bin/sh
#
# DROP "nodes" COLLECTION IN "test" DB
echo "db.nodes.drop()" | mongo test

mongo test "$(dirname "$0")"/createIndexes.js || echo aborting on error ; exit 1
mongo test "$(dirname "$0")"/loadNodes.js || echo aborting on error ; exit 1
