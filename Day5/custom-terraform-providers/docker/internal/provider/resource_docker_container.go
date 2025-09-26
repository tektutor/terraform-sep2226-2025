package provider

import (
	"context"
	"log"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerContainer() *schema.Resource {
	return &schema.Resource{
		Description: "This resource will support Create, Read, Update and Delete docker container resources via Terraform.",

		CreateContext: resourceCreateDockerContainer,
		ReadContext:   resourceReadDockerContainer,
		UpdateContext: resourceUpdateDockerContainer,
		DeleteContext: resourceDeleteDockerContainer,

		Schema: map[string]*schema.Schema{
			"container_name": {
				Description: "Name of the container",
				Type:        schema.TypeString,
				Required:    true,
			},

			"host_name": {
				Description: "Hostname of the container",
				Type:        schema.TypeString,
				Required:    true,
			},
			"image_name": {
				Description: "Name of the docker image",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	//Retrieve inputs from terraform manifest scripts provided by users
	imageName := d.Get("image_name").(string)
	containerName := d.Get("container_name").(string)
	containerHostname := d.Get("host_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Unable to communicate with docker engine")
		panic(err)
	}
	defer cli.Close()

	resp, err1 := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image:    imageName,
			Hostname: containerHostname,
		}, nil, nil, nil, containerName,
	)
	if err1 != nil {
		log.Printf("Unable to create docker container: %s", containerName)
		panic(err1)
	}

	if err2 := cli.ContainerStart(ctx, resp.ID, client.ContainerStartOptions{}); err2 != nil {
		log.Printf("Unable to start the docker container: %s", containerName)
		panic(err2)
	}

	d.Set("container_name", containerName)
	d.Set("host_name", containerHostname)
	d.SetId(resp.ID)

	return nil
}

func resourceReadDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceUpdateDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve inputs from terraform manifest scripts provided by users
	containerName := d.Get("container_name").(string)
	containerID := d.Id()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Unable to communicate with docker engine")
		panic(err)
	}
	defer cli.Close()

	err = cli.ContainerRename(ctx, containerID, containerName)
	if err != nil {
		log.Printf("Unable to rename the docker container: %s", containerName)
		panic(err)
	}

	return nil
}

func resourceDeleteDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve inputs from terraform manifest scripts provided by users
	containerName := d.Get("container_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Unable to communicate with docker engine")
		panic(err)
	}
	defer cli.Close()

	err = cli.ContainerStop(ctx, containerName, client.ContainerStopOptions{})
	if err != nil {
		log.Printf("Unable to stop the docker container: %s", containerName)
		panic(err)
	}

	err = cli.ContainerRemove(ctx, containerName, client.ContainerRemoveOptions{RemoveVolumes: true, Force: true})
	if err != nil {
		log.Printf("Unable to delete container: %s", containerName)
		panic(err)
	}

	return nil
}
