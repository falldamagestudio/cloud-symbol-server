
namespace ClientAPI
{
    public static class Helpers
    {
        public static BackendAPI.Api.DefaultApi CreateApi(string ServiceURL, string Email, string PAT) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);
            return api;
        }
    }
}