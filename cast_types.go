package tdenginebuilder

import (
	"fmt"

	"github.com/tkcrm/tdenginebuilder/tdenginefuncs"
)

func castType(value any) (string, error) {
	switch v := value.(type) {
	case float32, float64, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint:
		return fmt.Sprintf("%d", v), nil
	case string:
		return fmt.Sprint("'" + v + "'"), nil
	case bool:
		return fmt.Sprintf("%v", v), nil
	case tdenginefuncs.TDEngineFunc:
		return v.String(), nil
	default:
		return "", fmt.Errorf("undefined type: %T", v)
	}
}
