namespace KicksharingApp.Models.Accessories;
using KicksharingApp.Models.Base;
using KicksharingApp.Models.Interfaces;

public class ChargingStation : Accessory, IEnergyConsumer
{
    public double DailyEnergyConsumption {get; }
    public ChargingStation(int number, string title, double dailyEnergyConsumption) : base(number, title)
    {
        DailyEnergyConsumption = dailyEnergyConsumption;
    }
}