package app

import (
	"github.com/atlassian/ctrl"
	rep_v1 "github.com/atlassian/voyager/pkg/apis/reporter/v1"
	reporterClient "github.com/atlassian/voyager/pkg/reporter/client"
	"github.com/atlassian/voyager/pkg/reporterreporter"
	"github.com/atlassian/voyager/pkg/util"
)

type ControllerConstructor struct {
	ConfigFile  string
	SplitterURI string
}

func (cc *ControllerConstructor) AddFlags(flagset ctrl.FlagSet) {
	flagset.StringVar(&cc.ConfigFile, "config", "config.yaml", "config file")
	flagset.StringVar(&cc.SplitterURI, "spitter-uri", "https://micros2-analytics-slurper.us-east-1.prod.atl-paas.net/", "Remote data endpoint")
}

func (cc *ControllerConstructor) New(config *ctrl.Config, cctx *ctrl.Context) (*ctrl.Constructed, error) {
	opts, err := readAndValidateOptions(cc.ConfigFile)
	if err != nil {
		return nil, err
	}

	reporter, err := reporterClient.NewForConfig(config.RestConfig)
	if err != nil {
		return nil, err
	}

	return &ctrl.Constructed{
		Server: &reporterreporter.Report{
			RemoteURI:        cc.SplitterURI,
			Cluster:          opts.Cluster,
			HTTPClient:       util.HTTPClient(),
			ReporterClient:   reporter,
			KubernetesClient: config.MainClient,
			Logger:           config.Logger,
		},
	}, nil
}

func (cc *ControllerConstructor) Describe() ctrl.Descriptor {
	return ctrl.Descriptor{
		Gvk: rep_v1.ReportGvk,
	}
}
