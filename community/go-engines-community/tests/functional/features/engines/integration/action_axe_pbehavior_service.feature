Feature: update service when alarm is updated by action pbehavior
  I need to be able to update service when action pbehavior is applied to alarm.

  Scenario: given entity service and scenario with pbehavior action for dependency should update service alarm
    Given I am admin
    When I do POST /api/v4/entityservices:
    """json
    {
      "name": "test-entityservice-action-axe-pbehavior-service-1-name",
      "output_template": "All: {{ `{{.All}}` }}; Active: {{ `{{.Active}}` }}; Acknowledged: {{ `{{.Acknowledged}}` }}; NotAcknowledged: {{ `{{.NotAcknowledged}}` }}; AcknowledgedUnderPbh: {{ `{{.AcknowledgedUnderPbh}}` }}; StateCritical: {{ `{{.State.Critical}}` }}; StateMajor: {{ `{{.State.Major}}` }}; StateMinor: {{ `{{.State.Minor}}` }}; StateOk: {{ `{{.State.Ok}}` }}; Pbehaviors: {{ `{{.PbehaviorCounters}}` }}; UnderPbehavior: {{ `{{.UnderPbehavior}}` }};",
      "impact_level": 1,
      "enabled": true,
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-resource-action-axe-pbehavior-service-1"
            }
          }
        ]
      ],
      "sli_avail_state": 0
    }
    """
    Then the response code should be 201
    When I save response serviceID={{ .lastResponse._id }}
    When I wait the end of 2 events processing
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-pbehavior-service-1-name",
      "priority": 10043,
      "enabled": true,
      "triggers": ["cancel"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-pbehavior-service-1"
                }
              }
            ]
          ],
          "type": "pbehavior",
          "parameters": {
            "name": "test-pbehavior-action-axe-pbehavior-service-1",
            "tstart": {{ now }},
            "tstop": {{ nowAdd "1h" }},
            "type": "test-maintenance-type-to-engine",
            "reason": "test-reason-to-engine"
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-action-axe-pbehavior-service-1",
      "connector_name": "test-connector-name-action-axe-pbehavior-service-1",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-action-axe-pbehavior-service-1",
      "resource": "test-resource-action-axe-pbehavior-service-1",
      "state": 1,
      "output": "test-output-action-axe-pbehavior-service-1"
    }
    """
    When I wait the end of 2 events processing
    When I send an event:
    """json
    {
      "connector": "test-connector-action-axe-pbehavior-service-1",
      "connector_name": "test-connector-name-action-axe-pbehavior-service-1",
      "source_type": "resource",
      "event_type": "cancel",
      "component":  "test-component-action-axe-pbehavior-service-1",
      "resource": "test-resource-action-axe-pbehavior-service-1",
      "output": "test-output-action-axe-pbehavior-service-1"
    }
    """
    When I wait the end of 2 events processing
    When I do GET /api/v4/alarms?search={{ .serviceID }}
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "component": "{{ .serviceID }}",
            "connector": "service",
            "connector_name": "service",
            "state": {
              "val": 0
            },
            "status": {
              "val": 0
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 1
      }
    }
    """
    When I do POST /api/v4/alarm-details:
    """json
    [
      {
        "_id": "{{ (index .lastResponse.data 0)._id }}",
        "steps": {
          "page": 1
        }
      }
    ]
    """
    Then the response code should be 207
    Then the response body should contain:
    """json
    [
      {
        "status": 200,
        "data": {
          "steps": {
            "data": [
              {
                "_t": "stateinc",
                "a": "service.service",
                "m": "All: 1; Active: 1; Acknowledged: 0; NotAcknowledged: 1; AcknowledgedUnderPbh: 0; StateCritical: 0; StateMajor: 0; StateMinor: 1; StateOk: 0; Pbehaviors: map[]; UnderPbehavior: 0;",
                "val": 1
              },
              {
                "_t": "statusinc",
                "a": "service.service",
                "m": "All: 1; Active: 1; Acknowledged: 0; NotAcknowledged: 1; AcknowledgedUnderPbh: 0; StateCritical: 0; StateMajor: 0; StateMinor: 1; StateOk: 0; Pbehaviors: map[]; UnderPbehavior: 0;",
                "val": 1
              },
              {
                "_t": "statedec",
                "a": "service.service",
                "m": "All: 1; Active: 0; Acknowledged: 0; NotAcknowledged: 0; AcknowledgedUnderPbh: 0; StateCritical: 0; StateMajor: 0; StateMinor: 0; StateOk: 1; Pbehaviors: map[test-maintenance-type-to-engine:1]; UnderPbehavior: 1;",
                "val": 0
              },
              {
                "_t": "statusdec",
                "a": "service.service",
                "m": "All: 1; Active: 0; Acknowledged: 0; NotAcknowledged: 0; AcknowledgedUnderPbh: 0; StateCritical: 0; StateMajor: 0; StateMinor: 0; StateOk: 1; Pbehaviors: map[test-maintenance-type-to-engine:1]; UnderPbehavior: 1;",
                "val": 0
              }
            ],
            "meta": {
              "page": 1,
              "page_count": 1,
              "per_page": 10,
              "total_count": 4
            }
          }
        }
      }
    ]
    """