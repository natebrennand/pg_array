
# PG Array

[![Build Status](https://travis-ci.org/natebrennand/pg_array.svg)](https://travis-ci.org/natebrennand/pg_array)

A simple datatype that can be used to scan in array in Postgres.

database/sql does not currently support array types.
This allows arrays to be read out of the database.


NOTE: allocation and formation of the arrays is done very inefficiently with repeated calls with `append`.

### Arrays Supported

- int
- string


