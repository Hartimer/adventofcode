package _test

import (
	"adventofcode/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const emptyFileID = -1

type file struct {
	fileID int
	length int
}

func (f file) fill(f2 file) ([]file, bool) {
	if f.fileID != emptyFileID {
		panic("not an empty slot")
	}
	if f.length < f2.length {
		return []file{
			{
				fileID: f2.fileID,
				length: f.length,
			},
			{
				fileID: f2.fileID,
				length: f2.length - f.length,
			},
		}, true
	}
	if f.length == f2.length {
		return []file{f2}, false
	}

	return []file{
		f2,
		{
			fileID: emptyFileID,
			length: f.length - f2.length,
		},
	}, false
}

type filesystem []file

func (fs filesystem) compact() filesystem {
	idx := 1
	for idx < len(fs) {
		emptySlot := fs[idx]
		if emptySlot.fileID != emptyFileID {
			idx++
			continue
		}
		remaining := fs[idx+1:]
		if len(remaining) == 0 {
			break
		}
		for j := len(remaining) - 1; j >= 0; j-- {
			file := remaining[j]
			if file.fileID == emptyFileID {
				continue
			}
			newBlocks, hasRemainingFile := emptySlot.fill(file)
			if hasRemainingFile {
				fs[idx] = newBlocks[0]
				remaining[j] = newBlocks[1]
			} else {
				var newFs filesystem
				newFs = append(newFs, fs[:idx]...)
				newFs = append(newFs, newBlocks...)
				newFs = append(newFs, fs[idx+1:idx+j+1]...)
				fs = newFs
			}
			break
		}
	}
	return fs
}

func (fs filesystem) compact2() filesystem {
	idx := 1
	for idx < len(fs) {
		emptySlot := fs[idx]
		if emptySlot.fileID != emptyFileID {
			idx++
			continue
		}
		remaining := fs[idx+1:]
		if len(remaining) == 0 {
			break
		}
		for j := len(remaining) - 1; j >= 0; j-- {
			remainingFile := remaining[j]
			if remainingFile.fileID == emptyFileID || remainingFile.length > emptySlot.length {
				continue
			}
			newBlocks, hasRemainingFile := emptySlot.fill(remainingFile)
			if hasRemainingFile {
				panic("should not happen")
			}
			remaining[j] = file{
				fileID: emptyFileID,
				length: newBlocks[0].length,
			}
			if len(newBlocks) == 1 {
				fs[idx] = newBlocks[0]
			} else {
				var newFs filesystem
				newFs = append(newFs, fs[:idx]...)
				newFs = append(newFs, newBlocks...)
				newFs = append(newFs, fs[idx+1:]...)
				fs = newFs
			}
			break
		}
		idx++
	}
	return fs
}

func (fs filesystem) hash() int {
	total := 0
	idx := 0
	for _, file := range fs {
		if file.fileID == emptyFileID {
			idx += file.length
			continue
		}
		for i := 0; i < file.length; i++ {
			total += (idx * file.fileID)
			idx++
		}
	}
	return total
}

func TestDay9_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day9.1.1.input",
			expected: 1928,
		},
		{
			filename: "day9.1.input",
			expected: 6216544403458,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			var files filesystem
			for fileLine := range helper.FileLineReader(input.filename) {
				fileMode := true
				fileID := 0
				for _, cell := range fileLine {
					n := int(cell - '0')
					if fileMode {
						files = append(files, file{
							fileID: fileID,
							length: n,
						})
						fileID++
					} else if n > 0 {
						files = append(files, file{
							fileID: emptyFileID,
							length: n,
						})
					}
					fileMode = !fileMode
				}
			}
			files = files.compact()
			require.Equal(t, input.expected, files.hash())
		})
	}
}
func TestDay9_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day9.1.1.input",
			expected: 2858,
		},
		{
			filename: "day9.1.input",
			expected: 6237075041489,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			var files filesystem
			for fileLine := range helper.FileLineReader(input.filename) {
				fileMode := true
				fileID := 0
				for _, cell := range fileLine {
					n := int(cell - '0')
					if fileMode {
						files = append(files, file{
							fileID: fileID,
							length: n,
						})
						fileID++
					} else if n > 0 {
						files = append(files, file{
							fileID: emptyFileID,
							length: n,
						})
					}
					fileMode = !fileMode
				}
			}
			files = files.compact2()
			require.Equal(t, input.expected, files.hash())
		})
	}
}
