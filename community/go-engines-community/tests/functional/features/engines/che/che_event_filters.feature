Feature: modify event on event filter
  I need to be able to modify event on event filter

  Scenario: given check event and drop event filter should drop event
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "drop",
      "event_pattern":[[
        {
          "field": "event_type",
          "cond": {
            "type": "eq",
            "value": "check"
          }
        },
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-1"
          }
        }
      ]],
      "description": "test-event-filter-che-event-filters-1-description",
      "priority": 1,
      "enabled": true
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-1",
      "connector_name": "test-connector-name-che-event-filters-1",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-1",
      "resource": "test-resource-che-event-filters-1",
      "state": 2,
      "output": "test-output-che-event-filters-1"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/alarms?search=che-event-filters-1
    Then the response code should be 200
    Then the response body should be:
    """json
    {
      "data": [],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 0
      }
    }
    """

  Scenario: given check event and break event filter should not process next event filters
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "break",
      "event_pattern":[[
        {
          "field": "event_type",
          "cond": {
            "type": "eq",
            "value": "check"
          }
        },
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-2"
          }
        }
      ]],
      "description": "test-event-filter-che-event-filters-2-description",
      "priority": 1,
      "enabled": true
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "drop",
      "event_pattern":[[
        {
          "field": "event_type",
          "cond": {
            "type": "eq",
            "value": "check"
          }
        },
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-2"
          }
        }
      ]],
      "description": "test-event-filter-che-event-filters-1-description",
      "priority": 2,
      "enabled": true
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-2",
      "connector_name": "test-connector-name-che-event-filters-2",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-2",
      "resource": "test-resource-che-event-filters-2",
      "state": 2,
      "output": "test-output-che-event-filters-2"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-2
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-2"
        },
        {
          "_id": "test-connector-che-event-filters-2/test-connector-name-che-event-filters-2"
        },
        {
          "_id": "test-resource-che-event-filters-2/test-component-che-event-filters-2"
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given check event and enrichment event filter with set_entity_info_from_template action
    should update event and entity
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern":[[
        {
          "field": "event_type",
          "cond": {
            "type": "eq",
            "value": "check"
          }
        },
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-3"
          }
        }
      ]],
      "config": {
        "actions": [
          {
            "type": "set_entity_info_from_template",
            "name": "customer",
            "description": "Client",
            "value": "{{ `{{ .Event.ExtraInfos.customer }}` }}"
          },
          {
            "type": "set_entity_info_from_template",
            "name": "manager",
            "description": "Manager",
            "value": "{{ `{{ .Event.ExtraInfos.manager }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-3-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern":[[
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-3"
          }
        }
      ]],
      "config": {
        "actions": [
          {
            "type": "set_field_from_template",
            "name": "Output",
            "value": "{{ `{{ .Event.Output }}` }} (client: {{ `{{ .Event.Entity.Infos.customer.Value }}` }})"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-3-description",
      "enabled": true,
      "priority": 3
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern":[[
        {
          "field": "event_type",
          "cond": {
            "type": "eq",
            "value": "check"
          }
        },
        {
          "field": "component",
          "cond": {
            "type": "eq",
            "value": "test-component-che-event-filters-3"
          }
        }
      ]],
      "config": {
        "actions": [
          {
            "type": "set_entity_info_from_template",
            "name": "output",
            "description": "Output",
            "value": "{{ `{{ .Event.Output }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-3-description",
      "enabled": true,
      "priority": 4
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-3",
      "connector_name": "test-connector-name-che-event-filters-3",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-3",
      "resource": "test-resource-che-event-filters-3",
      "state": 2,
      "output": "test-output-che-event-filters-3",
      "customer": "test-customer-che-event-filters-3",
      "manager": "test-manager-che-event-filters-3"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-3
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-3"
        },
        {
          "_id": "test-connector-che-event-filters-3/test-connector-name-che-event-filters-3"
        },
        {
          "_id": "test-resource-che-event-filters-3/test-component-che-event-filters-3",
          "infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-3"
            },
            "manager": {
              "description": "Manager",
              "name": "manager",
              "value": "test-manager-che-event-filters-3"
            },
            "output": {
              "description": "Output",
              "name": "output",
              "value": "test-output-che-event-filters-3 (client: test-customer-che-event-filters-3)"
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-3",
      "connector_name": "test-connector-name-che-event-filters-3",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-3",
      "resource": "test-resource-che-event-filters-3",
      "state": 2,
      "output": "test-output-che-event-filters-3",
      "customer": "test-customer-che-event-filters-3",
      "manager": "test-manager-che-event-filters-3-updated"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entitybasics?_id=test-resource-che-event-filters-3/test-component-che-event-filters-3
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-resource-che-event-filters-3/test-component-che-event-filters-3",
      "infos": {
        "customer": {
          "description": "Client",
          "name": "customer",
          "value": "test-customer-che-event-filters-3"
        },
        "manager": {
          "description": "Manager",
          "name": "manager",
          "value": "test-manager-che-event-filters-3-updated"
        },
        "output": {
          "description": "Output",
          "name": "output",
          "value": "test-output-che-event-filters-3 (client: test-customer-che-event-filters-3)"
        }
      }
    }
    """

  Scenario: given resource event should fill component infos
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-4"
            }
          },
          {
            "field": "source_type",
            "cond": {
              "type": "eq",
              "value": "component"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.customer",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_entity_info",
            "name": "customer",
            "description": "Client",
            "value": "test-customer-che-event-filters-4"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-4-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-4",
      "connector_name": "test-connector-name-che-event-filters-4",
      "source_type": "component",
      "event_type": "check",
      "component": "test-component-che-event-filters-4",
      "state": 2,
      "output": "test-output-che-event-filters-4"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-4",
      "connector_name": "test-connector-name-che-event-filters-4",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-4",
      "resource": "test-resource-che-event-filters-4",
      "state": 2,
      "output": "test-output-che-event-filters-4"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-4
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-4",
          "infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-4"
            }
          }
        },
        {
          "_id": "test-connector-che-event-filters-4/test-connector-name-che-event-filters-4"
        },
        {
          "_id": "test-resource-che-event-filters-4/test-component-che-event-filters-4",
          "component_infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-4"
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given component event should fill component infos of resource entity
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-5"
            }
          },
          {
            "field": "source_type",
            "cond": {
              "type": "eq",
              "value": "component"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.customer",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_entity_info",
            "name": "customer",
            "description": "Client",
            "value": "test-customer-che-event-filters-5"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-5-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-5",
      "connector_name": "test-connector-name-che-event-filters-5",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-5",
      "resource": "test-resource-che-event-filters-5",
      "state": 2,
      "output": "test-output-che-event-filters-5"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-5",
      "connector_name": "test-connector-name-che-event-filters-5",
      "source_type": "component",
      "event_type": "check",
      "component": "test-component-che-event-filters-5",
      "state": 2,
      "output": "test-output-che-event-filters-5"
    }
    """
    When I wait the end of 2 events processing
    When I do GET /api/v4/entities?search=che-event-filters-5
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-5",
          "infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-5"
            }
          }
        },
        {
          "_id": "test-connector-che-event-filters-5/test-connector-name-che-event-filters-5"
        },
        {
          "_id": "test-resource-che-event-filters-5/test-component-che-event-filters-5",
          "component_infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-5"
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given check event and enrichment event filter should not update disabled entity
    Given I am admin
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-6",
      "connector_name": "test-connector-name-che-event-filters-6",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-6",
      "resource": "test-resource-che-event-filters-6",
      "state": 2,
      "output": "test-output-che-event-filters-6"
    }
    """
    When I wait the end of event processing
    When I do PUT /api/v4/entitybasics?_id=test-resource-che-event-filters-6/test-component-che-event-filters-6:
    """json
    {
      "enabled": false,
      "impact_level": 1,
      "sli_avail_state": 0,
      "infos": [],
      "impact": [
        "test-component-che-event-filters-6"
      ],
      "depends": [
        "test-connector-che-event-filters-6/test-connector-name-che-event-filters-6"
      ]
    }
    """
    Then the response code should be 200
    When I wait the end of event processing
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-6"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.customer",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_entity_info",
            "name": "customer",
            "description": "Client",
            "value": "test-customer-che-event-filters-6"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-6-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-6",
      "connector_name": "test-connector-name-che-event-filters-6",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-6",
      "resource": "test-resource-che-event-filters-6",
      "state": 2,
      "output": "test-output-che-event-filters-6"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-6
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-6"
        },
        {
          "_id": "test-connector-che-event-filters-6/test-connector-name-che-event-filters-6"
        },
        {
          "_id": "test-resource-che-event-filters-6/test-component-che-event-filters-6",
          "infos": {}
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given check event and enrichment event filter with set_entity_info action
  should update event and entity
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-7"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.testdate",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_entity_info",
            "name": "testdate",
            "description": "Date",
            "value": 1592215337
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-7-2-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-7",
      "connector_name": "test-connector-name-che-event-filters-7",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-7",
      "resource": "test-resource-che-event-filters-7",
      "state": 2,
      "output": "test-output-che-event-filters-7"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-7
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-7"
        },
        {
          "_id": "test-connector-che-event-filters-7/test-connector-name-che-event-filters-7"
        },
        {
          "_id": "test-resource-che-event-filters-7/test-component-che-event-filters-7",
          "infos": {
            "testdate": {
              "description": "Date",
              "name": "testdate",
              "value": 1592215337
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given check event and enrichment event filter with copy_to_entity_info action
  should update event and entity
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-8"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.customer",
            "cond": {
              "type": "exist",
              "value": false
            }
          },
          {
            "field": "infos.testdate",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "copy_to_entity_info",
            "name": "customer",
            "description": "Client",
            "value": "Event.ExtraInfos.customer"
          },
          {
            "type": "copy_to_entity_info",
            "name": "testdate",
            "description": "Date",
            "value": "Event.ExtraInfos.testdate"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-8-2-description",
      "enabled": true,
      "priority": 2
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-8"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_field_from_template",
            "name": "Output",
            "value": "{{ `{{ .Event.Output }}` }} (client: {{ `{{ .Event.Entity.Infos.customer.Value }}` }})"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-8-3-description",
      "enabled": true,
      "priority": 3
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-8"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "infos.output",
            "cond": {
              "type": "exist",
              "value": false
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "copy_to_entity_info",
            "name": "output",
            "description": "Output",
            "value": "Event.Output"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "priority": 4,
      "description": "test-event-filter-che-event-filters-8-4-description",
      "enabled": true
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-8",
      "connector_name": "test-connector-name-che-event-filters-8",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-8",
      "resource": "test-resource-che-event-filters-8",
      "state": 2,
      "output": "test-output-che-event-filters-8",
      "customer": "test-customer-che-event-filters-8",
      "testdate": 1592215337
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=che-event-filters-8
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-component-che-event-filters-8"
        },
        {
          "_id": "test-connector-che-event-filters-8/test-connector-name-che-event-filters-8"
        },
        {
          "_id": "test-resource-che-event-filters-8/test-component-che-event-filters-8",
          "infos": {
            "customer": {
              "description": "Client",
              "name": "customer",
              "value": "test-customer-che-event-filters-8"
            },
            "testdate": {
              "description": "Date",
              "name": "testdate",
              "value": 1592215337
            },
            "output": {
              "description": "Output",
              "name": "output",
              "value": "test-output-che-event-filters-8 (client: test-customer-che-event-filters-8)"
            }
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 3
      }
    }
    """

  Scenario: given check event and drop event filter with several event patterns should drop event
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "drop",
      "event_pattern":[
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-14"
            }
          }
        ],
        [
          {
            "field": "resource",
            "cond": {
              "type": "eq",
              "value": "test-resource-che-event-filters-14"
            }
          }
        ]
      ],
      "description": "test-event-filter-che-event-filters-14-description",
      "priority": 1,
      "enabled": true
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-14",
      "connector_name": "test-connector-name-che-event-filters-14",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-14-another",
      "resource": "test-resource-che-event-filters-14",
      "state": 2,
      "output": "test-output-che-event-filters-14"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-14",
      "connector_name": "test-connector-name-che-event-filters-14",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-14",
      "resource": "test-resource-che-event-filters-14-another",
      "state": 2,
      "output": "test-output-che-event-filters-14"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/alarms?search=che-event-filters-14
    Then the response code should be 200
    Then the response body should be:
    """json
    {
      "data": [],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 0
      }
    }
    """

  Scenario: given check event and drop event filter with several event patterns should drop event
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "entity_pattern": [[
        {
          "field": "name",
          "cond": {
            "type": "regexp",
            "value": "CMDB:(?P<SI_CMDB>.*?)($|,)"
          }
        }
      ]],
      "config": {
        "actions": [
          {
            "type": "set_entity_info_from_template",
            "name": "test_template",
            "description": "test template",
            "value": "{{ `{{ .RegexMatch.Entity.Name.SI_CMDB }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      },
      "description": "test-event-filter-che-event-filters-15-description",
      "priority": 1,
      "enabled": true
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test_connector",
      "connector_name": "test_connector_name",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-che-event-filters-15",
      "resource": "CMDB:TEST_PROD",
      "state": 2
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entitybasics?_id=CMDB:TEST_PROD/test-component-che-event-filters-15
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "CMDB:TEST_PROD/test-component-che-event-filters-15",
      "infos": {
        "test_template": {
          "name": "test_template",
          "description": "test template",
          "value": "TEST_PROD"
        }
      }
    }
    """

  Scenario: given rule with old patterns format, backward compatibility test
    Given I am admin
    When I send an event:
    """json
    {
      "connector": "test-connector-eventfilter-to-backward-compatibility-1",
      "connector_name": "test-connector-name-eventfilter-to-backward-compatibility-1",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-component-eventfilter-to-backward-compatibility-1",
      "resource": "CMDD:TEST_PROD",
      "state": 2,
      "output": "test-output-eventfilter-to-backward-compatibility-1",
      "customer": "test-customer-eventfilter-to-backward-compatibility-1",
      "manager": "test-manager-eventfilter-to-backward-compatibility-1"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entitybasics?_id=CMDD:TEST_PROD/test-component-eventfilter-to-backward-compatibility-1
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "infos": {
        "customer": {
          "description": "customer",
          "name": "customer",
          "value": "TEST_PROD"
        },
        "manager": {
          "description": "manager",
          "name": "manager",
          "value": "TEST_PROD"
        }
      }
    }
    """

  Scenario: given check event and enrichment event filter should enrich from external api data where response is a document with an array
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "external_data": {
        "title": {
          "type": "api",
          "request": {
            "url": "http://localhost:3000/api/external_data_document_with_array",
            "method": "GET"
          }
        }
      },
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-eventfilter-assets-customer-3"
            }
          }
        ]
      ],
      "description": "test-event-filter-che-event-filters-16-description",
      "priority": 1,
      "enabled": true,
      "config": {
        "actions": [
          {
            "type": "set_entity_info_from_template",
            "name": "title",
            "value": "{{ `{{ index .ExternalData.title \"array.1.title\" }}` }}",
            "description": "title from external api"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-16",
      "connector_name": "test-connector-name-che-event-filters-16",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-eventfilter-assets-customer-3",
      "resource": "test-resource-che-event-filters-16",
      "state": 2,
      "output": "test-output-che-event-filters-16"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=test-resource-che-event-filters-16
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-resource-che-event-filters-16/test-eventfilter-assets-customer-3",
          "category": null,
          "component": "test-eventfilter-assets-customer-3",
          "depends": [
            "test-connector-che-event-filters-16/test-connector-name-che-event-filters-16"
          ],
          "enabled": true,
          "impact": [
            "test-eventfilter-assets-customer-3"
          ],
          "enable_history": [
            {{ .createTimestamp }}
          ],
          "impact_level": 1,
          "infos": {
            "title": {
              "name": "title",
              "description": "title from external api",
              "value": "test title 2"
            }
          },
          "measurements": null,
          "name": "test-resource-che-event-filters-16",
          "type": "resource"
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

  Scenario: given check event and enrichment event filter should enrich from external api data where response is an array
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "external_data": {
        "title": {
          "type": "api",
          "request": {
            "url": "http://localhost:3000/api/external_data_response_is_array",
            "method": "GET"
          }
        }
      },
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-eventfilter-assets-customer-4"
            }
          }
        ]
      ],
      "description": "test-event-filter-che-event-filters-17-description",
      "priority": 1,
      "enabled": true,
      "config": {
        "actions": [
          {
            "type": "set_entity_info_from_template",
            "name": "title",
            "value": "{{ `{{ index .ExternalData.title \"1.title\" }}` }}",
            "description": "title from external api"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-che-event-filters-17",
      "connector_name": "test-connector-name-che-event-filters-17",
      "source_type": "resource",
      "event_type": "check",
      "component": "test-eventfilter-assets-customer-4",
      "resource": "test-resource-che-event-filters-17",
      "state": 2,
      "output": "test-output-che-event-filters-17"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/entities?search=test-resource-che-event-filters-17
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-resource-che-event-filters-17/test-eventfilter-assets-customer-4",
          "category": null,
          "component": "test-eventfilter-assets-customer-4",
          "depends": [
            "test-connector-che-event-filters-17/test-connector-name-che-event-filters-17"
          ],
          "enabled": true,
          "impact": [
            "test-eventfilter-assets-customer-4"
          ],
          "enable_history": [
            {{ .createTimestamp }}
          ],
          "impact_level": 1,
          "infos": {
            "title": {
              "name": "title",
              "description": "title from external api",
              "value": "test title 2"
            }
          },
          "measurements": null,
          "name": "test-resource-che-event-filters-17",
          "type": "resource"
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
