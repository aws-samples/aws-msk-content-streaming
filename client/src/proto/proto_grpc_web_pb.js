/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */

const grpc = {};
grpc.web = require("grpc-web");

var google_protobuf_timestamp_pb = require("google-protobuf/google/protobuf/timestamp_pb.js");
const proto = {};
proto.proto = require("./proto_pb.js");

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.MicroClient = function(hostname, credentials, options) {
  if (!options) options = {};
  options["format"] = "text";

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.MicroPromiseClient = function(hostname, credentials, options) {
  if (!options) options = {};
  options["format"] = "text";

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Insert.Request,
 *   !proto.proto.Insert.Response>}
 */
const methodDescriptor_Micro_Insert = new grpc.web.MethodDescriptor(
  "/proto.Micro/Insert",
  grpc.web.MethodType.UNARY,
  proto.proto.Insert.Request,
  proto.proto.Insert.Response,
  /**
   * @param {!proto.proto.Insert.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Insert.Response.deserializeBinary
);

/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Insert.Request,
 *   !proto.proto.Insert.Response>}
 */
const methodInfo_Micro_Insert = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Insert.Response,
  /**
   * @param {!proto.proto.Insert.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Insert.Response.deserializeBinary
);

/**
 * @param {!proto.proto.Insert.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Insert.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Insert.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.MicroClient.prototype.insert = function(
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + "/proto.Micro/Insert",
    request,
    metadata || {},
    methodDescriptor_Micro_Insert,
    callback
  );
};

/**
 * @param {!proto.proto.Insert.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Insert.Response>}
 *     A native promise that resolves to the response
 */
proto.proto.MicroPromiseClient.prototype.insert = function(request, metadata) {
  return this.client_.unaryCall(
    this.hostname_ + "/proto.Micro/Insert",
    request,
    metadata || {},
    methodDescriptor_Micro_Insert
  );
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Update.Request,
 *   !proto.proto.Update.Response>}
 */
const methodDescriptor_Micro_Update = new grpc.web.MethodDescriptor(
  "/proto.Micro/Update",
  grpc.web.MethodType.UNARY,
  proto.proto.Update.Request,
  proto.proto.Update.Response,
  /**
   * @param {!proto.proto.Update.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Update.Response.deserializeBinary
);

/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Update.Request,
 *   !proto.proto.Update.Response>}
 */
const methodInfo_Micro_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Update.Response,
  /**
   * @param {!proto.proto.Update.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Update.Response.deserializeBinary
);

/**
 * @param {!proto.proto.Update.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Update.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Update.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.MicroClient.prototype.update = function(
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + "/proto.Micro/Update",
    request,
    metadata || {},
    methodDescriptor_Micro_Update,
    callback
  );
};

/**
 * @param {!proto.proto.Update.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Update.Response>}
 *     A native promise that resolves to the response
 */
proto.proto.MicroPromiseClient.prototype.update = function(request, metadata) {
  return this.client_.unaryCall(
    this.hostname_ + "/proto.Micro/Update",
    request,
    metadata || {},
    methodDescriptor_Micro_Update
  );
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.ListArticles.Request,
 *   !proto.proto.ListArticles.Response>}
 */
const methodDescriptor_Micro_ListArticles = new grpc.web.MethodDescriptor(
  "/proto.Micro/ListArticles",
  grpc.web.MethodType.SERVER_STREAMING,
  proto.proto.ListArticles.Request,
  proto.proto.ListArticles.Response,
  /**
   * @param {!proto.proto.ListArticles.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ListArticles.Response.deserializeBinary
);

/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.ListArticles.Request,
 *   !proto.proto.ListArticles.Response>}
 */
const methodInfo_Micro_ListArticles = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.ListArticles.Response,
  /**
   * @param {!proto.proto.ListArticles.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ListArticles.Response.deserializeBinary
);

/**
 * @param {!proto.proto.ListArticles.Request} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.ListArticles.Response>}
 *     The XHR Node Readable Stream
 */
proto.proto.MicroClient.prototype.listArticles = function(request, metadata) {
  return this.client_.serverStreaming(
    this.hostname_ + "/proto.Micro/ListArticles",
    request,
    metadata || {},
    methodDescriptor_Micro_ListArticles
  );
};

/**
 * @param {!proto.proto.ListArticles.Request} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.ListArticles.Response>}
 *     The XHR Node Readable Stream
 */
proto.proto.MicroPromiseClient.prototype.listArticles = function(
  request,
  metadata
) {
  return this.client_.serverStreaming(
    this.hostname_ + "/proto.Micro/ListArticles",
    request,
    metadata || {},
    methodDescriptor_Micro_ListArticles
  );
};

module.exports = proto.proto;
