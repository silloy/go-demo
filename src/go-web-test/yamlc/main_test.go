package main

import (
	"reflect"
	"testing"
)

//func main() {
//	var c conf
//	conf := c.getConf()
//	fmt.Println(conf.Host)
//}

type fields struct {
	Host   string
	User   string
	Pwd    string
	Dbname string
}

func Test_conf_getConf(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   *conf
	}{
		{
			name: "acc",
			fields: struct {
				Host   string
				User   string
				Pwd    string
				Dbname string
			}{Host: "h", User: "b", Pwd: "c", Dbname: "d"},
			want: &conf{Host: "a", User: "b", Pwd: "c", Dbname: "d"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &conf{
				Host:   tt.fields.Host,
				User:   tt.fields.User,
				Pwd:    tt.fields.Pwd,
				Dbname: tt.fields.Dbname,
			}
			if got := c.getConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
