using System;
using System.IO;
using System.Reflection.PortableExecutable;

namespace ClientAPI
{
    public class PEParser
    {

        // Compute hash for Windows Portable Executable files
        // Specification here: https://github.com/dotnet/symstore/blob/main/docs/specs/SSQP_Key_Conventions.md#pe-timestamp-filesize

        public static string GetCodeIdentifier(string pePath)
        {
            using (Stream stream = new FileStream(pePath, FileMode.Open))
            {
                try {
                    PEHeaders peHeaders = new PEHeaders(stream);
                    if (peHeaders.CoffHeader == null || peHeaders.PEHeader == null) {
                        return null;
                    }
                    // SSQP has a particular casing standard for PE file hashes; the first portion is uppercase, the second portion is lowercase
                    string hash = String.Format("{0:X8}{1:x}", peHeaders.CoffHeader.TimeDateStamp, peHeaders.PEHeader.SizeOfImage);
                    return hash;
                } catch (BadImageFormatException) {
                    return null;
                }
            }
        }
    }
}