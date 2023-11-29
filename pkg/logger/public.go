package logger

import (
	"log"
	"os"
)

func Debug(kv ...interface{}) {
	log.Println(body("debug", caller(2), kv...))
}
func Info(kv ...interface{}) {
	log.Println(Yellow(body("info", caller(2), kv...)))
}
func Warn(kv ...interface{}) {
	log.Println(Magenta(body("warn", caller(2), kv...)))
}
func Success(kv ...interface{}) {
	log.Println(Green(body("success", caller(2), kv...)))
}
func Error(kv ...interface{}) {
	log.Println(Red(body("error", caller(2), kv...)))
}
func Fatal(kv ...interface{}) {
	log.Println(Red(body("fatal", caller(2), kv...)))
	os.Exit(1)
}
