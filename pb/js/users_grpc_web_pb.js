/**
 * @fileoverview gRPC-Web generated client stub for users
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var types_pb = require('./types_pb.js')
const proto = {};
proto.users = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.users.V1UsersClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
proto.users.V1UsersPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.users.CreateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_V1Users_Create = new grpc.web.MethodDescriptor(
  '/users.V1Users/Create',
  grpc.web.MethodType.UNARY,
  proto.users.CreateUserRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.CreateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_V1Users_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/Create',
      request,
      metadata || {},
      methodDescriptor_V1Users_Create,
      callback);
};


/**
 * @param {!proto.users.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     A native promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/Create',
      request,
      metadata || {},
      methodDescriptor_V1Users_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.FindByIdRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_V1Users_FindById = new grpc.web.MethodDescriptor(
  '/users.V1Users/FindById',
  grpc.web.MethodType.UNARY,
  proto.users.FindByIdRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByIdRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.FindByIdRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_V1Users_FindById = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByIdRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.FindByIdRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.findById =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/FindById',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindById,
      callback);
};


/**
 * @param {!proto.users.FindByIdRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     A native promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.findById =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/FindById',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindById);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.FindByEmailRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_V1Users_FindByEmail = new grpc.web.MethodDescriptor(
  '/users.V1Users/FindByEmail',
  grpc.web.MethodType.UNARY,
  proto.users.FindByEmailRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByEmailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.FindByEmailRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_V1Users_FindByEmail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByEmailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.FindByEmailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.findByEmail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/FindByEmail',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindByEmail,
      callback);
};


/**
 * @param {!proto.users.FindByEmailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     A native promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.findByEmail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/FindByEmail',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindByEmail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.UpdateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_V1Users_Update = new grpc.web.MethodDescriptor(
  '/users.V1Users/Update',
  grpc.web.MethodType.UNARY,
  proto.users.UpdateUserRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.UpdateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_V1Users_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/Update',
      request,
      metadata || {},
      methodDescriptor_V1Users_Update,
      callback);
};


/**
 * @param {!proto.users.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     A native promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/Update',
      request,
      metadata || {},
      methodDescriptor_V1Users_Update);
};


module.exports = proto.users;

