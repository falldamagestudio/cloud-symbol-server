using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Security.Cryptography;

namespace ClientAPI
{
    using FileType = BackendAPI.Model.StoreFileBlobType;

    public class ComputeFileMetadata
    {
        public struct FileWithMetadata
        {
        	public string FileWithPath;
        	public string FileWithoutPath;
        	public string BlobIdentifier;
            public FileType Type;
            public Int64 Size;
            public string ContentHash;
        }

        private static string ComputeSHA256Hash(string filename)
        {
            using (FileStream file = File.OpenRead(filename))
            using (SHA256 sha256Hash = SHA256.Create())
            {
                byte[] hashBytes = sha256Hash.ComputeHash(file);
                StringBuilder builder = new StringBuilder();
                for (int i = 0; i < hashBytes.Length; i++)
                {
                    builder.Append(hashBytes[i].ToString("x2"));
                }
                return builder.ToString();
            }
        }

        private static FileWithMetadata GetFileWithMetadata(string fileWithPath, string fileWithoutPath)
        {
            FileInfo fileInfo = new FileInfo(fileWithPath);
            Int64 size = fileInfo.Length;
            string sha256Hash = ComputeSHA256Hash(fileWithPath);

            string pdbDebugIdentifier = PDBParser.GetDebugIdentifier(fileWithPath);
            if (pdbDebugIdentifier != null)
                return new FileWithMetadata {
                    FileWithPath = fileWithPath,
                    FileWithoutPath = fileWithoutPath,
                    BlobIdentifier = pdbDebugIdentifier,
                    Type = FileType.Pdb,
                    Size = size,
                    ContentHash = sha256Hash,
                };

            string peCodeIdentifier = PEParser.GetCodeIdentifier(fileWithPath);
            if (peCodeIdentifier != null)
                return new FileWithMetadata {
                    FileWithPath = fileWithPath,
                    FileWithoutPath = fileWithoutPath,
                    BlobIdentifier = peCodeIdentifier,
                    Type = FileType.Pe,
                    Size = size,
                    ContentHash = sha256Hash,
                };

            throw new ApplicationException($"File {fileWithPath} is not of a recognized format");
        }

        public static IEnumerable<FileWithMetadata> DoComputeFileMedatadata(IReadOnlyCollection<string> Files) {

            if (!Files.Any()) {
                throw new ArgumentException($"DoComputeFileMetadata requires at least one filename", nameof(Files));
            }

            IEnumerable<FileWithMetadata> filesWithMetadata = Files.Select(fileName =>
                GetFileWithMetadata(fileName, Path.GetFileName(fileName))
            ).Where(fileWithMetadata => fileWithMetadata.BlobIdentifier != null);

            return filesWithMetadata;
        }
    }
}
