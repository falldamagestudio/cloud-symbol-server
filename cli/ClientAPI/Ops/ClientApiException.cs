using System;

namespace ClientAPI
{
    public class ClientAPIException : Exception
    {
            public ClientAPIException(string message) : base(message) { }
    }
}