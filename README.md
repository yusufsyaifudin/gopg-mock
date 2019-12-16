# Go-PG Mock

This package is Proof of Concept that if we use [https://github.com/go-pg/pg](https://github.com/go-pg/pg) 
for our DB connection, it still can be tested. Since go-pg package already use abstraction for their DB query 
(using [`orm.DB`](https://github.com/go-pg/pg/blob/v9.1.0/orm/orm.go#L34-L58) interface), we can create new struct that implement it. 

At the same time, instead doing real query we can compare what queries we want to mock and what results will be.

## Usage

See in directory `examples`.

