package processor

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/victor-leee/portal-be/internal/config"
	"github.com/victor-leee/portal-be/internal/model"
	"strconv"
	"strings"
	"time"
)

type RPCServiceProcessor interface {
	Create(ctx context.Context, name string, hierarchyInfo []string, parentID uint64, isService bool) error
	Query(ctx context.Context, parentID uint64) ([]*model.RPCService, error)
}

type DefaultRPCServiceProcessor struct {
	ServiceDao model.RPCServiceDao
}

func (d *DefaultRPCServiceProcessor) Create(ctx context.Context, name string, hierarchyInfo []string, parentID uint64, isService bool) error {
	completePath := buildCompletePath(hierarchyInfo, name)
	serviceModel := &model.RPCService{
		Name:               name,
		IsService:          isService,
		ParentID:           parentID,
		UniqueCompletePath: completePath,
		ServiceKey:         buildServiceKey(completePath),
	}

	return d.ServiceDao.Insert(ctx, serviceModel)
}

func buildCompletePath(hierarchyInfo []string, service string) string {
	return service + config.ServiceHierarchySeparator + strings.Join(reverseStrings(hierarchyInfo), config.ServiceHierarchySeparator)
}

func reverseStrings(original []string) []string {
	for l, r := 0, len(original)-1; l < r; l, r = l+1, r-1 {
		original[l], original[r] = original[r], original[l]
	}

	return original
}

// buildServiceKey generates a digest based on complete path, current time in millis and a random number
// which are base64 encoded
func buildServiceKey(completePath string) string {
	randN, _ := rand.Int(rand.Reader, config.ServiceKeyMaxRandomNumber)
	keyComponents := []string{completePath, strconv.FormatInt(time.Now().UnixMilli(), 10), randN.String()}
	base64Encoded := []byte(base64.StdEncoding.EncodeToString([]byte(strings.Join(keyComponents, config.ServiceKeySeparator))))
	s := sha256.New()
	s.Write(base64Encoded)

	return string(s.Sum(nil))
}

func (d *DefaultRPCServiceProcessor) Query(ctx context.Context, parentID uint64) ([]*model.RPCService, error) {
	return d.ServiceDao.QueryByParentID(ctx, parentID)
}
