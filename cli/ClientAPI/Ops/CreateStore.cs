using System.Threading.Tasks;

namespace ClientAPI
{
    public class CreateStore
    {
        public static async Task<bool> DoCreateStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.CreateStoreAsync(StoreId);
        }
    }
}
