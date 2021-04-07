package Register

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
	xerrors "github.com/pkg/errors"
)

func RegisterService(namingClient naming_client.INamingClient) (bool, error) {
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.1.35",
		Port:        9600,
		ServiceName: "UserService",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
	if err != nil {
		return success, xerrors.Wrapf(err, "Register Fail")
	}

	return success, err
}
