package res

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const filePath string = "C:/mycode/go/goproject/leetcode/testingPractice1/res/monster.json"

func (monster *Monster) Store() bool {
	monsterByteSlice, err := json.Marshal(*monster)
	if err != nil {
		return false
	} else {
		monsterStr := string(monsterByteSlice)
		var fd *os.File
		fd, err = os.OpenFile(filePath, os.O_RDWR, 6)
		if err != nil {
			return false
		}
		defer fd.Close()
		writer := bufio.NewWriter(fd)
		writer.WriteString(monsterStr)
		writer.Flush()
		return true
	}
}

func (monster *Monster) Restore() (*Monster, error) {
	var newMonster *Monster = &Monster{}
	fd, err := os.OpenFile(filePath, os.O_RDWR, 6)
	if err != nil {
		return newMonster, err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return newMonster, err
	}
	err = json.Unmarshal([]byte(str), newMonster)
	if err != nil {
		return newMonster, err
	}
	return newMonster, nil
}
