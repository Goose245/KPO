namespace KicksharingApp.Models.Base;

using KicksharingApp.Models.Interfaces;

public abstract class Accessory : IInventoryItem
{
    public int InventoryNumber {get; }
    public string Title {get; }

    protected Accessory(int number , string title)
    {
        InventoryNumber = number;
        Title = title;
    }
}