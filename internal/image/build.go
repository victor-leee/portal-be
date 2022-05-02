package image

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"strings"
	"time"
)

var authCfgBase64Encoded string

func MustInitBuildCfg(cfg *config.Config) {
	authCfg := types.AuthConfig{
		ServerAddress: cfg.K8SCfg.RegistryIP,
	}
	b, _ := json.Marshal(authCfg)
	authCfgBase64Encoded = base64.URLEncoding.EncodeToString(b)
}

type Processor interface {
	BuildAndPush(base, buildFile string) (interface{}, error)
}

type Docker struct {
}

func (d *Docker) BuildAndPush(base, buildFile string) (interface{}, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	ctx, _ := archive.TarWithOptions(base, &archive.TarOptions{})
	randomTag := fmt.Sprintf("%s:%d",
		strings.ToLower(base64.URLEncoding.EncodeToString([]byte(base+buildFile))), time.Now().UnixMilli())
	imageBuildResponse, err := cli.ImageBuild(context.Background(),
		ctx,
		types.ImageBuildOptions{
			Context:    ctx,
			Dockerfile: buildFile,
			Tags:       []string{randomTag},
			Remove:     true,
		})
	if err != nil {
		return nil, err
	}
	defer imageBuildResponse.Body.Close()

	res := ""
	scanner := bufio.NewScanner(imageBuildResponse.Body)
	for scanner.Scan() {
		res += scanner.Text()
		logrus.Infof("res:%s", res)
	}

	return res, d.push(context.Background(), cli, randomTag)
}

func (d *Docker) push(ctx context.Context, cli *client.Client, tag string) error {
	logrus.Info("enter push")
	rd, err := cli.ImagePush(ctx, tag, types.ImagePushOptions{
		RegistryAuth: authCfgBase64Encoded,
	})
	if err != nil {
		return err
	}
	defer rd.Close()

	return nil
}
