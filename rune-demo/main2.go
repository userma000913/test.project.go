package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "{\"msg_type\":\"interactive\",\"card\":{\"header\":{\"template\":\"green\",\"title\":{\"content\":\"%s\",\"tag\":\"plain_text\"}},\"elements\":[{\"fields\":[{\"is_short\":true,\"text\":{\"content\":\"**TIME时间**\\n%s\",\"tag\":\"lark_md\"}},{\"is_short\":true,\"text\":{\"content\":\"**PROJECT项目**\\n%s\",\"tag\":\"lark_md\"}}],\"tag\":\"div\"},{\"tag\":\"hr\"},{\"tag\":\"div\",\"text\":{\"content\":%#v,\"tag\":\"lark_md\"}}]}}"
	//s = strings.Replace(s, "PROJECT", GetSpecialCharactersMap()["PROJECT"], 1)
	//s = strings.Replace(s, "TIME", GetSpecialCharactersMap()["TIME"], 1)
	//s = strings.Replace(s, "PEOPLE", GetSpecialCharactersMap()["PEOPLE"], 1)
	charactersMap := GetSpecialCharactersMap()
	for k, v := range charactersMap {
		if strings.Contains(s, k) {
			s = strings.Replace(s, k, v, 1)
		}
	}
	fmt.Println(s)
}
func GetSpecialCharactersMap() map[string]string {
	return map[string]string{
		"PROJECT": "📋",
		"TIME":    "🕐",
		"PEOPLE":  "👤",
	}
}
