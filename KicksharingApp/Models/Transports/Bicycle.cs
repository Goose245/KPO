namespace KicksharingApp.Models.Transports;
using KicksharingApp.Models.Base;

public class Bicycle : Transport
{
    public Bicycle (int number, string title, int simplicity, bool sound = true) : base(number, title, simplicity, sound)
    {
        
    }
}