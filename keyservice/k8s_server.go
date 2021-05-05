package keyservice

import (
	context "context"

	"github.com/sirupsen/logrus"
	"go.mozilla.org/sops/v3/version"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	k8skmsapi "k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v1beta1"
)

const (
	/* Current version for the protocol interface definition. */
	kmsapiVersion = "v1beta1"
)

type K8sServer struct {
	Log *logrus.Logger
}

func (s K8sServer) Version(ctx context.Context, req *k8skmsapi.VersionRequest) (*k8skmsapi.VersionResponse, error) {
	s.Log.Infof("Received request for Version: %v", req)
	return &k8skmsapi.VersionResponse{Version: kmsapiVersion, RuntimeName: "sopsKMS", RuntimeVersion: version.Version}, nil
}

func (s K8sServer) Decrypt(ctx context.Context, req *k8skmsapi.DecryptRequest) (*k8skmsapi.DecryptResponse, error) {
	s.Log.Infof("Received request for Decrypt: %v", req)
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func (s K8sServer) Encrypt(ctx context.Context, req *k8skmsapi.EncryptRequest) (*k8skmsapi.EncryptResponse, error) {
	s.Log.Infof("Received request for Decrypt: %v", req)
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
