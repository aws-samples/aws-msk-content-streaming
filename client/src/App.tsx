import React from 'react';
import logo from './logo.svg';
import './App.css';
import * as grpcWeb from 'grpc-web';
import {MonologClient } from './proto/monolog_grpc_web_pb'
import { ListArticles } from './proto/monolog_pb'

function App() {
  const mono = new MonologClient('http://localhost:8080');
  const req = new ListArticles.Request()

  const resp = mono.listArticles(req);
  resp.on("data", (msg: any) => {
    console.log(msg)
  });

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
