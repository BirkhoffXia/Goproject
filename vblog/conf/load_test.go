package conf_test

import (
	"testing"

	"gitlab.com/go-course-project/go13/vblog/conf"
)

func TestLoadFromFile(t *testing.T) {
	err := conf.LoadFromFile("etc/application.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadFromEnv(t *testing.T) {
	// os.Setenv("DATASOURCE_HOST", "PAIBO.aliyun.com")
	err := conf.LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}
