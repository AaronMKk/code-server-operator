/*
Copyright 2019 tommylikehu@gmail.com.

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

package controllers

import (
	"k8s.io/apimachinery/pkg/types"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CodeServerOption struct {
	DomainName          string
	VSExporterImage     string
	ProbeInterval       int
	MaxProbeRetry       int
	HttpsSecretName     string
	LxdClientSecretName string
	EnableUserIngress   bool
	MaxConcurrency      int
}

type WatchType string

const (
	AddInactiveWatch    WatchType = "AddInactiveWatch"
	DeleteInactiveWatch WatchType = "DeleteInactiveWatch"
	AddRecycleWatch     WatchType = "AddRecycleWatch"
	DeleteRecycleWatch  WatchType = "DeleteRecycleWatch"
)

type CodeServerRequest struct {
	resource               types.NamespacedName
	duration               int64
	operate                WatchType
	endpoint               string
	inactiveTime           metav1.Time
	increaseRecycleSeconds bool
}
