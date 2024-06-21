package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStdErrLogger(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)
	err = os.WriteFile(tmpDir+"/2020-01-01", []byte("test"), 0o644)
	require.NoError(t, err)

	logger, err := NewStdErrLogger(tmpDir+"/stderr.log", time.Second)
	require.NoError(t, err)
	os.Stderr.Write([]byte("test"))
	time.Sleep(time.Second)
	files, err := os.ReadDir(tmpDir)
	require.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Equal(t, "2020-01-01", files[0].Name())
	assert.Equal(t, "stderr.log", files[1].Name())
	actual, err := os.ReadFile(tmpDir + "/stderr.log")
	require.NoError(t, err)
	assert.Equal(t, "test", string(actual))
	err = logger.Close()
	assert.NoError(t, err)

	logger, err = NewStdErrLogger(tmpDir+"/stderr.log", time.Second)
	require.NoError(t, err)
	os.Stderr.Write([]byte("test2"))
	files, err = os.ReadDir(tmpDir)
	require.NoError(t, err)
	assert.Len(t, files, 1)
	assert.Equal(t, "stderr.log", files[0].Name())
	actual, err = os.ReadFile(tmpDir + "/stderr.log")
	require.NoError(t, err)
	assert.Equal(t, "test2", string(actual))
	err = logger.Close()
	assert.NoError(t, err)
}
