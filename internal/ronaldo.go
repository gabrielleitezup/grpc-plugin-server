package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gabrielleitezup/grpc-plugin-server/pkg/datasource"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

type RonaldoServer struct {
	datasource.UnimplementedDatasourceServer
	savedMetricList []*datasource.MetricList // read-only after initialized
}

func notGetPrometheusApiClient(datasourceConfiguration []byte) (api.Client, error) {
	var prometheusConfig map[string]string
	_ = json.Unmarshal(datasourceConfiguration, &prometheusConfig)

	apiConf := api.Config{
		Address: prometheusConfig["url"],
	}

	return api.NewClient(apiConf)
}

func (s *RonaldoServer) GetMetrics(ctx context.Context, datasourceConfiguration *datasource.MetricsRequest) (*datasource.MetricList, error) {
	fmt.Println("RONALD, BRILHA MUITO NO CORINTHIANS")
	apiClient, err := getPrometheusApiClient(datasourceConfiguration.Configuration)
	if err != nil {
		return nil, err
	}

	v1Api := v1.NewAPI(apiClient)
	namedLabels := "__name__"
	labelValues, _, err := v1Api.LabelValues(context.Background(), namedLabels, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	var metricList []string
	for _, label := range labelValues {
		metricList = append(metricList, string(label))
	}

	return &datasource.MetricList{Metrics: metricList}, nil
}