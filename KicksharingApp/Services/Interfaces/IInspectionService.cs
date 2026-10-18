namespace KicksharingApp.Services.Interfaces;
using KicksharingApp.Models.Base;

public interface   IInspectionService 
{
    bool Inspect(Transport transport);
}