import React, { useState, useEffect } from "react";
import logo from "./logo.svg";
import "./App.css";
import { grpc } from "@improbable-eng/grpc-web";
import { MicroClient } from "./proto/proto_grpc_web_pb";
import { ListArticles, Article, Insert, Item } from "./proto/proto_pb";
import {
  Button,
  Pane,
  Text,
  TextInputField,
  FormField,
  Heading
} from "evergreen-ui";
import ArticlesList from "./Articles";

function App() {
  // saving input state
  const [inputs, setInputs] = useState({});
  // save articles state, this is the list of provided articles
  const [articles, setArticles] = useState<Array<Article>>([]);
  const mono = new MicroClient("http://localhost:8080");

  // handling the have to the form fields,
  // which is the title and the description of the article
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.persist();
    setInputs(inputs => ({ ...inputs, [e.target.name]: e.target.value }));
  };

  // handling the submit event to the form
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    if (e) {
      e.preventDefault();
    }

    const req = new Insert.Request();

    const article = new Article();
    article.setBody((inputs as any).body);

    const item = new Item();
    item.setArticle(article);
    req.setItem(item);

    mono.insert(req, {}, (err, resp) => {
      console.log(err, resp);
    });
  };

  const addArticles = (newArticles: Array<Article>) =>
    setArticles(state => [...state, ...newArticles]);

  // the subscription is a side-effect for the rendering,
  // we want to avoid to do a new request with every render
  useEffect(() => {
    const req = new ListArticles.Request();

    const resp = mono.listArticles(req);
    resp.on("data", (resp: ListArticles.Response) => {
      const list = resp.getArticlesList();
      addArticles(list);
    });
  }, [true]);

  return (
    <div className="App">
      <Pane
        display="flex"
        border="default"
        margin={16}
        padding={16}
        alignItems="center"
        justifyContent="center"
      >
        <ArticlesList articles={articles} />
      </Pane>
      <Pane
        display="flex"
        border="default"
        justifyContent="center"
        margin={16}
        padding={16}
      >
        <form onSubmit={handleSubmit}>
          <TextInputField
            name="body"
            label="Create new item"
            placeholder="Body of the item ..."
            onChange={handleInputChange}
          />
        </form>
      </Pane>
    </div>
  );
}

export default App;
