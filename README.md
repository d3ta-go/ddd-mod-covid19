# ddd-mod-covid19

DDD Module: Covid19 (Core Subdomain)

As a part of `Simple Implementation of Modular DDD Technical Architecture Patterns in Go`.

## Diagram v 0.2.2-Modular

![DDD-Technical-Architecture-Patterns-Golang-0.2.2-DDD Covid19 Module](docs/img/DDD-Technical-Architecture-Patterns-Golang-0.2.2-DDD_Covid19_Module.png)

## Components

A. Interface Layer (None)

B. DDD Modules:

1. Covid19 - using DDD Layered Architecture (Contract, Adapters) [ [d3ta-go/ddd-mod-covid19](https://github.com/d3ta-go/ddd-mod-covid19) ]

C. Common System Libraries [ [d3ta-go/system](https://github.com/d3ta-go/system) ]:

1. Configuration - using yaml
1. Identity & Securities - using JWT, Casbin (RBAC)
1. Initializer
1. Email Sender - using SMTP
1. Handler
1. Migrations
1. Utils

D. Databases

1. MySQL (tested)
2. PostgreSQL (untested)
3. SQLServer (untested)
4. SQLite3 (untested)

E. Providers (Connectors) [ [d3ta-go/connector-\*](https://github.com/d3ta-go/connector-covid19) ]:

1. data.covid19.go.id (Official Covid19 Website - Indonesia)
2. covid19.who.it (Official Covid19 Website - WHO)

F. Persistent Caches

1. Session/Token/JWT Cache (Redis, File, DB, etc) [tested: Redis]

G. Messaging [to-do]

H. Logs [to-do]

### Development

**1. Clone**

```shell
$ git clone https://github.com/d3ta-go/ddd-mod-covid19.git
```

**2. Setup**

```
a. copy `conf/config-sample.yaml` to `conf/config.yaml`
b. copy `conf/data/test-data-sample.yaml` to `conf/data/test-data.yaml`
c. setup your dependencies/requirements (e.g: database, redis, smtp, etc.)
```

**3. Runing TDD on Development Stage**

3.1. TDD: DB Migration Test

```shell
$ cd ddd-mod-covid19
$ sh tdd/clean-testcache.sh

$ sh tdd/ut.db.migration.run-001.sh

$ sh tdd/ut.db.migration.rollback-001.sh
```

3.2. TDD: Functionality Test (unit test)

```shell
$ cd ddd-mod-covid19
$ sh tdd/clean-testcache.sh

$ sh tdd/ut.db.migration.run-001.sh

$ sh tdd/ut.pkg.infra-layer.adapter.covid19.who-001.sh

$ sh tdd/ut.pkg.infra-layer.adapter.covid19.covid19goid-001.sh

$ sh tdd/ut.pkg.infra-layer.repository-001.sh

$ sh tdd/ut.pkg.app-layer.service-001.sh

$ sh tdd/ut.pkg.app-layer.application-001.sh
```

OR

```shell
$ cd ddd-mod-covid19
$ sh tdd/run.tdd.sh
```

**TDD Result Sample:**

- MySQL Database (Casbin Rule):

![MySQL Database Migration Result](docs/img/covid19-sample-db-migration-mysql-casbin.png)

- Redis (Cache):

![Redis (Cache)](docs/img/covid19-sample-cache-redis-server.png)
