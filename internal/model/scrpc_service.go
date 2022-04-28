package model

import "context"

type RPCService struct {
	ID                 uint64 `json:"ID" gorm:"column:id"`
	Name               string `json:"name" gorm:"column:name"`
	ServiceKey         string `json:"serviceKey" gorm:"column:service_key"`
	IsService          bool   `json:"isService" gorm:"column:is_service"`
	ParentID           uint64 `json:"parentID" gorm:"column:parent_id"`
	UniqueCompletePath string `json:"completePath" gorm:"complete_path"`
}

type RPCServiceDao interface {
	Insert(ctx context.Context, service *RPCService) error
	QueryByParentID(ctx context.Context, parentID uint64) ([]*RPCService, error)
}

type RPCServiceDaoImpl struct {
}

func (R *RPCServiceDaoImpl) QueryByParentID(ctx context.Context, parentID uint64) ([]*RPCService, error) {
	var services []*RPCService
	queryMap := map[string]interface{}{
		"parent_id": parentID,
	}
	if err := GetMysql(ctx).Model(&RPCService{}).Find(&services).Where(queryMap).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (R *RPCServiceDaoImpl) Insert(ctx context.Context, service *RPCService) error {
	return GetMysql(ctx).Create(service).Error
}
