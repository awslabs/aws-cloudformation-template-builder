package main

import "testing"

func TestNameFromAWSType(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"AWS::S3::Bucket", "AWSS3Bucket"},
		{"AWS::CodeBuild::Project.Artifacts", "AWSCodeBuildProject_Artifacts"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nameFromAWSType(tt.name); got != tt.want {
				t.Errorf("nameFromAWSType() = %v, want %v", got, tt.want)
			}
		})
	}
}
