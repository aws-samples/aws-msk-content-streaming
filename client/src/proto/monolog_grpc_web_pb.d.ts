import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {
  Request,
  Response,
  Request,
  Response,
  Request,
  Response} from './monolog_pb';

export class MonologClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  insert(
    request: Insert.Request,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Insert.Response) => void
  ): grpcWeb.ClientReadableStream<Insert.Response>;

  update(
    request: Update.Request,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Update.Response) => void
  ): grpcWeb.ClientReadableStream<Update.Response>;

  listArticles(
    request: ListArticles.Request,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<ListArticles.Response>;

}

export class MonologPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  insert(
    request: Insert.Request,
    metadata?: grpcWeb.Metadata
  ): Promise<Insert.Response>;

  update(
    request: Update.Request,
    metadata?: grpcWeb.Metadata
  ): Promise<Update.Response>;

  listArticles(
    request: ListArticles.Request,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<ListArticles.Response>;

}

