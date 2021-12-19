package day18_test

import (
	"adventofcode/2021/day18"
	"adventofcode/testhelper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsing(t *testing.T) {
	inputs := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}
	for _, i := range inputs {
		n, _, err := day18.ParseRemaining(i)
		require.NoError(t, err)
		require.Equal(t, i, n.String())
	}
}

func TestExplode(t *testing.T) {
	inputs := []struct {
		start    string
		expected string
	}{
		{
			start:    "[[[[[9,8],1],2],3],4]",
			expected: "[[[[0,9],2],3],4]",
		},
		{
			start:    "[7,[6,[5,[4,[3,2]]]]]",
			expected: "[7,[6,[5,[7,0]]]]",
		},
		{
			start:    "[[6,[5,[4,[3,2]]]],1]",
			expected: "[[6,[5,[7,0]]],3]",
		},
		{
			start:    "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			start:    "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, i := range inputs {
		t.Run(fmt.Sprintf("%q should explode into %q", i.start, i.expected), func(t *testing.T) {
			n, _, err := day18.ParseRemaining(i.start)
			require.NoError(t, err)
			n, _, _, exploded := n.Explode(0)
			require.True(t, exploded)
			require.Equal(t, i.expected, n.String())
		})
	}
}

func TestFullExample(t *testing.T) {
	expr1 := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	expr2 := "[1,1]"
	n1, _, err := day18.ParseRemaining(expr1)
	require.NoError(t, err)
	n2, _, err := day18.ParseRemaining(expr2)
	require.NoError(t, err)

	// Add
	n := n1.Add(n2)
	require.Equal(t, "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", n.String())

	// Explode
	n, _, _, exploded := n.Explode(0)
	require.True(t, exploded)
	require.Equal(t, "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", n.String())

	// Explode
	n, _, _, exploded = n.Explode(0)
	require.True(t, exploded)
	require.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", n.String())

	// Split
	n, split := n.Split()
	require.True(t, split)
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", n.String())

	// Split
	n, split = n.Split()
	require.True(t, split)
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", n.String())

	// Explode
	n, _, _, exploded = n.Explode(0)
	require.True(t, exploded)
	require.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", n.String())
}

func TestFullTarget(t *testing.T) {
	inputs := []struct {
		numbers  []string
		expected string
	}{
		// {
		// 	numbers:  []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"},
		// 	expected: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		// },
		// {
		// 	numbers:  []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"},
		// 	expected: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		// },
		// {
		// 	numbers:  []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"},
		// 	expected: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		// },
		{
			numbers: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			},
			expected: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		// {
		// 	numbers: []string{
		// 		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		// 		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		// 		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		// 		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		// 		"[7,[5,[[3,8],[1,4]]]]",
		// 		"[[2,[2,2]],[8,[8,1]]]",
		// 		"[2,9]",
		// 		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		// 		"[[[5,[7,4]],7],1]",
		// 		"[[[[4,2],2],6],[8,7]]",
		// 	},
		// 	expected: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		// },
	}
	for _, i := range inputs {
		input := i.numbers
		problemNumber, _, err := day18.ParseRemaining(input[0])
		require.NoError(t, err)
		input = input[1:]

		for len(input) > 0 {
			nextN, _, err := day18.ParseRemaining(input[0])
			require.NoError(t, err)
			input = input[1:]
			problemNumber = problemNumber.Add(nextN)
			for {
				problemNumber, _, _, _ = problemNumber.Explode(0)
				newN, Split := problemNumber.Split()
				if Split {
					problemNumber = newN
					continue
				}
				break
			}
		}
		require.Equal(t, i.expected, problemNumber.String())
	}
}

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 4140,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day18.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 112,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day18.Solve2, input)
}
