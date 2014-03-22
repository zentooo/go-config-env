# go-config-env

## Description
Config::ENV like config loader in go

## Usage

```
// config.go
package main

import (
    "env"
)

func init() {
    env.Init("MY_APP_ENV")
    env.Common(env.C{
        "appname": "MyApp",
    })
    env.Config("local", env.C{
        "dsn": "root:@unix(/tmp/mysql.sock)/mydb",
    })
    env.Config("production", env.C{
        "dsn": "mysql_user:@10.55.1.16/mydb",
    })
}

// main.go
package main

import (
    "env"
)

func main() {
    env.GetString("foo") => "bar"
    env.GetInt("bar") => 1
    env.GetBool("baz") => true
}

// run

MY_APP_ENV=local go run main.go config.go

```
