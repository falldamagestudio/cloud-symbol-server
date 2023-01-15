using Xunit;
using Xunit.Abstractions;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    private readonly ITestOutputHelper output;

    public TestCommands(ITestOutputHelper output)
    {
        this.output = output;
    }
}
