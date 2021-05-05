package keyservice_k8s

import (
	context "context"

	"github.com/sirupsen/logrus"
	"go.mozilla.org/sops/v3/aes"
	"go.mozilla.org/sops/v3/cmd/sops/common"
	"go.mozilla.org/sops/v3/cmd/sops/formats"
	"go.mozilla.org/sops/v3/stores/json"
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

	tree, err := common.StoreForFormat(formats.Yaml).LoadEncryptedFile(req.Cipher)
	if err != nil {
		s.Log.Errorf("Cannot load data: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot load data: %v", err)
	}

	_, err = common.DecryptTree(common.DecryptTreeOpts{
		Cipher:    aes.NewCipher(),
		IgnoreMac: false,
		Tree:      &tree,
	})
	if err != nil {
		s.Log.Errorf("Cannot decrypt data: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot decrypt data: %v", err)
	}

	outputStore := &json.BinaryStore{}
	data, err := outputStore.EmitPlainFile(tree.Branches)
	if err != nil {
		s.Log.Errorf("Cannot emit plain file: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot emit plain file: %v", err)
	}

	return &k8skmsapi.DecryptResponse{Plain: data}, nil
}

func (s K8sServer) Encrypt(ctx context.Context, req *k8skmsapi.EncryptRequest) (*k8skmsapi.EncryptResponse, error) {
	s.Log.Infof("Received request for Decrypt: %v", req)
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
