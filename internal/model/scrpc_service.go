package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type RPCService struct {
	ID                 uint64 `json:"ID" gorm:"column:id"`
	Name               string `json:"name" gorm:"column:name"`
	ServiceKey         string `json:"serviceKey" gorm:"column:service_key"`
	IsService          bool   `json:"isService" gorm:"column:is_service"`
	ParentID           uint64 `json:"parentID" gorm:"column:parent_id"`
	UniqueCompletePath string `json:"completePath" gorm:"column:complete_path"`
	GitRepo            string `json:"gitRepo" gorm:"column:git_repo"`
	BuildFileRelPath   string `json:"buildFileRelPath" gorm:"column:build_file_rel_path"`
	Type               string `json:"type" gorm:"column:app_type"`
	// CustomPort is valid only if Type == config.AppTypeHTTP
	CustomPort uint16 `json:"customPort" gorm:"column:custom_port"`
	// PrefixMapping is valid only if Type == config.AppTypeHTTP
	PrefixMapping string `json:"prefixMapping" gorm:"column:prefix_mapping"`
}

func (m *RPCService) TableName() string {
	return "service_tab"
}

type RPCServiceDao interface {
	Insert(ctx context.Context, service *RPCService) error
	QueryByID(ctx context.Context, id uint64) (*RPCService, error)
	QueryByParentID(ctx context.Context, parentID uint64) ([]*RPCService, error)
}

type RPCServiceDaoImpl struct {
}

func (R *RPCServiceDaoImpl) QueryByID(ctx context.Context, id uint64) (*RPCService, error) {
	s := &RPCService{}
	if err := GetMysql(ctx).Model(&RPCService{}).Where(map[string]interface{}{
		"id": id,
	}).First(s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return s, nil
}

func (R *RPCServiceDaoImpl) QueryByParentID(ctx context.Context, parentID uint64) ([]*RPCService, error) {
	var services []*RPCService
	queryMap := map[string]interface{}{
		"parent_id": parentID,
	}
	if err := GetMysql(ctx).Model(&RPCService{}).Where(queryMap).Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (R *RPCServiceDaoImpl) Insert(ctx context.Context, service *RPCService) error {
	return GetMysql(ctx).Create(service).Error
}
