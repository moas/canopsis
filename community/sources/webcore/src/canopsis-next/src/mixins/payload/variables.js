import {
  ALARM_PAYLOADS_VARIABLES,
  DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES,
  DECLARE_TICKET_PAYLOAD_PREVIOUS_STEP_VARIABLES,
} from '@/constants';

export const payloadVariablesMixin = {
  props: {
    hasPrevious: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    alarmPayloadSubVariables() {
      return [
        {
          value: ALARM_PAYLOADS_VARIABLES.component,
          text: this.$t('common.component'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.resource,
          text: this.$t('common.resource'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.stateMessage,
          text: this.$t('alarm.stateMessage'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.stateValue,
          text: this.$t('alarm.stateValue'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.statusValue,
          text: this.$t('common.status'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.ticketAuthor,
          text: this.$t('alarm.ticketAuthor'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.ticketValue,
          text: this.$t('alarm.ticketId'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.ticketMessage,
          text: this.$t('alarm.ticketMessage'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.ackAuthor,
          text: this.$t('alarm.ackAuthor'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.ackMessage,
          text: this.$t('alarm.ackMessage'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.lastCommentAuthor,
          text: this.$t('alarm.lastCommentAuthor'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.lastCommentMessage,
          text: this.$t('alarm.lastCommentMessage'),
        },
        {
          value: ALARM_PAYLOADS_VARIABLES.entityInfosValue,
          text: this.$t('common.infos'),
        },
      ];
    },

    alarmPayloadRangeVariables() {
      return [{
        value: ALARM_PAYLOADS_VARIABLES.alarms,
        enumerable: true,
        variables: this.alarmPayloadSubVariables,
      }];
    },

    alarmPayloadVariables() {
      return this.alarmPayloadSubVariables.map(
        variable => (
          variable.value !== ALARM_PAYLOADS_VARIABLES.entityInfosValue
            ? ({
              ...variable,
              value: `${ALARM_PAYLOADS_VARIABLES.alarm}${variable.value}`,
            })
            : variable
        ),
      );
    },

    payloadVariablesFromPreviousStep() {
      return [
        {
          value: DECLARE_TICKET_PAYLOAD_PREVIOUS_STEP_VARIABLES.header,
          text: this.$t('declareTicket.headerFieldFromPreviousSteps'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_PREVIOUS_STEP_VARIABLES.response,
          text: this.$t('declareTicket.responseFieldFromPreviousSteps'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_PREVIOUS_STEP_VARIABLES.headerByStep,
          text: this.$t('declareTicket.headerFieldFromStep'),
        },
      ];
    },

    additionalDataVariables() {
      return [
        {
          value: DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES.author,
          text: this.$t('common.username'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES.user,
          text: this.$t('declareTicket.userId'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES.alarmChangeType,
          text: this.$tc('common.trigger'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES.initiator,
          text: this.$t('declareTicket.actionInitiator'),
        },
        {
          value: DECLARE_TICKET_PAYLOAD_ADDITIONAL_DATA_VARIABLES.output,
          text: this.$t('declareTicket.triggerEventMessage'),
        },
      ];
    },

    payloadVariables() {
      const variables = [...this.alarmPayloadRangeVariables];

      if (this.hasPrevious) {
        variables.push(...this.payloadVariablesFromPreviousStep);
      }

      return variables;
    },
  },
};
