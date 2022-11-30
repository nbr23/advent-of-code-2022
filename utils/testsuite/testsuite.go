package testsuite

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
)

func RunTestforPart(t *testing.T, part1 utils.Resolver, part2 utils.Resolver, part int, day int) {
	items, _ := ioutil.ReadDir(fmt.Sprintf("../inputs/test/day%02d", day))
	var result string
	for _, item := range items {
		if item.IsDir() {
			t.Logf("Testing with : %s\n", fmt.Sprintf("../inputs/test/day%02d/%s/input.txt", day, item.Name()))
			binput, err := os.ReadFile(fmt.Sprintf("../inputs/test/day%02d/%s/input.txt", day, item.Name()))
			if err != nil {
				t.Logf("Test input not found %s\n", fmt.Sprintf("../inputs/test/day%02d/%s/input.txt", day, item.Name()))
				t.Fail()
				continue
			}
			bresult, err := os.ReadFile(fmt.Sprintf("../inputs/test/day%02d/%s/result_p%d.txt", day, item.Name(), part))
			if err != nil {
				t.Logf("Test result not found %s\n", fmt.Sprintf("../inputs/test/day%02d/%s/result_p%d.txt: skipping…", day, item.Name(), part))
				continue
			}
			if part == 1 {
				result = fmt.Sprintf("%v", part1(string(binput)))
			} else {
				result = fmt.Sprintf("%v", part2(string(binput)))
			}
			t.Logf("Found %s, expected %s\n", result, string(bresult))
			if result != string(bresult) {
				t.Fail()
			}
		}
	}
}
