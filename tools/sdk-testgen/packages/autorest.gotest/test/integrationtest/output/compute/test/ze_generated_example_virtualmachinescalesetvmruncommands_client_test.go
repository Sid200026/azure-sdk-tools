//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package test_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateOrUpdateVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := test.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		test.VirtualMachineRunCommand{
			Location: to.StringPtr("<location>"),
			Properties: &test.VirtualMachineRunCommandProperties{
				AsyncExecution: to.BoolPtr(false),
				Parameters: []*test.RunCommandInputParameter{
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					},
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					}},
				RunAsPassword: to.StringPtr("<run-as-password>"),
				RunAsUser:     to.StringPtr("<run-as-user>"),
				Source: &test.VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("<script>"),
				},
				TimeoutInSeconds: to.Int32Ptr(3600),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res.VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/UpdateVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := test.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		test.VirtualMachineRunCommandUpdate{
			Properties: &test.VirtualMachineRunCommandProperties{
				Source: &test.VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("<script>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res.VirtualMachineScaleSetVMRunCommandsClientUpdateResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/DeleteVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := test.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := test.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		&test.VirtualMachineScaleSetVMRunCommandsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res.VirtualMachineScaleSetVMRunCommandsClientGetResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := test.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&test.VirtualMachineScaleSetVMRunCommandsClientListOptions{Expand: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			// TODO: use page item
			_ = v
		}
	}
}