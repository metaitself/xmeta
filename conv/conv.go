package conv

import (
	"fmt"
	"time"
)

// ToPtr convert to pointer
func ToPtr[E any](t E) *E {
	return &t
}

// ToAny converts one type to another type.
func ToAny[T any](a any) T {
	v, _ := ToAnyE[T](a)
	return v
}

// ToAnyE converts one type to another and returns an error if error occurred.
func ToAnyE[T any](a any) (T, error) {
	var t T
	switch any(t).(type) {
	case bool:
		v, err := ToBoolE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int:
		v, err := ToIntE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int8:
		v, err := ToInt8E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int16:
		v, err := ToInt16E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int32:
		v, err := ToInt32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int64:
		v, err := ToInt64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint:
		v, err := ToUintE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint8:
		v, err := ToUint8E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint16:
		v, err := ToUint16E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint32:
		v, err := ToUint32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint64:
		v, err := ToUint64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float32:
		v, err := ToFloat32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float64:
		v, err := ToFloat64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case time.Duration:
		v, err := ToDurationE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case string:
		v, err := ToStringE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	default:
		return t, fmt.Errorf("the type %T isn't supported", t)
	}
	return t, nil
}

// ToBool casts an interface to a bool type.
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToTime casts an interface to a time.Time type.
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

func ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time {
	v, _ := ToTimeInDefaultLocationE(i, location)
	return v
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToUint casts an interface to a uint type.
func ToUint(i interface{}) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i interface{}) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i interface{}) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint8 casts an interface to a uint8 type.
func ToUint8(i interface{}) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToString casts an interface to a string type.
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}
