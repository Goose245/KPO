namespace KicksharingApp.Models.Transports;
using KicksharingApp.Models.Base;
using KicksharingApp.Models.Interfaces;

public  class  ElectricBicycle : Transport , IEnergyConsumer
{
    public double DailyEnergyConsumption {get; }

    public ElectricBicycle(int number, string title, int simplicity, double dailyEnergyConsumption, bool sound = true) : base(number, title, simplicity, sound)
    {
        DailyEnergyConsumption = dailyEnergyConsumption;
    }
}