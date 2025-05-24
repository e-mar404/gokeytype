package test

type status int
type Test struct {
	text []byte
	status []status
	position int
  windowWidth int
  windowHeight int
}

const (
  Empty status = iota
	INCORRECT
	CORRECT
)

func New(width, height int) Test {
  testText, statusSlice := createText()
	return Test {
		text: []byte(testText),
		status: statusSlice, 
		position: 0,
    windowWidth: width,
    windowHeight: height,
	}
}

func createText() ([]byte, []status) {
  text := []byte("this is a long test")
  statusSlice:= make([]status, len(text))
  for i := range(statusSlice) {
    statusSlice[i] = Empty 
  }
  return text, statusSlice
}
