#!/bin/sh

rm db.created

mongo test "$(dirname "$0")"/createIndexes.js
mongo test "$(dirname "$0")"/loadNodes.js

touch db.created
