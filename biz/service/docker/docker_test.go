package docker

import (
	"context"
	"testing"
)

func TestStartDockerImage(t *testing.T) {
	ctx := context.Background()
	image := "lyleshaw/gradio-test:latest"
	port := "8080"

	containerInfo, err := StartDockerImage(ctx, image, port)
	if err != nil {
		t.Errorf("startDockerImage() error: %v", err)
	}
	if containerInfo.ID == "" {
		t.Error("startDockerImage(): empty container ID")
	}

	err = StopDockerImage(ctx, containerInfo.ID)
	if err != nil {
		t.Errorf("stopDockerImage() error: %v", err)
	}
}

func TestStopDockerImage(t *testing.T) {
	ctx := context.Background()
	image := "lyleshaw/gradio-test:latest"
	port := "8081"

	containerInfo, err := StartDockerImage(ctx, image, port)

	err = StopDockerImage(ctx, containerInfo.ID)
	if err != nil {
		t.Errorf("stopDockerImage() error: %v", err)
	}
}

func TestReloadDockerImage(t *testing.T) {
	ctx := context.Background()
	image := "lyleshaw/gradio-test:latest"
	port := "8081"

	containerInfo, err := StartDockerImage(ctx, image, port)

	containerInfo, err = ReloadDockerImage(ctx, containerInfo.ID, image, port)
	if err != nil {
		t.Errorf("reloadDockerImage() error: %v", err)
	}

	err = StopDockerImage(ctx, containerInfo.ID)
	if err != nil {
		t.Errorf("stopDockerImage() error: %v", err)
	}
}
