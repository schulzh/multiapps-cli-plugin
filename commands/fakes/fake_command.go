// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/SAP/cf-mta-plugin/commands"

	"github.com/cloudfoundry/cli/plugin"
)

type FakeCommand struct {
	GetPluginCommandStub        func() plugin.Command
	getPluginCommandMutex       sync.RWMutex
	getPluginCommandArgsForCall []struct{}
	getPluginCommandReturns     struct {
		result1 plugin.Command
	}
	InitializeStub        func(name string, cliConnection plugin.CliConnection)
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		name          string
		cliConnection plugin.CliConnection
	}
	ExecuteStub        func(args []string) commands.ExecutionStatus
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		args []string
	}
	executeReturns struct {
		result1 commands.ExecutionStatus
	}
}

func (fake *FakeCommand) GetPluginCommand() plugin.Command {
	fake.getPluginCommandMutex.Lock()
	fake.getPluginCommandArgsForCall = append(fake.getPluginCommandArgsForCall, struct{}{})
	fake.getPluginCommandMutex.Unlock()
	if fake.GetPluginCommandStub != nil {
		return fake.GetPluginCommandStub()
	} else {
		return fake.getPluginCommandReturns.result1
	}
}

func (fake *FakeCommand) GetPluginCommandCallCount() int {
	fake.getPluginCommandMutex.RLock()
	defer fake.getPluginCommandMutex.RUnlock()
	return len(fake.getPluginCommandArgsForCall)
}

func (fake *FakeCommand) GetPluginCommandReturns(result1 plugin.Command) {
	fake.GetPluginCommandStub = nil
	fake.getPluginCommandReturns = struct {
		result1 plugin.Command
	}{result1}
}

func (fake *FakeCommand) Initialize(name string, cliConnection plugin.CliConnection) {
	fake.initializeMutex.Lock()
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		name          string
		cliConnection plugin.CliConnection
	}{name, cliConnection})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		fake.InitializeStub(name, cliConnection)
	}
}

func (fake *FakeCommand) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeCommand) InitializeArgsForCall(i int) (string, plugin.CliConnection) {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return fake.initializeArgsForCall[i].name, fake.initializeArgsForCall[i].cliConnection
}

func (fake *FakeCommand) Execute(args []string) commands.ExecutionStatus {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		args []string
	}{args})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(args)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeCommand) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeCommand) ExecuteArgsForCall(i int) []string {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].args
}

func (fake *FakeCommand) ExecuteReturns(result1 commands.ExecutionStatus) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 commands.ExecutionStatus
	}{result1}
}

var _ commands.Command = new(FakeCommand)
