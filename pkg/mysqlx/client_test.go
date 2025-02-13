package mysqlx

import (
	"testing"

	"github.com/andys920605/meme-coin/pkg/conf"
)

func TestNewClient(t *testing.T) {
	cfg := &conf.Config{}
	cfg.MySQL.Host = "localhost"
	cfg.MySQL.Database = "hr"
	cfg.MySQL.Username = "root"
	cfg.MySQL.Password = "root"
	cfg.MySQL.Port = 3306
	cfg.MySQL.MaxIdle = 5
	cfg.MySQL.MaxOpen = 10

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("連線資料庫失敗: %v", err)
	}

	var version string
	if err := client.Raw("SELECT VERSION()").Scan(&version).Error; err != nil {
		t.Fatalf("查詢 MySQL 版本失敗: %v", err)
	}
	t.Logf("連線成功，MySQL 版本：%s", version)

	sqlDB, err := client.DB.DB()
	if err != nil {
		t.Fatalf("取得底層 sql.DB 失敗: %v", err)
	}
	sqlDB.Close()
}
