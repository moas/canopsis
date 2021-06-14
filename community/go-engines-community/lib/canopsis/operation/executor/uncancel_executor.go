package executor

import (
	"fmt"
	"git.canopsis.net/canopsis/go-engines/lib/canopsis/config"
	operationlib "git.canopsis.net/canopsis/go-engines/lib/canopsis/operation"
	"git.canopsis.net/canopsis/go-engines/lib/canopsis/types"
	"git.canopsis.net/canopsis/go-engines/lib/utils"
)

func NewUncancelExecutor(configProvider config.AlarmConfigProvider) operationlib.Executor {
	return &uncancelExecutor{configProvider: configProvider}
}

type uncancelExecutor struct {
	configProvider config.AlarmConfigProvider
}

func (e *uncancelExecutor) Exec(
	operation types.Operation,
	alarm *types.Alarm,
	time types.CpsTime,
	role, initiator string,
) (types.AlarmChangeType, error) {
	var params types.OperationParameters
	var ok bool
	if params, ok = operation.Parameters.(types.OperationParameters); !ok {
		return "", fmt.Errorf("invalid parameters")
	}

	if alarm.Value.Canceled == nil {
		return "", nil
	}

	alarmConfig := e.configProvider.Get()
	err := alarm.PartialUpdateUncancel(
		time,
		params.Author,
		utils.TruncateString(params.Output, alarmConfig.OutputLength),
		role,
		initiator,
		alarmConfig,
	)
	if err != nil {
		return "", err
	}

	return types.AlarmChangeTypeUncancel, nil
}
