using Microsoft.Extensions.DependencyInjection;
using KicksharingApp.Models.Transports;
using KicksharingApp.Models.Accessories;
using KicksharingApp.Services.Interfaces;
using KicksharingApp.Services;
using KicksharingApp.Models.Base;
using KicksharingApp.Models.Interfaces;

var services = new ServiceCollection();
services.AddTransient<IInspectionService, InspectionService>();
services.AddSingleton<IFleetManager, FleetManager>();
var serviceProvider = services.BuildServiceProvider();
var fleet = serviceProvider.GetRequiredService<IFleetManager>();
Scooter scooter = new Scooter(101 , "Vishka 300", 8 , 1.2, true);
ElectricBicycle electricBicycle = new ElectricBicycle(102, "Vishka 400" , 7 , 2.0, false);
Bicycle bicycle = new Bicycle(103 , "Vishka 500" , 4 , true);

if (fleet.RegisterTransport(scooter))
{
    Console.WriteLine("Scooter успешно принят в парк!");
}
else
{
    Console.WriteLine("Scooter отклонен тех осмотром!");
}
if (fleet.RegisterTransport(electricBicycle))
{
    Console.WriteLine("ElectricBicycle успешно принят в парк!");
}
else
{
    Console.WriteLine("ElectricBicycle отклонен тех осмотром!");
}
if (fleet.RegisterTransport(bicycle))
{
    Console.WriteLine("Bicycle успешно принят в парк!");
}
else
{
    Console.WriteLine("Bicycle отклонен тех осмотром!");
}

Helmet helmet = new Helmet(201, "Helmet 201");
ChargingStation chargingStation = new ChargingStation(301 , "ChargingStation 201" , 20.0);
fleet.RegisterAccessory(helmet);
fleet.RegisterAccessory(chargingStation);

double energy = fleet.GetTotalDailyEnergyConsumption();
Console.WriteLine($"Суммарное суточное потребление парка: {energy} кВт*ч");

IEnumerable<Transport> transports =  fleet.GetVehiclesForBeginners();
foreach(Transport transport in transports)
{
    Console.WriteLine($"Номер: {transport.InventoryNumber} , Название:{transport.Title}, Уровень: {transport.SimplicityLevel}");
}

IEnumerable<IInventoryItem> items =  fleet.GetAllInventoryItems();
foreach (IInventoryItem item in items)
{
    Console.WriteLine($"Номер: {item.InventoryNumber} , Название:{item.Title}");
}
