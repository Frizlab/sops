package keyservice_k8s

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.mozilla.org/sops/v3"
	"go.mozilla.org/sops/v3/keyservice"
	"go.mozilla.org/sops/v3/keyservice_k8s"
	"go.mozilla.org/sops/v3/logging"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	k8skmsapi "k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v1beta1"
)

var log *logrus.Logger

func init() {
	log = logging.NewLogger("KEYSERVICE-K8S")
}

// Opts are the options the key service server can take
type Opts struct {
	Path        string
	KeyGroups   []sops.KeyGroup
	KeyServices []keyservice.KeyServiceClient
}

// Run runs a SOPS key service server
func Run(opts Opts) error {
	lis, err := net.Listen("unix" /* Only network supported by k8s for now */, opts.Path)
	if err != nil {
		return err
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	k8skmsapi.RegisterKeyManagementServiceServer(grpcServer, keyservice_k8s.K8sServer{
		Log:         log,
		KeyGroups:   opts.KeyGroups,
		KeyServices: opts.KeyServices,
	})
	log.Infof("Listening on unix://%s", opts.Path)

	// Close socket if we get killed
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		log.Infof("Caught signal %s: shutting down.", sig)
		lis.Close()
		os.Exit(0)
	}(sigc)
	return grpcServer.Serve(lis)
}
