package keyservice

import (
	context "context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v1beta1"
)

type K8sServer struct {
}

func (K8sServer) Version(ctx context.Context, req *v1beta1.VersionRequest) (*v1beta1.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (K8sServer) Decrypt(ctx context.Context, req *v1beta1.DecryptRequest) (*v1beta1.DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (K8sServer) Encrypt(ctx context.Context, req *v1beta1.EncryptRequest) (*v1beta1.EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
