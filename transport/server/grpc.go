package server

import (
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	Server      *grpc.Server
	Descriptors map[string]*desc.MethodDescriptor
	listener    net.Listener
	notify      chan error
}

//func NewGrpcServer(development bool) (*GrpcServer, error) {
//	grpcServer := new(GrpcServer)
//	grpcServer.Descriptors = make(map[string]*desc.MethodDescriptor)

//	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Config{}, config.Grpc.Port))
//	if err != nil {
//		return nil, err
//	}
//
//	opts := []grpc.ServerOption{
//		middlewares.New(
//			middlewares.GrpcRecovery(),
//			middlewares.MethodDescriptors(grpcServer.Descriptors),
//			middlewares.GrpcSentryPerformance(logger.GetSentryClient()),
//			middlewares.GrpcValidator(errHandler),
//			middlewares.GrpcJwt(permissionDescriptor, serviceInfo, errHandler, publicSecret, privateSecret, googleCaptchaSecretKey),
//			middlewares.GrpcLogger(logger),
//			grpc_prometheus.UnaryServerInterceptor,
//		),
//	}
//
//	if len(config.Grpc.CertFilePath) != 0 {
//		cred, err := config.LoadGrpcServerCredentials()
//		if err != nil {
//			return nil, err
//		}
//		opts = append(opts, grpc.Creds(cred))
//	}
//
//	srv := grpc.NewServer(opts...)
//
//	grpc_prometheus.Register(srv)
//
//	if development {
//		reflection.Register(srv)
//	}
//
//	grpcServer.listener = listener
//	grpcServer.Server = srv
//	grpcServer.notify = make(chan error)
//
//	return grpcServer, err
//}
//
//func (g *GrpcServer) Start() {
//	go func() {
//		g.notify <- g.Server.Serve(g.listener)
//		close(g.notify)
//	}()
//}
//
//func (g *GrpcServer) Notify() <-chan error {
//	return g.notify
//}
//
//func (g *GrpcServer) Shutdown() {
//	g.Server.GracefulStop()
//}
//
//func permissionDescriptor(ctx context.Context) ([]int32, bool, bool, bool) {
//	descriptor := new(desc.MethodDescriptor)
//	if d, ok := ctx.Value("desc").(*desc.MethodDescriptor); ok {
//		descriptor = d
//	}
//
//	perms := make([]int32, 0)
//	if permission, ok := proto.GetExtension(descriptor.GetMethodOptions(), ipminterPB.E_Permission).(*ipminterPB.Permission); ok {
//		if permission != nil {
//			for _, perm := range permission.Permissions {
//				perms = append(perms, int32(perm))
//			}
//			return perms, permission.Optional, permission.ValidatePermissions, permission.Captcha
//		}
//	}
//	return nil, false, false, false
//}
