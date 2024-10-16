package model

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/shunsuke-kawata/socket-learning/logic"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID         uint           `gorm:"primaryKey"`
	StatusName string         `gorm:"size:50;not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"default:null"` // デフォルトで NULL
}

type Address struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	IPAddress string         `gorm:"size:39;not null;unique"` // IPアドレスに変更
	ColorCode string         `gorm:"size:7;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"default:null"` // デフォルトで NULL
}

type Task struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"size:100;not null"`
	Description string         `gorm:"type:text"`
	StatusID    uint           `gorm:"not null"`             // 外部キー
	Status      Status         `gorm:"foreignKey:StatusID"`  // リレーション
	AddressID   uint           `gorm:"not null"`             // 外部キー
	Address     Address        `gorm:"foreignKey:AddressID"` // リレーション
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"default:null"` // デフォルトで NULL
}

var db *gorm.DB
var err error

func init() {
	mysqlRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbContainerName := os.Getenv("DB_CONTAINER_NAME")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", mysqlRootPassword, dbContainerName, mysqlPort, mysqlDatabase)

	maxRetries := 10                 // リトライ回数
	retryInterval := 2 * time.Second // リトライ間隔

	//データベースの接続をリトライする
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("database access succeeded")
			break
		}

		// 最大リトライ回数に達した場合
		if i == maxRetries-1 {
			fmt.Println("maximum retry attempts reached, exiting.")
			return
		}
		// リトライ間隔を待つ
		time.Sleep(retryInterval)
	}

}

// ステータス新規作成
func CreateStatus(statusName string) (*Status, error) {
	newStatus := &Status{StatusName: statusName}
	if err := db.Debug().Create(newStatus).Error; err != nil {
		return nil, err
	}
	return newStatus, nil
}

// IPアドレス新規登録
func CreateAddress(iPAddress string) (*Address, error) {
	fmt.Println("model.address", iPAddress)
	address := &Address{}
	// IPアドレスでAddressを検索
	if err := db.First(address, "ip_address = ?", iPAddress).Error; err != nil {
		// アドレスが存在しない場合、新規にレコードを追加
		colorCodeString := logic.CreateRandomColorCode()
		newAddress := &Address{IPAddress: iPAddress, ColorCode: colorCodeString}
		if err := db.Debug().Create(newAddress).Error; err != nil {
			return nil, err
		}
		return newAddress, nil
	}
	// アドレスがすでに存在する場合、そのアドレスを返す
	return address, nil
}

// タスク新規登録
func CreateTask(title string, description string, iPAddress string) (*Task, error) {
	fmt.Println("model.task", iPAddress)
	address, err := CreateAddress(iPAddress)
	if err != nil {
		return nil, err
	}

	if address == nil {
		return nil, errors.New("failed to create or retrieve address")
	}

	status := &Status{}
	if err := db.First(&status, "status_name = ?", "Pending").Error; err != nil {
		return nil, err
	}

	newTask := &Task{
		Title:       title,
		Description: description,
		StatusID:    status.ID,
		AddressID:   address.ID,
	}

	if err := db.Debug().Create(newTask).Error; err != nil {
		return nil, err
	}

	return newTask, nil
}

func ReadStatus() ([]Status, error) {
	statuses := []Status{}

	//mysqlからデータ一覧の取得
	if err := db.Debug().Find(&statuses).Error; err != nil {
		return nil, err
	}

	return statuses, nil
}

func ReadAddress() ([]Address, error) {
	addresses := []Address{}
	if err := db.Debug().Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func ReadAddressByColorCode(colorCode string) (Address, error) {
	address := Address{}
	if err := db.First(&address, "color_code = ?", colorCode).Error; err != nil {
		return address, err
	}
	return address, nil
}

func ReadAddressByIPAddress(ipAddress string) (Address, error) {
	address := Address{}
	if err := db.First(&address, "ip_address = ?", ipAddress).Error; err != nil {
		return address, err
	}
	return address, nil
}

func ReadTask() ([]Task, error) {
	tasks := []Task{}
	if err := db.Debug().Preload("Status").Preload("Address").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
