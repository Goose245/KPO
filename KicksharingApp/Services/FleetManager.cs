namespace KicksharingApp.Services;

using KicksharingApp.Models.Base;
using KicksharingApp.Models.Interfaces;
using KicksharingApp.Services.Interfaces;

public class FleetManager : IFleetManager
{
    private readonly IInspectionService _inspectionService;
    public FleetManager(IInspectionService inspectionService)
    {
        _inspectionService = inspectionService;
    }

    private readonly List<Transport> transports = new();
    private readonly List<Accessory> accessories = new();

    public bool RegisterTransport(Transport transport)
    {
        if (_inspectionService.Inspect(transport) == true)
        {
            transports.Add(transport);
            return true;
        }
        return false;
    }
    public void RegisterAccessory(Accessory accessory)
    {
        accessories.Add(accessory);
    }
    public double GetTotalDailyEnergyConsumption()
    {
        double DailyEnergyConsumption = 0.0;
        foreach (Transport transport in transports)
        {
            if (transport is IEnergyConsumer consumer)
            {
                DailyEnergyConsumption += consumer.DailyEnergyConsumption;
            }
        }
        foreach (Accessory accessory in accessories)
        {
            if (accessory is IEnergyConsumer consumer1)
            {
                DailyEnergyConsumption += consumer1.DailyEnergyConsumption;
            }
        }
        return DailyEnergyConsumption;
    }
    public IEnumerable<Transport> GetVehiclesForBeginners()
    {
        List<Transport> result = new List<Transport>();
        foreach(Transport transport in transports)
        {
            if (transport.IsSafeForBeginners)
            {
                result.Add(transport);
            }
        }
        return result;
    }
    public IEnumerable<IInventoryItem> GetAllInventoryItems()
    {
        List<IInventoryItem> result = new List<IInventoryItem>();
        foreach (Transport transport in transports)
        {
            result.Add(transport);
        }
        foreach (Accessory accessory in accessories)
        {

            result.Add(accessory);
        }
        return result;
    }
}