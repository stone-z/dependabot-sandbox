package main

import (
	"github.com/giantswarm/k8sclient"
	v6 "github.com/giantswarm/k8scloudconfig/v6/v_6_0_0"
	"github.com/giantswarm/microerror"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

type Config struct {
	// Dependencies.
	Logger micrologger.Logger

	Viper *viper.Viper
}

func main() {
	c := k8sclient.ClientsConfig{}

	k8sClient, err := k8sclient.NewClients(c)
	if err != nil {
		panic(microerror.Mask(err))
	}

	k8scc := v6.Params{}

	mc := microserver.Config{}

	prom := prometheus.GaugeVec{}

	kapi := v1.Pod{}
	kapim := runtime.Object{}
	kcgo := rest.Config{}

}
