package day14

type binaryNumber [36]string

type assignment struct {
	memoryAddress int
	value         binaryNumber
}

func (a *assignment) mask(m binaryNumber) {
	for i := 0; i < len(a.value); i++ {

	}
}

type program struct {
	mask        binaryNumber
	assignments []assignment
}

func Solve1(inputFilePath string) (int, error) {
	panic("not implemented")
}

func Solve2(inputFilePath string) (int, error) {
	panic("not implemented")
}
