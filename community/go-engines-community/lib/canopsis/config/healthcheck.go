package config

import "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis"

type HealthCheckConf struct {
	EngineOrder []EngineOrder         `toml:"engine_order" bson:"engine_order" json:"engine_order"`
	Parameters  HealthCheckParameters `bson:"parameters" json:"parameters"`
}

type EngineOrder struct {
	From string `toml:"from" bson:"from" json:"from"`
	To   string `toml:"to" bson:"to" json:"to"`
}

type EngineParameters struct {
	Minimal int  `bson:"minimal" json:"minimal"`
	Optimal int  `bson:"optimal" json:"optimal" binding:"gtefield=Minimal"`
	Enabled bool `bson:"enabled" json:"enabled"`
}

type QueueParameters struct {
	Limit   int  `bson:"limit" json:"limit"`
	Enabled bool `bson:"enabled" json:"enabled"`
}

type HealthCheckParameters struct {
	Queue        QueueParameters  `bson:"queue" json:"queue"`
	Fifo         EngineParameters `bson:"engine-fifo" json:"engine-fifo"`
	Che          EngineParameters `bson:"engine-che" json:"engine-che"`
	PBehavior    EngineParameters `bson:"engine-pbehavior" json:"engine-pbehavior"`
	Axe          EngineParameters `bson:"engine-axe" json:"engine-axe"`
	Correlation  EngineParameters `bson:"engine-correlation" json:"engine-correlation"`
	Remediation  EngineParameters `bson:"engine-remediation" json:"engine-remediation"`
	Service      EngineParameters `bson:"engine-service" json:"engine-service"`
	DynamicInfos EngineParameters `bson:"engine-dynamic-infos" json:"engine-dynamic-infos"`
	Action       EngineParameters `bson:"engine-action" json:"engine-action"`
	Webhook      EngineParameters `bson:"engine-webhook" json:"engine-webhook"`
}

func (h HealthCheckParameters) GetEngineParameters(name string) EngineParameters {
	switch name {
	case canopsis.FIFOEngineName:
		return h.Fifo
	case canopsis.CheEngineName:
		return h.Che
	case canopsis.PBehaviorEngineName:
		return h.PBehavior
	case canopsis.AxeEngineName:
		return h.Axe
	case canopsis.CorrelationEngineName:
		return h.Correlation
	case canopsis.RemediationEngineName:
		return h.Remediation
	case canopsis.ServiceEngineName:
		return h.Service
	case canopsis.DynamicInfosEngineName:
		return h.DynamicInfos
	case canopsis.ActionEngineName:
		return h.Action
	case canopsis.WebhookEngineName:
		return h.Webhook
	}

	return EngineParameters{}
}
