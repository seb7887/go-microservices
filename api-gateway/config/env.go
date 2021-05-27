package config

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"sync"
	"reflect"
	"strings"
	"strconv"
)

type Configuration struct {
	Port int `env:"PORT" default:"9000" json:"port"`
	UsersHost string `env:"USERS_HOST" default:"localhost" json:"usersHost"`
	UsersPort string `env:"USERS_PORT" default:"7001" json:"usersPort"`
	JWTSecret string `env:"JWT_SECRET" default:"jwtSecret" json:"jwtSecret"`
	RedisHost string `env:"REDIS_HOST" default:":6379" json:"redisHost"`
}

var _config *Configuration
var _once sync.Once

func GetConfig() *Configuration {
	_once.Do(func() {
		Setup()
	})
	return _config
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load .env file, using environment variables. Error: %v", err)
	}
}

func Setup() {
	loadEnv()
	_config = &Configuration{}

	configReflect := reflect.ValueOf(_config).Elem()
	err := loadConfig(configReflect, configReflect.Type())

	if err != nil {
		log.Println("Error in reading go-bank configurations.")
	}

}

func internalField(fieldDef reflect.StructField) bool {
	return "true" == fieldDef.Tag.Get("internal")
}

func requiredField(fieldDef reflect.StructField) bool {
	return "true" == fieldDef.Tag.Get("required")
}

func loadConfigField(field reflect.Value, fieldDef reflect.StructField) error {
	var err error
	configField := fieldDef.Tag.Get("env")
	defaultValue := fieldDef.Tag.Get("default")
	configValue := os.Getenv(configField)
	if len(configValue) == 0 {
		configValue = defaultValue
	}
	if len(configValue) == 0 {
		if requiredField(fieldDef) {
			log.Printf("Field %s missing required configuration", configField)
		}
	} else {
		switch field.Type().Kind() {
		case reflect.Slice:
			values := strings.Split(configValue, ",")
			field.Set(reflect.ValueOf(values))
			if !internalField(fieldDef) {
				log.Println("Loaded configuration")
			}
		case reflect.String:
			field.SetString(configValue)
			if !internalField(fieldDef) {
				log.Println("Loaded configuration")
			}
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(configValue)
			if err != nil {
				log.Println("Invalid configuration")
			} else {
				field.SetBool(boolValue)
				if !internalField(fieldDef) {
					log.Println("Loaded configuration")
				}
			}
		case reflect.Int:
			intValue, err := strconv.Atoi(configValue)
			if err != nil {
				log.Println("Invalid configuration")
			} else {
				// Make sure the configured value meets the minimum requirements
				field.SetInt(int64(intValue))
				if !internalField(fieldDef) {
					log.Println("Loaded configuration")
				}
			}
		}
	}

	return err
}

func loadConfig(configValue reflect.Value, configValueType reflect.Type) error {
	var err error
	for i := 0; i < configValue.NumField(); i++ {
		field := configValue.Field(i)
		if field.Type().Kind() == reflect.Struct {
			err = loadConfig(field, field.Type())
		} else {
			err = loadConfigField(field, configValueType.Field(i))
		}
	}
	return err
}