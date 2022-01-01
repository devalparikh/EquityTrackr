import React from "react";
import { Button } from "@windmill/react-ui";
import { GithubIcon } from "../../icons";

function GenerateCodeButton({ generatedCode, handleGenerateCode }) {
  return (
    <div className="my-6 flex justify-center">
      <Button
        layout={generatedCode ? "primary" : "outline"}
        onClick={handleGenerateCode}
        iconRight={GithubIcon}
      >
        Generate Code
      </Button>
    </div>
  );
}
export default GenerateCodeButton;
