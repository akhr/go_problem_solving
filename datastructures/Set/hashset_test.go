package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set_Add(t *testing.T) {
	set := NewSet()
	set.Add("akhash")
	set.Add("janani")
	set.Add("diya")
	assert.True(t, set.Contains("akhash"))
	assert.True(t, set.Contains("janani"))
	assert.True(t, set.Contains("diya"))
	assert.False(t, set.Contains("rajesh"))
}

func Test_Set_Delete(t *testing.T) {
	set := NewSet()
	assert.True(t, set.Add("akhash"))
	assert.False(t, set.Add("akhash"))
	assert.True(t, set.Add("janani"))

	assert.True(t, set.Contains("akhash"))
	assert.True(t, set.Contains("janani"))

	assert.True(t, set.Delete("akhash"))
	assert.False(t, set.Contains("akhash"))

	assert.False(t, set.Delete("rajesh"))
}
