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

package eiffelevents

// TODO: Make sure this variable is generated by a program.
var eventRoundTripTestTable = []struct {
	Filename     string
	ExpectedType interface{}
}{
	{"protocol/examples/events/EiffelActivityCanceledEvent/simple.json", &ActivityCanceledV3{}},
	{"protocol/examples/events/EiffelActivityFinishedEvent/simple.json", &ActivityFinishedV3{}},
	{"protocol/examples/events/EiffelActivityStartedEvent/simple.json", &ActivityStartedV4{}},
	{"protocol/examples/events/EiffelActivityTriggeredEvent/simple-customdata.json", &ActivityTriggeredV3{}},
	{"protocol/examples/events/EiffelActivityTriggeredEvent/simple.json", &ActivityTriggeredV3{}},
	{"protocol/examples/events/EiffelAnnouncementPublishedEvent/simple.json", &AnnouncementPublishedV3{}},
	{"protocol/examples/events/EiffelArtifactCreatedEvent/backend.json", &ArtifactCreatedV3{}},
	{"protocol/examples/events/EiffelArtifactCreatedEvent/dependent.json", &ArtifactCreatedV3{}},
	{"protocol/examples/events/EiffelArtifactCreatedEvent/interface.json", &ArtifactCreatedV3{}},
	{"protocol/examples/events/EiffelArtifactCreatedEvent/simple.json", &ArtifactCreatedV3{}},
	{"protocol/examples/events/EiffelArtifactPublishedEvent/multifile.json", &ArtifactPublishedV3{}},
	{"protocol/examples/events/EiffelArtifactPublishedEvent/simple.json", &ArtifactPublishedV3{}},
	{"protocol/examples/events/EiffelArtifactReusedEvent/simple.json", &ArtifactReusedV3{}},
	{"protocol/examples/events/EiffelCompositionDefinedEvent/simple.json", &CompositionDefinedV3{}},
	{"protocol/examples/events/EiffelConfidenceLevelModifiedEvent/simple.json", &ConfidenceLevelModifiedV3{}},
	{"protocol/examples/events/EiffelEnvironmentDefinedEvent/host.json", &EnvironmentDefinedV3{}},
	{"protocol/examples/events/EiffelEnvironmentDefinedEvent/image.json", &EnvironmentDefinedV3{}},
	{"protocol/examples/events/EiffelEnvironmentDefinedEvent/runtime-env-link.json", &EnvironmentDefinedV3{}},
	{"protocol/examples/events/EiffelEnvironmentDefinedEvent/uri.json", &EnvironmentDefinedV3{}},
	{"protocol/examples/events/EiffelFlowContextDefinedEvent/simple.json", &FlowContextDefinedV3{}},
	{"protocol/examples/events/EiffelIssueDefinedEvent/simple.json", &IssueDefinedV3{}},
	{"protocol/examples/events/EiffelIssueVerifiedEvent/simple.json", &IssueVerifiedV4{}},
	{"protocol/examples/events/EiffelSourceChangeCreatedEvent/simple.json", &SourceChangeCreatedV4{}},
	{"protocol/examples/events/EiffelSourceChangeSubmittedEvent/simple.json", &SourceChangeSubmittedV3{}},
	{"protocol/examples/events/EiffelTestCaseCanceledEvent/simple.json", &TestCaseCanceledV3{}},
	{"protocol/examples/events/EiffelTestCaseFinishedEvent/simple.json", &TestCaseFinishedV3{}},
	{"protocol/examples/events/EiffelTestCaseStartedEvent/simple.json", &TestCaseStartedV3{}},
	{"protocol/examples/events/EiffelTestCaseTriggeredEvent/simple-2.0.0.json", &TestCaseTriggeredV2{}},
	{"protocol/examples/events/EiffelTestCaseTriggeredEvent/simple.json", &TestCaseTriggeredV3{}},
	{"protocol/examples/events/EiffelTestExecutionRecipeCollectionCreatedEvent/batches-1.0.0.json", &TestExecutionRecipeCollectionCreatedV1{}},
	{"protocol/examples/events/EiffelTestExecutionRecipeCollectionCreatedEvent/batches.json", &TestExecutionRecipeCollectionCreatedV4{}},
	{"protocol/examples/events/EiffelTestExecutionRecipeCollectionCreatedEvent/batchesUri.json", &TestExecutionRecipeCollectionCreatedV4{}},
	{"protocol/examples/events/EiffelTestSuiteFinishedEvent/simple.json", &TestSuiteFinishedV3{}},
	{"protocol/examples/events/EiffelTestSuiteStartedEvent/simple.json", &TestSuiteStartedV3{}},
}
