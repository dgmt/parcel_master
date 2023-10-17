package parcelcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParcelCalc(t *testing.T) {
	// 0
	result := roughPack(0)
	assert.Equal(t, 0, len(result))
	resultOpt := packageNumOptimization(result)
	assert.Equal(t, 0, len(resultOpt))

	// 1
	result = roughPack(1)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 1, len(resultOpt))

	// 200
	result = roughPack(200)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 1, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[250])

	// 250
	result = roughPack(250)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 1, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[250])

	// 251
	result = roughPack(251)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, int64(2), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 1, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[500])

	// 501
	result = roughPack(501)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, int64(1), result[500])
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 2, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[500])
	assert.Equal(t, int64(1), resultOpt[250])

	// 2500
	result = roughPack(2501)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, int64(1), result[2000])
	assert.Equal(t, int64(1), result[500])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 3, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[2000])
	assert.Equal(t, int64(1), resultOpt[500])

	// 2501
	result = roughPack(2501)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, int64(1), result[2000])
	assert.Equal(t, int64(1), result[500])
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 3, len(resultOpt))
	assert.Equal(t, int64(1), resultOpt[2000])
	assert.Equal(t, int64(1), resultOpt[500])
	assert.Equal(t, int64(1), resultOpt[250])

	// 12001
	result = roughPack(12001)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, int64(2), result[5000])
	assert.Equal(t, int64(1), result[2000])
	assert.Equal(t, int64(1), result[250])
	resultOpt = packageNumOptimization(result)
	assert.Equal(t, 3, len(resultOpt))
	assert.Equal(t, int64(2), resultOpt[5000])
	assert.Equal(t, int64(1), resultOpt[2000])
	assert.Equal(t, int64(1), resultOpt[250])
}
