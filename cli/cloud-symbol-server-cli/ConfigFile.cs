using Newtonsoft.Json;

namespace CLI
{
    public static class ConfigFile
    {
        private static InternalConfigFile? internalConfigFile;

        public static void Init()
        {
            internalConfigFile = new InternalConfigFile();
        }

        public static void Init(string path)
        {
            internalConfigFile = new InternalConfigFile(path);
        }

        public static string GetOrDefault(string key, string defaultValue)
        {
            if (internalConfigFile != null) {
                return internalConfigFile.GetOrDefault(key, defaultValue);
            } else {
                throw new ApplicationException("ConfigFile has not yet been initialized");
            }
        }

        public class InternalConfigFile {

            private Dictionary<string, string> config;

            public InternalConfigFile()
            {
                config = new Dictionary<string, string>();
            }

            public InternalConfigFile(string path)
            {
                try {
                    string json = File.ReadAllText(path);
                    config = JsonConvert.DeserializeObject<Dictionary<string, string>>(json);
                } catch {
                    config = new Dictionary<string, string>();
                    throw;
                }
            }

            public string GetOrDefault(string key, string defaultValue)
            {
                if (config.ContainsKey(key)) {
                    return config[key];
                } else {
                    return defaultValue;
                }
            }
        }
    }
}
