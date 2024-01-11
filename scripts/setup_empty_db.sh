#!/bin/sh
#
# DROP "nodes" COLLECTION IN "test" DB
echo "db.nodes.drop()" | mongo test

mongo test "$(dirname "$0")"/createIndexes.js
mongo test "$(dirname "$0")"/loadNodes.js
