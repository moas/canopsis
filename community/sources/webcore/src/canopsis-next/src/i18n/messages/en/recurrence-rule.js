import { PBEHAVIOR_RRULE_PERIODS_RANGES } from '@/constants';

export default {
  advancedHint: 'Separate numbers with a comma',
  freq: 'Frequency',
  until: 'Until',
  byweekday: 'By week day',
  count: 'Repeat',
  interval: 'Interval',
  wkst: 'Week start',
  bymonth: 'By month',
  bysetpos: 'By set position',
  bymonthday: 'By month day',
  byyearday: 'By year day',
  byweekno: 'By week n°',
  byhour: 'By hour',
  byminute: 'By minute',
  bysecond: 'By second',
  tabs: {
    simple: 'Simple',
    advanced: 'Advanced',
  },
  errors: {
    main: 'Please note that the recurrence rule you chose is not valid. We strongly advise you to modify it before saving changes.',
  },
  periodsRanges: {
    [PBEHAVIOR_RRULE_PERIODS_RANGES.thisWeek]: 'This week',
    [PBEHAVIOR_RRULE_PERIODS_RANGES.nextWeek]: 'Next week',
    [PBEHAVIOR_RRULE_PERIODS_RANGES.next2Weeks]: 'Next 2 weeks',
    [PBEHAVIOR_RRULE_PERIODS_RANGES.thisMonth]: 'This month',
    [PBEHAVIOR_RRULE_PERIODS_RANGES.nextMonth]: 'Next month',
  },
  tooltips: {
    bysetpos: 'If given, it must be one or many integers, positive or negative. Each given integer will specify an occurrence number, corresponding to the nth occurrence of the rule inside the frequency period. For example, a \'bysetpos\' of -1 if combined with a monthly frequency, and a \'byweekday\' of (Monday, Tuesday, Wednesday, Thursday, Friday), will result in the last work day of every month.',
    bymonthday: 'If given, it must be one or many integers, meaning the month days to apply the recurrence to.',
    byyearday: 'If given, it must be one or many integers, meaning the year days to apply the recurrence to.',
    byweekno: 'If given, it must be on or many integers, meaning the week numbers to apply the recurrence to. Week numbers have the meaning described in ISO8601, that is, the first week of the year is that containing at least four days of the new year.',
    byhour: 'If given, it must be one or many integers, meaning the hours to apply the recurrence to.',
    byminute: 'If given, it must be one or many integers, meaning the minutes to apply the recurrence to.',
    bysecond: 'If given, it must be one or many integers, meaning the seconds to apply the recurrence to.',
  },
};
