package runtime

import "github.com/elliott5/gohaxelib/hxbuiltin"

func Gosched()          {} // an empty function here works fine to enable other goroutines to be scheduled
func NumGoroutine() int { return int(hxbuiltin.HAXE("Scheduler.NumGoroutine();")) }
