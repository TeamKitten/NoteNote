import React from "react";
import ReactDOM from "react-dom";
import "./assets/styles/reset.css";
import AppRouter from "./AppRouter";

ReactDOM.render(
  <React.StrictMode>
    <AppRouter />
  </React.StrictMode>,
  document.getElementById("root")
);
