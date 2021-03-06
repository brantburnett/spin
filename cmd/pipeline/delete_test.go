// Copyright (c) 2018, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package pipeline

import (
	"os"
	"testing"
)

// TODO(jacobkiefer): This test overlaps heavily with pipeline_save_test.go,
// consider factoring common testing code out.
func TestPipelineDelete_basic(t *testing.T) {
	ts := GateServerSuccess()
	defer ts.Close()

	args := []string{"pipeline", "delete", "--application", "app", "--name", "one", "--gate-endpoint", ts.URL}
	currentCmd := NewDeleteCmd(pipelineOptions{})
	rootCmd := getRootCmdForTest()
	pipelineCmd := NewPipelineCmd(os.Stdout)
	pipelineCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestPipelineDelete_fail(t *testing.T) {
	ts := GateServerFail()
	defer ts.Close()

	args := []string{"pipeline", "delete", "--application", "app", "--name", "one", "--gate-endpoint", ts.URL}
	currentCmd := NewDeleteCmd(pipelineOptions{})
	rootCmd := getRootCmdForTest()
	pipelineCmd := NewPipelineCmd(os.Stdout)
	pipelineCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestPipelineDelete_flags(t *testing.T) {
	ts := GateServerSuccess()
	defer ts.Close()

	args := []string{"pipeline", "delete", "--gate-endpoint", ts.URL} // Missing pipeline app and name.
	currentCmd := NewDeleteCmd(pipelineOptions{})
	rootCmd := getRootCmdForTest()
	pipelineCmd := NewPipelineCmd(os.Stdout)
	pipelineCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command failed with: %s", err)
	}
}

func TestPipelineDelete_missingname(t *testing.T) {
	ts := GateServerSuccess()
	defer ts.Close()

	args := []string{"pipeline", "delete", "--application", "app", "--gate-endpoint", ts.URL}
	currentCmd := NewDeleteCmd(pipelineOptions{})
	rootCmd := getRootCmdForTest()
	pipelineCmd := NewPipelineCmd(os.Stdout)
	pipelineCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command errantly succeeded. %s", err)
	}
}

func TestPipelineDelete_missingapp(t *testing.T) {
	ts := GateServerSuccess()
	defer ts.Close()

	args := []string{"pipeline", "delete", "--name", "one", "--gate-endpoint", ts.URL}
	currentCmd := NewDeleteCmd(pipelineOptions{})
	rootCmd := getRootCmdForTest()
	pipelineCmd := NewPipelineCmd(os.Stdout)
	pipelineCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(pipelineCmd)

	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Command errantly succeeded. %s", err)
	}
}
