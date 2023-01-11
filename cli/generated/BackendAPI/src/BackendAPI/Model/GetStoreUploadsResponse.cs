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
    /// GetStoreUploadsResponse
    /// </summary>
    [DataContract(Name = "getStoreUploadsResponse")]
    public partial class GetStoreUploadsResponse : IEquatable<GetStoreUploadsResponse>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreUploadsResponse" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected GetStoreUploadsResponse() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="GetStoreUploadsResponse" /> class.
        /// </summary>
        /// <param name="uploads">uploads (required).</param>
        /// <param name="pagination">pagination (required).</param>
        public GetStoreUploadsResponse(List<GetStoreUploadResponse> uploads = default(List<GetStoreUploadResponse>), PaginationResponse pagination = default(PaginationResponse))
        {
            // to ensure "uploads" is required (not null)
            if (uploads == null)
            {
                throw new ArgumentNullException("uploads is a required property for GetStoreUploadsResponse and cannot be null");
            }
            this.Uploads = uploads;
            // to ensure "pagination" is required (not null)
            if (pagination == null)
            {
                throw new ArgumentNullException("pagination is a required property for GetStoreUploadsResponse and cannot be null");
            }
            this.Pagination = pagination;
        }

        /// <summary>
        /// Gets or Sets Uploads
        /// </summary>
        [DataMember(Name = "uploads", IsRequired = true, EmitDefaultValue = true)]
        public List<GetStoreUploadResponse> Uploads { get; set; }

        /// <summary>
        /// Gets or Sets Pagination
        /// </summary>
        [DataMember(Name = "pagination", IsRequired = true, EmitDefaultValue = true)]
        public PaginationResponse Pagination { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GetStoreUploadsResponse {\n");
            sb.Append("  Uploads: ").Append(Uploads).Append("\n");
            sb.Append("  Pagination: ").Append(Pagination).Append("\n");
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
            return this.Equals(input as GetStoreUploadsResponse);
        }

        /// <summary>
        /// Returns true if GetStoreUploadsResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of GetStoreUploadsResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(GetStoreUploadsResponse input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Uploads == input.Uploads ||
                    this.Uploads != null &&
                    input.Uploads != null &&
                    this.Uploads.SequenceEqual(input.Uploads)
                ) && 
                (
                    this.Pagination == input.Pagination ||
                    (this.Pagination != null &&
                    this.Pagination.Equals(input.Pagination))
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
                if (this.Uploads != null)
                {
                    hashCode = (hashCode * 59) + this.Uploads.GetHashCode();
                }
                if (this.Pagination != null)
                {
                    hashCode = (hashCode * 59) + this.Pagination.GetHashCode();
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
