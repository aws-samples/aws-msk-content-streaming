import * as jspb from "google-protobuf";

import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Item extends jspb.Message {
  getArticle(): Article | undefined;
  setArticle(value?: Article): void;
  hasArticle(): boolean;
  clearArticle(): void;

  getImage(): Image | undefined;
  setImage(value?: Image): void;
  hasImage(): boolean;
  clearImage(): void;

  getItemCase(): Item.ItemCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Item.AsObject;
  static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
  static serializeBinaryToWriter(
    message: Item,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Item;
  static deserializeBinaryFromReader(
    message: Item,
    reader: jspb.BinaryReader
  ): Item;
}

export namespace Item {
  export type AsObject = {
    article?: Article.AsObject;
    image?: Image.AsObject;
  };

  export enum ItemCase {
    ITEM_NOT_SET = 0,
    ARTICLE = 10,
    IMAGE = 11
  }
}

export class Article extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Article.AsObject;
  static toObject(includeInstance: boolean, msg: Article): Article.AsObject;
  static serializeBinaryToWriter(
    message: Article,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Article;
  static deserializeBinaryFromReader(
    message: Article,
    reader: jspb.BinaryReader
  ): Article;
}

export namespace Article {
  export type AsObject = {
    uuid: string;
    body: string;
  };
}

export class Image extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Image.AsObject;
  static toObject(includeInstance: boolean, msg: Image): Image.AsObject;
  static serializeBinaryToWriter(
    message: Image,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Image;
  static deserializeBinaryFromReader(
    message: Image,
    reader: jspb.BinaryReader
  ): Image;
}

export namespace Image {
  export type AsObject = {
    uuid: string;
    url: string;
  };
}

export class Insert extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Insert.AsObject;
  static toObject(includeInstance: boolean, msg: Insert): Insert.AsObject;
  static serializeBinaryToWriter(
    message: Insert,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Insert;
  static deserializeBinaryFromReader(
    message: Insert,
    reader: jspb.BinaryReader
  ): Insert;
}

export namespace Insert {
  export type AsObject = {};

  export class Request extends jspb.Message {
    getItem(): Item | undefined;
    setItem(value?: Item): void;
    hasItem(): boolean;
    clearItem(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(
      message: Request,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(
      message: Request,
      reader: jspb.BinaryReader
    ): Request;
  }

  export namespace Request {
    export type AsObject = {
      item?: Item.AsObject;
    };
  }

  export class Response extends jspb.Message {
    getUuid(): string;
    setUuid(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(
      message: Response,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(
      message: Response,
      reader: jspb.BinaryReader
    ): Response;
  }

  export namespace Response {
    export type AsObject = {
      uuid: string;
    };
  }
}

export class Update extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Update.AsObject;
  static toObject(includeInstance: boolean, msg: Update): Update.AsObject;
  static serializeBinaryToWriter(
    message: Update,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Update;
  static deserializeBinaryFromReader(
    message: Update,
    reader: jspb.BinaryReader
  ): Update;
}

export namespace Update {
  export type AsObject = {};

  export class Request extends jspb.Message {
    getItem(): Item | undefined;
    setItem(value?: Item): void;
    hasItem(): boolean;
    clearItem(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(
      message: Request,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(
      message: Request,
      reader: jspb.BinaryReader
    ): Request;
  }

  export namespace Request {
    export type AsObject = {
      item?: Item.AsObject;
    };
  }

  export class Response extends jspb.Message {
    getUuid(): string;
    setUuid(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(
      message: Response,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(
      message: Response,
      reader: jspb.BinaryReader
    ): Response;
  }

  export namespace Response {
    export type AsObject = {
      uuid: string;
    };
  }
}

export class ListArticles extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListArticles.AsObject;
  static toObject(
    includeInstance: boolean,
    msg: ListArticles
  ): ListArticles.AsObject;
  static serializeBinaryToWriter(
    message: ListArticles,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): ListArticles;
  static deserializeBinaryFromReader(
    message: ListArticles,
    reader: jspb.BinaryReader
  ): ListArticles;
}

export namespace ListArticles {
  export type AsObject = {};

  export class Request extends jspb.Message {
    getCount(): number;
    setCount(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(
      message: Request,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(
      message: Request,
      reader: jspb.BinaryReader
    ): Request;
  }

  export namespace Request {
    export type AsObject = {
      count: number;
    };
  }

  export class Response extends jspb.Message {
    getArticlesList(): Array<Article>;
    setArticlesList(value: Array<Article>): void;
    clearArticlesList(): void;
    addArticles(value?: Article, index?: number): Article;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(
      message: Response,
      writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(
      message: Response,
      reader: jspb.BinaryReader
    ): Response;
  }

  export namespace Response {
    export type AsObject = {
      articlesList: Array<Article.AsObject>;
    };
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(
    message: Empty,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(
    message: Empty,
    reader: jspb.BinaryReader
  ): Empty;
}

export namespace Empty {
  export type AsObject = {};
}
