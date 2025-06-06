package ubisoft

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/encoding/protowire"
	"gopkg.in/yaml.v3"
)

// Parser for `Ubisoft Game Launcher\cache\configuration\configurations` file.
// https://github.com/JosefNemec/Playnite/issues/196
// Looks like a protobuf:
// (u)int32: Launch/Game ID
// (u)int32: Installation ID (the same in most cases as launch id based on couple games I checked)
// string: YAML string with all the game information

type UplayProtobuf struct {
	GameID    uint32
	InstallID uint32
	GameYaml  string
}

// Game data from the YAML string.
// Each yaml string has tons of properties,
// majority of them are not needed.
type UplayGame struct {
	Root struct {
		Name      string `yaml:"name"`
		StartGame struct {
			Online struct {
				Executables []struct {
					Path struct {
						Relative string `yaml:"relative"`
					} `yaml:"path"`
					WorkingDirectory struct {
						Register string `yaml:"register"`
					} `yaml:"working_directory"`
				} `yaml:"executables"`
			} `yaml:"online"`
			Offline struct {
				Executables []struct {
					Path struct {
						Relative string `yaml:"relative"`
					} `yaml:"path"`
					WorkingDirectory struct {
						Register string `yaml:"register"`
					} `yaml:"working_directory"`
				} `yaml:"executables"`
			} `yaml:"offline"`
		} `yaml:"start_game"`
	} `yaml:"root"`
}

func parseConfigurationsFile(base_path string) ([]UplayProtobuf, error) {
	conf_file := filepath.Join(base_path, "cache", "configuration", "configurations")
	if _, err := os.Stat(conf_file); err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(conf_file)
	if err != nil {
		return nil, err
	}

	return decodeConfigurationsFile(raw)
}

func decodeConfigurationsFile(data []byte) ([]UplayProtobuf, error) {
	var games []UplayProtobuf

	for len(data) > 0 {
		num, typ, n := protowire.ConsumeTag(data)
		if n < 0 {
			return nil, fmt.Errorf("failed to read tag: %v", protowire.ParseError(n))
		}
		data = data[n:]

		if typ != protowire.BytesType || num != 1 {
			return nil, fmt.Errorf("unexpected field number or type: got %d (type %v)", num, typ)
		}

		msg_bytes, m := protowire.ConsumeBytes(data)
		if m < 0 {
			return nil, fmt.Errorf("failed to read embedded message: %v", protowire.ParseError(m))
		}
		data = data[m:]

		game, err := parseUplayGame(msg_bytes)
		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}

func parseUplayGame(data []byte) (UplayProtobuf, error) {
	var game UplayProtobuf

	for len(data) > 0 {
		num, typ, n := protowire.ConsumeTag(data)
		if n < 0 {
			return game, fmt.Errorf("failed to read tag: %v", protowire.ParseError(n))
		}
		data = data[n:]

		switch num {
		case 1: // Game ID
			if typ != protowire.VarintType {
				return game, fmt.Errorf("unexpected type for uplay_id: %v", typ)
			}
			val, vlen := protowire.ConsumeVarint(data)
			if vlen < 0 {
				return game, fmt.Errorf("invalid varint for uplay_id")
			}

			game.GameID = uint32(val)
			data = data[vlen:]

		case 2: // Install ID
			if typ != protowire.VarintType {
				return game, fmt.Errorf("unexpected type for install_id: %v", typ)
			}
			val, vlen := protowire.ConsumeVarint(data)
			if vlen < 0 {
				return game, fmt.Errorf("invalid varint for install_id")
			}

			game.InstallID = uint32(val)
			data = data[vlen:]

		case 3: // Game YAML data
			if typ != protowire.BytesType {
				return game, fmt.Errorf("unexpected type for game_info: %v", typ)
			}
			str, slen := protowire.ConsumeBytes(data)
			if slen < 0 {
				return game, fmt.Errorf("invalid string")
			}

			game.GameYaml = string(str)
			data = data[slen:]

		default: // Ignore unknown fields
			skip := protowire.ConsumeFieldValue(num, typ, data)
			if skip < 0 {
				return game, fmt.Errorf("failed to skip field")
			}
			data = data[skip:]
		}
	}

	return game, nil
}

func parseGameYaml(data []byte) (*UplayGame, error) {
	var game UplayGame
	if err := yaml.Unmarshal(data, &game); err != nil {
		return nil, err
	}

	return &game, nil
}
