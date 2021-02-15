package offer

import (
	"regexp"
	"strings"
)

func isNumber(s string) bool {
	var re = `^ *[-+]?[0-9]+\.?[0-9]* *$|^ *[-+]?[0-9]*\.?[0-9]+ *$|^ *[-+]?[0-9]+\.?[0-9]*[Ee][-+]?[0-9]+ *$|^ *[-+]?[0-9]*\.?[0-9]+[Ee][-+]?[0-9]+ *$`
	return regexp.MustCompile(re).MatchString(strings.TrimSpace(s))
}
