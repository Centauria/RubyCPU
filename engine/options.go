package engine

import "fmt"

// Options available:
// Hash: int
// UCI_AnalyseMode: bool
var (
	Options = map[string]interface{}{
		"Hash":            1,
		"UCI_AnalyseMode": false,
		"UCI_Chess960":    false,
	}
	//TODO: add `OptionConditions` which includes limits of the Options
)

func CheckOptions() (err error) {
	if Options["Hash"].(int) < 1 || Options["Hash"].(int) > 64 {
		err = fmt.Errorf("hash size must be 1~64 integer")
	}
	return
}
