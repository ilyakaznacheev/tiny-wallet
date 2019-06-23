// Package config contains application configuration structures and data read logic
package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fileData := `
server:
  host: "localhost"
  port: "8080"

database:
  host: "localhost"
  port: "5433"
  database: "wallet"
  username: "postgres"
  password: "postgres"
  ssl: "disable"
  conn-wait: yes
  conn-pool: 5`

	tests := []struct {
		name       string
		conf       string
		want       *MainConfig
		createFile bool
		wantErr    bool
	}{
		{
			name: "valid file",
			conf: fileData,
			want: &MainConfig{
				Server: ServerConfig{
					Host: "localhost",
					Port: "8080",
				},
				Database: DatabaseConfig{
					Host:           "localhost",
					Port:           "5433",
					Database:       "wallet",
					Username:       "postgres",
					Password:       "postgres",
					SSL:            "disable",
					ConnectionWait: true,
					ConnectionPool: 5,
				},
			},

			createFile: true,
			wantErr:    false,
		},

		{
			name:       "invalid file",
			conf:       "12345",
			want:       nil,
			createFile: true,
			wantErr:    true,
		},
		{
			name:       "no file",
			conf:       "",
			want:       nil,
			createFile: false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		var (
			fileName string
			tmpFile  *os.File
			err      error
		)

		if tt.createFile {
			tmpFile, err = ioutil.TempFile(os.TempDir(), "readconfigtmp_")
			if err != nil {
				t.Fatal("cannot create temporary file:", err)
			}

			text := []byte(tt.conf)
			if _, err = tmpFile.Write(text); err != nil {
				t.Fatal("failed to write to temporary file:", err)
			}
			fileName = tmpFile.Name()
		} else {
			fileName = "test/path.yml"
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unexpected error behavior: error %v, but expected %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrong config list %+v, want %v", got, tt.want)
			}
		})

		if tt.createFile {
			os.Remove(tmpFile.Name())
		}
	}
}
