package Register

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	xerrors "github.com/pkg/errors"
)

func InitNacos() (naming_client.INamingClient, error) {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.35",
			Port:   8848,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "D:\\GoPro\\UserService\\Log\\nacos",
		CacheDir:            "D:\\GoPro\\UserService\\Log\\cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, xerrors.Wrapf(err, "Init Nacos Fail")
	}
	return namingClient, err
}
