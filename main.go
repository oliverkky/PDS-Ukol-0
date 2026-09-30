package main

import (
	"math/rand"
	"time"
)

var globalCounter int

func main() {
	spawnThreads(500)
}

// funkce, která dle zadaného parametru spustí daný počet vláken. Každé vlánko inkrementuje čístač.
func spawnThreads(count int) {
	for i := 0; i < count; i++ {
		go func() {
			globalCounter++
			println("globalCounter in spawnThreads", globalCounter)
		}()
	}
}

// funkce, která náhodný čas počká, zkopíruje hodnotu z globálního čítače do jeho lokální kopie, náhodný čas počká, inkrementuje lokální kopii čítače, uloží hodnotu z lokální kopie do globálního čítače.
func threadFunc() {
	localCounter := globalCounter
	randomSleep()
	localCounter++
	randomSleep()
	globalCounter = localCounter
}

func randomSleep() {
	time.Sleep(time.Duration(rand.Intn(1000)))
}
