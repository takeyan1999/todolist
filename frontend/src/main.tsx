import React from "react";
import { createRoot } from "react-dom/client";
import "./css/reset.css";
import "./css/style.css";

import App from "./App";

const rootElem = document.getElementById("root");
if (rootElem) {
    const root = createRoot(rootElem);
    root.render(<App />);
}
