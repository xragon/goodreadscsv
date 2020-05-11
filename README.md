# GoodReads CSV Importer

Import a subset the goodreads CSV export into a postgres database

## Development Plan
- [x] Hard coded import of CSV file - DONE
- [x] Handle the imported CSV so that I can pick and choose the fields I want - DONE
- [x] Command line control so I can execute the code on whichever file I choose - Done
- [ ] Write chosen files to a postgres DB - In Progress

## Setup

### Create Database
`flyway -configFiles=migration/conf/flyway.conf -locations=filesystem:./migration/sql/ migrate`
