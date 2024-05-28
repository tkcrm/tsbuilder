package tsbuilder

import (
	"fmt"

	"github.com/tkcrm/tsbuilder/tsfuncs"
)

func castType(value any) (string, error) {
	switch v := value.(type) {
	case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case string:
		return fmt.Sprint("'" + v + "'"), nil
	case bool:
		return fmt.Sprintf("%v", v), nil
	case tsfuncs.TDEngineFunc:
		return v.String(), nil
	case nil:
		return "NULL", nil
	default:
		return "", fmt.Errorf("undefined type: %T", v)
	}
}
