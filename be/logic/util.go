package logic

import "golang.org/x/exp/rand"

var colorCodeArray = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

// ランダムかカラーコードを生成する関数
func CreateRandomColorCode() string {
	colorCodeString := "#"

	for i := 0; i < 6; i++ {
		colorCodeString += colorCodeArray[rand.Intn(15)]
	}

	return colorCodeString
}
