using KicksharingApp.Models.Transports;
using KicksharingApp.Services;
using Xunit;

namespace KicksharingApp.Tests;

public class InspectionServiceTests
{
    [Fact]
    public void Inspect_WhenTransportIsSound_ShouldReturnTrue()
    {
        var service = new InspectionService();
        var scooter = new Scooter(1, "Vishka 42", 8, 1.0, true);
        var result = service.Inspect(scooter);
        Assert.True(result);
    }
    [Fact]
    public void Inspect_WhenTransportIsNotSound_ShouldReturnFalse()
    {
        var service = new InspectionService();
        var scooter = new Scooter(1, "Vishka 42", 8, 1.0, false);
        var result = service.Inspect(scooter);
        Assert.False(result);
    }
}