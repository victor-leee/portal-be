package image

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/google/uuid"
	"github.com/victor-leee/portal-be/internal/config"
	"google.golang.org/protobuf/proto"
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
	BuildAndPush(base, buildFile string) (*string, error)
}

type Docker struct {
}

func (d *Docker) BuildAndPush(base, buildFile string) (*string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	ctx, _ := archive.TarWithOptions(base, &archive.TarOptions{})
	tagUID, _ := uuid.NewUUID()
	randomTag := tagUID.String()
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

	// TODO visualize the progress
	scanner := bufio.NewScanner(imageBuildResponse.Body)
	for scanner.Scan() {
		scanner.Text()
	}

	return proto.String(randomTag), d.push(context.Background(), cli, randomTag)
}

func (d *Docker) push(ctx context.Context, cli *client.Client, tag string) error {
	rd, err := cli.ImagePush(ctx, tag, types.ImagePushOptions{
		RegistryAuth: authCfgBase64Encoded,
	})
	if err != nil {
		return err
	}
	defer rd.Close()

	return nil
}
