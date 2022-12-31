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
    /// GetTokenResponse
    /// </summary>
    [DataContract(Name = "getTokenResponse")]
    public partial class GetTokenResponse : IEquatable<GetTokenResponse>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GetTokenResponse" /> class.
        /// </summary>
        /// <param name="token">Personal Access Token This token can be used for authentication when accessing non-token related APIs.</param>
        /// <param name="description">Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used.</param>
        /// <param name="creationTimestamp">Creation timestamp, in RFC3339 format.</param>
        public GetTokenResponse(string token = default(string), string description = default(string), string creationTimestamp = default(string))
        {
            this.Token = token;
            this.Description = description;
            this.CreationTimestamp = creationTimestamp;
        }

        /// <summary>
        /// Personal Access Token This token can be used for authentication when accessing non-token related APIs
        /// </summary>
        /// <value>Personal Access Token This token can be used for authentication when accessing non-token related APIs</value>
        [DataMember(Name = "token", EmitDefaultValue = false)]
        public string Token { get; set; }

        /// <summary>
        /// Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used
        /// </summary>
        /// <value>Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used</value>
        [DataMember(Name = "description", EmitDefaultValue = false)]
        public string Description { get; set; }

        /// <summary>
        /// Creation timestamp, in RFC3339 format
        /// </summary>
        /// <value>Creation timestamp, in RFC3339 format</value>
        [DataMember(Name = "creationTimestamp", EmitDefaultValue = false)]
        public string CreationTimestamp { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GetTokenResponse {\n");
            sb.Append("  Token: ").Append(Token).Append("\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  CreationTimestamp: ").Append(CreationTimestamp).Append("\n");
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
            return this.Equals(input as GetTokenResponse);
        }

        /// <summary>
        /// Returns true if GetTokenResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of GetTokenResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(GetTokenResponse input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Token == input.Token ||
                    (this.Token != null &&
                    this.Token.Equals(input.Token))
                ) && 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && 
                (
                    this.CreationTimestamp == input.CreationTimestamp ||
                    (this.CreationTimestamp != null &&
                    this.CreationTimestamp.Equals(input.CreationTimestamp))
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
                if (this.Token != null)
                {
                    hashCode = (hashCode * 59) + this.Token.GetHashCode();
                }
                if (this.Description != null)
                {
                    hashCode = (hashCode * 59) + this.Description.GetHashCode();
                }
                if (this.CreationTimestamp != null)
                {
                    hashCode = (hashCode * 59) + this.CreationTimestamp.GetHashCode();
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
