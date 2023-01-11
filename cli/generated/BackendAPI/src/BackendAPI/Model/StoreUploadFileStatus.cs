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
    /// Defines storeUploadFileStatus
    /// </summary>
    [JsonConverter(typeof(StringEnumConverter))]
    public enum StoreUploadFileStatus
    {
        /// <summary>
        /// Enum Unknown for value: unknown
        /// </summary>
        [EnumMember(Value = "unknown")]
        Unknown = 1,

        /// <summary>
        /// Enum AlreadyPresent for value: already_present
        /// </summary>
        [EnumMember(Value = "already_present")]
        AlreadyPresent = 2,

        /// <summary>
        /// Enum Pending for value: pending
        /// </summary>
        [EnumMember(Value = "pending")]
        Pending = 3,

        /// <summary>
        /// Enum Completed for value: completed
        /// </summary>
        [EnumMember(Value = "completed")]
        Completed = 4,

        /// <summary>
        /// Enum Aborted for value: aborted
        /// </summary>
        [EnumMember(Value = "aborted")]
        Aborted = 5,

        /// <summary>
        /// Enum Expired for value: expired
        /// </summary>
        [EnumMember(Value = "expired")]
        Expired = 6

    }

}
