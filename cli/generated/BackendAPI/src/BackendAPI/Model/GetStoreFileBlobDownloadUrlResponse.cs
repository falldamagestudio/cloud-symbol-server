/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = BackendAPI.Client.OpenAPIDateConverter;

namespace BackendAPI.Model
{
    /// <summary>
    /// GetStoreFileBlobDownloadUrlResponse
    /// </summary>
    [DataContract(Name = "getStoreFileBlobDownloadUrlResponse")]
    public partial class GetStoreFileBlobDownloadUrlResponse : IEquatable<GetStoreFileBlobDownloadUrlResponse>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreFileBlobDownloadUrlResponse" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected GetStoreFileBlobDownloadUrlResponse() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreFileBlobDownloadUrlResponse" /> class.
        /// </summary>
        /// <param name="method">method (required).</param>
        /// <param name="url">url (required).</param>
        public GetStoreFileBlobDownloadUrlResponse(string method = default(string), string url = default(string))
        {
            // to ensure "method" is required (not null)
            if (method == null)
            {
                throw new ArgumentNullException("method is a required property for GetStoreFileBlobDownloadUrlResponse and cannot be null");
            }
            this.Method = method;
            // to ensure "url" is required (not null)
            if (url == null)
            {
                throw new ArgumentNullException("url is a required property for GetStoreFileBlobDownloadUrlResponse and cannot be null");
            }
            this.Url = url;
        }

        /// <summary>
        /// Gets or Sets Method
        /// </summary>
        [DataMember(Name = "method", IsRequired = true, EmitDefaultValue = true)]
        public string Method { get; set; }

        /// <summary>
        /// Gets or Sets Url
        /// </summary>
        [DataMember(Name = "url", IsRequired = true, EmitDefaultValue = true)]
        public string Url { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GetStoreFileBlobDownloadUrlResponse {\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
            sb.Append("  Url: ").Append(Url).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as GetStoreFileBlobDownloadUrlResponse);
        }

        /// <summary>
        /// Returns true if GetStoreFileBlobDownloadUrlResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of GetStoreFileBlobDownloadUrlResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(GetStoreFileBlobDownloadUrlResponse input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Method == input.Method ||
                    (this.Method != null &&
                    this.Method.Equals(input.Method))
                ) && 
                (
                    this.Url == input.Url ||
                    (this.Url != null &&
                    this.Url.Equals(input.Url))
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.Method != null)
                {
                    hashCode = (hashCode * 59) + this.Method.GetHashCode();
                }
                if (this.Url != null)
                {
                    hashCode = (hashCode * 59) + this.Url.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}