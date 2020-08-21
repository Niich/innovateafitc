package modelsx

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func copyToMatchingFields(in interface{}, out interface{}) {

	reflectIn := reflect.ValueOf(in)
	reflectOut := reflect.ValueOf(out)

	for reflectOut.Kind() == reflect.Ptr || reflectOut.Kind() == reflect.Interface {
		reflectOut = reflectOut.Elem()
	}
	for reflectIn.Kind() == reflect.Ptr || reflectIn.Kind() == reflect.Interface {
		reflectIn = reflectIn.Elem()
	}

	if reflectOut.IsValid() && reflectIn.IsValid() {
		for fi := 0; fi < reflectOut.NumField(); fi++ {
			f := reflectOut.Field(fi)
			if reflectOut.Field(fi).Kind() != reflect.Interface && reflectOut.Field(fi).Kind() != reflect.Ptr {
				if reflectIn.FieldByName(reflectOut.Type().Field(fi).Name).IsValid() {
					if f.IsValid() {
						// A Value can be changed only if it is
						// addressable and was not obtained by
						// the use of unexported struct fields.
						if f.CanSet() {
							// change value of N
							valueToSet := reflectIn.FieldByName(reflectOut.Type().Field(fi).Name)
							if !valueToSet.IsZero() {
								f.Set(reflectIn.FieldByName(reflectOut.Type().Field(fi).Name))
							}
						}
					}
				} else if reflectIn.FieldByName("R").IsValid() {
					// add something for fileds that are realations but not pointers
				}
			} else if reflectOut.Field(fi).Kind() == reflect.Ptr {
				if reflectIn.FieldByName("R").IsValid() && !reflectIn.FieldByName("R").IsZero() {
					nestRelations(in, out)
				}
			}
		}
	}
}

func nestRelations(in interface{}, out interface{}) {
	s := reflect.ValueOf(in)
	x := reflect.ValueOf(out)

	for s.Kind() == reflect.Ptr || s.Kind() == reflect.Interface {
		s = s.Elem()
	}
	for x.Kind() == reflect.Ptr || x.Kind() == reflect.Interface {
		x = x.Elem()
	}

	if s.Kind() == reflect.Struct && x.Kind() == reflect.Struct {
		sRels := s.FieldByName("R")
		if sRels.IsValid() {
			for sRels.Kind() == reflect.Ptr {
				sRels = sRels.Elem()
			}
			if sRels.IsValid() && sRels.Kind() == reflect.Struct {
				for fi := 0; fi < sRels.NumField(); fi++ {
					xField := x.FieldByName(sRels.Type().Field(fi).Name)
					if xField.CanSet() {
						val := sRels.Field(fi).Interface()
						var fldType reflect.Type
						var reflectRes reflect.Value
						fldType = x.FieldByName(sRels.Type().Field(fi).Name).Type()
						if x.FieldByName(sRels.Type().Field(fi).Name).Kind() == reflect.Ptr {
							res := reflect.New(fldType.Elem())
							copyToMatchingFields(val, res.Interface())
							reflectRes = reflect.ValueOf(res.Interface())

						} else {
							res := reflect.New(fldType)
							copyToMatchingFields(val, res.Interface())
							reflectRes = reflect.ValueOf(res.Elem().Interface())
						}
						if !reflectRes.IsZero() {
							xField.Set(reflectRes)
						}
					}
				}
			}
		}
	}
}

func unNestRelations(in interface{}, out interface{}) {
	s := reflect.ValueOf(in)
	x := reflect.ValueOf(out)

	for s.Kind() == reflect.Ptr || s.Kind() == reflect.Interface {
		s = s.Elem()
	}
	for x.Kind() == reflect.Ptr || x.Kind() == reflect.Interface {
		x = x.Elem()
	}

	if s.Kind() == reflect.Struct && x.Kind() == reflect.Struct {
		sRels := s.FieldByName("R")
		if sRels.IsValid() {
			for sRels.Kind() == reflect.Ptr {
				sRels = sRels.Elem()
			}
			if sRels.IsValid() && sRels.Kind() == reflect.Struct {
				for fi := 0; fi < sRels.NumField(); fi++ {
					xField := x.FieldByName(sRels.Type().Field(fi).Name)
					if xField.CanSet() {
						val := sRels.Field(fi).Interface()
						var fldType reflect.Type
						var res reflect.Value
						fldType = x.FieldByName(sRels.Type().Field(fi).Name).Type()
						if x.FieldByName(sRels.Type().Field(fi).Name).Kind() == reflect.Ptr {
							res = reflect.New(fldType.Elem())
							copyToMatchingFields(val, res.Interface())
							reflectRes := reflect.ValueOf(res.Interface())
							xField.Set(reflectRes)

						} else {
							res = reflect.New(fldType)
							copyToMatchingFields(val, res.Interface())
							reflectRes := reflect.ValueOf(res.Elem().Interface())
							xField.Set(reflectRes)
						}
					} else {
						println("Cant Set: ", sRels.Type().Field(fi).Name)
					}
				}
			} else {
				println("not Valid: ", sRels.Kind())
			}
		} else {
			println("Rels not Valid: ", sRels.Kind())
			for fi := 0; fi < s.NumField(); fi++ {
				println("Rels Fields: ", s.Type().Field(fi).Name)
			}
		}
	} else {
		println("not struct pointers, cant relate")
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

func SendModelsxObj(c *gin.Context, obj interface{}) error {
	api := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 string("out"),
	}.Froze()

	data, err := api.Marshal(obj)

	if err != nil {
		return err
	}

	c.Data(http.StatusOK, "application/json", data)

	return nil
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
