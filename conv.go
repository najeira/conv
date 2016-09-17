package conv

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

const (
	Version = "0.2.0"
)

type Stringer interface {
	String() string
}

type Inter interface {
	Int64() int64
}

type Floater interface {
	Float64() float64
}

type Booler interface {
	Bool() bool
}

func TryString(v interface{}) (string, bool) {
	if v == nil {
		return "", false
	}
	switch d := v.(type) {
	case string:
		return d, true
	case *string:
		return *d, true
	case sql.NullString:
		return d.String, true
	case *sql.NullString:
		return d.String, true
	case json.Number:
		return d.String(), true
	case *json.Number:
		return d.String(), true
	case Stringer:
		return d.String(), true
	case []byte:
		return string(d), true
	case rune:
		return string(d), true
	case []rune:
		return string(d), true
	}
	return "", false
}

func String(v interface{}, args ...string) string {
	if v == nil {
		return stringGetDefault(args)
	}
	if d, ok := TryString(v); ok {
		return d
	}
	switch d := v.(type) {
	case int:
		return strconv.FormatInt(int64(d), 10)
	case *int:
		return strconv.FormatInt(int64(*d), 10)
	case int8:
		return strconv.FormatInt(int64(d), 10)
	case *int8:
		return strconv.FormatInt(int64(*d), 10)
	case int16:
		return strconv.FormatInt(int64(d), 10)
	case *int16:
		return strconv.FormatInt(int64(*d), 10)
	case int32:
		return strconv.FormatInt(int64(d), 10)
	case *int32:
		return strconv.FormatInt(int64(*d), 10)
	case int64:
		return strconv.FormatInt(d, 10)
	case *int64:
		return strconv.FormatInt(*d, 10)
	case uint:
		return strconv.FormatInt(int64(d), 10)
	case *uint:
		return strconv.FormatInt(int64(*d), 10)
	case uint8:
		return strconv.FormatInt(int64(d), 10)
	case *uint8:
		return strconv.FormatInt(int64(*d), 10)
	case uint16:
		return strconv.FormatInt(int64(d), 10)
	case *uint16:
		return strconv.FormatInt(int64(*d), 10)
	case uint32:
		return strconv.FormatInt(int64(d), 10)
	case *uint32:
		return strconv.FormatInt(int64(*d), 10)
	case uint64:
		return strconv.FormatInt(int64(d), 10)
	case *uint64:
		return strconv.FormatInt(int64(*d), 10)
	case float32:
		return strconv.FormatFloat(float64(d), 'f', -1, 32)
	case *float32:
		return strconv.FormatFloat(float64(*d), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(d, 'f', -1, 64)
	case *float64:
		return strconv.FormatFloat(*d, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	case *bool:
		return strconv.FormatBool(*d)
	case sql.NullInt64:
		if d.Valid {
			return strconv.FormatInt(d.Int64, 10)
		}
	case *sql.NullInt64:
		if d.Valid {
			return strconv.FormatInt(d.Int64, 10)
		}
	case sql.NullFloat64:
		if d.Valid {
			return strconv.FormatFloat(d.Float64, 'f', -1, 64)
		}
	case *sql.NullFloat64:
		if d.Valid {
			return strconv.FormatFloat(d.Float64, 'f', -1, 64)
		}
	case sql.NullBool:
		if d.Valid {
			return strconv.FormatBool(d.Bool)
		}
	case *sql.NullBool:
		if d.Valid {
			return strconv.FormatBool(d.Bool)
		}
	case Inter:
		return strconv.FormatInt(d.Int64(), 10)
	case Floater:
		return strconv.FormatFloat(d.Float64(), 'f', -1, 64)
	case Booler:
		return strconv.FormatBool(d.Bool())
	}
	return stringGetDefault(args)
}

func Int(v interface{}, args ...int64) int64 {
	if v == nil {
		return intGetDefault(args)
	}
	switch d := v.(type) {
	case int:
		return int64(d)
	case *int:
		return int64(*d)
	case int8:
		return int64(d)
	case *int8:
		return int64(*d)
	case int16:
		return int64(d)
	case *int16:
		return int64(*d)
	case int32:
		return int64(d)
	case *int32:
		return int64(*d)
	case int64:
		return d
	case *int64:
		return *d
	case uint:
		return int64(d)
	case *uint:
		return int64(*d)
	case uint8:
		return int64(d)
	case *uint8:
		return int64(*d)
	case uint16:
		return int64(d)
	case *uint16:
		return int64(*d)
	case uint32:
		return int64(d)
	case *uint32:
		return int64(*d)
	case uint64:
		return int64(d)
	case *uint64:
		return int64(*d)
	case float32:
		return int64(d)
	case *float32:
		return int64(*d)
	case float64:
		return int64(d)
	case *float64:
		return int64(*d)
	case bool:
		if d {
			return 1
		} else {
			return 0
		}
	case *bool:
		if *d {
			return 1
		} else {
			return 0
		}
	case sql.NullInt64:
		if d.Valid {
			return d.Int64
		}
		return intGetDefault(args)
	case *sql.NullInt64:
		if d.Valid {
			return d.Int64
		}
		return intGetDefault(args)
	case sql.NullFloat64:
		if d.Valid {
			return int64(d.Float64)
		}
		return intGetDefault(args)
	case *sql.NullFloat64:
		if d.Valid {
			return int64(d.Float64)
		}
		return intGetDefault(args)
	case sql.NullBool:
		if d.Valid {
			if d.Bool {
				return 1
			} else {
				return 0
			}
		}
		return intGetDefault(args)
	case *sql.NullBool:
		if d.Valid {
			if d.Bool {
				return 1
			} else {
				return 0
			}
		}
		return intGetDefault(args)
	case json.Number:
		if n, err := d.Int64(); err == nil {
			return n
		}
	case *json.Number:
		if n, err := d.Int64(); err == nil {
			return n
		}
	case Inter:
		return d.Int64()
	case Floater:
		return int64(d.Float64())
	case Booler:
		if d.Bool() {
			return 1
		} else {
			return 0
		}
	}
	if s, ok := TryString(v); ok && s != "" {
		if n, err := strconv.ParseInt(s, 10, 64); err == nil {
			return n
		}
	}
	return intGetDefault(args)
}

func Float(v interface{}, args ...float64) float64 {
	if v == nil {
		return floatGetDefault(args)
	}
	switch d := v.(type) {
	case int:
		return float64(d)
	case *int:
		return float64(*d)
	case int8:
		return float64(d)
	case *int8:
		return float64(*d)
	case int16:
		return float64(d)
	case *int16:
		return float64(*d)
	case int32:
		return float64(d)
	case *int32:
		return float64(*d)
	case int64:
		return float64(d)
	case *int64:
		return float64(*d)
	case uint:
		return float64(d)
	case *uint:
		return float64(*d)
	case uint8:
		return float64(d)
	case *uint8:
		return float64(*d)
	case uint16:
		return float64(d)
	case *uint16:
		return float64(*d)
	case uint32:
		return float64(d)
	case *uint32:
		return float64(*d)
	case uint64:
		return float64(d)
	case *uint64:
		return float64(*d)
	case float32:
		return float64(d)
	case *float32:
		return float64(*d)
	case float64:
		return d
	case *float64:
		return *d
	case bool:
		if d {
			return 1
		} else {
			return 0
		}
	case *bool:
		if *d {
			return 1
		} else {
			return 0
		}
	case sql.NullInt64:
		if d.Valid {
			return float64(d.Int64)
		}
		return floatGetDefault(args)
	case *sql.NullInt64:
		if d.Valid {
			return float64(d.Int64)
		}
		return floatGetDefault(args)
	case sql.NullFloat64:
		if d.Valid {
			return d.Float64
		}
		return floatGetDefault(args)
	case *sql.NullFloat64:
		if d.Valid {
			return d.Float64
		}
		return floatGetDefault(args)
	case sql.NullBool:
		if d.Valid {
			if d.Bool {
				return 1
			} else {
				return 0
			}
		}
		return floatGetDefault(args)
	case *sql.NullBool:
		if d.Valid {
			if d.Bool {
				return 1
			} else {
				return 0
			}
		}
		return floatGetDefault(args)
	case json.Number:
		if n, err := d.Float64(); err == nil {
			return n
		}
		return floatGetDefault(args)
	case *json.Number:
		if n, err := d.Float64(); err == nil {
			return n
		}
		return floatGetDefault(args)
	case Inter:
		return float64(d.Int64())
	case Floater:
		return d.Float64()
	case Booler:
		if d.Bool() {
			return 1
		} else {
			return 0
		}
	}
	if s, ok := TryString(v); ok && s != "" {
		if n, err := strconv.ParseFloat(s, 64); err == nil {
			return n
		}
	}
	return floatGetDefault(args)
}

func Bool(v interface{}, args ...bool) bool {
	if v != nil {
		switch d := v.(type) {
		case bool:
			return d
		case *bool:
			return *d
		case int:
			return d != 0
		case *int:
			return *d != 0
		case int8:
			return d != 0
		case *int8:
			return *d != 0
		case int16:
			return d != 0
		case *int16:
			return *d != 0
		case int32:
			return d != 0
		case *int32:
			return *d != 0
		case int64:
			return d != 0
		case *int64:
			return *d != 0
		case uint:
			return d != 0
		case *uint:
			return *d != 0
		case uint8:
			return d != 0
		case *uint8:
			return *d != 0
		case uint16:
			return d != 0
		case *uint16:
			return *d != 0
		case uint32:
			return d != 0
		case *uint32:
			return *d != 0
		case uint64:
			return d != 0
		case *uint64:
			return *d != 0
		case float32:
			return d != 0
		case *float32:
			return *d != 0
		case float64:
			return d != 0
		case *float64:
			return *d != 0
		case sql.NullInt64:
			return d.Int64 != 0
		case *sql.NullInt64:
			return d.Int64 != 0
		case sql.NullFloat64:
			return d.Float64 != 0
		case *sql.NullFloat64:
			return d.Float64 != 0
		case sql.NullBool:
			return d.Bool
		case *sql.NullBool:
			return d.Bool
		case json.Number:
			if n, err := strconv.ParseBool(d.String()); err == nil {
				return n
			}
		case *json.Number:
			if n, err := strconv.ParseBool(d.String()); err == nil {
				return n
			}
		case Inter:
			return d.Int64() != 0
		case Floater:
			return d.Float64() != 0
		case Booler:
			return d.Bool()
		}
		if s, ok := TryString(v); ok && s != "" {
			if n, err := strconv.ParseBool(s); err == nil {
				return n
			}
		}
	}
	if len(args) > 0 {
		return args[0]
	}
	return false
}

func stringGetDefault(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func intGetDefault(args []int64) int64 {
	if len(args) > 0 {
		return args[0]
	}
	return 0
}

func floatGetDefault(args []float64) float64 {
	if len(args) > 0 {
		return args[0]
	}
	return 0
}
