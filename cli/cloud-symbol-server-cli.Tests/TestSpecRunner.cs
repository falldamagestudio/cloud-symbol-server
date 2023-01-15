using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Xunit;
using Xunit.Abstractions;

namespace cloud_symbol_server_cli.Tests;

public static class TestSpecRunner
{
    private class Spec
    {
        [JsonProperty(Required = Required.Always)]
        public List<string> CommandArguments;
        
        [JsonProperty(Required = Required.Always)]
        public int ExitCode;
        
        public int SkipStdoutLines;
        public int SkipStderrLines;
        
        public string Stdout;
        public string Stderr;
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

    private static string SkipLines(string source, int linesToSkip)
    {
        int position = 0;

        for (int lineBreaksEncountered = 0; lineBreaksEncountered < linesToSkip; lineBreaksEncountered++)
        {
            int nextLineBreakPosition = source.IndexOf('\n', position);
            if (nextLineBreakPosition == -1)
                break;
            position = nextLineBreakPosition;
        }

        return source.Substring(position);
    }

    private static void ValidateResult(Spec spec, Helpers.CLICommandResult cliCommandResult, ITestOutputHelper output)
    {
        string expectedStdoutContent = SkipLines(spec.Stdout, spec.SkipStdoutLines);
        string actualStdoutContent = SkipLines(cliCommandResult.Stdout, spec.SkipStdoutLines);

        string expectedStderrContent = SkipLines(spec.Stderr, spec.SkipStderrLines);
        string actualStderrContent = SkipLines(cliCommandResult.Stderr, spec.SkipStderrLines);

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