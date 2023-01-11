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
    /// GetStoreUploadResponse
    /// </summary>
    [DataContract(Name = "getStoreUploadResponse")]
    public partial class GetStoreUploadResponse : IEquatable<GetStoreUploadResponse>, IValidatableObject
    {

        /// <summary>
        /// Gets or Sets Status
        /// </summary>
        [DataMember(Name = "status", IsRequired = true, EmitDefaultValue = true)]
        public StoreUploadStatus Status { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreUploadResponse" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected GetStoreUploadResponse() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreUploadResponse" /> class.
        /// </summary>
        /// <param name="description">description (required).</param>
        /// <param name="buildId">buildId (required).</param>
        /// <param name="timestamp">timestamp (required).</param>
        /// <param name="files">files (required).</param>
        /// <param name="status">status (required).</param>
        public GetStoreUploadResponse(string description = default(string), string buildId = default(string), string timestamp = default(string), List<GetStoreUploadFileResponse> files = default(List<GetStoreUploadFileResponse>), StoreUploadStatus status = default(StoreUploadStatus))
        {
            // to ensure "description" is required (not null)
            if (description == null)
            {
                throw new ArgumentNullException("description is a required property for GetStoreUploadResponse and cannot be null");
            }
            this.Description = description;
            // to ensure "buildId" is required (not null)
            if (buildId == null)
            {
                throw new ArgumentNullException("buildId is a required property for GetStoreUploadResponse and cannot be null");
            }
            this.BuildId = buildId;
            // to ensure "timestamp" is required (not null)
            if (timestamp == null)
            {
                throw new ArgumentNullException("timestamp is a required property for GetStoreUploadResponse and cannot be null");
            }
            this.Timestamp = timestamp;
            // to ensure "files" is required (not null)
            if (files == null)
            {
                throw new ArgumentNullException("files is a required property for GetStoreUploadResponse and cannot be null");
            }
            this.Files = files;
            this.Status = status;
        }

        /// <summary>
        /// Gets or Sets Description
        /// </summary>
        [DataMember(Name = "description", IsRequired = true, EmitDefaultValue = true)]
        public string Description { get; set; }

        /// <summary>
        /// Gets or Sets BuildId
        /// </summary>
        [DataMember(Name = "buildId", IsRequired = true, EmitDefaultValue = true)]
        public string BuildId { get; set; }

        /// <summary>
        /// Gets or Sets Timestamp
        /// </summary>
        [DataMember(Name = "timestamp", IsRequired = true, EmitDefaultValue = true)]
        public string Timestamp { get; set; }

        /// <summary>
        /// Gets or Sets Files
        /// </summary>
        [DataMember(Name = "files", IsRequired = true, EmitDefaultValue = true)]
        public List<GetStoreUploadFileResponse> Files { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GetStoreUploadResponse {\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  BuildId: ").Append(BuildId).Append("\n");
            sb.Append("  Timestamp: ").Append(Timestamp).Append("\n");
            sb.Append("  Files: ").Append(Files).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
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
            return this.Equals(input as GetStoreUploadResponse);
        }

        /// <summary>
        /// Returns true if GetStoreUploadResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of GetStoreUploadResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(GetStoreUploadResponse input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && 
                (
                    this.BuildId == input.BuildId ||
                    (this.BuildId != null &&
                    this.BuildId.Equals(input.BuildId))
                ) && 
                (
                    this.Timestamp == input.Timestamp ||
                    (this.Timestamp != null &&
                    this.Timestamp.Equals(input.Timestamp))
                ) && 
                (
                    this.Files == input.Files ||
                    this.Files != null &&
                    input.Files != null &&
                    this.Files.SequenceEqual(input.Files)
                ) && 
                (
                    this.Status == input.Status ||
                    this.Status.Equals(input.Status)
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
                if (this.Description != null)
                {
                    hashCode = (hashCode * 59) + this.Description.GetHashCode();
                }
                if (this.BuildId != null)
                {
                    hashCode = (hashCode * 59) + this.BuildId.GetHashCode();
                }
                if (this.Timestamp != null)
                {
                    hashCode = (hashCode * 59) + this.Timestamp.GetHashCode();
                }
                if (this.Files != null)
                {
                    hashCode = (hashCode * 59) + this.Files.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Status.GetHashCode();
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
