import React from "react";
import { Pane, Card } from "evergreen-ui";
import { Article } from "./proto/proto_pb";

export default function(props: any) {
  // deconstructing articles into a list and unwrap into aticle
  const articles = props.articles;
  const listArticles = articles.map((article: Article) => (
    <Card border="default" padding={16} margin={8} key={article.getUuid()}>
      {article.getBody()}
    </Card>
  ));

  // render a pane with the articles
  return <Pane>{articles.length > 0 ? listArticles : "No items"}</Pane>;
}
