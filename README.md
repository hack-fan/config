# config
Load config into go struct from shell environment and docker/k8s secrets.

## Install
```shell script
go get github.com/hyacinthus/config
```

## Quick Start
```go
package main

import (
    "fmt"

	"github.com/hyacinthus/config"
)

type Settings struct {
	AppName string `default:"default_app"` // testing snake env APP_NAME
	DB      struct {
		Name     string `default:"default_name"`             // testing default
		User     string `default:"default_user"`             // testing env
		Password string `default:"default_pwd"`              // testing secret
		Port     int    `default:"3306" env:"MYSQL_DB_PORT"` // testing int and custom env name
	}
}

func main() {
    var settings = new(Settings)
    config.MustLoad(settings)
    fmt.Printf("%+v",settings)
}
```