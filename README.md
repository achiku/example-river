## Setup

Install libraries.

```
$ go mod init github.com/achiku/example-river
$ go get github.com/riverqueue/river
$ go get github.com/riverqueue/river/riverdriver/riverpgxv5
```

Install CLI.

```
$ go install github.com/riverqueue/river/cmd/river@latest
```

```
(~/work/achiku/example-river)[main|…3] 
❯❯❯ which river
/Users/achiku/sdk/go1.22.0/bin/river
```

Create user and database.

```
achiku:template1:[local]:dev 15:47 =# create user river;
CREATE ROLE
Time: 7.493 ms
achiku:template1:[local]:dev 15:48 =# create database river owner river;
CREATE DATABASE
Time: 686.168 ms
achiku:template1:[local]:dev 15:48 =#
```

```
❯❯❯ psql -U river -d river
Timing is on.
Expanded display is used automatically.
Null display is "[NULL]".
psql (14.11 (Homebrew), server 16.2 (Homebrew))
WARNING: psql major version 14, server major version 16.
         Some psql features might not work.
Type "help" for help.

river:river:[local]:dev 15:49 => create table t1 (id bigint);
CREATE TABLE
Time: 10.284 ms
river:river:[local]:dev 15:49 => \dt
       List of relations
 Schema | Name | Type  | Owner
--------+------+-------+-------
 public | t1   | table | river
(1 row)

river:river:[local]:dev 15:49 => drop table t1;
DROP TABLE
Time: 1.731 ms
```

Migrate tables.

```
(~/work/achiku/example-river)[main|…3]
❯❯❯ river migrate-up --database-url "postgres://river@localhost:5432/river"

applied migration 001 [up] create river migration  [44.82ms]
applied migration 002 [up] initial schema          [276.12ms]
applied migration 003 [up] river job tags non null [4.7ms]
applied migration 004 [up] pending and more        [10.76ms]
applied migration 005 [up] migration unique client [40.68ms]
applied migration 006 [up] bulk unique             [3.7ms]
```

Check tables.

```
river:river:[local]:dev 15:52 => \dt
              List of relations
 Schema |        Name        | Type  | Owner
--------+--------------------+-------+-------
 public | river_client       | table | river
 public | river_client_queue | table | river
 public | river_job          | table | river
 public | river_leader       | table | river
 public | river_migration    | table | river
 public | river_queue        | table | river
(6 rows)
```

## code

- https://github.com/achiku/example-river/blob/main/river.go

## test

- https://github.com/achiku/example-river/blob/main/river_test.go