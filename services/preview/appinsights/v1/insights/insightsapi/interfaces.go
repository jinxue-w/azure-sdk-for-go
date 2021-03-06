package insightsapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/appinsights/v1/insights"
)

// MetricsClientAPI contains the set of methods on the MetricsClient type.
type MetricsClientAPI interface {
	Get(ctx context.Context, appID string, metricID insights.MetricID, timespan string, interval *string, aggregation []insights.MetricsAggregation, segment []insights.MetricsSegment, top *int32, orderby string, filter string) (result insights.MetricsResult, err error)
	GetMetadata(ctx context.Context, appID string) (result insights.SetObject, err error)
	GetMultiple(ctx context.Context, appID string, body []insights.MetricsPostBodySchema) (result insights.ListMetricsResultsItem, err error)
}

var _ MetricsClientAPI = (*insights.MetricsClient)(nil)

// EventsClientAPI contains the set of methods on the EventsClient type.
type EventsClientAPI interface {
	Get(ctx context.Context, appID string, eventType insights.EventType, eventID string, timespan string) (result insights.EventsResults, err error)
	GetByType(ctx context.Context, appID string, eventType insights.EventType, timespan string, filter string, search string, orderby string, selectParameter string, skip *int32, top *int32, formatParameter string, count *bool, apply string) (result insights.EventsResults, err error)
	GetOdataMetadata(ctx context.Context, appID string) (result insights.SetObject, err error)
}

var _ EventsClientAPI = (*insights.EventsClient)(nil)

// QueryClientAPI contains the set of methods on the QueryClient type.
type QueryClientAPI interface {
	Execute(ctx context.Context, appID string, body insights.QueryBody) (result insights.QueryResults, err error)
}

var _ QueryClientAPI = (*insights.QueryClient)(nil)

// PostClientAPI contains the set of methods on the PostClient type.
type PostClientAPI interface {
	Metadata(ctx context.Context, appID string) (result insights.MetadataResults, err error)
}

var _ PostClientAPI = (*insights.PostClient)(nil)

// GetClientAPI contains the set of methods on the GetClient type.
type GetClientAPI interface {
	Metadata(ctx context.Context, appID string) (result insights.MetadataResults, err error)
}

var _ GetClientAPI = (*insights.GetClient)(nil)
