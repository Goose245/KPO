using KicksharingApp.Models.Base;

namespace KicksharingApp.Tests.Fakes;
using KicksharingApp.Services.Interfaces;
public class FakeInspectionService : IInspectionService
{
    public bool ResultToReturn { get; set; }

    public bool Inspect(Transport transport)
    {
        return ResultToReturn;
    }
    public FakeInspectionService(bool resultToReturn)
    {
        ResultToReturn = resultToReturn;
    }
}