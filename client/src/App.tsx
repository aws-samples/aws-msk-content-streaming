import "./App.css";
import { ListArticles, Article } from "./proto/proto_pb";
import { MicroClient } from "./proto/proto_grpc_web_pb";
import { Pane } from "evergreen-ui";
import { toaster } from "evergreen-ui";
import ArticlesList from "./Articles";
import InsertForm from "./Form";
import React, { useState, useEffect } from "react";

declare var process: {
  env: {
    REACT_APP_ENDPOINT: string;
  };
};

const mono = new MicroClient(process.env.REACT_APP_ENDPOINT);

function App() {
  // save articles state, this is the list of provided articles
  const [articles, setArticles] = useState<Array<Article>>([]);

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

    resp.on("error", err => {
      toaster.danger("Argh! Please reload the browser.", {
        duration: 60 * 60,
        hasCloseButton: true
      });
    });

    resp.on("status", status => {
      toaster.notify(`The client status changed to ${status}.`);
    });

    resp.on("end", function() {
      toaster.notify("Argh! The client has disconnected.");
    });
  }, []);

  return (
    <div className="App">
      <Pane
        border="default"
        margin={16}
        padding={16}
        alignItems="center"
        justifyContent="center"
      >
        <ArticlesList articles={articles} />
      </Pane>
      <Pane border="default" justifyContent="center" margin={16} padding={16}>
        <InsertForm client={mono} />
      </Pane>
    </div>
  );
}

export default App;
