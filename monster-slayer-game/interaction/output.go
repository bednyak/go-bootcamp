package interaction

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong atack againts monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	asciiFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)
	asciiFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {

	// May fix problem if compiled go script don't save in correct place. This solution won't run if you just launch 'go run .'

	/*exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing log file failed. Exiting!")
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")*/

	file, err := os.Create("/gamelog.txt")

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, value := range *rounds {

		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)

		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log")
}
