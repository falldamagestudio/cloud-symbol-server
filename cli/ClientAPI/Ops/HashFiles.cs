using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace ClientAPI
{
    public class HashFiles
    {
        public struct FileWithHash
        {
        	public string FileWithPath;
        	public string FileWithoutPath;
        	public string Hash;
        }

        private static string GetHash(string fileName)
        {
            string pdbHash = PDBParser.GetHash(fileName);
            if (pdbHash != null)
                return pdbHash;

            string peHash = PEParser.GetHash(fileName);
            if (peHash != null)
                return peHash;

            throw new ApplicationException($"File {fileName} is not of a recognized format");
        }

        public static IEnumerable<FileWithHash> GetFilesWithHashes(IEnumerable<string> fileNames)
        {
            IEnumerable<FileWithHash> filesWithHashes = fileNames.Select(fileName => new FileWithHash {
                FileWithPath = fileName,
                FileWithoutPath = Path.GetFileName(fileName),
                Hash = GetHash(fileName)
            }).Where(fileWithHash => fileWithHash.Hash != null);

            return filesWithHashes;
        }

        public static IEnumerable<FileWithHash> DoHashFiles(IReadOnlyCollection<string> Files) {

            if (!Files.Any()) {
                throw new ArgumentException($"HashFiles requires at least one filename", nameof(Files));
            }

            IEnumerable<FileWithHash> filesWithHashes = GetFilesWithHashes(Files);

            return filesWithHashes;
        }
    }
}
