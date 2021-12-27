package axe

import (
	"context"
	"testing"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	mock_alarm "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/canopsis/alarm"
	mock_alarmstatus "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/canopsis/alarmstatus"
	mock_config "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/canopsis/config"
	mock_idlealarm "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/canopsis/idlealarm"
	"github.com/golang/mock/gomock"
)

func TestPeriodicalWorker_Work(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockAlarmService := mock_alarm.NewMockService(ctrl)
	mockAlarmAdapter := mock_alarm.NewMockAdapter(ctrl)
	mockIdleAlarmService := mock_idlealarm.NewMockService(ctrl)
	mockAlarmConfigProvider := mock_config.NewMockAlarmConfigProvider(ctrl)
	mockAlarmStatusService := mock_alarmstatus.NewMockService(ctrl)

	interval := time.Minute
	worker := periodicalWorker{
		PeriodicalInterval:  interval,
		AlarmService:        mockAlarmService,
		AlarmAdapter:        mockAlarmAdapter,
		AlarmStatusService:  mockAlarmStatusService,
		IdleAlarmService:    mockIdleAlarmService,
		AlarmConfigProvider: mockAlarmConfigProvider,
	}

	alarmConfig := config.AlarmConfig{
		StealthyInterval:         time.Second,
		EnableLastEventDate:      true,
		CancelAutosolveDelay:     time.Second,
		OutputLength:             10,
		TimeToKeepResolvedAlarms: time.Second,
	}
	mockAlarmConfigProvider.EXPECT().Get().Return(alarmConfig)
	mockAlarmAdapter.EXPECT().DeleteResolvedAlarms(gomock.Any(), gomock.Any())
	mockAlarmService.EXPECT().ResolveClosed(gomock.Any())
	mockAlarmService.EXPECT().ResolveSnoozes(gomock.Any(), gomock.Eq(alarmConfig))
	mockAlarmService.EXPECT().ResolveCancels(gomock.Any(), gomock.Eq(alarmConfig))
	mockAlarmService.EXPECT().ResolveDone(gomock.Any())
	mockAlarmService.EXPECT().UpdateFlappingAlarms(gomock.Any())
	mockIdleAlarmService.EXPECT().Process(gomock.Any())
	mockAlarmStatusService.EXPECT().Load(gomock.Any())

	_ = worker.Work(ctx)
}
