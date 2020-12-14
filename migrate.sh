#!/bin/sh
/usr/local/bin/migrate -source file://./migrations/ -database $MYSQLDSN up
