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
        originalError = Console.Error;
        Console.SetError(stringWriter);
    }

    public string GetError()
    {
        return stringWriter.ToString();
    }

    public void Dispose()
    {
        Console.SetError(originalError);
        stringWriter.Dispose();
    }
}