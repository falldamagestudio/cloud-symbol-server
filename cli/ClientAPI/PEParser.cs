using System;
using System.IO;
using System.Reflection.PortableExecutable;

namespace ClientAPI
{
    public class PEParser
    {

        public static string GetHash(string pePath)
        {
            using (Stream stream = new FileStream(pePath, FileMode.Open))
            {
                try {
                    PEHeaders peHeaders = new PEHeaders(stream);
                    if (peHeaders.CoffHeader == null || peHeaders.PEHeader == null) {
                        return null;
                    }
                    string hash = String.Format("{0:X}{1:X}", peHeaders.CoffHeader.TimeDateStamp, peHeaders.PEHeader.SizeOfImage);
                    return hash;
                } catch (BadImageFormatException) {
                    return null;
                }
            }
        }
    }
}