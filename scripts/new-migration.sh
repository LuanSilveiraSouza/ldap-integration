#!/bin/bash

timestamp=$(date +%s)
touch ./database/migrations/"${timestamp}_$1.up.sql"
touch ./database/migrations/"${timestamp}_$1.down.sql"
