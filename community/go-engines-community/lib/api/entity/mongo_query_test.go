package entity

import (
	"context"
	"fmt"
	"testing"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/pagination"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pattern"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/savedpattern"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/view"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	mock_mongo "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/mocks/lib/mongo"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenPaginationRequest_ShouldBuildQueryWithLookupsAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	request := ListRequestWithPagination{
		Query: pagination.Query{
			Page:     2,
			Limit:    10,
			Paginate: true,
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 10},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithEntityAndPbehaviorInfoPatterns_ShouldBuildQueryWithLookupsAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	filter := view.WidgetFilter{
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "name",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"name": bson.M{"$eq": "test-resource"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAllPatterns_ShouldBuildQueryWithEntityLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	filter := view.WidgetFilter{
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
				},
			},
		},
		PbehaviorPatternFields: savedpattern.PbehaviorPatternFields{
			PbehaviorPattern: pattern.PbehaviorInfo{
				{
					{
						Field:     "canonical_type",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "pause"),
					},
				},
			},
		},
		EntityPatternFields: savedpattern.EntityPatternFields{
			EntityPattern: pattern.Entity{
				{
					{
						Field:     "category",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-category"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"category": bson.M{"$eq": "test-category"}},
		}}}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"pbehavior_info.canonical_type": bson.M{"$eq": "pause"}},
		}}}}},
	}
	expected = append(expected, getAlarmLookup()...)
	expected = append(expected,
		bson.M{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"alarm.v.resource": bson.M{"$eq": "test-resource"}},
		}}}}},
		bson.M{"$project": bson.M{"alarm": 0}},
		bson.M{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		bson.M{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	)

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAlarmPatternWithDurationField_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	durationCond, err := pattern.NewDurationCondition(pattern.ConditionGT, types.DurationWithUnit{
		Value: 10,
		Unit:  "m",
	})
	if err != nil {
		panic(err)
	}
	filter := view.WidgetFilter{
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
					{
						Field:     "v.duration",
						Condition: durationCond,
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := getAlarmLookup()
	expected = append(expected, []bson.M{
		{"$addFields": bson.M{"alarm.v.duration": getDurationField(now)}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"alarm.v.resource": bson.M{"$eq": "test-resource"}},
			{"alarm.v.duration": bson.M{"$gt": 600}},
		}}}}},
		{"$project": bson.M{"alarm": 0}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithAlarmPatternWithInfosField_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	filter := view.WidgetFilter{
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.resource",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-resource"),
					},
					{
						Field:     "v.infos.info_name",
						FieldType: pattern.FieldTypeInt,
						Condition: pattern.NewIntCondition(pattern.ConditionEqual, 3),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := getAlarmLookup()
	expected = append(expected, []bson.M{
		{"$addFields": bson.M{"alarm.v.infos_array": bson.M{"$objectToArray": "$alarm.v.infos"}}},
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"alarm.v.resource": bson.M{"$eq": "test-resource"}},
			{"$and": []bson.M{
				{"alarm.v.infos_array.v.info_name": bson.M{"$type": bson.A{"long", "int", "decimal"}}},
				{"alarm.v.infos_array.v.info_name": bson.M{"$eq": 3}},
			}},
		}}}}},
		{"$project": bson.M{"alarm": 0}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithWidgetFilterWithOldQuery_ShouldBuildQueryWithLookupsAfterLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	filter := view.WidgetFilter{
		OldMongoQuery: `{"$and": [
			{"type": "resource"},
			{"category": "test-category"}
		]}`,
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$match": bson.M{"$and": []bson.M{
			{"type": "resource"},
			{"category": "test-category"},
		}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithAlarmSort_ShouldBuildQueryWithLookupsBeforeLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterId := "test-filter"
	filter := view.WidgetFilter{
		AlarmPatternFields: savedpattern.AlarmPatternFields{
			AlarmPattern: pattern.Alarm{
				{
					{
						Field:     "v.component",
						Condition: pattern.NewStringCondition(pattern.ConditionEqual, "test-component"),
					},
				},
			},
		},
	}
	mockDbClient := createMockDbClientWithFilterFetching(ctrl, filterId, filter)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Filter: filterId,
			},
			SortBy: "alarm.v.resource",
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{
			{Key: "alarm.v.resource", Value: 1},
			{Key: "_id", Value: 1},
		}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := getAlarmLookup()
	expected = append(expected, []bson.M{
		{"$match": bson.M{"$or": []bson.M{{"$and": []bson.M{
			{"alarm.v.component": bson.M{"$eq": "test-component"}},
		}}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}...)

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearch_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	search := "test-search"
	searchRegexp := primitive.Regex{
		Pattern: fmt.Sprintf(".*%s.*", search),
		Options: "i",
	}
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Search: search,
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$match": bson.M{"$and": []bson.M{
			{"$or": []bson.M{
				{"_id": searchRegexp},
				{"name": searchRegexp},
				{"type": searchRegexp},
			}},
		}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func TestMongoQueryBuilder_CreateListAggregationPipeline_GivenRequestWithSearchExpression_ShouldBuildQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockDbClient := createMockDbClient(ctrl)
	request := ListRequestWithPagination{
		Query: pagination.GetDefaultQuery(),
		ListRequest: ListRequest{
			BaseFilterRequest: BaseFilterRequest{
				Search: "infos.test1.value LIKE \"test val\"",
			},
		},
	}
	now := types.NewCpsTime()
	expectedDataPipeline := []bson.M{
		{"$sort": bson.D{{Key: "_id", Value: 1}}},
		{"$skip": 0},
		{"$limit": 10},
	}
	expectedDataPipeline = append(expectedDataPipeline, getAlarmLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getCategoryLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getPbehaviorInfoTypeLookup()...)
	expectedDataPipeline = append(expectedDataPipeline, getEventStatsLookup(now)...)
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$addFields": getComputedFields(),
	})
	expectedDataPipeline = append(expectedDataPipeline, bson.M{
		"$project": bson.M{
			"alarm":               0,
			"event_stats":         0,
			"pbehavior_info_type": 0,
		},
	})
	expected := []bson.M{
		{"$match": bson.M{"$and": []bson.M{{"infos.test1.value": bson.M{"$regex": "test val"}}}}},
		{"$facet": bson.M{
			"data":        expectedDataPipeline,
			"total_count": []bson.M{{"$count": "count"}},
		}},
		{"$addFields": bson.M{
			"total_count": bson.M{"$sum": "$total_count.count"},
		}},
	}

	b := NewMongoQueryBuilder(mockDbClient)
	result, err := b.CreateListAggregationPipeline(ctx, request, now)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if diff := pretty.Compare(result, expected); diff != "" {
		t.Errorf("unexpected result: %s", diff)
	}
}

func createMockDbClient(ctrl *gomock.Controller) mongo.DbClient {
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		default:
			return nil
		}
	}).AnyTimes()

	return mockDbClient
}

func createMockDbClientWithFilterFetching(ctrl *gomock.Controller, filterId string, filter view.WidgetFilter) mongo.DbClient {
	mockSingleResult := mock_mongo.NewMockSingleResultHelper(ctrl)
	mockSingleResult.EXPECT().Decode(gomock.Any()).Do(func(v *view.WidgetFilter) {
		*v = filter
	})
	mockFilterDbCollection := mock_mongo.NewMockDbCollection(ctrl)
	mockFilterDbCollection.EXPECT().FindOne(gomock.Any(), gomock.Eq(bson.M{"_id": filterId})).Return(mockSingleResult)
	mockDbClient := mock_mongo.NewMockDbClient(ctrl)
	mockDbClient.EXPECT().Collection(gomock.Any()).DoAndReturn(func(name string) mongo.DbCollection {
		switch name {
		case mongo.WidgetFiltersMongoCollection:
			return mockFilterDbCollection
		default:
			return nil
		}
	}).AnyTimes()

	return mockDbClient
}
