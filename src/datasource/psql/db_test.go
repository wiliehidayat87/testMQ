package db

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestInitPSQL(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitPSQL(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitPSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitPSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
