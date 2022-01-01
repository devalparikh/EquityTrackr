import React from "react";
import Editor from "react-simple-code-editor";
import { highlight, languages } from "prismjs/components/prism-core";
import "prismjs/components/prism-clike";
import "prismjs/components/prism-python";
import "prismjs/themes/prism.css"; //Example style, you can use another

function CodeEditor({ code }) {
  return (
    <div className="mx-24 my-6 flex bg-gray-100 rounded-lg">
      <Editor
        value={code}
        onValueChange={(code) => console.log(code)}
        highlight={(code) => highlight(code, languages.py)}
        padding={10}
        style={{
          fontFamily: '"Fira code", "Fira Mono", monospace',
          fontSize: 12,
          width: "100%",
        }}
      />
    </div>
  );
}
export default CodeEditor;
