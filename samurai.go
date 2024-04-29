package main

import (
	"fmt"
)

func jumlahKalahkan(strength, speed, opponentStrength, opponentSpeed int) int {
	var numEnemiesKilled int

	// Kenshin's strength is greater than or equal to opponent's strength
	if strength >= opponentStrength {
		// Kenshin's speed is greater than opponent's speed
		if speed > opponentSpeed {
			numEnemiesKilled = 1
			speed -= 6
		}
	} else if strength < opponentStrength {
		// Kenshin's speed is greater than opponent's speed
		if speed > opponentSpeed {
			numEnemiesKilled = 1
			strength -= 6
		}
	}

	return numEnemiesKilled
}

func main() {
	fmt.Print("Enter the number of samurais (n): ")
	var n int
	fmt.Scan(&n)

	fmt.Print("Enter Kenshin's strength and speed: ")
	var kenshinStrength, kenshinSpeed int
	fmt.Scan(&kenshinStrength, &kenshinSpeed)

	var totalEnemiesKilled int = 0

	for i := 1; i <= n; i++ {
		fmt.Printf("Enter samurai %d's strength and speed: ", i)
		var enemyStrength, enemySpeed int
		fmt.Scan(&enemyStrength, &enemySpeed)

		enemiesKilled := jumlahKalahkan(kenshinStrength, kenshinSpeed, enemyStrength, enemySpeed)

		totalEnemiesKilled += enemiesKilled
	}

	fmt.Printf("\n%d enemies have been killed\n", totalEnemiesKilled)
	fmt.Printf("Kenshin now has %d strength and %d speed left.\n", kenshinStrength, kenshinSpeed)
}
