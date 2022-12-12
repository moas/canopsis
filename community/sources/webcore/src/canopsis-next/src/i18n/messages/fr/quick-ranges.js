import { QUICK_RANGES } from '@/constants';

export default {
  title: 'Valeurs usuelles',
  timeField: 'Champ de temps',
  types: {
    [QUICK_RANGES.custom.value]: 'Personnalisé',
    [QUICK_RANGES.last15Minutes.value]: '15 dernières minutes',
    [QUICK_RANGES.last30Minutes.value]: '30 dernières minutes',
    [QUICK_RANGES.last1Hour.value]: 'Dernière heure',
    [QUICK_RANGES.last3Hour.value]: '3 dernières heures',
    [QUICK_RANGES.last6Hour.value]: '6 dernières heures',
    [QUICK_RANGES.last12Hour.value]: '12 dernières heures',
    [QUICK_RANGES.last24Hour.value]: '24 dernières heures',
    [QUICK_RANGES.last2Days.value]: '2 derniers jours',
    [QUICK_RANGES.last7Days.value]: '7 derniers jours',
    [QUICK_RANGES.last30Days.value]: '30 derniers jours',
    [QUICK_RANGES.last1Year.value]: 'Dernière année',
    [QUICK_RANGES.yesterday.value]: 'Hier',
    [QUICK_RANGES.previousWeek.value]: 'Dernière semaine',
    [QUICK_RANGES.previousMonth.value]: 'Dernier mois',
    [QUICK_RANGES.today.value]: 'Aujourd\'hui',
    [QUICK_RANGES.todaySoFar.value]: 'Aujourd\'hui jusqu\'à maintenant',
    [QUICK_RANGES.thisWeek.value]: 'Cette semaine',
    [QUICK_RANGES.thisWeekSoFar.value]: 'Cette semaine jusqu\'à maintenant',
    [QUICK_RANGES.thisMonth.value]: 'Ce mois',
    [QUICK_RANGES.thisMonthSoFar.value]: 'Ce mois jusqu\'à maintenant',
  },
};
