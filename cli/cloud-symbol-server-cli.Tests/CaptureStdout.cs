using System;
using System.IO;

namespace cloud_symbol_server_cli.Tests;

public class CaptureStdout : IDisposable
{
    private StringWriter stringWriter;
    private TextWriter originalOutput;

    public CaptureStdout()
    {
        stringWriter = new StringWriter();
        originalOutput = Console.Out;
        Console.SetOut(stringWriter);
    }

    public string GetOutput()
    {
        return stringWriter.ToString();
    }

    public void Dispose()
    {
        Console.SetOut(originalOutput);
        stringWriter.Dispose();
    }
}