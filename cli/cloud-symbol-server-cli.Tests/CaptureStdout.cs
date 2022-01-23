using System;
using System.IO;

namespace cloud_symbol_server_cli.Tests;

public class CaptureStderr : IDisposable
{
    private StringWriter stringWriter;
    private TextWriter originalError;

    public CaptureStderr()
    {
        stringWriter = new StringWriter();
        originalError = Console.Out;
        Console.SetOut(stringWriter);
    }

    public string GetOutput()
    {
        return stringWriter.ToString();
    }

    public void Dispose()
    {
        Console.SetOut(originalError);
        stringWriter.Dispose();
    }
}