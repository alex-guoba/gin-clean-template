
# Gin Clean Template

Clean Architecture template for Golang Gin services

## Overview 

This is an example of implementation of Gin framework (Golang) service.

- Services are split into multiple layers according to the DDD hierarchical architecture.
- Testabe. Domain or DAO layer can be testable without Database. 
- Database Independence. used GORM to hide the implementation of database, it can be replace by [other](https://gorm.io/docs/connecting_to_the_database.html) DB.


## Quick start

```sh

$ go build

# install mysql, see Makefile

# Run app with defualt configuaration in configs/config.yml
$ ./gin-clean-template
```

## Features

- [x] DB migration
Support database schema migration from migration source('db/migration'). see [migrate](https://github.com/golang-migrate/migrate) for more details.
Use this command to generate some seed data.

``` shell
» ./gin-clean-template seed --count 100
```

- [x] Overload Protection
With the overload protection capabilities of [sentinel](https://github.com/alibaba/sentinel-golang)，system adaptive overload protection capabilities are implemented.

- [x] Load Testing

``` shell
» ./scripts/load-testing/start_load_testing.sh

Text output:
Requests      [total, rate, throughput]         156889, 15686.32, 4762.42
Duration      [total, attack, wait]             10.004s, 10.002s, 1.894ms
Latencies     [min, mean, 50, 90, 95, 99, max]  32.917µs, 12.25ms, 3.367ms, 28.519ms, 45.603ms, 152.698ms, 1.101s
Bytes In      [total, mean]                     22006503, 140.27
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           30.37%
Status Codes  [code:count]                      0:65  200:47641  429:109183
Error Set:
Get "http://127.0.0.1:8000/api/v1/tags": read tcp 127.0.0.1:60416->127.0.0.1:8000: read: connection reset by peer
Get "http://127.0.0.1:8000/api/v1/tags": read tcp 127.0.0.1:60417->127.0.0.1:8000: read: connection reset by peer
429 Too Many Requests

Histogram output:
Bucket           #       %       Histogram
[0s,     10ms]   115958  73.91%  #######################################################
[10ms,   40ms]   31313   19.96%  ##############
[40ms,   80ms]   6390    4.07%   ###
[80ms,   200ms]  2141    1.36%   #
[200ms,  500ms]  1023    0.65%
[500ms,  1s]     63      0.04%
[1s,     +Inf]   1       0.00%
```

## Layered Architecture 

There are 4 layers in this project:

### API

Also known as the application service layer, it is the external service expression of the application layer. It is responsible for adapting external terminals and channels, and for internal parameter verification and protocol conversion, reducing the cost of error handling at lower levels.

### Service

Business layer, semantically speaking at the product level, is responsible for the logical expression of a business scenario and coordinates between business scenarios and domain knowledge. Typically, a business scenario requires collaboration from multiple domains to be completed, such as an e-commerce purchase scenario, which involves multiple domains like account (VIP/anti-fraud verification), payment (order, debit), inventory (deduction), and transaction flow (reconciliation, audit).

The application layer should be kept as simple as possible and should not contain business rules or knowledge. It should not retain the status of business objects but only the status of application task's progress. It focuses more on the display of business capabilities or business processes. It mainly coordinates with domain layers to accomplish its tasks.


### Domain

Domain is used to identify scope and boundaries, assigning specific business issues within specific boundaries. It is mainly responsible for expressing business concepts, business status information, and business rules. Based on importance and functionality, domains can be divided into three sub-domains:

- **Core domain**: Responsible for the rules and implementation of core business scenarios, such as products, transactions, inventory, etc., in e-commerce systems.
- **Generic domain**: Addresses more general requirements, such as authentication and authorization.
- **Support domain**: Non-core functionality that is also not generic logic. It is used to support specific functions, such as a dictionary table for a specific feature.


### DAO

Storage layer, facing the logical expression of operations on storage objects, including DB storage and external RPC. The core value of DAO lies in assembling SQL, maintaining DB connections, and handling transactions, allowing upper layers to focus solely on business logic. Business logic should not be concerned with the form of data storage.


## TODO

### Optimization
- [ ] Remove global: [Why is Global State so Evil?](https://softwareengineering.stackexchange.com/questions/148108/why-is-global-state-so-evil)
- [ ] Enable Docker 
- [ ] setup command to init db, including faking data generated.

### New features
- [ ] Resource upload support, like S3, minio.
- [ ] Auth with OIDC.
- [ ] Enable Cache
- [ ] Enable MQ

## Useful projects
- [ddd-oriented-microservice](https://learn.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/ddd-oriented-microservice)
- [blog-service](https://github.com/go-programming-tour-book/blog-service)