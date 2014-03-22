package env

import (
    "os"
)

type C map[string] interface{}

var envName string
var commons C
var configs map[string] C

func Init(e string) {
    envName = e
    commons = nil
    configs = nil
}

func Common(values map[string] interface{}) {
    commons = values
}

func Config(environment string, values map[string] interface{}) {
    if configs == nil {
        configs = make(map[string] C)
    }
    configs[environment] = values
}

func Get(key string) interface{} {
    value := configs[os.Getenv(envName)][key]
    if value == nil {
        value = commons[key]
    }
    return value
}

func GetInt(key string) int {
    return Get(key).(int)
}

func GetString(key string) string {
    return Get(key).(string)
}

func GetBool(key string) bool {
    return Get(key).(bool)
}
