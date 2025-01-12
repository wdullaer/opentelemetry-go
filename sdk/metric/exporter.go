// Copyright The OpenTelemetry Authors
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

//go:build go1.18
// +build go1.18

package metric // import "go.opentelemetry.io/otel/sdk/metric"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// ErrExporterShutdown is returned if Export or Shutdown are called after an
// Exporter has been Shutdown.
var ErrExporterShutdown = fmt.Errorf("exporter is shutdown")

// Exporter handles the delivery of metric data to external receivers. This is
// the final component in the metric push pipeline.
type Exporter interface {
	// Export serializes and transmits metric data to a receiver.
	//
	// This is called synchronously, there is no concurrency safety
	// requirement. Because of this, it is critical that all timeouts and
	// cancellations of the passed context be honored.
	//
	// All retry logic must be contained in this function. The SDK does not
	// implement any retry logic. All errors returned by this function are
	// considered unrecoverable and will be reported to a configured error
	// Handler.
	Export(context.Context, metricdata.ResourceMetrics) error

	// ForceFlush flushes any metric data held by an exporter.
	//
	// The deadline or cancellation of the passed context must be honored. An
	// appropriate error should be returned in these situations.
	ForceFlush(context.Context) error

	// Shutdown flushes all metric data held by an exporter and releases any
	// held computational resources.
	//
	// The deadline or cancellation of the passed context must be honored. An
	// appropriate error should be returned in these situations.
	//
	// After Shutdown is called, calls to Export will perform no operation and
	// instead will return an error indicating the shutdown state.
	Shutdown(context.Context) error
}
