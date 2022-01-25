namespace CLI
{
    public class GlobalOptions {
        public string? ServiceUrl { get; set; }
        public string? Email { get; set; }
        public string? Pat { get; set; }

        public bool Validate() {
            return !string.IsNullOrEmpty(ServiceUrl)
            && !string.IsNullOrEmpty(Email)
            && !string.IsNullOrEmpty(Pat);
        }
    }
}