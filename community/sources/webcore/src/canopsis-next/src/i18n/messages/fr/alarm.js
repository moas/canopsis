import { EVENT_ENTITY_TYPES, ALARM_METRIC_PARAMETERS } from '@/constants';

export default {
  eventsCount: 'Les événements comptent',
  liveReporting: 'Définir un intervalle de dates',
  ackAuthor: 'Confirmer l\'auteur',
  lastCommentAuthor: 'Auteur du dernier commentaire',
  lastCommentMessage: 'Message du dernier commentaire',
  metaAlarm: 'Meta-alarmes',
  acknowledge: 'Acquitter',
  ackResources: 'Acquitter les ressources',
  ackResourcesQuestion: 'Voulez-vous acquitter les ressources liées ?',
  actionsRequired: 'Actions required',
  acknowledgeAndDeclareTicket: 'Acquitter et déclarer un ticket',
  acknowledgeAndAssociateTicket: 'Acquitter et associer un ticket',
  advancedSearch: '<span>Aide sur la recherche avancée :</span>\n'
    + '<p>- [ NOT ] &lt;NomColonne&gt; &lt;Opérateur&gt; &lt;Valeur&gt;</p> [ AND|OR [ NOT ] &lt;NomColonne&gt; &lt;Opérateur&gt; &lt;Valeur&gt; ]\n'
    + '<p>Le "-" avant la recherche est obligatoire</p>\n'
    + '<p>Opérateurs:\n'
    + '    <=, <,=, !=,>=, >, LIKE (Pour les expressions régulières MongoDB)</p>\n'
    + '<p>Les types de valeurs : Chaîne de caractères entre guillemets doubles, Booléen ("TRUE", "FALSE"), Entier, Nombre flottant, "NULL"</p>\n'
    + '<dl><dt>Exemples :</dt><dt>- Connector = "connector_1"</dt>\n'
    + '    <dd>Alarmes dont le connecteur est "connector_1"</dd><dt>- Connector="connector_1" AND Resource="resource_3"</dt>\n'
    + '    <dd>Alarmes dont le connecteur est "connector_1" et la ressource est "resource_3"</dd><dt>- Connector="connector_1" OR Resource="resource_3"</dt>\n'
    + '    <dd>Alarmes dont le connecteur est "connector_1" ou la ressource est "resource_3"</dd><dt>- Connector LIKE 1 OR Connector LIKE 2</dt>\n'
    + '    <dd>Alarmes dont le connecteur contient 1 ou 2</dd><dt>- NOT Connector = "connector_1"</dt>\n'
    + '    <dd>Alarmes dont le connecteur n\'est pas "connector_1"</dd>\n'
    + '</dl>',
  otherTickets: 'D\'autres tickets sont disponibles dans le panneau d\'expansion',
  noAlarmFound: 'Aucune alarme n\'est trouvée selon les modèles définis',
  associateTicketResources: 'Ticket associé pour les ressources',
  followLink: 'Suivez le lien "{title}"',
  actions: {
    titles: {
      ack: 'Acquitter',
      fastAck: 'Acquitter rapidement',
      ackRemove: 'Annuler l\'acquittement',
      pbehavior: 'Définir un comportement périodique',
      snooze: 'Mettre en veille',
      declareTicket: 'Déclarer un incident',
      associateTicket: 'Associer un ticket',
      cancel: 'Annuler l\'alarme',
      changeState: 'Changer et verrouiller la criticité',
      variablesHelp: 'Liste des variables disponibles',
      history: 'Historique',
      groupRequest: 'Proposition de regroupement pour méta-alarmes',
      createManualMetaAlarm: 'Gestion manuelle des méta-alarmes',
      removeAlarmsFromManualMetaAlarm: 'Dissocier l\'alarme de la méta-alarme manuelle',
      comment: 'Commenter l\'alarme',
    },
    iconsTitles: {
      ack: 'Acquitter',
      declareTicket: 'Déclarer un incident',
      canceled: 'Annulé',
      snooze: 'Mettre en veille',
      pbehaviors: 'Définir un comportement périodique',
      grouping: 'Meta-alarmes',
      comment: 'Commentaire',
    },
    iconsFields: {
      ticketNumber: 'Numéro de ticket',
      parents: 'Causes',
      children: 'Conséquences',
    },
  },
  timeLine: {
    titlePaths: {
      by: 'par',
    },
    stateCounter: {
      header: 'Criticités compressées (depuis le dernier changement de statut)',
      stateIncreased: 'Criticité augmentée',
      stateDecreased: 'Criticité diminuée',
    },
    types: {
      [EVENT_ENTITY_TYPES.ack]: 'Acquittement',
      [EVENT_ENTITY_TYPES.ackRemove]: 'Suppression d\'acquittement',
      [EVENT_ENTITY_TYPES.stateinc]: 'Augmentation de la criticité',
      [EVENT_ENTITY_TYPES.statedec]: 'Diminution de la criticité',
      [EVENT_ENTITY_TYPES.statusinc]: 'Augmentation du statut',
      [EVENT_ENTITY_TYPES.statusdec]: 'Diminution du statut',
      [EVENT_ENTITY_TYPES.assocTicket]: 'Association d\'un ticket',
      [EVENT_ENTITY_TYPES.declareTicket]: 'Le ticket est déclaré avec succès',
      [EVENT_ENTITY_TYPES.declareTicketFail]: 'La déclaration du ticket a échoué',
      [EVENT_ENTITY_TYPES.webhookStart]: 'Webhook a démarré par la racine',
      [EVENT_ENTITY_TYPES.webhookComplete]: 'Webhook est exécuté avec succès',
      [EVENT_ENTITY_TYPES.webhookFail]: 'Webhook a échoué',
      [EVENT_ENTITY_TYPES.snooze]: 'Alarme mise en veille',
      [EVENT_ENTITY_TYPES.unsooze]: 'Alarme sortie de veille',
      [EVENT_ENTITY_TYPES.changeState]: 'Changement et verrouillage de la criticité',
      [EVENT_ENTITY_TYPES.pbhenter]: 'Comportement périodique activé',
      [EVENT_ENTITY_TYPES.pbhleave]: 'Comportement périodique désactivé',
      [EVENT_ENTITY_TYPES.cancel]: 'Alarme annulée',
      [EVENT_ENTITY_TYPES.comment]: 'Alarme commentée',
      [EVENT_ENTITY_TYPES.metaalarmattach]: 'Alarme liée à la méta alarme',
      [EVENT_ENTITY_TYPES.instructionStart]: 'L\'exécution de la consigne a été déclenchée',
      [EVENT_ENTITY_TYPES.instructionPause]: 'L\'exécution de la consigne a été mise en pause',
      [EVENT_ENTITY_TYPES.instructionResume]: 'L\'exécution de la consigne a été reprise',
      [EVENT_ENTITY_TYPES.instructionComplete]: 'L\'exécution de la consigne a été terminée',
      [EVENT_ENTITY_TYPES.instructionAbort]: 'L\'exécution de la consigne a été abandonnée',
      [EVENT_ENTITY_TYPES.instructionFail]: 'L\'exécution de la consigne a échoué',
      [EVENT_ENTITY_TYPES.instructionJobStart]: 'L\'exécution d\'une tâche de remédiation a été démarrée',
      [EVENT_ENTITY_TYPES.instructionJobComplete]: 'L\'exécution de la tâche de remédiation est terminée',
      [EVENT_ENTITY_TYPES.instructionJobAbort]: 'L\'exécution de la tâche de remédiation a été abandonnée',
      [EVENT_ENTITY_TYPES.instructionJobFail]: 'L\'exécution de la tâche de remédiation a échouée',
      [EVENT_ENTITY_TYPES.autoInstructionStart]: 'La consigne automatique a été lancée',
      [EVENT_ENTITY_TYPES.autoInstructionComplete]: 'La consigne automatique a été complétée',
      [EVENT_ENTITY_TYPES.autoInstructionFail]: 'La consigne automatique a échoué',
      [EVENT_ENTITY_TYPES.autoInstructionAlreadyRunning]: 'La consigne automatique a déjà été lancée pour une autre alarme',
      [EVENT_ENTITY_TYPES.junitTestSuiteUpdate]: 'La suite de tests a été mise à jour',
      [EVENT_ENTITY_TYPES.junitTestCaseUpdate]: 'Le cas de test a été mis à jour',
    },
  },
  tabs: {
    moreInfos: 'Plus d\'infos',
    timeLine: 'Chronologie',
    alarmsChildren: 'Alarmes liées',
    trackSource: 'Cause racine',
    impactChain: 'Chaîne d\'impact',
    entityGantt: 'Diagramme de Gantt',
    ticketsDeclared: 'Tickets déclarés',
  },
  moreInfos: {
    defineATemplate: 'Pour définir le template de cette fenêtre, rendez-vous dans les paramètres du bac à alarmes.',
  },
  infoPopup: 'Info popup',
  tooltips: {
    priority: 'La priorité est égale à la sévérité multipliée par le niveau d\'impact de l\'entité sur laquelle l\'alarme est déclenchée',
    runningManualInstructions: 'Consigne manuelle <strong>{title}</strong> en cours | Consignes manuelles <strong>{title}</strong> en cours',
    runningAutoInstructions: 'Consigne automatique <strong>{title}</strong> en cours | Consignes automatiques <strong>{title}</strong> en cours',
    successfulManualInstructions: 'Consigne manuelle <strong>{title}</strong> réussie | Consignes manuelles <strong>{title}</strong> réussies',
    successfulAutoInstructions: 'Consigne automatique <strong>{title}</strong> réussies | Consignes automatiques <strong>{title}</strong> réussies',
    failedManualInstructions: 'Consigne manuelle <strong>{title}</strong> en échec | Consignes manuelles <strong>{title}</strong> en échec',
    failedAutoInstructions: 'Consigne automatique <strong>{title}</strong> en échec | Consignes automatiques <strong>{title}</strong> en échec',
    hasManualInstruction: 'Il y a une consigne manuelle associée | Il y a des consignes manuelles associées',
  },
  metrics: {
    [ALARM_METRIC_PARAMETERS.createdAlarms]: 'Nombre d\'alarmes créées',
    [ALARM_METRIC_PARAMETERS.activeAlarms]: 'Nombre d\'alarmes actives',
    [ALARM_METRIC_PARAMETERS.nonDisplayedAlarms]: 'Nombre d\'alarmes non affichées',
    [ALARM_METRIC_PARAMETERS.instructionAlarms]: 'Nombre d\'alarmes en cours de remédiation automatique',
    [ALARM_METRIC_PARAMETERS.pbehaviorAlarms]: 'Nombre d\'alarmes avec PBehavior',
    [ALARM_METRIC_PARAMETERS.correlationAlarms]: 'Nombre d\'alarmes corrélées',
    [ALARM_METRIC_PARAMETERS.ackAlarms]: 'Nombre d\'alarmes avec acquittement',
    [ALARM_METRIC_PARAMETERS.ackActiveAlarms]: 'Number of active alarms with acks',
    [ALARM_METRIC_PARAMETERS.cancelAckAlarms]: 'Nombre d\'accusés de réception annulés',
    [ALARM_METRIC_PARAMETERS.ticketActiveAlarms]: 'Nombre d\'alarmes actives avec tickets',
    [ALARM_METRIC_PARAMETERS.withoutTicketActiveAlarms]: 'Nombre d\'alarmes actives sans tickets',
    [ALARM_METRIC_PARAMETERS.ratioCorrelation]: '% d\'alarmes corrélées',
    [ALARM_METRIC_PARAMETERS.ratioInstructions]: '% d\'alarmes avec remédiation automatique',
    [ALARM_METRIC_PARAMETERS.ratioTickets]: '% d\'alarmes avec tickets créés',
    [ALARM_METRIC_PARAMETERS.ratioRemediatedAlarms]: '% d\'alarmes corrigées manuellement',
    [ALARM_METRIC_PARAMETERS.ratioNonDisplayed]: '% d\'alarmes non affichées',
    [ALARM_METRIC_PARAMETERS.averageAck]: 'Délai moyen d\'acquittement des alarmes',
    [ALARM_METRIC_PARAMETERS.averageResolve]: 'Temps moyen pour résoudre les alarmes',
    [ALARM_METRIC_PARAMETERS.manualInstructionExecutedAlarms]: 'Nombre d\'alarmes corrigées manuellement',
    [ALARM_METRIC_PARAMETERS.manualInstructionAssignedAlarms]: 'Nombre d\'alarmes avec instructions manuelles',
    [ALARM_METRIC_PARAMETERS.notAckedAlarms]: 'Nombre d\'alarmes non acquittées',
    [ALARM_METRIC_PARAMETERS.notAckedInHourAlarms]: 'Nombre d\'alarmes non acquittées avec durée 1-4h',
    [ALARM_METRIC_PARAMETERS.notAckedInFourHoursAlarms]: 'Nombre d\'alarmes non acquittées avec durée 4-24h',
    [ALARM_METRIC_PARAMETERS.notAckedInDayAlarms]: 'Nombre d\'alarmes non acquittées de plus de 24h',
  },
  fields: {
    displayName: 'Nom simplifié (DisplayName)',
    initialOutput: 'Sortie initiale longue',
    initialLongOutput: 'Sortie longue initiale',
    lastComment: 'Dernier commentaire',
    ackBy: 'Acquis par',
    ackMessage: 'Message de l\'acquittement',
    ackInitiator: 'Origine de l\'acquittement',
    stateMessage: 'Message d\'état',
    statusMessage: 'Message de statut',
    totalStateChanges: 'Changements d\'état totaux',
    stateValue: 'Valeur d\'état',
    statusValue: 'valeur de statut',
    ackAt: 'Acquitté à',
    stateAt: 'État changé à',
    statusAt: 'Le statut a changé à',
    resolved: 'Résolue à',
    activationDate: 'Date d\'activation',
    currentStateDuration: 'Durée de l\'état actuel',
    snoozeDuration: 'Durée de sommeil',
    pbhInactiveDuration: 'Pbehavior durée d\'inactivité',
    activeDuration: 'Durée active',
    eventsCount: 'Les événements comptent',
    extraDetails: 'Détails supplémentaires',
    ticketAuthor: 'Auteur du ticket',
    ticketId: 'ID du ticket',
    ticketMessage: 'Message du ticket',
    entityId: 'ID d\'entité',
    entityName: 'Nom de l\'entité',
    entityCategoryName: 'Nom de la catégorie d\'entité',
    entityType: 'Type d\'entité',
    entityComponent: 'Composant d\'entité',
    entityConnector: 'Connecteur d\'entité',
    entityImpactLevel: 'Niveau d\'impact de l\'entité',
    entityKoEvents: 'Événements d\'entité KO',
    entityOkEvents: 'Événements d\'entité OK',
    entityInfos: 'Informations sur l\'entité',
    entityComponentInfos: 'Informations sur les composants de l\'entité',
    entityLastPbehaviorDate: 'Date du dernier comportement de l\'entité',
  },
};