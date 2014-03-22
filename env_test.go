package env

import (
    "testing"
)

func TestInit(t *testing.T) {
    Init("MY_APP_ENV")
    if envName != "MY_APP_ENV" {
        t.Errorf("envName not correctly set")
    }
}

func TestCommon(t *testing.T) {
    Init("MY_APP_ENV")
    Common(C{
        "foo": "foo",
    })

    if commons["foo"] != "foo" {
        t.Errorf("common config value not correctly set")
    }
}

func TestConfig(t *testing.T) {
    Init("MY_APP_ENV")
    Config("local", C{
        "foo": "foo",
    })

    if configs["local"]["foo"] != "foo" {
        t.Errorf("env-specific config value not correctly set")
    }
}

func TestGet(t *testing.T) {
    Init("MY_APP_ENV")
    Common(C{
        "foo": "foo",
        "bar": 1,
    })
    Config("local", C{
        "foo": "bar",
    })

    if Get("foo").(string) != "bar" {
        t.Errorf("common value is not overridden by env-specific config value: expect = %s, got = %s", "bar", Get("foo").(string))
    }
    if Get("bar").(int) != 1 {
        t.Errorf("could not get common value: expect = %v, got = %v", 1, Get("bar").(int))
    }
}

func TestGetInt(t *testing.T) {
    Init("MY_APP_ENV")
    Common(C{
        "bar": 1,
    })

    if GetInt("bar") != 1 {
        t.Errorf("could not get common int value: expect = %v, got = %v", 1, GetInt("bar"))
    }
}

func TestGetString(t *testing.T) {
    Init("MY_APP_ENV")
    Common(C{
        "foo": "foo",
    })

    if GetString("foo") != "foo" {
        t.Errorf("could not get common string value: expect = %v, got = %v", "foo", GetInt("foo"))
    }
}

func TestGetBool(t *testing.T) {
    Init("MY_APP_ENV")
    Common(C{
        "baz": true,
    })

    if GetBool("baz") != true {
        t.Errorf("could not get common bool value: expect = %v, got = %v", true, GetBool("baz"))
    }
}
