package main

import (
	"log"
	"strconv"
	"time"
)

func defaultUsers() {
	conn, err := driver.OpenPool()
	if err != nil {
		log.Printf("\nutil.go >> defaultUsers() >> driver.OpenPool() >> %v\n", err)
		return
	}
	defer conn.Close()

	_, err = conn.ExecNeo("merge (user:User { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} })", map[string]interface{}{
		"userId":        "0",
		"userFirstName": "admin",
		"userLastName":  "admin",
		"userEmail":     "admin",
		"userPassword":  "admin",
		"userActive":    true,
		"userRole":      "ADMIN",
	})

	if err != nil {
		log.Printf("\nutil.go >> defaultUsers() >> conn.ExecNeo() >> %v\n", err)
		return
	}
}

func NewId() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}

/*func FormToStruct(ptr interface{}, vals url.Values, start string) {
	var strct reflect.Value
	if reflect.TypeOf(ptr) == reflect.TypeOf(reflect.Value{}) {
		strct = ptr.(reflect.Value)
	} else {
		strct = reflect.ValueOf(ptr).Elem()
	}
	strctType := strct.Type()
	for i := 0; i < strct.NumField(); i++ {
		fld := strct.Field(i)
		if ok, v := GetVal(ToLowerFirst(start+strctType.Field(i).Name), vals); ok || fld.Kind() == reflect.Struct {
			switch fld.Kind() {
			case reflect.String:
				strct.Field(i).SetString(v)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				in, _ := strconv.ParseInt(v, 10, 64)
				strct.Field(i).SetInt(in)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				u, _ := strconv.ParseUint(v, 10, 64)
				strct.Field(i).SetUint(u)
			case reflect.Float32, reflect.Float64:
				f, _ := strconv.ParseFloat(v, 64)
				strct.Field(i).SetFloat(f)
			case reflect.Bool:
				b, _ := strconv.ParseBool(v)
				strct.Field(i).SetBool(b)
			case reflect.Map:
				strct.Field(i).Set(reflect.MakeMap(strct.Field(i).Type()))
			case reflect.Slice:
				ss := reflect.MakeSlice(strct.Field(i).Type(), 0, 0)
				strct.Field(i).Set(genSlice(ss, v))
			case reflect.Struct:
				//st := reflect.Indirect(reflect.New(strct.Field(i).Type()))
				st := reflect.Indirect(strct.Field(i))
				FormToStruct(st, vals, start+ToLowerFirst(strctType.Field(i).Name)+".")
				strct.Field(i).Set(st)
			}
		}
	}
}

func genSlice(sl reflect.Value, val string) reflect.Value {
	vs := strings.Split(val, ",")
	for _, v := range vs {
		switch sl.Type().String() {
		case "[]string":
			sl = reflect.Append(sl, reflect.ValueOf(v))
		case "[]int":
			in, _ := strconv.ParseInt(v, 10, 0)
			sl = reflect.Append(sl, reflect.ValueOf(int(in)))
		case "[]int8":
			in, _ := strconv.ParseInt(v, 10, 8)
			sl = reflect.Append(sl, reflect.ValueOf(int8(in)))
		case "[]int16":
			in, _ := strconv.ParseInt(v, 10, 16)
			sl = reflect.Append(sl, reflect.ValueOf(int16(in)))
		case "[]int32":
			in, _ := strconv.ParseInt(v, 10, 32)
			sl = reflect.Append(sl, reflect.ValueOf(int32(in)))
		case "[]int64":
			in, _ := strconv.ParseInt(v, 10, 64)
			sl = reflect.Append(sl, reflect.ValueOf(int64(in)))
		case "[]uint":
			in, _ := strconv.ParseUint(v, 10, 0)
			sl = reflect.Append(sl, reflect.ValueOf(uint(in)))
		case "[]uint8":
			in, _ := strconv.ParseUint(v, 10, 8)
			sl = reflect.Append(sl, reflect.ValueOf(uint8(in)))
		case "[]uint16":
			in, _ := strconv.ParseUint(v, 10, 16)
			sl = reflect.Append(sl, reflect.ValueOf(uint16(in)))
		case "[]uint32":
			in, _ := strconv.ParseUint(v, 10, 32)
			sl = reflect.Append(sl, reflect.ValueOf(uint32(in)))
		case "[]uint64":
			in, _ := strconv.ParseUint(v, 10, 64)
			sl = reflect.Append(sl, reflect.ValueOf(uint64(in)))
		case "[]float32":
			in, _ := strconv.ParseFloat(v, 32)
			sl = reflect.Append(sl, reflect.ValueOf(float32(in)))
		case "[]float64":
			in, _ := strconv.ParseFloat(v, 64)
			sl = reflect.Append(sl, reflect.ValueOf(float64(in)))
		case "[]bool":
			b, _ := strconv.ParseBool(v)
			sl = reflect.Append(sl, reflect.ValueOf(b))
		}
	}
	return sl
}

func GetVal(key string, v url.Values) (bool, string) {
	if v == nil {
		return false, ""
	}
	vs, ok := v[key]
	if !ok || len(vs) == 0 {
		return false, ""
	}
	return true, vs[0]
}

func ToLowerFirst(s string) string {
	return strings.ToLower(string(s[0])) + s[1:len(s)]
}*/
