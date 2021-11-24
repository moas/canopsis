Feature: update alarm on pbehavior
  I need to be able to create pbehavior for alarm

  Scenario: given pbehavior should create alarm with pbehavior info
    Given I am admin
    When I send an event:
    """json
    {
      "connector" : "test-connector-pbehavior-1",
      "connector_name" : "test-connector-name-pbehavior-1",
      "source_type" : "resource",
      "event_type" : "check",
      "component" : "test-component-pbehavior-1",
      "resource" : "test-resource-pbehavior-1",
      "state" : 0,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do POST /api/v4/pbehaviors:
    """json
    {
      "enabled": true,
      "name": "test-pbehavior-1",
      "tstart": {{ now }},
      "tstop": {{ nowAdd "10m" }},
      "type": "test-maintenance-type-to-engine",
      "reason": "test-reason-to-engine",
      "filter":{
        "$and":[
          {
            "name": "test-resource-pbehavior-1"
          }
        ]
      }
    }
    """
    Then the response code should be 201
    When I wait 1s
    When I send an event:
    """json
    {
      "connector" : "test-connector-pbehavior-1",
      "connector_name" : "test-connector-name-pbehavior-1",
      "source_type" : "resource",
      "event_type" : "check",
      "component" : "test-component-pbehavior-1",
      "resource" : "test-resource-pbehavior-1",
      "state" : 1,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/alarms?filter={"$and":[{"v.resource":"test-resource-pbehavior-1"}]}&with_steps=true
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "pbehavior_info": {
              "name": "test-pbehavior-1"
            },
            "connector" : "test-connector-pbehavior-1",
            "connector_name" : "test-connector-name-pbehavior-1",
            "component" : "test-component-pbehavior-1",
            "resource" : "test-resource-pbehavior-1",
            "steps": [
              {
                "_t": "stateinc",
                "val": 1
              },
              {
                "_t": "statusinc",
                "val": 1
              },
              {
                "_t": "pbhenter",
                "m": "Pbehavior test-pbehavior-1. Type: Engine maintenance. Reason: Test Engine"
              }
            ]
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

  Scenario: given pbehavior and alarm should update alarm pbehavior info
    Given I am admin
    When I send an event:
    """json
    {
      "connector" : "test-connector-pbehavior-2",
      "connector_name" : "test-connector-name-pbehavior-2",
      "source_type" : "resource",
      "event_type" : "check",
      "component" : "test-component-pbehavior-2",
      "resource" : "test-resource-pbehavior-2",
      "state" : 1,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do POST /api/v4/pbehaviors:
    """json
    {
      "enabled": true,
      "name": "test-pbehavior-2",
      "tstart": {{ now }},
      "tstop": {{ nowAdd "10m" }},
      "type": "test-maintenance-type-to-engine",
      "reason": "test-reason-to-engine",
      "filter":{
        "$and":[
          {
            "name": "test-resource-pbehavior-2"
          }
        ]
      }
    }
    """
    Then the response code should be 201
    When I wait the end of event processing
    When I do GET /api/v4/alarms?filter={"$and":[{"v.resource":"test-resource-pbehavior-2"}]}&with_steps=true
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "pbehavior_info": {
              "name": "test-pbehavior-2"
            },
            "connector" : "test-connector-pbehavior-2",
            "connector_name" : "test-connector-name-pbehavior-2",
            "component" : "test-component-pbehavior-2",
            "resource" : "test-resource-pbehavior-2",
            "steps": [
              {
                "_t": "stateinc",
                "val": 1
              },
              {
                "_t": "statusinc",
                "val": 1
              },
              {
                "_t": "pbhenter",
                "m": "Pbehavior test-pbehavior-2. Type: Engine maintenance. Reason: Test Engine"
              }
            ]
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

  Scenario: given pbehavior should update last alarm date of pbehavior
    Given I am admin
    When I send an event:
    """json
    {
      "connector" : "test-connector-pbehavior-3",
      "connector_name" : "test-connector-name-pbehavior-3",
      "source_type" : "resource",
      "event_type" : "check",
      "component" : "test-component-pbehavior-3",
      "resource" : "test-resource-pbehavior-3",
      "state" : 1,
      "output" : "test-output-pbehavior-3"
    }
    """
    When I wait the end of event processing
    When I do POST /api/v4/pbehaviors:
    """json
    {
      "enabled": true,
      "name": "test-pbehavior-3",
      "tstart": {{ now }},
      "tstop": {{ nowAdd "10m" }},
      "type": "test-maintenance-type-to-engine",
      "reason": "test-reason-to-engine",
      "filter":{
        "$and":[
          {
            "name": "test-resource-pbehavior-3"
          }
        ]
      }
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "last_alarm_date": null
    }
    """
    When I save response pbehaviorID={{ .lastResponse._id }}
    When I wait the end of event processing
    When I do GET /api/v4/pbehaviors/{{ .pbehaviorID }}
    Then the response code should be 200
    Then the response key "last_alarm_date" should not be "null"

  Scenario: given pbehavior should update alarm pbehavior duration
    Given I am admin
    When I send an event:
    """json
    {
      "connector" : "test-connector-pbehavior-4",
      "connector_name" : "test-connector-name-pbehavior-4",
      "source_type" : "resource",
      "event_type" : "check",
      "component" : "test-component-pbehavior-4",
      "resource" : "test-resource-pbehavior-4",
      "state" : 1,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do POST /api/v4/pbehaviors:
    """json
    {
      "enabled": true,
      "name": "test-pbehavior-4",
      "tstart": {{ now }},
      "tstop": {{ nowAdd "2s" }},
      "type": "test-maintenance-type-to-engine",
      "reason": "test-reason-to-engine",
      "filter":{
        "$and":[
          {
            "name": "test-resource-pbehavior-4"
          }
        ]
      }
    }
    """
    Then the response code should be 201
    When I wait the end of 2 events processing
    When I do GET /api/v4/alarms?filter={"$and":[{"v.resource":"test-resource-pbehavior-4"}]}&with_steps=true
    Then the response code should be 200
    Then the response key "data.0.v.pbh_inactive_duration" should not be "0"
