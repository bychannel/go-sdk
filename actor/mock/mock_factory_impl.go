/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mock

import (
	"context"

	"github.com/dapr/go-sdk/actor"
)

func ActorImplFactory() actor.Server {
	return &ActorImpl{}
}

type ActorImpl struct {
	actor.ServerImplBase
}

func (t *ActorImpl) Activate(invokeName string) error {
	return nil
}

func (t *ActorImpl) Deactivate() error {
	return nil
}

func (t *ActorImpl) Type() string {
	return "testActorType"
}

func (t *ActorImpl) Invoke(ctx context.Context, req string) (string, error) {
	return req, nil
}

func (t *ActorImpl) ReminderCall(reminderName string, state []byte, dueTime string, period string) {
}

func NotReminderCalleeActorFactory() actor.Server {
	return &NotReminderCalleeActor{}
}

type NotReminderCalleeActor struct {
	actor.ServerImplBase
}

func (t *NotReminderCalleeActor) Activate() error {
	return nil
}

func (t *NotReminderCalleeActor) Deactivate() error {
	return nil
}

func (t *NotReminderCalleeActor) Type() string {
	return "testActorNotReminderCalleeType"
}
