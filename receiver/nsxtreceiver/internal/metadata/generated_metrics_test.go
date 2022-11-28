// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), component.BuildInfo{}, WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["nsxt.node.cpu.utilization"] = true
	mb.RecordNsxtNodeCPUUtilizationDataPoint(ts, 1, AttributeClass(1))

	enabledMetrics["nsxt.node.filesystem.usage"] = true
	mb.RecordNsxtNodeFilesystemUsageDataPoint(ts, 1, AttributeDiskState(1))

	enabledMetrics["nsxt.node.filesystem.utilization"] = true
	mb.RecordNsxtNodeFilesystemUtilizationDataPoint(ts, 1)

	enabledMetrics["nsxt.node.memory.cache.usage"] = true
	mb.RecordNsxtNodeMemoryCacheUsageDataPoint(ts, 1)

	enabledMetrics["nsxt.node.memory.usage"] = true
	mb.RecordNsxtNodeMemoryUsageDataPoint(ts, 1)

	enabledMetrics["nsxt.node.network.io"] = true
	mb.RecordNsxtNodeNetworkIoDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["nsxt.node.network.packet.count"] = true
	mb.RecordNsxtNodeNetworkPacketCountDataPoint(ts, 1, AttributeDirection(1), AttributePacketType(1))

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		NsxtNodeCPUUtilization:        MetricSettings{Enabled: true},
		NsxtNodeFilesystemUsage:       MetricSettings{Enabled: true},
		NsxtNodeFilesystemUtilization: MetricSettings{Enabled: true},
		NsxtNodeMemoryCacheUsage:      MetricSettings{Enabled: true},
		NsxtNodeMemoryUsage:           MetricSettings{Enabled: true},
		NsxtNodeNetworkIo:             MetricSettings{Enabled: true},
		NsxtNodeNetworkPacketCount:    MetricSettings{Enabled: true},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))

	mb.RecordNsxtNodeCPUUtilizationDataPoint(ts, 1, AttributeClass(1))
	mb.RecordNsxtNodeFilesystemUsageDataPoint(ts, 1, AttributeDiskState(1))
	mb.RecordNsxtNodeFilesystemUtilizationDataPoint(ts, 1)
	mb.RecordNsxtNodeMemoryCacheUsageDataPoint(ts, 1)
	mb.RecordNsxtNodeMemoryUsageDataPoint(ts, 1)
	mb.RecordNsxtNodeNetworkIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordNsxtNodeNetworkPacketCountDataPoint(ts, 1, AttributeDirection(1), AttributePacketType(1))

	metrics := mb.Emit(WithDeviceID("attr-val"), WithNsxtNodeID("attr-val"), WithNsxtNodeName("attr-val"), WithNsxtNodeType("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("device.id")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("nsxt.node.id")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("nsxt.node.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("nsxt.node.type")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "nsxt.node.cpu.utilization":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The average amount of CPU being used by the node.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("class")
			assert.True(t, ok)
			assert.Equal(t, "datapath", attrVal.Str())
			validatedMetrics["nsxt.node.cpu.utilization"] = struct{}{}
		case "nsxt.node.filesystem.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of storage space used by the node.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "used", attrVal.Str())
			validatedMetrics["nsxt.node.filesystem.usage"] = struct{}{}
		case "nsxt.node.filesystem.utilization":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The percentage of storage space utilized.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["nsxt.node.filesystem.utilization"] = struct{}{}
		case "nsxt.node.memory.cache.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The size of the node's memory cache.", ms.At(i).Description())
			assert.Equal(t, "KBy", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["nsxt.node.memory.cache.usage"] = struct{}{}
		case "nsxt.node.memory.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The memory usage of the node.", ms.At(i).Description())
			assert.Equal(t, "KBy", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["nsxt.node.memory.usage"] = struct{}{}
		case "nsxt.node.network.io":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of bytes which have flowed through the network interface.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "received", attrVal.Str())
			validatedMetrics["nsxt.node.network.io"] = struct{}{}
		case "nsxt.node.network.packet.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of packets which have flowed through the network interface on the node.", ms.At(i).Description())
			assert.Equal(t, "{packets}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "received", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "dropped", attrVal.Str())
			validatedMetrics["nsxt.node.network.packet.count"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		NsxtNodeCPUUtilization:        MetricSettings{Enabled: false},
		NsxtNodeFilesystemUsage:       MetricSettings{Enabled: false},
		NsxtNodeFilesystemUtilization: MetricSettings{Enabled: false},
		NsxtNodeMemoryCacheUsage:      MetricSettings{Enabled: false},
		NsxtNodeMemoryUsage:           MetricSettings{Enabled: false},
		NsxtNodeNetworkIo:             MetricSettings{Enabled: false},
		NsxtNodeNetworkPacketCount:    MetricSettings{Enabled: false},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))
	mb.RecordNsxtNodeCPUUtilizationDataPoint(ts, 1, AttributeClass(1))
	mb.RecordNsxtNodeFilesystemUsageDataPoint(ts, 1, AttributeDiskState(1))
	mb.RecordNsxtNodeFilesystemUtilizationDataPoint(ts, 1)
	mb.RecordNsxtNodeMemoryCacheUsageDataPoint(ts, 1)
	mb.RecordNsxtNodeMemoryUsageDataPoint(ts, 1)
	mb.RecordNsxtNodeNetworkIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordNsxtNodeNetworkPacketCountDataPoint(ts, 1, AttributeDirection(1), AttributePacketType(1))

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
