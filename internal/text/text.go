package text

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
  
  var textList []string
  for i := range(wordCount) {
    textList = append(textList, wordBank[i])
  }

  return strings.Join(textList, " "), nil
}
