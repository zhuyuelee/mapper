package mapper

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var zeroValue reflect.Value

const tagName string = "mapper"

func init() {
	zeroValue = reflect.Value{}
}

//Map to struct or Slice
func Map(source, target interface{}) (err error) {
	startTime := time.Now()
	defer func() {
		fmt.Println("mapper time=", time.Now().Sub(startTime).Seconds())
	}()
	defer func() {
		if mapErr := recover(); mapErr != nil {
			err = fmt.Errorf("map error %v", mapErr)
		}
	}()

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must pointer")
	}
	if targetValue.IsNil() {
		return errors.New("target must not nil")
	}
	var sources = reflect.ValueOf(source)
	if sources.Kind() == reflect.Ptr {
		sources = sources.Elem()
	}
	switch targetValue.Elem().Kind() {
	case reflect.Slice:
		return toSlice(sources, targetValue)
	case reflect.Struct:
		sourceMap, err := toMap(sources)
		if err == nil {
			return toStruct(sourceMap, targetValue, "")
		}
		return err
	}
	err = errors.New("data type only supported struct or Slice")

	return
}

//mapToSlice to Slice
func toSlice(source, target reflect.Value) error {
	//remove pointer
	target = target.Elem()
	targetType := target.Type()
	len := source.Len()
	targetSlice := reflect.MakeSlice(targetType, len, len)
	for i := 0; i < len; i++ {
		value := reflect.New(targetType.Elem())
		structMap, err := toMap(source.Index(i))
		if err == nil {
			toStruct(structMap, value, "")
			targetSlice.Index(i).Set(value.Elem())
		}
	}
	target.Set(targetSlice)
	return nil
}

//toStruct map to Struct
func toStruct(source map[string]reflect.Value, target reflect.Value, parentTag string) (err error) {
	//remove pointer
	target = target.Elem()
	for i := 0; i < target.NumField(); i++ {
		vField := target.Field(i)
		tField := target.Type().Field(i)
		tag, ok := tField.Tag.Lookup(tagName)

		if ok && tag == "-" {
			continue
		}
		if tField.Anonymous {
			nValue := reflect.New(tField.Type)
			toStruct(source, nValue, tag)
			vField.Set(nValue.Elem())
			return
		}
		if !ok {
			tag = tField.Name
		}
		if parentTag != "" {
			tag = fmt.Sprintf("%s.%s", parentTag, tag)
		}
		if value, ok := source[tag]; ok {
			sType := value.Type()
			vType := tField.Type

			if sType.Kind() == reflect.Ptr {
				sType = sType.Elem()
			}
			if vType.Kind() == reflect.Ptr {
				vType = vType.Elem()
			}
			//type of Struct
			if sType.Kind() == reflect.Struct && sType != vType {

				nValue := reflect.New(vType)
				childMap, err := toMap(value)
				if err == nil {
					toStruct(childMap, nValue, "")
					vField.Set(nValue.Elem())
				}
				//type of Slice
			} else if sType.Kind() == reflect.Slice && sType != vType {
				nValue := reflect.New(vType)
				toSlice(value, nValue)
				vField.Set(nValue.Elem())
			} else {
				vField.Set(value)
			}
		}
	}
	return
}

//toMap struct to map
func toMap(source reflect.Value) (map[string]reflect.Value, error) {
	m := make(map[string]reflect.Value)
	t := source.Type()

	if source == zeroValue {
		return nil, errors.New("no exists this value")
	}
	for i := 0; i < source.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tagName)
		if !ok || tag == "-" {
			tag = field.Name
		}
		//Anonymous Field
		if field.Anonymous {
			childMap, err := toMap(source.Field(i))
			if err == nil && len(childMap) > 0 {
				for key, val := range childMap {
					if ok {
						key = fmt.Sprintf("%s.%s", tag, key)
					}
					m[key] = val
				}
			}
		} else {
			m[tag] = source.Field(i)
		}
	}
	return m, nil
}
