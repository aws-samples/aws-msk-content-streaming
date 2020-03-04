import React, { useState, useEffect } from "react";
import "./App.css";
import { MicroClient } from "./proto/proto_grpc_web_pb";
import { ListArticles, Article, Insert, Item } from "./proto/proto_pb";
import { Button, Pane, TextInputField } from "evergreen-ui";
import ArticlesList from "./Articles";

declare var process: {
  env: {
    REACT_APP_ENDPOINT: string;
  };
};

function App() {
  // saving input state
  const [inputs, setInputs] = useState({});
  // save articles state, this is the list of provided articles
  const [articles, setArticles] = useState<Array<Article>>([]);
  const mono = new MicroClient(process.env.REACT_APP_ENDPOINT);

  // handling the have to the form fields,
  // which is the title and the description of the article
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.persist();
    setInputs(inputs => ({ ...inputs, [e.target.name]: e.target.value }));
    console.log(inputs);
  };

  // handling the submit event to the form
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    if (e) {
      e.preventDefault();
    }

    const req = new Insert.Request();

    const article = new Article();
    article.setBody((inputs as any).body);
    article.setTitle((inputs as any).title);

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
    resp.on("status", status => {
      console.log(status);
    });
    resp.on("end", function() {
      console.log("end");
    });
  }, [true]);

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
        <form>
          <TextInputField
            name="title"
            label="Title"
            placeholder="What is the title of your post?"
            onChange={handleInputChange}
          />
          <TextInputField
            name="body"
            label="Content"
            placeholder="What do you want to say?"
            onChange={handleInputChange}
          />
          <Button onClick={handleSubmit}>Create Post</Button>
        </form>
      </Pane>
    </div>
  );
}

export default App;
