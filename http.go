package goopenmeteo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func httpGet(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
	} else if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	return io.ReadAll(resp.Body)
}

func urlValues(s interface{}) url.Values {
	values := url.Values{}
	if s == nil {
		return values
	}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return values
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return values
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue
		}
		parts := strings.Split(tag, ",")
		name := parts[0]
		omitEmpty := len(parts) > 1 && parts[1] == "omitempty"
		var strValue string
		isSet := true
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
			isSet = strValue != ""
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if value.Int() != 0 || !omitEmpty {
				strValue = strconv.FormatInt(value.Int(), 10)
			} else {
				isSet = false
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if value.Uint() != 0 || !omitEmpty {
				strValue = strconv.FormatUint(value.Uint(), 10)
			} else {
				isSet = false
			}
		case reflect.Float32, reflect.Float64:
			if value.Float() != 0 || !omitEmpty {
				strValue = strconv.FormatFloat(value.Float(), 'f', -1, 64)
			} else {
				isSet = false
			}
		case reflect.Bool:
			if value.Bool() || !omitEmpty {
				strValue = strconv.FormatBool(value.Bool())
			} else {
				isSet = false
			}
		case reflect.Ptr:
			if value.IsNil() {
				isSet = false
			} else {
				elem := value.Elem()
				switch elem.Kind() {
				case reflect.Bool:
					strValue = strconv.FormatBool(elem.Bool())
				case reflect.String:
					strValue = elem.String()
					isSet = strValue != ""
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					strValue = strconv.FormatInt(elem.Int(), 10)
				default:
					isSet = false
				}
			}
		case reflect.Slice:
			if value.Len() > 0 || !omitEmpty {
				var sliceValues []string
				for j := 0; j < value.Len(); j++ {
					elem := value.Index(j)
					var elemStr string
					switch elem.Kind() {
					case reflect.String:
						elemStr = elem.String()
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						elemStr = strconv.FormatInt(elem.Int(), 10)
					}
					if elemStr != "" {
						sliceValues = append(sliceValues, elemStr)
					}
				}
				if len(sliceValues) > 0 {
					strValue = strings.Join(sliceValues, ",")
					isSet = true
				} else {
					isSet = false
				}
			} else {
				isSet = false
			}
		case reflect.Struct:
			if value.Type() == reflect.TypeOf(time.Time{}) {
				t := value.Interface().(time.Time)
				if !t.IsZero() || !omitEmpty {
					strValue = t.Format(time.RFC3339)
				} else {
					isSet = false
				}
			}
		}
		if isSet && strValue != "" {
			values.Add(name, strValue)
		}
	}
	return values
}

func urlQuery(s interface{}) string {
	values := urlValues(s)
	if len(values) == 0 {
		return ""
	}

	var buf strings.Builder
	first := true
	for k, vs := range values {
		for _, v := range vs {
			if !first {
				buf.WriteByte('&')
			}
			first = false
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			parts := strings.Split(v, ",")
			for i, part := range parts {
				if i > 0 {
					buf.WriteByte(',')
				}
				buf.WriteString(url.QueryEscape(part))
			}
		}
	}
	return buf.String()
}
