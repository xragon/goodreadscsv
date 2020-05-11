# GoodReads CSV Importer

Import a subset the goodreads CSV export into a postgres database

## Development Plan
- [x] Hard coded import of CSV file
- [x] Handle the imported CSV so that I can pick and choose the fields I want
- [x] Command line control so I can execute the code on whichever file I choose
- [x] Write chosen files to a postgres DB
- [ ] ????

## Setup

### Create Database
`$flyway -configFiles=migration/conf/flyway.conf -locations=filesystem:./migration/sql/ migrate`

## Use

`$go run main.go -path path/to/file.csv`

If no file path is provided it will default to the working directory and the default download filename "goodreads_library_export.csv".