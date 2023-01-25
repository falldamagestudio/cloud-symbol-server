using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Xunit;
using Xunit.Abstractions;

namespace cloud_symbol_server_cli.Tests;

public static class SpecRunner
{
    private class Spec
    {
        [JsonProperty(Required = Required.Always)]
        public List<string> CommandArguments = new List<string>();
        
        [JsonProperty(Required = Required.Always)]
        public int ExitCode = 0;
        
        public string Stdout = "";
        public string Stderr = "";
    }

    private static Spec ReadSpec(string directory)
    {
        const string specFileName = "spec.json";
        const string stdoutFileName = "stdout.txt";
        const string stderrFileName = "stderr.txt";

        try {
            string specJson = File.ReadAllText(System.IO.Path.Combine(directory, specFileName));
            Spec spec = JsonConvert.DeserializeObject<Spec>(specJson);
            spec.Stdout = File.ReadAllText(System.IO.Path.Combine(directory, stdoutFileName));
            spec.Stderr = File.ReadAllText(System.IO.Path.Combine(directory, stderrFileName));
            return spec;
        } catch (Exception e) {
            throw new ApplicationException($"Error reading {specFileName}: {e.Message}");
        }
    }

    private static async Task<Helpers.CLICommandResult> InvokeSpecCLICommand(Spec spec)
    {
        List<string> commandArgs = new List<string>{
                "--service-url", Helpers.GetBackendAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
        };

        commandArgs.AddRange(spec.CommandArguments);

        Helpers.CLICommandResult cliCommandResult = await Helpers.RunCLICommand(commandArgs.ToArray());

        return cliCommandResult;
    }

    private static string ReplaceRoot(string source)
    {
        const string relativeRootLocation = "../../../..";
        const string rootToken = "{CLI_ROOT}";
        string absoluteRootLocation = Path.GetFullPath(relativeRootLocation);

        return source.Replace(absoluteRootLocation, rootToken);
    }

    private static void ValidateResult(Spec spec, Helpers.CLICommandResult cliCommandResult, ITestOutputHelper output)
    {
        string expectedStdoutContent = spec.Stdout;
        string actualStdoutContent = ReplaceRoot(cliCommandResult.Stdout);

        string expectedStderrContent = spec.Stderr;
        string actualStderrContent = ReplaceRoot(cliCommandResult.Stderr);

        int expectedExitCode = spec.ExitCode;
        int actualExitCode = cliCommandResult.ExitCode;

        if (expectedStdoutContent != actualStdoutContent) {
            output.WriteLine("========== Expected stdout content ==============");
            output.WriteLine(expectedStdoutContent);
            output.WriteLine("========== Actual stdout content ================");
            output.WriteLine(actualStdoutContent);
            output.WriteLine("=================================================");
        }

        if (expectedStderrContent != actualStderrContent) {
            output.WriteLine("========== Expected stderr content ==============");
            output.WriteLine(expectedStderrContent);
            output.WriteLine("========== Actual stderr content ================");
            output.WriteLine(actualStderrContent);
            output.WriteLine("=================================================");
        }
        if (expectedExitCode != actualExitCode) {
            output.WriteLine("========== Expected exitcode ====================");
            output.WriteLine($"{expectedExitCode}");
            output.WriteLine("========== Actual exitcode ======================");
            output.WriteLine($"{actualExitCode}");
            output.WriteLine("=================================================");
        }

        Assert.Equal(expectedStdoutContent, actualStdoutContent);
        Assert.Equal(expectedStderrContent, actualStderrContent);
        Assert.Equal(expectedExitCode, actualExitCode);
    }

    public static async Task RunSpecCommand(string directory, ITestOutputHelper output)
    {
        Spec spec = ReadSpec(directory);
        Helpers.CLICommandResult cliCommandResult = await InvokeSpecCLICommand(spec);
        ValidateResult(spec, cliCommandResult, output);
    }
}