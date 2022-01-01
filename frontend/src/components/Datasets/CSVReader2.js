import React from "react";
import { CSVReader } from "react-papaparse";
import SectionTitle from "../Typography/SectionTitle";

const CSVReader2 = ({ updateDataset }) => {
  const handleOnDrop = (data) => {
    console.log("---------------------------");
    console.log(data);
    updateDataset(data);
    console.log("---------------------------");
  };

  const handleOnError = (err, file, inputElem, reason) => {
    console.log(err);
  };

  const handleOnRemoveFile = (data) => {
    console.log("---------------------------");
    console.log(data);
    updateDataset([]);
    console.log("---------------------------");
  };

  return (
    <>
      <SectionTitle>Import Dataset</SectionTitle>
      <div className="p-4 mb-8 shadow-2xl rounded-lg h-18 bg-white dark:bg-gray-800 dark:text-white">
        <CSVReader
          onDrop={handleOnDrop}
          onError={handleOnError}
          addRemoveButton
          onRemoveFile={handleOnRemoveFile}
        >
          <span>Drop CSV file here or click to upload.</span>
        </CSVReader>
      </div>
    </>
  );
};
export default CSVReader2;
