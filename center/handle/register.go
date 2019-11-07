package handle

import (
	"net/http"

	"github.com/chaos007/easy/center/handle/controller"
)

// Init Init
func Init() {
	// grpc.RegisterProtocol("pb.UpGetConfigToCenter", controller.UpGetConfigToCenter)

}

// HTTPInit 后台
func HTTPInit() {
	http.HandleFunc("/BackendLogin", controller.BackendLogin)
	http.HandleFunc("/UpGetBackendLoginInfo", controller.UpGetBackendLoginInfo)
}
