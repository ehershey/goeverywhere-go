#!/bin/sh

mongo test "$(dirname "$0")"/createIndexes.js
mongo test "$(dirname "$0")"/loadNodes.js
