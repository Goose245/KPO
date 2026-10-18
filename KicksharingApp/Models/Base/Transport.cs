namespace KicksharingApp.Models.Base;

using KicksharingApp.Models.Interfaces;

public abstract class Transport : IInventoryItem
{
    public int InventoryNumber {get; }
    public string Title {get; }

    public int SimplicityLevel {get; }
    public bool IsSafeForBeginners => SimplicityLevel >= 6;

    public bool IsTechnicallySound{get; set;}

    protected Transport(int number , string title , int simplicity , bool sound = true)
    {
        InventoryNumber = number;
        Title = title;
        SimplicityLevel = simplicity;
        IsTechnicallySound = sound;
    }
}