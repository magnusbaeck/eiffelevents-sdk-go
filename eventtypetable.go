// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED AND MUST NOT BE EDITED BY HAND.

package eiffelevents

import "reflect"

// eventTypeTable maps the major version of each event to a type reference to
// the Go type used to represent the event.
var eventTypeTable = map[string]map[int64]reflect.Type{
	"EiffelActivityCanceledEvent": {
		1: reflect.TypeOf(ActivityCanceledV1{}),
		2: reflect.TypeOf(ActivityCanceledV2{}),
		3: reflect.TypeOf(ActivityCanceledV3{}),
	},
	"EiffelActivityFinishedEvent": {
		1: reflect.TypeOf(ActivityFinishedV1{}),
		2: reflect.TypeOf(ActivityFinishedV2{}),
		3: reflect.TypeOf(ActivityFinishedV3{}),
	},
	"EiffelActivityStartedEvent": {
		1: reflect.TypeOf(ActivityStartedV1{}),
		2: reflect.TypeOf(ActivityStartedV2{}),
		3: reflect.TypeOf(ActivityStartedV3{}),
		4: reflect.TypeOf(ActivityStartedV4{}),
	},
	"EiffelActivityTriggeredEvent": {
		1: reflect.TypeOf(ActivityTriggeredV1{}),
		2: reflect.TypeOf(ActivityTriggeredV2{}),
		3: reflect.TypeOf(ActivityTriggeredV3{}),
		4: reflect.TypeOf(ActivityTriggeredV4{}),
	},
	"EiffelAnnouncementPublishedEvent": {
		1: reflect.TypeOf(AnnouncementPublishedV1{}),
		2: reflect.TypeOf(AnnouncementPublishedV2{}),
		3: reflect.TypeOf(AnnouncementPublishedV3{}),
	},
	"EiffelArtifactCreatedEvent": {
		1: reflect.TypeOf(ArtifactCreatedV1{}),
		2: reflect.TypeOf(ArtifactCreatedV2{}),
		3: reflect.TypeOf(ArtifactCreatedV3{}),
	},
	"EiffelArtifactPublishedEvent": {
		1: reflect.TypeOf(ArtifactPublishedV1{}),
		2: reflect.TypeOf(ArtifactPublishedV2{}),
		3: reflect.TypeOf(ArtifactPublishedV3{}),
	},
	"EiffelArtifactReusedEvent": {
		1: reflect.TypeOf(ArtifactReusedV1{}),
		2: reflect.TypeOf(ArtifactReusedV2{}),
		3: reflect.TypeOf(ArtifactReusedV3{}),
	},
	"EiffelCompositionDefinedEvent": {
		1: reflect.TypeOf(CompositionDefinedV1{}),
		2: reflect.TypeOf(CompositionDefinedV2{}),
		3: reflect.TypeOf(CompositionDefinedV3{}),
	},
	"EiffelConfidenceLevelModifiedEvent": {
		1: reflect.TypeOf(ConfidenceLevelModifiedV1{}),
		2: reflect.TypeOf(ConfidenceLevelModifiedV2{}),
		3: reflect.TypeOf(ConfidenceLevelModifiedV3{}),
	},
	"EiffelEnvironmentDefinedEvent": {
		1: reflect.TypeOf(EnvironmentDefinedV1{}),
		2: reflect.TypeOf(EnvironmentDefinedV2{}),
		3: reflect.TypeOf(EnvironmentDefinedV3{}),
	},
	"EiffelFlowContextDefinedEvent": {
		1: reflect.TypeOf(FlowContextDefinedV1{}),
		2: reflect.TypeOf(FlowContextDefinedV2{}),
		3: reflect.TypeOf(FlowContextDefinedV3{}),
	},
	"EiffelIssueDefinedEvent": {
		1: reflect.TypeOf(IssueDefinedV1{}),
		2: reflect.TypeOf(IssueDefinedV2{}),
		3: reflect.TypeOf(IssueDefinedV3{}),
	},
	"EiffelIssueVerifiedEvent": {
		1: reflect.TypeOf(IssueVerifiedV1{}),
		2: reflect.TypeOf(IssueVerifiedV2{}),
		3: reflect.TypeOf(IssueVerifiedV3{}),
		4: reflect.TypeOf(IssueVerifiedV4{}),
	},
	"EiffelSourceChangeCreatedEvent": {
		1: reflect.TypeOf(SourceChangeCreatedV1{}),
		2: reflect.TypeOf(SourceChangeCreatedV2{}),
		3: reflect.TypeOf(SourceChangeCreatedV3{}),
		4: reflect.TypeOf(SourceChangeCreatedV4{}),
	},
	"EiffelSourceChangeSubmittedEvent": {
		1: reflect.TypeOf(SourceChangeSubmittedV1{}),
		2: reflect.TypeOf(SourceChangeSubmittedV2{}),
		3: reflect.TypeOf(SourceChangeSubmittedV3{}),
	},
	"EiffelTestCaseCanceledEvent": {
		1: reflect.TypeOf(TestCaseCanceledV1{}),
		2: reflect.TypeOf(TestCaseCanceledV2{}),
		3: reflect.TypeOf(TestCaseCanceledV3{}),
	},
	"EiffelTestCaseFinishedEvent": {
		1: reflect.TypeOf(TestCaseFinishedV1{}),
		2: reflect.TypeOf(TestCaseFinishedV2{}),
		3: reflect.TypeOf(TestCaseFinishedV3{}),
	},
	"EiffelTestCaseStartedEvent": {
		1: reflect.TypeOf(TestCaseStartedV1{}),
		2: reflect.TypeOf(TestCaseStartedV2{}),
		3: reflect.TypeOf(TestCaseStartedV3{}),
	},
	"EiffelTestCaseTriggeredEvent": {
		1: reflect.TypeOf(TestCaseTriggeredV1{}),
		2: reflect.TypeOf(TestCaseTriggeredV2{}),
		3: reflect.TypeOf(TestCaseTriggeredV3{}),
	},
	"EiffelTestExecutionRecipeCollectionCreatedEvent": {
		1: reflect.TypeOf(TestExecutionRecipeCollectionCreatedV1{}),
		2: reflect.TypeOf(TestExecutionRecipeCollectionCreatedV2{}),
		3: reflect.TypeOf(TestExecutionRecipeCollectionCreatedV3{}),
		4: reflect.TypeOf(TestExecutionRecipeCollectionCreatedV4{}),
	},
	"EiffelTestSuiteFinishedEvent": {
		1: reflect.TypeOf(TestSuiteFinishedV1{}),
		2: reflect.TypeOf(TestSuiteFinishedV2{}),
		3: reflect.TypeOf(TestSuiteFinishedV3{}),
	},
	"EiffelTestSuiteStartedEvent": {
		1: reflect.TypeOf(TestSuiteStartedV1{}),
		2: reflect.TypeOf(TestSuiteStartedV2{}),
		3: reflect.TypeOf(TestSuiteStartedV3{}),
	},
}
