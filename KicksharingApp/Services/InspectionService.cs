namespace KicksharingApp.Services;

using KicksharingApp.Models.Base;
using KicksharingApp.Services.Interfaces;

public class InspectionService : IInspectionService
{
    public bool Inspect(Transport transport)
    {
        if (transport.IsTechnicallySound)
        {
            return true;
        }
        return false;
    }
}