package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func StartDockerImage(ctx context.Context, image string, port string) (types.ContainerJSON, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return types.ContainerJSON{}, err
	}

	_, err = cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return types.ContainerJSON{}, err
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("7007" + "/tcp"): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}

	config := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			nat.Port(port + "/tcp"): struct{}{},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return types.ContainerJSON{}, err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return types.ContainerJSON{}, err
	}

	containerInfo, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return containerInfo, nil
}

func StopDockerImage(ctx context.Context, containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	timeout := 5
	err = cli.ContainerStop(ctx, containerID, container.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		return err
	}

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ReloadDockerImage(ctx context.Context, containerID string, image string, port string) (types.ContainerJSON, error) {
	// Step 1: Stop the current container
	err := StopDockerImage(ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	// Step 2: Start a new container with the latest image
	containerInfo, err := StartDockerImage(ctx, image, port)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return containerInfo, nil
}
