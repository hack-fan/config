# Config Loader for Golang
Load config into go struct from shell environment and docker/k8s secrets.

## Install
```shell script
go get github.com/hyacinthus/config
```

## Features
- [x] Load from shell environment variables
- [x] Load from Docker/Kubernetes secrets
- [x] Default values support
- [x] Required check support
- [x] Simple and easy to use, no other features

## Load Order
`Default` -> `ENV` -> `Secret` -> `Value exists in struct`

Right side will overwrite left side.

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/hyacinthus/config"
)

type Settings struct {
    AppName string `default:"app"` // env APP_NAME will overwrite default value
    DB      struct {
        Name     string
        User     string `required:"true"`
        Password string `secret:"mysql_db_password"` // default secret name is 'db_password',change it use tag
        Port     int    `default:"3306" env:"MYSQL_DB_PORT"` // default env name is 'DB_PORT',change it use tag
    }
}

func main() {
    var settings = new(Settings)
    config.MustLoad(settings)
    fmt.Printf("%+v",settings)
}
```

## Name Conversion

- `ENV` will use ALL_CAP_SNAKE_CASE
- `Secret` will use snake_case

## Tags

- `default` set default value
- `env` custom shell env variable names
- `secret` custom secret file name
- `required` set attr as required
