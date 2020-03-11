import React from "react";
import { useFormik } from "formik";
import { MicroClient } from "./proto/proto_grpc_web_pb";
import { Article, Insert, Item } from "./proto/proto_pb";
import { Button, TextInputField, InlineAlert, toaster } from "evergreen-ui";

const defaultFields = {
  body: "",
  title: ""
};

type Props = { client: MicroClient };
type Fields = { body: string; title: string } & typeof defaultFields;
interface Errors {
  title?: string;
  body?: string;
}

// A custom validation function. This must return an object
// which keys are symmetrical to our values/initialValues
const validate = (values: Fields) => {
  const errors: Errors = {};

  if (!values.title) {
    errors.title = "You have to give your post a nice title";
  }

  if (!values.body) {
    errors.body = "You have to provide some content";
  }

  return errors;
};

const InsertForm = ({ client }: Props) => {
  const formik = useFormik({
    initialValues: {
      body: "",
      title: ""
    },
    validate,
    onSubmit: values => {
      const req = new Insert.Request();

      const article = new Article();
      article.setBody(values.body);
      article.setTitle(values.title);

      const item = new Item();
      item.setArticle(article);
      req.setItem(item);

      client.insert(req, {}, (err, _) => {
        if (err) {
          toaster.danger(`Upps! Could not publish your article.`);

          return;
        }

        toaster.success("Yeah! Your article was published.");
      });
    }
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <TextInputField
        name="title"
        label="Title"
        placeholder="What is the title of your post?"
        onChange={formik.handleChange}
      />
      {formik.errors.title ? (
        <InlineAlert intent="danger" marginBottom={16}>
          {formik.errors.title}
        </InlineAlert>
      ) : null}
      <TextInputField
        name="body"
        label="Content"
        placeholder="What do you want to say?"
        onChange={formik.handleChange}
      />
      {formik.errors.body ? (
        <InlineAlert intent="danger" marginBottom={16}>
          {formik.errors.body}
        </InlineAlert>
      ) : null}
      <Button type="submit">Create Post</Button>
    </form>
  );
};

export default InsertForm;
