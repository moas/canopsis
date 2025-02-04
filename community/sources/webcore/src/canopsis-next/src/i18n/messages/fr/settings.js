import {
  STATS_CRITICITY,
  ENTITY_TYPES,
  SIDE_BARS,
  ALARMS_OPENED_VALUES,
} from '@/constants';

export default {
  titles: {
    [SIDE_BARS.alarmSettings]: 'Paramètres du bac à alarmes',
    [SIDE_BARS.contextSettings]: 'Paramètres de l\'explorateur de contexte',
    [SIDE_BARS.serviceWeatherSettings]: 'Paramètres de la météo des services',
    [SIDE_BARS.statsCalendarSettings]: 'Paramètres du calendrier',
    [SIDE_BARS.textSettings]: 'Paramètres du widget texte',
    [SIDE_BARS.counterSettings]: 'Paramètres du widget compteur',
    [SIDE_BARS.testingWeatherSettings]: 'Paramètres du widget scénario des tests',
    [SIDE_BARS.mapSettings]: 'Paramètres du widget de cartographie',
  },
  openedTypes: {
    [ALARMS_OPENED_VALUES.opened]: 'Alarmes ouvertes',
    [ALARMS_OPENED_VALUES.resolved]: 'Alarmes résolues',
    [ALARMS_OPENED_VALUES.all]: 'Alarmes ouvertes et récemment résolues',
  },
  advancedSettings: 'Paramètres avancés',
  entityDisplaySettings: 'Paramètres d\'affichage des entités',
  entitiesUnderPbehaviorEnabled: 'Entités sous type PBh inactives, Pause, Maintenance display',
  widgetTitle: 'Titre du widget',
  columnName: 'Nom de la colonne',
  defaultSortColumn: 'Colonne de tri par défaut',
  sortColumnNoData: 'Appuyez sur <kbd>enter</kbd> pour en créer une nouvelle',
  columnNames: 'Nom des colonnes',
  exportColumnNames: 'Nom des colonnes à exporter',
  groupColumnNames: 'Nom des colonnes des méta-alarmes',
  trackColumnNames: 'Colonnes pour le suivi de cause racine',
  treeOfDependenciesColumnNames: 'Nom des colonnes pour l\'arborescence des dépendances',
  orderBy: 'Trier par',
  periodicRefresh: 'Rafraichissement périodique',
  defaultNumberOfElementsPerPage: 'Nombre d\'élements par page par défaut',
  elementsPerPage: 'Élements par page',
  filterOnOpenResolved: 'Filtre sur Ouverte/Résolue',
  open: 'Ouverte',
  filters: 'Filtres',
  filterEditor: 'Éditeur de filtre',
  isAckNoteRequired: 'Champ \'Note\' requis lors d\'un acquittement ?',
  isSnoozeNoteRequired: 'Champ \'Note\' requis lors d\'une mise en veille ?',
  inlineLinksCount: 'Nombre de liens en ligne',
  isMultiAckEnabled: 'Acquittement multiple',
  isMultiDeclareTicketEnabled: 'Déclarer un ticket multiple',
  fastAckOutput: 'Commentaire d\'acquittement rapide',
  isHtmlEnabledOnTimeLine: 'HTML activé dans la chronologie ?',
  isCorrelationEnabled: 'Corrélation activée ?',
  duration: 'Durée',
  tstop: 'Date de fin',
  periodsNumber: 'Nombre d\'étapes',
  yesNoMode: 'Mode Oui/Non',
  selectAFilter: 'Sélectionner un filtre',
  lockedFilter: 'Filtre verrouillé dans les paramètres du widget',
  exportAsCsv: 'Exporter les données du widget en csv',
  criticityLevels: 'Niveaux de criticité',
  isPriorityEnabled: 'Afficher la priorité',
  clearFilterDisabled: 'Désactiver la possibilité d\'effacer le filtre sélectionné',
  alarmsColumns: 'Colonnes de la liste des alarmes',
  resolvedAlarmsColumns: 'Noms de colonne pour les alarmes résolues',
  activeAlarmsColumns: 'Noms de colonne pour les alarmes actives',
  entitiesColumns: 'Colonnes de l\'explorateur de contexte',
  entityInfoPopup: 'Fenêtre contextuelle d\'informations sur l\'entité',
  modal: '(Modal)',
  exportCsv: {
    title: 'Exporter CSV',
    fields: {
      separator: 'Séparateur',
      datetimeFormat: 'Format date/heure',
    },
  },
  colorsSelector: {
    title: 'Sélecteur de couleur',
    statsCriticity: {
      [STATS_CRITICITY.ok]: 'ok',
      [STATS_CRITICITY.minor]: 'mineur',
      [STATS_CRITICITY.major]: 'majeur',
      [STATS_CRITICITY.critical]: 'critique',
    },
  },
  infoPopup: {
    title: 'Info popup',
    fields: {
      column: 'Colonne',
    },
  },
  rowGridSize: {
    title: 'Taille du widget',
    noData: 'Aucune ligne correspondante. Appuyez sur <kbd>enter</kbd> pour en créer une nouvelle',
    fields: {
      row: 'Ligne',
    },
  },
  moreInfosModal: 'Fenêtre "Plus d\'infos"',
  expandGridRangeSize: 'Largeur-position "Plus d\'infos / chronologie"',
  weatherTemplate: 'Template - Tuiles',
  modalTemplate: 'Template - Modale',
  entityTemplate: 'Template - Entités',
  blockTemplate: 'Template - Tuiles',
  columnMobile: 'Colonnes - Mobiles',
  columnTablet: 'Colonnes - Tablette',
  columnDesktop: 'Colonnes - Bureau',
  limit: 'Limite',
  height: 'Hauteur',
  margin: {
    title: 'Marges',
    top: 'Marge - Haut',
    right: 'Marge - Droite',
    bottom: 'Marge - Bas',
    left: 'Marge - Gauche',
  },
  contextTypeOfEntities: {
    title: 'Type d\'entité',
    fields: {
      [ENTITY_TYPES.component]: 'Composant',
      [ENTITY_TYPES.connector]: 'Type de connecteur',
      [ENTITY_TYPES.resource]: 'Ressource',
      [ENTITY_TYPES.service]: 'Service',
    },
  },
  considerPbehaviors: {
    title: 'Prendre en compte les comportements périodiques ?',
  },
  serviceWeatherModalTypes: {
    title: 'Type de modale',
    fields: {
      moreInfo: 'Plus d\'infos',
      alarmList: 'Bac à alarmes',
      both: 'Les deux',
    },
  },
  columns: {
    customLabel: 'Étiquette personnalisée',
    isHtml: 'Est-ce du HTML ?',
    withTemplate: 'Modèle personnalisé',
    isState: 'Affiché comme une criticité ?',
    onlyIcon: 'Afficher uniquement les icônes de liens',
  },
  liveReporting: {
    title: 'Suivi personnalisé',
  },
  counterLevels: {
    title: 'Niveaux',
    fields: {
      counter: 'Compteur',
    },
  },
  counters: 'Compteurs',
  pbehaviorCounters: 'Compteurs de pbehavior',
  entityStateCounters: 'Compteurs d\'états d\'entité',
  remediationInstructionsFilters: 'Filtres de consignes',
  colorIndicator: {
    title: 'Indicateur de couleur',
    fields: {
      displayAsSeverity: 'Afficher comme gravité',
      displayAsPriority: 'Afficher en priorité',
    },
  },
  receiveByApi: 'Réponse de l\'API',
  serverStorage: 'Stockage serveur',
  filenameRecognition: 'Reconnaissance du nom de fichier',
  resultDirectory: 'Stockage des résultats de test',
  screenshotDirectories: {
    title: 'Paramètres de stockage des captures d\'écran',
    helpText: 'Définir où les captures d\'écran sont stockées',
  },
  screenshotMask: {
    title: 'Règle de nommage des fichiers capture d\'écran',
    helpText: '<dl>'
      + '<dt>Définissez la règle de nommage des fichiers dont les captures d\'écran sont créées à l\'aide des variables suivantes:<dt>\n'
      + '<dd>- nom du cas de test %test_case%</dd>\n'
      + '<dd>- date (YYYY, MM, DD)</dd>\n'
      + '<dd>- temps d\'exécution (hh, mm, ss)</dd>'
      + '</dl>',
  },
  videoDirectories: {
    title: 'Paramètres de stockage vidéo',
    helpText: 'Définir où la vidéo est stockée',
  },
  videoMask: {
    title: 'Règle de nommage des fichiers vidéo',
    helpText: '<dl>'
      + '<dt>Définissez la règle de nommage des fichiers dont les vidéos sont créées à l\'aide des variables suivantes:<dt>\n'
      + '<dd>- nom du cas de test %test_case%</dd>\n'
      + '<dd>- date (YYYY, MM, DD)</dd>\n'
      + '<dd>- temps d\'exécution (hh, mm, ss)</dd>'
      + '</dl>',
  },
  stickyHeader: 'En-tête collant',
  reportFileRegexp: {
    title: 'Masque de fichier de rapport',
    helpText: '<dl>'
      + '<dt>Définir le nom de fichier regexp de quel rapport:<dt>\n'
      + '<dd>Par exemple:</dd>\n'
      + '<dd>"^(?P&lt;name&gt;\\\\w+)_(.+)\\\\.xml$"</dd>\n'
      + '</dl>',
  },
  density: {
    title: 'Vue par défaut',
    comfort: 'Vue confort',
    compact: 'Vue compacte',
    ultraCompact: 'Vue ultra compacte',
  },
};
