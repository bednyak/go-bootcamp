package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue

	return dmgValue
}

func HealPlayer() int {

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healDiff := PLAYER_HEALTH - currentPlayerHealth

	if healDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healDiff
	}
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dmgValue

	return dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
