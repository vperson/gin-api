package tb

import (
	"errors"
	"gin-api/config"
	"gorm.io/gorm"
)

type Common struct {
	Model
	DataType string `json:"data_type" gorm:"index:idx_data_type_key_name;size:50"` // 数据类型
	KeyName  string `json:"key_name" gorm:"index:idx_data_type_key_name;size:100"` // 数据键
	Value    string `json:"value" gorm:"size:255"`                                 // 数据值
}

type CommonDB struct {
	db *gorm.DB
}

func NewCommon(db *gorm.DB) *CommonDB {
	return &CommonDB{db: db}
}

func (d *CommonDB) Create(record *Common) error {
	return d.db.Create(&record).Error
}

func (d *CommonDB) GetByTypeAndName(tp, name string) (*Common, error) {
	var record Common
	err := d.db.Where("data_type = ? AND key_name = ?", tp, name).First(&record).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &record, nil
}

// CreateIfNotExist 初始化的时候用
func (d *CommonDB) CreateIfNotExist() error {
	record, err := d.GetByTypeAndName(CommonTypeSystem, CommonKeySystemName)
	if err != nil {
		return err
	}

	if record.ID == 0 {
		newRecord := Common{
			DataType: CommonTypeSystem,
			KeyName:  CommonKeySystemName,
			Value:    config.Get().Name,
		}
		err = d.Create(&newRecord)
		if err != nil {
			return err
		}
	}

	return nil
}
