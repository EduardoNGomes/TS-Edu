package configs

import "os"

func createBiome() error {
	fileContent := `{
	"vcs": {
		"enabled": false,
		"clientKind": "git",
		"useIgnoreFile": false
	},
	"files": {
		"ignoreUnknown": false,
		"includes": ["**", "!!**/dist"]
	},
	"linter": {
		"enabled": true,
		"rules": {
			"recommended": true,
			"correctness": {
				"noUnusedVariables": {
					"level": "warn",
					"fix": "safe"
				}
			},
			"complexity": {
				"noForEach": "off"
			},
			"style": {
				"useNodejsImportProtocol": {
					"level": "error",
					"fix": "safe"
				}
			}
		}
	},
	"formatter": {
		"enabled": true,
		"indentStyle": "space",
		"indentWidth": 4,
		"lineWidth": 80,
		"lineEnding": "lf"
	},
	"javascript": {
		"formatter": {
			"quoteStyle": "double",
			"trailingCommas": "none",
			"semicolons": "always",
			"jsxQuoteStyle": "double"
		}
	},
	"json": {
		"formatter": {
			"trailingCommas": "none",
			"indentStyle": "tab"
		},
		"parser": {
			"allowComments": true
		}
	}
}
`
	if err := os.WriteFile("biome.json", []byte(fileContent), os.ModePerm); err != nil {

		return err
	}

	return nil
}
