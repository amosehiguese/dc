package corejsonb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]any

func (jb *JSONB) Value() (driver.Value, error) {
	json, err := json.Marshal(jb)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func (jb *JSONB) Scan(src any) error {
	var datasource []byte
	_m := make(map[string]any)

	switch src.(type) {
	case []uint8:
		datasource = []byte(src.([]uint8))
	case string:
		datasource = []byte(src.(string))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for StringInterfaceMap")
	}

	err := json.Unmarshal(datasource, &_m)
	if err != nil {
		return err
	}

	*jb = JSONB(_m)

	return nil
}
