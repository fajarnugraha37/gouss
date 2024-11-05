package util

import (
	"log"
	"time"
    "os"
    "path/filepath"
)

func HangingExecution() {
	go SleepInSecs(1, func() {
		log.Printf("\n\n[*] Waiting for logs. To exit press CTRL+C")
	})
	
	var forever chan struct{}
	<-forever
}

func FailOnErrorWithMessage(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func FailOnError(err error) {
	if err != nil {
		panic("Unexpected Error: " + err.Error())
	}
}

func Coalesce[T any](one *T, two *T) *T {
	if one != nil {
		return one
	}

	return two
}

func GetEwd() string {
    ex, err := os.Executable()
    FailOnErrorWithMessage(err, "failed to get executeble")
	
	return filepath.Dir(ex)
}

func GetCwd() string {
    pwd, err := os.Getwd()
    FailOnErrorWithMessage(err, "failed to get executeble")
	
	return pwd
}

func SleepInMins(mins time.Duration, callback func()) {
	Sleep(mins * time.Minute, callback)
}

func SleepInSecs(secs time.Duration, callback func()) {
	Sleep(secs * time.Second, callback)
}

func Sleep(mills time.Duration, callback func()) {
	time.Sleep(mills * time.Millisecond) 
	callback()
}

func Ternary[T any](cond bool, vtrue, vfalse T) T {
    if cond {
        return vtrue
    }
    return vfalse
}