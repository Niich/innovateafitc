package modelsx

import (
	"reflect"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func copyToMatchingFields(in interface{}, out interface{}) {

	refIn := reflect.ValueOf(in)
	refOut := reflect.ValueOf(out)
	x := refIn.Elem()
	s := refOut.Elem()
	for fi := 0; fi < s.NumField(); fi++ {
		f := s.Field(fi)
		if s.Field(fi).Kind() != reflect.Interface || s.Field(fi).Kind() != reflect.Ptr {
			if x.FieldByName(s.Type().Field(fi).Name).IsValid() {
				// fmt.Println("Name: ", s.Type().Field(fi).Name, " Val: ", x.FieldByName(s.Type().Field(fi).Name).Interface())
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						// change value of N
						f.Set(x.FieldByName(s.Type().Field(fi).Name))
					}
				}
			}
		}
	}
}

func ParseToObj(c *gin.Context, obj interface{}) error {

	data, err := c.GetRawData()

	if err != nil {
		return err
	}

	api := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 string("in"),
	}.Froze()

	if err := api.Unmarshal(data, obj); err != nil {
		return err
	}

	return nil
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
