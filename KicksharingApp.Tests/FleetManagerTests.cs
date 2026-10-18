namespace KicksharingApp.Tests;

using System.Linq;
using Xunit;
using KicksharingApp.Services;
using KicksharingApp.Tests.Fakes;
using KicksharingApp.Models.Transports;
using KicksharingApp.Models.Accessories;

public class FleetManagerTests
{
    [Fact]
    public void RegisterTransport_WhenInspectionFails_ShouldNotAddTransport()
    {
        var fakeService = new FakeInspectionService(resultToReturn: false);
        FleetManager fleet = new FleetManager(fakeService);
        Scooter scooter = new Scooter(342 , "Vishka 67", 9 , 1.6, true);
        var result = fleet.RegisterTransport(scooter);
        Assert.False(result);
        Assert.Empty(fleet.GetAllInventoryItems());
    }
    [Fact]
    public void RegisterTransport_WhenInspectionPasses_ShouldAddTransport()
    {
        var fakeService = new FakeInspectionService(resultToReturn: true);
        FleetManager fleet = new FleetManager(fakeService);
        Scooter scooter = new Scooter(342 , "Vishka 67", 9 , 1.6, true);
        var result = fleet.RegisterTransport(scooter);
        Assert.True(result);
        Assert.Single(fleet.GetAllInventoryItems());
    }
    [Fact]
    public void GetVehiclesForBeginners_ShouldFilterBySimplicityLevel()
    {
        var fakeService = new FakeInspectionService(resultToReturn: true);
        FleetManager fleet = new FleetManager(fakeService);
        Scooter scooter = new Scooter(676 , "Vishka 123", 4 , 1.6, true);
        Bicycle bicycle = new Bicycle(232 , "Vishka 777",  6, true);
        fleet.RegisterTransport(scooter);
        fleet.RegisterTransport(bicycle);
        var beginnersVehicles = fleet.GetVehiclesForBeginners().ToList();
        Assert.Single(beginnersVehicles);
        Assert.Equal(6, beginnersVehicles[0].SimplicityLevel);
    }
    [Fact]
    public void GetTotalDailyEnergyConsumption_ShouldSumOnlyEnergyConsumers()
    {
        var fakeService = new FakeInspectionService(resultToReturn: true);
        FleetManager fleet = new FleetManager(fakeService);
        Scooter scooter = new Scooter(66 , "Vishka 987", 4 , 1.5, true);
        Bicycle bicycle = new Bicycle(232 , "Vishka 777",  6, true);
        Helmet helmet = new Helmet(45, "Helmet 32");
        ChargingStation chargingStation = new ChargingStation(90 , "Chargingstation 78", 10.0);
        fleet.RegisterTransport(scooter);
        fleet.RegisterTransport(bicycle);
        fleet.RegisterAccessory(helmet);
        fleet.RegisterAccessory(chargingStation);
        var totalEnergy = fleet.GetTotalDailyEnergyConsumption();
        Assert.Equal(11.5, totalEnergy, precision: 2);
    }
    [Fact]

    public void RegisterAccessory_ShouldAddAccessoryToInventoryWithoutInspection()
    {
        var fakeService = new FakeInspectionService(resultToReturn: false);
        FleetManager fleet = new FleetManager(fakeService);
        Helmet helmet = new Helmet(45, "Helmet 32");
        fleet.RegisterAccessory(helmet);
        var item = fleet.GetAllInventoryItems().ToList();
        Assert.Single(item);
        Assert.Equal(45, item[0].InventoryNumber);
        Assert.Equal("Helmet 32", item[0].Title);
    }

}
