package configer

import (
	"fmt"
	"path/filepath"

	"encoding/json"

	"github.com/bigkevmcd/go-configparser"

	"strings"
)

func GetConfigVal(conf string, section string, key string, convertToMap bool) any {
	p, err := configparser.NewConfigParserFromFile(filepath.Join("configs", conf, fmt.Sprintf("%s.cfg", conf)))
	if err != nil {
		panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
	}
	val, err := p.GetInterpolated(section, key)
	if err != nil {
		panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
	}

	if convertToMap {
		var data map[string]string
		err := json.Unmarshal([]byte(val), &data)
		if err != nil {
			panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
		}
		return data
	}

	return val
}

func GetFullConf(conf string) map[string]interface{} {
	p, err := configparser.NewConfigParserFromFile(filepath.Join("configs", conf, fmt.Sprintf("%s.cfg", conf)))
	if err != nil {
		panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
	}
	sections := p.Sections()

	data := make(map[string]interface{})

	for _, sec := range sections {
		keyVals, err := p.Items(sec)
		if err != nil {
			panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
		}
		data[sec] = keyVals
	}
	return data
}

func GetSection(conf string, section string) map[string]interface{} {
	p, err := configparser.NewConfigParserFromFile(filepath.Join("configs", conf, fmt.Sprintf("%s.cfg", conf)))
	if err != nil {
		panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
	}
	sections := p.Sections()

	data := make(map[string]interface{})

	for _, sec := range sections {
		if sec == section {
			keyVals, err := p.Items(sec)
			if err != nil {
				panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
			}

			for k, v := range keyVals {
				if !strings.HasPrefix(v, "{") {
					data[k] = v
				} else {
					var val map[string]string
					err := json.Unmarshal([]byte(v), &val)
					if err != nil {
						panic(fmt.Sprintf("ERROR: CONFIG --- %s", err))
					}
					data[k] = val
				}

			}
		}
	}
	return data
}
