package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":8081")
	envoy_api_v2.RegisterRouteDiscoveryServiceServer(grpcServer, rds{})
	//envoy_api_v2.RegisterEndpointDiscoveryServiceServer(grpcServer, eds{})

	if err := grpcServer.Serve(lis); err != nil {
		// error handling
	}
}

type rds struct{}

func (r rds) StreamRoutes(ls envoy_api_v2.RouteDiscoveryService_StreamRoutesServer) error {

	go func() {
		m, err := ls.Recv()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Msg: %v \n", m)

	}()

	rc := []*envoy_api_v2.RouteConfiguration{
		&envoy_api_v2.RouteConfiguration{
			Name: "route_1",
			VirtualHosts: []*route.VirtualHost{
				&route.VirtualHost{
					Name: "server_1",
					Domains: []string{
						"*",
					},
					Routes: []*route.Route{
						&route.Route{
							Match: &route.RouteMatch{
								PathSpecifier: &route.RouteMatch_Prefix{Prefix: "/logout"},
							},
							Action: &route.Route_Route{Route: &route.RouteAction{
								ClusterSpecifier: &route.RouteAction_Cluster{Cluster: "some_service"},
							}},
						},
					},
				},
			},
		},
	}
 
	var resource []*any.Any

	for _, rca := range rc {
		data, err := proto.Marshal(rca)
		if err != nil {
			return err
		}

		resource = append(resource, &any.Any{
			TypeUrl: "type.googleapis.com/envoy.api.v2.RouteConfiguration",
			Value:   data,
		})
	}

	for {
		ls.Send(&envoy_api_v2.DiscoveryResponse{
			VersionInfo: "1",
			Resources:   resource,
			Canary:      false,
			TypeUrl:     "type.googleapis.com/envoy.api.v2.RouteConfiguration",
			Nonce:       time.Now().String(),
		})

		time.Sleep(10 * time.Second)
	}

}

func (r rds) DeltaRoutes(envoy_api_v2.RouteDiscoveryService_DeltaRoutesServer) error {
	panic("implement me")
}

func (r rds) FetchRoutes(context.Context, *envoy_api_v2.DiscoveryRequest) (*envoy_api_v2.DiscoveryResponse, error) {
	panic("implement me")
}

//type eds struct{}
//
//func (e eds) StreamEndpoints(ls envoy_api_v2.EndpointDiscoveryService_StreamEndpointsServer) error {
//	fmt.Println("stream")
//	ca := []*envoy_api_v2.ClusterLoadAssignment{
//		&envoy_api_v2.ClusterLoadAssignment{
//			ClusterName: "some_service",
//			Endpoints: []*endpoint.LocalityLbEndpoints{
//				&endpoint.LocalityLbEndpoints{
//					LbEndpoints: []*endpoint.LbEndpoint{
//						&endpoint.LbEndpoint{
//							HostIdentifier: &endpoint.LbEndpoint_Endpoint{
//								Endpoint: &endpoint.Endpoint{
//									Address: &core.Address{
//										Address: &core.Address_SocketAddress{
//											SocketAddress: &core.SocketAddress{
//												Protocol:      core.SocketAddress_TCP,
//												Address:       "192.168.0.8",
//												PortSpecifier: &core.SocketAddress_PortValue{PortValue: 3000},
//											},
//										},
//									},
//									HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//										PortValue: 3000,
//									},
//								},
//							},
//							HealthStatus: core.HealthStatus_UNKNOWN,
//						},
//						&endpoint.LbEndpoint{
//							HostIdentifier: &endpoint.LbEndpoint_Endpoint{
//								Endpoint: &endpoint.Endpoint{
//									Address: &core.Address{
//										Address: &core.Address_SocketAddress{
//											SocketAddress: &core.SocketAddress{
//												Protocol:      core.SocketAddress_TCP,
//												Address:       "192.168.0.18",
//												PortSpecifier: &core.SocketAddress_PortValue{PortValue: 80},
//											},
//										},
//									},
//									HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//										PortValue: 80,
//									},
//								},
//							},
//							HealthStatus: core.HealthStatus_UNKNOWN,
//						},
//					},
//				},
//			},
//			//NamedEndpoints: map[string]*endpoint.Endpoint{
//			//	"some_service": &endpoint.Endpoint{
//			//		Address: &core.Address{
//			//			Address: &core.Address_SocketAddress{
//			//				SocketAddress: &core.SocketAddress{
//			//					Protocol:      core.SocketAddress_TCP,
//			//					Address:       "192.168.0.8",
//			//					PortSpecifier: &core.SocketAddress_PortValue{PortValue: 3000},
//			//				},
//			//			},
//			//		},
//			//		HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//			//			PortValue: 3000,
//			//		},
//			//	},
//			//},
//			//Policy:&envoy_api_v2.ClusterLoadAssignment_Policy{
//			//
//			//},
//
//		},
//	}
//
//	var resources []*any.Any
//	for _, cLA := range ca {
//		_d, _ := json.Marshal(cLA)
//		fmt.Println(string(_d))
//		data, err := proto.Marshal(cLA)
//		if err != nil {
//			return err
//		}
//
//		resources = append(resources, &any.Any{
//			TypeUrl: "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment",
//			Value:   data,
//		})
//	}
//
//	for {
//		ls.Send(&envoy_api_v2.DiscoveryResponse{
//			VersionInfo: "0",
//			Resources:   resources,
//			Canary:      false,
//			TypeUrl:     "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment",
//			Nonce:       time.Now().String(),
//		})
//
//		time.Sleep(5 * time.Second)
//		fmt.Printf("Trigger \n")
//	}
//
//	return nil
//}
//
//func (e eds) DeltaEndpoints(ls envoy_api_v2.EndpointDiscoveryService_DeltaEndpointsServer) error {
//	ca := []*envoy_api_v2.ClusterLoadAssignment{
//		&envoy_api_v2.ClusterLoadAssignment{
//			ClusterName: "some_service",
//			Endpoints: []*endpoint.LocalityLbEndpoints{
//				&endpoint.LocalityLbEndpoints{
//					LbEndpoints: []*endpoint.LbEndpoint{
//						&endpoint.LbEndpoint{
//							HostIdentifier: &endpoint.LbEndpoint_Endpoint{
//								Endpoint: &endpoint.Endpoint{
//									Address: &core.Address{
//										Address: &core.Address_SocketAddress{
//											SocketAddress: &core.SocketAddress{
//												Protocol:      core.SocketAddress_TCP,
//												Address:       "192.168.0.8",
//												PortSpecifier: &core.SocketAddress_PortValue{PortValue: 3000},
//											},
//										},
//									},
//									HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//										PortValue: 3000,
//									},
//								},
//							},
//							HealthStatus: core.HealthStatus_UNKNOWN,
//						},
//						&endpoint.LbEndpoint{
//							HostIdentifier: &endpoint.LbEndpoint_Endpoint{
//								Endpoint: &endpoint.Endpoint{
//									Address: &core.Address{
//										Address: &core.Address_SocketAddress{
//											SocketAddress: &core.SocketAddress{
//												Protocol:      core.SocketAddress_TCP,
//												Address:       "192.168.0.18",
//												PortSpecifier: &core.SocketAddress_PortValue{PortValue: 80},
//											},
//										},
//									},
//									HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//										PortValue: 80,
//									},
//								},
//							},
//							HealthStatus: core.HealthStatus_UNKNOWN,
//						},
//					},
//				},
//			},
//			//NamedEndpoints: map[string]*endpoint.Endpoint{
//			//	"some_service": &endpoint.Endpoint{
//			//		Address: &core.Address{
//			//			Address: &core.Address_SocketAddress{
//			//				SocketAddress: &core.SocketAddress{
//			//					Protocol:      core.SocketAddress_TCP,
//			//					Address:       "192.168.0.8",
//			//					PortSpecifier: &core.SocketAddress_PortValue{PortValue: 3000},
//			//				},
//			//			},
//			//		},
//			//		HealthCheckConfig: &endpoint.Endpoint_HealthCheckConfig{
//			//			PortValue: 3000,
//			//		},
//			//	},
//			//},
//			//Policy:&envoy_api_v2.ClusterLoadAssignment_Policy{
//			//
//			//},
//
//		},
//	}
//
//	var resources []*any.Any
//	for _, cLA := range ca {
//		_d, _ := json.Marshal(cLA)
//		fmt.Println(string(_d))
//		data, err := proto.Marshal(cLA)
//		if err != nil {
//			return err
//		}
//
//		resources = append(resources, &any.Any{
//			TypeUrl: "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment",
//			Value:   data,
//		})
//	}
//
//	for {
//		time.Sleep(30 * time.Second)
//		fmt.Printf("Detla Trigger \n")
//		ls.Send(&envoy_api_v2.DeltaDiscoveryResponse{
//			SystemVersionInfo: "0",
//			Resources: []*envoy_api_v2.Resource{
//				&envoy_api_v2.Resource{
//					Name:     "some_service",
//					Version:  "1",
//					Resource: resources[0],
//				},
//			},
//			//Canary:      false,
//			TypeUrl: "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment",
//			Nonce:   time.Now().String(),
//		})
//
//	}
//
//	return nil
//}
//
//func (e eds) FetchEndpoints(context.Context, *envoy_api_v2.DiscoveryRequest) (*envoy_api_v2.DiscoveryResponse, error) {
//	fmt.Println("FetchEndpoints")
//	return nil, nil
//}
