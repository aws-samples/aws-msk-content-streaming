import React from "react";
import { Pane, Card, Paragraph, Heading } from "evergreen-ui";
import { Article } from "./proto/proto_pb";

export default function(props: any) {
  // deconstructing articles into a list and unwrap into aticle
  const articles = props.articles;
  const listArticles = articles.map((article: Article) => (
    <Card border="default" padding={16} margin={8} key={article.getUuid()}>
      <Heading>{article.getTitle()}</Heading>
      <Paragraph>{article.getBody()}</Paragraph>
    </Card>
  ));

  // render a pane with the articles
  return <Pane>{articles.length > 0 ? listArticles : "No items"}</Pane>;
}
