namespace KicksharingApp.Services.Interfaces;
using KicksharingApp.Models.Base;
using KicksharingApp.Models.Interfaces;

public interface IFleetManager
{
    bool RegisterTransport(Transport transport);
    void RegisterAccessory(Accessory accessory);
    double GetTotalDailyEnergyConsumption();
    IEnumerable<Transport> GetVehiclesForBeginners();
    IEnumerable<IInventoryItem> GetAllInventoryItems();
} 