package regex_util

import "regexp"

func FindStringCaptured(re *regexp.Regexp, text string) map[string]string {
	names := re.SubexpNames()
	result := make(map[string]string)
	for i, str := range re.FindStringSubmatch(text) {
		if i == 0 {
			continue
		}
		result[names[i]] = str
	}
	return result
}
