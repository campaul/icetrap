#!/usr/bin/env sh
psql -d $DATABASE_URL -a -f $1
