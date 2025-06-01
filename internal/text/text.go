package text

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const wordBankFile = ".english.words"

func Generate(wordCount int) (string, error) {
  file, err := os.Open(wordBankFile)   
  if err != nil {
    return "", err
  } 

  scanner := bufio.NewScanner(file)

  scanner.Scan()
  lengthStr := scanner.Text()
  length, _ := strconv.ParseInt(lengthStr, 10, 32)
  wordBank := make([]string, length) 

  count := 0
  for scanner.Scan() {
    wordBank[count] = scanner.Text()
    count++
  }
  
  r := rand.New(rand.NewSource(time.Now().Local().UnixMilli()))
  wordBankSize := len(wordBank)
  var textList []string
  for range(wordCount) {
    textList = append(textList, wordBank[r.Intn(wordBankSize)])
  }

  return strings.Join(textList, " "), nil
}
